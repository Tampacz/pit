// Command pit turns a month of Stripe activity into the revenue documents for the accountant.
//
// Each payment is converted to PLN at the NBP table A mid rate from the last business day before
// it reached the Stripe balance, in Warsaw time (art. 11a ustawy o PIT). The gap to the PLN amount
// Stripe actually credited is reported as an exchange-rate difference (art. 24c). Refunds and
// chargebacks are booked when they happen, at the rate of the sale they reverse. PayPal processed
// through Stripe is part of this data, so it must not be added again from PayPal.
//
// It also watches two VAT thresholds: EUR 10 000 (42 000 zł) a year of sales to consumers in other
// EU countries, this year or last, above which VAT is due in the customer's country (OSS), and the
// subject exemption limit of 240 000 zł, prorated by days in the year sales start (art. 113).
// Patreon sales are outside both: they are services to Patreon, Inc. in the USA.
//
// Patreon has no API for earnings, so its row for the month comes from the newest earnings export in
// ~/Downloads (Insights → Earnings → Export data; Patreon emails the CSV), already in PLN.
//
// The key comes from STRIPE_SECRET_KEY in the environment or in ./.env; a restricted key with Read
// on "Balance Transaction Sources", "Charges and Refunds" and "Payment Disputes" is enough.
// Usage: go run . 2026-09
package main

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"maps"
	"math"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	euThreshold     = 42_000_00  // grosze
	exemptionLimit  = 240_000_00 // grosze
	patreonInsights = "https://www.patreon.com/insights/earnings"
)

var (
	warsaw = func() *time.Location {
		loc, err := time.LoadLocation("Europe/Warsaw")
		if err != nil {
			log.Fatal(err)
		}
		return loc
	}()
	client    = &http.Client{Timeout: 30 * time.Second}
	stripeAPI = "https://api.stripe.com/v1/" // swapped in tests

	// EU member states other than Poland.
	eu = strings.Fields("AT BE BG CY CZ DE DK EE ES FI FR GR HR HU IE IT LT LU LV MT NL PT RO SE SI SK")
	// Stripe currencies without minor units that NBP table A quotes; Stripe amounts in them are whole units.
	zeroDecimal = []string{"CLP", "JPY", "KRW"}
)

type rate struct {
	No   string  `json:"no"`
	Date string  `json:"effectiveDate"`
	Mid  float64 `json:"mid"`
}

type place struct {
	Country string `json:"country"`
}

type stripeCharge struct {
	ID                   string `json:"id"`
	Created              int64  `json:"created"`
	PaymentMethodDetails struct {
		Type   string `json:"type"`
		Card   place  `json:"card"`
		Link   place  `json:"link"` // country of the card behind a Link payment
		Paypal place  `json:"paypal"`
	} `json:"payment_method_details"`
	BillingDetails struct {
		Address place `json:"address"`
	} `json:"billing_details"`
}

type balanceTx struct {
	ID       string `json:"id"`
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
	Created  int64  `json:"created"`
	Source   *struct {
		stripeCharge
		Object   string `json:"object"`
		Amount   int64  `json:"amount"`
		Currency string `json:"currency"`
		Charge   string `json:"charge"` // refunds and disputes point at the charge they reverse
	} `json:"source"`
}

type entry struct {
	at, sold  time.Time // when it hit the Stripe balance and when the sale it belongs to did (they differ for refunds), Warsaw time
	id, kind  string
	method    string
	country   string
	currency  string
	amount    int64 // minor units of currency, negative for refunds
	stripePLN int64 // grosze Stripe credited
	rate      rate
	pln       int64 // grosze at the NBP rate
}

type sum struct {
	n           int
	amount      map[string]int64 // per currency
	pln, stripe int64
}

func (s *sum) add(e entry) {
	if s.amount == nil {
		s.amount = map[string]int64{}
	}
	s.n++
	s.amount[e.currency] += e.amount
	s.pln += e.pln
	s.stripe += e.stripePLN
}

func main() {
	log.SetFlags(0)
	key := secret("STRIPE_SECRET_KEY")
	if len(os.Args) != 2 || key == "" {
		log.Fatal("użycie: go run . RRRR-MM (klucz Stripe w STRIPE_SECRET_KEY, w środowisku albo w .env)")
	}
	month := os.Args[1]
	start, err := time.ParseInLocation("2006-01", month, warsaw)
	if err != nil {
		log.Fatal(err)
	}
	end := start.AddDate(0, 1, 0)
	yearStart := time.Date(start.Year(), 1, 1, 0, 0, 0, 0, warsaw)
	if time.Now().Before(end) {
		fmt.Printf("UWAGA: %s jeszcze trwa, dane są niepełne.\n\n", month)
	}

	// Last year and this year so far feed the VAT thresholds; only the month is converted and written.
	// ponytail: thresholds sum Stripe's PLN, within ~0.3% of NBP, enough to see them coming
	var entries []entry
	var euLastYear, euThisYear, plThisYear int64
	soldLastYear, firstSale := false, end
	for _, e := range fetchEntries(key, yearStart.AddDate(-1, 0, 0), end) {
		if e.at.Before(yearStart) {
			soldLastYear = true
			if region(e.country) == "UE" {
				euLastYear += e.stripePLN
			}
			continue
		}
		if e.at.Before(firstSale) {
			firstSale = e.at
		}
		switch region(e.country) {
		case "UE":
			euThisYear += e.stripePLN
		case "Polska":
			plThisYear += e.stripePLN
		}
		if !e.at.Before(start) {
			entries = append(entries, e)
		}
	}
	if len(entries) == 0 {
		log.Fatalf("brak płatności w %s", month)
	}
	sort.Slice(entries, func(i, j int) bool { return entries[i].at.Before(entries[j].at) })

	// A refund is corrected at the rate of the sale it reverses, so rates start before the oldest sale.
	first, last := entries[0].sold, entries[0].sold
	for _, e := range entries {
		if e.sold.Before(first) {
			first = e.sold
		}
		if e.sold.After(last) {
			last = e.sold
		}
	}
	rates := map[string][]rate{}
	for i := range entries {
		e := &entries[i]
		if rates[e.currency] == nil {
			rates[e.currency] = fetchRates(e.currency, day(first.AddDate(0, 0, -10)), day(last))
		}
		r, err := rateBefore(rates[e.currency], day(e.sold))
		if err != nil {
			log.Fatal(err)
		}
		e.rate, e.pln = r, toPLN(e.amount, e.currency, r.Mid)
	}

	var total sum
	daily, byMethod, byRegion := map[string]*sum{}, map[string]*sum{}, map[string]*sum{}
	txRows := [][]string{{"Data i godzina (Warszawa)", "ID płatności Stripe", "Metoda", "Rodzaj", "Kraj", "Obszar VAT", "Kwota", "Waluta",
		"Tabela NBP", "Data kursu", "Kurs NBP", "Przychód PLN (kurs NBP)", "Wpływ PLN (Stripe)", "Różnica kursowa PLN"}}
	for _, e := range entries {
		txRows = append(txRows, []string{e.at.Format("2006-01-02 15:04:05"), e.id, method(e.method), e.kind, e.country, region(e.country),
			units(e.amount, e.currency), e.currency, e.rate.No, e.rate.Date, strings.Replace(strconv.FormatFloat(e.rate.Mid, 'f', -1, 64), ".", ",", 1),
			money(e.pln), money(e.stripePLN), money(e.stripePLN - e.pln)})
		total.add(e)
		addTo(daily, day(e.at), e)
		addTo(byMethod, method(e.method), e)
		addTo(byRegion, region(e.country), e)
	}
	txRows = append(txRows, []string{"RAZEM", "", "", "", "", "", "", "", "", "", "", money(total.pln), money(total.stripe), money(total.stripe - total.pln)})

	dayRows := [][]string{{"Data", "Liczba transakcji", "Przychód PLN (kurs NBP)", "Wpływ PLN (Stripe)", "Różnica kursowa PLN"}}
	for _, d := range slices.Sorted(maps.Keys(daily)) {
		s := daily[d]
		dayRows = append(dayRows, []string{d, strconv.Itoa(s.n), money(s.pln), money(s.stripe), money(s.stripe - s.pln)})
	}
	dayRows = append(dayRows, []string{"RAZEM", strconv.Itoa(total.n), money(total.pln), money(total.stripe), money(total.stripe - total.pln)})

	if err := os.MkdirAll("out", 0o755); err != nil {
		log.Fatal(err)
	}
	txPath := filepath.Join("out", month+"_stripe_transakcje.csv")
	dayPath := filepath.Join("out", month+"_stripe_dzienne.csv")
	writeCSV(txPath, txRows)
	writeCSV(dayPath, dayRows)

	fmt.Printf("Stripe %s (czas warszawski), kurs NBP z dnia roboczego poprzedzającego płatność:\n", month)
	for _, group := range []map[string]*sum{byMethod, byRegion} {
		for _, k := range slices.Sorted(maps.Keys(group)) {
			printSum(k, group[k])
		}
	}
	printSum("razem", &total)

	fmt.Printf("\nKonsumenci z UE poza Polską od 1 stycznia: %s zł z progu %s zł (10 000 euro)\n", money(euThisYear), money(euThreshold))
	domestic := plThisYear + euThisYear // ponytail: keeps EU sales past the EU threshold in, which only warns earlier
	switch {
	case euLastYear > euThreshold:
		fmt.Println("UWAGA: próg UE przekroczony w zeszłym roku, więc od 1 stycznia sprzedaż do UE podlega VAT w kraju klienta (OSS).")
		domestic = plThisYear
	case euThisYear*10 >= euThreshold*8:
		fmt.Println("UWAGA: ponad 80% progu UE. Powyżej VAT płaci się w kraju klienta (OSS), ustal to z księgową.")
	}
	limit, prorated := exemptionLimitFor(soldLastYear, firstSale, yearStart)
	fmt.Printf("Sprzedaż w Polsce (PL i UE do progu) od 1 stycznia: %s zł z limitu zwolnienia z VAT %s zł%s\n", money(domestic), money(limit), prorated)
	if domestic*10 >= limit*8 {
		fmt.Println("UWAGA: ponad 80% limitu zwolnienia z VAT. Po jego przekroczeniu trzeba się zarejestrować jako czynny podatnik VAT.")
	}

	files := []string{txPath, dayPath}
	patreon, ok := findPatreon(month)
	switch {
	case ok:
		path := filepath.Join("out", month+"_patreon.csv")
		writeCSV(path, [][]string{
			{"Miesiąc", "Przychód PLN (brutto po zwrotach)", "Prowizje Patreon PLN", "Netto PLN", "Źródło"},
			{month, money(patreon.revenue), money(patreon.costs), money(patreon.net), filepath.Base(patreon.file)},
		})
		files = append(files, path)
		fmt.Printf("\nPatreon %s: przychód %s zł (prowizje %s zł, netto %s zł)\n", month, money(patreon.revenue), money(patreon.costs), money(patreon.net))
	case patreon.file == "":
		fmt.Printf("\nPatreon: brak eksportu w Pobranych. Wyeksportuj go: %s → Export data (CSV przyjdzie mailem).\n", patreonInsights)
	default:
		fmt.Printf("\nPatreon: brak %s w %s. Jeśli nie było wpływów, to w porządku; inaczej wyeksportuj ponownie: %s\n",
			month, filepath.Base(patreon.file), patreonInsights)
	}
	fmt.Printf("Przychód %s razem: Stripe %s zł + różnice kursowe %s zł + Patreon %s zł = %s zł\n", month,
		money(total.pln), money(total.stripe-total.pln), money(patreon.revenue), money(total.stripe+patreon.revenue))

	fmt.Println("\nDo wysłania księgowej:")
	for _, f := range files {
		fmt.Println("  " + f)
	}
	fmt.Println("  faktura Stripe za prowizje: https://dashboard.stripe.com/settings/plans-and-fees/invoice-history (do 10. dnia następnego miesiąca)")
}

type patreonMonth struct {
	file                string
	revenue, costs, net int64 // grosze
}

// findPatreon reads month from the newest Patreon earnings export in ~/Downloads; ok is false without a row for it.
func findPatreon(month string) (p patreonMonth, ok bool) {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	exports, _ := filepath.Glob(filepath.Join(home, "Downloads", "patreon_earnings_breakdown_by_month_*.csv"))
	var newest time.Time
	for _, f := range exports {
		if fi, err := os.Stat(f); err == nil && fi.ModTime().After(newest) {
			p.file, newest = f, fi.ModTime()
		}
	}
	if p.file == "" {
		return p, false
	}
	recs := readCSV(p.file)
	col := map[string]int{}
	for i, h := range recs[0] {
		col[h] = i
	}
	for _, r := range recs[1:] {
		if r[0] != month { // first column: "Month – successful transactions"
			continue
		}
		field := func(name string) string {
			i, ok := col[name]
			if !ok {
				log.Fatalf("%s: brak kolumny %q, Patreon zmienił format eksportu", p.file, name)
			}
			return r[i]
		}
		amount := func(names ...string) (sum int64) {
			for _, n := range names {
				sum += cents(field(n))
			}
			return sum
		}
		if c := field("Currency"); c != "PLN" {
			log.Fatalf("%s: waluta %s, oczekiwano PLN", p.file, c)
		}
		p.revenue = amount("Total gross earnings", "Refunds", "Patreon adjustments", "Recovered payments")
		p.costs = amount("Patreon platform fees", "Total payment processing fees", "Merch costs (items + delivery)")
		p.net = amount("Total net earnings")
		if p.revenue+p.costs != p.net {
			log.Fatalf("%s: kolumny za %s nie sumują się do netto, sprawdź eksport", p.file, month)
		}
		return p, true
	}
	return p, false
}

func readCSV(path string) [][]string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	recs, err := csv.NewReader(f).ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return recs
}

func cents(s string) int64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatalf("kwota %q: %v", s, err)
	}
	return int64(math.Round(f * 100))
}

// exemptionLimitFor is the subject VAT exemption limit, prorated by days in the year sales start (art. 113 ust. 9).
func exemptionLimitFor(soldLastYear bool, firstSale, yearStart time.Time) (int64, string) {
	if soldLastYear {
		return exemptionLimit, ""
	}
	from := time.Date(firstSale.Year(), firstSale.Month(), firstSale.Day(), 0, 0, 0, 0, warsaw)
	yearEnd := yearStart.AddDate(1, 0, 0)
	days := func(a, b time.Time) int64 { return int64(math.Round(b.Sub(a).Hours() / 24)) }
	return exemptionLimit * days(from, yearEnd) / days(yearStart, yearEnd), " (proporcjonalnie od pierwszej sprzedaży " + day(from) + ")"
}

// fetchEntries lists the money customers moved (sales, refunds, chargebacks) that hit the Stripe balance in [from, to).
func fetchEntries(key string, from, to time.Time) []entry {
	q := url.Values{"limit": {"100"}, "expand[]": {"data.source"},
		"created[gte]": {strconv.FormatInt(from.Unix(), 10)}, "created[lt]": {strconv.FormatInt(to.Unix(), 10)}}
	var entries []entry
	for {
		var page struct {
			Data    []balanceTx `json:"data"`
			HasMore bool        `json:"has_more"`
		}
		getJSON(stripeAPI+"balance_transactions?"+q.Encode(), key, &page)
		for _, bt := range page.Data {
			if bt.Source == nil {
				continue
			}
			at := time.Unix(bt.Created, 0).In(warsaw)
			e := entry{at: at, sold: at, kind: "sprzedaż", currency: strings.ToUpper(bt.Source.Currency), amount: bt.Source.Amount, stripePLN: bt.Amount}
			sale := bt.Source.stripeCharge
			switch bt.Source.Object {
			case "charge":
			case "refund":
				e.kind = "zwrot"
			case "dispute":
				e.kind = "chargeback"
			default:
				continue // fees, payouts, transfers
			}
			if e.kind != "sprzedaż" { // the reversed sale sets the rate, country and method
				getJSON(stripeAPI+"charges/"+bt.Source.Charge, key, &sale)
				e.sold = time.Unix(sale.Created, 0).In(warsaw)
			}
			if bt.Currency != "pln" {
				log.Fatalf("%s: saldo Stripe w %q, oczekiwano pln", bt.ID, bt.Currency)
			}
			if bt.Amount < 0 {
				e.amount = -e.amount
			}
			e.id, e.method, e.country = sale.ID, sale.PaymentMethodDetails.Type, countryOf(sale)
			entries = append(entries, e)
		}
		if !page.HasMore {
			return entries
		}
		q.Set("starting_after", page.Data[len(page.Data)-1].ID)
	}
}

// countryOf is the payer's country as the payment provider reports it (card issuer, PayPal account),
// which alone documents the place of supply below EUR 100 000 a year; the billing address is the fallback.
func countryOf(c stripeCharge) string {
	pm := c.PaymentMethodDetails
	for _, cc := range []string{pm.Card.Country, pm.Link.Country, pm.Paypal.Country, c.BillingDetails.Address.Country} {
		if cc != "" {
			return cc
		}
	}
	return ""
}

func region(country string) string {
	switch {
	case country == "PL":
		return "Polska"
	case slices.Contains(eu, country):
		return "UE"
	case country == "":
		return "?"
	}
	return "poza UE"
}

func method(t string) string {
	switch t {
	case "card":
		return "karta"
	case "link":
		return "Link"
	case "paypal":
		return "PayPal"
	}
	return t
}

func addTo(m map[string]*sum, key string, e entry) {
	if m[key] == nil {
		m[key] = &sum{}
	}
	m[key].add(e)
}

func printSum(label string, s *sum) {
	var amounts []string
	for _, c := range slices.Sorted(maps.Keys(s.amount)) {
		amounts = append(amounts, units(s.amount[c], c)+" "+c)
	}
	fmt.Printf("  %-8s %4d płatności  %s  →  %s PLN wg NBP (Stripe: %s PLN, różnica kursowa %s PLN)\n",
		label, s.n, strings.Join(amounts, ", "), money(s.pln), money(s.stripe), money(s.stripe-s.pln))
}

// secret reads name from the environment, falling back to a NAME=value line in ./.env.
func secret(name string) string {
	if v := os.Getenv(name); v != "" {
		return v
	}
	data, err := os.ReadFile(".env")
	if err != nil && !errors.Is(err, fs.ErrNotExist) {
		log.Fatal(err)
	}
	for line := range strings.SplitSeq(string(data), "\n") {
		k, v, ok := strings.Cut(strings.TrimSpace(line), "=")
		if ok && strings.TrimSpace(strings.TrimPrefix(k, "export ")) == name {
			return strings.Trim(strings.TrimSpace(v), `"'`)
		}
	}
	return ""
}

func fetchRates(code, from, to string) []rate {
	var body struct{ Rates []rate }
	getJSON(fmt.Sprintf("https://api.nbp.pl/api/exchangerates/rates/a/%s/%s/%s/?format=json", strings.ToLower(code), from, to), "", &body)
	return body.Rates
}

// getJSON decodes a GET response; key, when set, authorizes the request to Stripe.
func getJSON(rawURL, key string, out any) {
	req, err := http.NewRequest(http.MethodGet, rawURL, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "pit/1.0") // NBP answers 403 to Go's default User-Agent
	if key != "" {
		req.Header.Set("Authorization", "Bearer "+key)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		log.Fatalf("%s: %s\n%s", req.URL.Path, resp.Status, body)
	}
	if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
		log.Fatal(err)
	}
}

// rateBefore returns the last table published strictly before day (YYYY-MM-DD); rates are sorted by date.
func rateBefore(rates []rate, day string) (rate, error) {
	i := sort.Search(len(rates), func(i int) bool { return rates[i].Date >= day })
	if i == 0 {
		return rate{}, fmt.Errorf("brak tabeli NBP przed %s", day)
	}
	return rates[i-1], nil
}

// toPLN converts an amount in the currency's minor units to grosze at a per-unit rate, rounding half away from zero.
// The rate is scaled by 10^8 because NBP quotes small currencies (IDR, KRW) to 7 decimal places.
func toPLN(minor int64, currency string, mid float64) int64 {
	div := int64(1_000_000)
	if !slices.Contains(zeroDecimal, currency) {
		div *= 100
	}
	g := minor * int64(math.Round(mid*1e8))
	if g < 0 {
		return -((-g + div/2) / div)
	}
	return (g + div/2) / div
}

// writeCSV writes semicolon-separated UTF-8 with a BOM, which Polish Excel opens as is.
func writeCSV(path string, rows [][]string) {
	f, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.WriteString("\xef\xbb\xbf"); err != nil {
		log.Fatal(err)
	}
	w := csv.NewWriter(f)
	w.Comma = ';'
	if err := w.WriteAll(rows); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func day(t time.Time) string { return t.Format("2006-01-02") }

// money formats minor units with a decimal comma.
func money(v int64) string {
	sign := ""
	if v < 0 {
		sign, v = "-", -v
	}
	return fmt.Sprintf("%s%d,%02d", sign, v/100, v%100)
}

// units formats an amount in the currency's minor units.
func units(v int64, currency string) string {
	if slices.Contains(zeroDecimal, currency) {
		return strconv.FormatInt(v, 10)
	}
	return money(v)
}
