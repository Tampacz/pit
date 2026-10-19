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
// The summary ends with a rough check of the month's taxes on the tax scale (PIT advance, health
// contribution, VAT-9M) to sanity-check the accountant's figures. Purchase invoices for it go into
// ./koszty.csv as data,opis,kwota,waluta,import, with import "tak" when the invoice has no VAT.
//
// The key comes from STRIPE_SECRET_KEY in the environment or in ./.env; a restricted key with Read
// on "Balance Transaction Sources", "Charges and Refunds" and "Payment Disputes" is enough.
// Usage: go run . [RRRR-MM] for a month (default: the previous one), go run . RRRR for the annual summary.
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
	minHealth       = 432_54     // grosze; ponytail: minimum for Feb 2026 – Jan 2027, update every February
	jdgStart        = "2026-09"  // first month booked as JDG; earlier months were unregistered activity
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
	nbpAPI    = "https://api.nbp.pl/api/"    // swapped in tests

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
	ID                string `json:"id"`
	Amount            int64  `json:"amount"`
	Fee               int64  `json:"fee"`
	Currency          string `json:"currency"`
	Created           int64  `json:"created"`
	ReportingCategory string `json:"reporting_category"`
	Source            *struct {
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
	now := time.Now().In(warsaw)
	month := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, warsaw).AddDate(0, -1, 0).Format("2006-01")
	if len(os.Args) == 2 {
		month = os.Args[1]
	}
	if len(os.Args) > 2 || key == "" {
		log.Fatal("użycie: go run . [RRRR-MM | RRRR] (bez argumentu: poprzedni miesiąc; klucz Stripe w STRIPE_SECRET_KEY, w środowisku albo w .env)")
	}
	if len(month) == 4 {
		runYear(key, month)
		return
	}
	start, err := time.ParseInLocation("2006-01", month, warsaw)
	if err != nil {
		log.Fatal(err)
	}
	end := start.AddDate(0, 1, 0)
	yearStart := time.Date(start.Year(), 1, 1, 0, 0, 0, 0, warsaw)
	if time.Now().Before(end) {
		fmt.Printf("UWAGA: %s jeszcze trwa, dane są niepełne.\n\n", month)
	}

	// Last year and this year so far feed the VAT thresholds and the tax check; only the month is written.
	all, fees := fetchEntries(key, yearStart.AddDate(-1, 0, 0), end)
	var entries []entry
	for _, e := range all {
		if !e.at.Before(start) {
			entries = append(entries, e)
		}
	}
	if len(entries) == 0 {
		log.Fatalf("brak płatności w %s", month)
	}
	sort.Slice(entries, func(i, j int) bool { return entries[i].at.Before(entries[j].at) })
	convert(entries)

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

	printVATThresholds(all, yearStart)

	patreon, ok := findPatreon(month)
	switch {
	case ok:
		path := filepath.Join("out", month+"_patreon.csv")
		writeCSV(path, [][]string{
			{"Miesiąc", "Przychód PLN (brutto po zwrotach)", "Prowizje Patreon PLN", "Netto PLN", "Źródło"},
			{month, money(patreon.revenue), money(patreon.costs), money(patreon.net), filepath.Base(patreon.file)},
		})
		fmt.Printf("\nPatreon %s: przychód %s zł (prowizje %s zł, netto %s zł)\n", month, money(patreon.revenue), money(patreon.costs), money(patreon.net))
	case patreon.file == "":
		fmt.Printf("\nPatreon: brak eksportu w Pobranych. Wyeksportuj go: %s → Export data (CSV przyjdzie mailem).\n", patreonInsights)
	default:
		fmt.Printf("\nPatreon: brak %s w %s. Jeśli nie było wpływów, to w porządku; inaczej wyeksportuj ponownie: %s\n",
			month, filepath.Base(patreon.file), patreonInsights)
	}
	fmt.Printf("Przychód %s razem: Stripe %s zł + różnice kursowe %s zł + Patreon %s zł = %s zł\n", month,
		money(total.pln), money(total.stripe-total.pln), money(patreon.revenue), money(total.stripe+patreon.revenue))

	// A rough check of the month's taxes on the tax scale, to compare with the accountant's figures.
	cumStart := taxYearStart(yearStart)
	if start.Before(cumStart) {
		fmt.Printf("\n%s to działalność nierejestrowana: przychód wchodzi do rocznego PIT, bez zaliczek.\n", month)
	} else {
		costs := readCosts(cumStart, end)
		var t monthTax
		var before, cum int64 // income since cumStart up to the previous month and up to this one
		for m := cumStart; m.Before(end); m = m.AddDate(0, 1, 0) {
			t = monthTaxFor(m.Format("2006-01"), all, fees, costs)
			before, cum = cum, cum+t.income
		}
		pit := pitScale(cum) - pitScale(before)
		health := max((t.income*9+50)/100, minHealth)
		vat := ((t.fees+t.imported)*23 + 5000) / 10000 * 100 // full złoty
		fmt.Printf("\nKontrolka %s (szacunek do porównania z księgową, nie zastępuje jej):\n", month)
		fmt.Printf("  przychód  %s zł (Stripe z różnicami kursowymi %s, Patreon %s)\n", money(t.stripe+t.patreon), money(t.stripe), money(t.patreon))
		fmt.Printf("  koszty    %s zł (prowizje Stripe i Patreon %s, faktury z koszty.csv %s)\n", money(t.fees+t.purchases), money(t.fees), money(t.purchases))
		fmt.Printf("  dochód    %s zł (od %s narastająco %s zł)\n", money(t.income), cumStart.Format("2006-01"), money(cum))
		fmt.Printf("  zaliczka PIT (skala, z kwotą zmniejszającą 3 600 zł): %s zł do 20. dnia następnego miesiąca\n", money(pit))
		fmt.Printf("  składka zdrowotna od tego dochodu (9%%, min. %s zł): %s zł w ZUS za następny miesiąc\n", money(minHealth), money(health))
		fmt.Printf("  VAT-9M (import usług, najwyżej 23%% od prowizji i faktur bez VAT): %s zł do 25. dnia następnego miesiąca\n", money(vat))
		fmt.Println("  założenia: prowizje Stripe w miesiącu pobrania, bez składek społecznych (ulga na start)")
	}

	// Anything else saved as out/RRRR-MM_* (e.g. purchase invoices) goes into the same package.
	fmt.Println("\nDo wysłania księgowej:")
	files, _ := filepath.Glob(filepath.Join("out", month+"_*"))
	for _, f := range files {
		fmt.Println("  " + f)
	}
	fmt.Println("  faktura Stripe za prowizje: https://dashboard.stripe.com/settings/plans-and-fees/invoice-history (do 10. dnia następnego miesiąca)")
}

// runYear writes out/RRRR_roczne.csv for the annual PIT-36: each month's revenue, costs and income, with the months
// before jdgStart marked as unregistered activity (other sources), and a rough scale tax on the whole year.
func runYear(key, year string) {
	yearStart, err := time.ParseInLocation("2006", year, warsaw)
	if err != nil {
		log.Fatal(err)
	}
	end := yearStart.AddDate(1, 0, 0)
	if time.Now().Before(end) {
		fmt.Printf("UWAGA: %s jeszcze trwa, dane są niepełne.\n\n", year)
	}
	all, fees := fetchEntries(key, yearStart.AddDate(-1, 0, 0), end)
	var entries []entry
	for _, e := range all {
		if !e.at.Before(yearStart) {
			entries = append(entries, e)
		}
	}
	convert(entries)
	costs := readCosts(yearStart, end)

	type total struct{ nbp, fx, patreon, costs, income int64 }
	periods := map[string]*total{"nierejestrowana": {}, "JDG": {}}
	rows := [][]string{{"Miesiąc", "Okres", "Przychód Stripe PLN (kurs NBP)", "Różnice kursowe PLN", "Przychód Patreon PLN", "Koszty PLN", "Dochód PLN"}}
	row := func(label, period string, t *total) []string {
		return []string{label, period, money(t.nbp), money(t.fx), money(t.patreon), money(t.costs), money(t.income)}
	}
	for m := yearStart; m.Before(end); m = m.AddDate(0, 1, 0) {
		ms := m.Format("2006-01")
		t := monthTaxFor(ms, entries, fees, costs)
		if t.stripe == 0 && t.patreon == 0 && t.fees == 0 && t.purchases == 0 {
			continue
		}
		mt := total{patreon: t.patreon, costs: t.fees + t.purchases, income: t.income}
		for _, e := range entries {
			if e.at.Format("2006-01") == ms {
				mt.nbp += e.pln
			}
		}
		mt.fx = t.stripe - mt.nbp
		period := "JDG"
		if m.Before(taxYearStart(yearStart)) {
			period = "nierejestrowana"
		}
		rows = append(rows, row(ms, period, &mt))
		p := periods[period]
		p.nbp, p.fx, p.patreon, p.costs, p.income = p.nbp+mt.nbp, p.fx+mt.fx, p.patreon+mt.patreon, p.costs+mt.costs, p.income+mt.income
	}

	fmt.Printf("Rok %s, kurs NBP z dnia roboczego poprzedzającego płatność:\n", year)
	for _, period := range []string{"nierejestrowana", "JDG"} {
		p := periods[period]
		if *p == (total{}) {
			continue
		}
		rows = append(rows, row("RAZEM", period, p))
		fmt.Printf("  %-15s przychód %s zł (Stripe %s + różnice kursowe %s + Patreon %s), koszty %s zł, dochód %s zł\n",
			period, money(p.nbp+p.fx+p.patreon), money(p.nbp), money(p.fx), money(p.patreon), money(p.costs), money(p.income))
	}
	if periods["nierejestrowana"].income != 0 {
		fmt.Println("  dochód z działalności nierejestrowanej idzie do PIT-36 jako inne źródła, razem z dochodem z JDG")
	}
	tax := pitScale(periods["nierejestrowana"].income + periods["JDG"].income)
	fmt.Printf("  podatek wg skali od łącznego dochodu (z kwotą zmniejszającą, bez innych dochodów, np. z etatu): %s zł, minus zapłacone zaliczki\n", money(tax))
	printVATThresholds(all, yearStart)

	if err := os.MkdirAll("out", 0o755); err != nil {
		log.Fatal(err)
	}
	path := filepath.Join("out", year+"_roczne.csv")
	writeCSV(path, rows)
	fmt.Printf("\nZapisano %s\n", path)
}

// taxYearStart is where PIT advances start counting: 1 January, or the JDG start in its first year.
func taxYearStart(yearStart time.Time) time.Time {
	jdg, err := time.ParseInLocation("2006-01", jdgStart, warsaw)
	if err != nil {
		log.Fatal(err)
	}
	if jdg.Before(yearStart) {
		return yearStart
	}
	return jdg
}

// convert sets each entry's NBP rate from the last business day before its sale and its PLN value.
// A refund is corrected at the rate of the sale it reverses, so rates start before the oldest sale.
func convert(entries []entry) {
	spans := map[string][2]time.Time{}
	for _, e := range entries {
		s, ok := spans[e.currency]
		if !ok || e.sold.Before(s[0]) {
			s[0] = e.sold
		}
		if !ok || e.sold.After(s[1]) {
			s[1] = e.sold
		}
		spans[e.currency] = s
	}
	rates := map[string][]rate{}
	for cur, s := range spans {
		rates[cur] = fetchRates(cur, s[0].AddDate(0, 0, -10), s[1])
	}
	for i := range entries {
		e := &entries[i]
		r, err := rateBefore(rates[e.currency], day(e.sold))
		if err != nil {
			log.Fatal(err)
		}
		e.rate, e.pln = r, toPLN(e.amount, e.currency, r.Mid)
	}
}

// printVATThresholds sums sales by VAT area since 1 January (and last year) against the EU threshold and the
// exemption limit. ponytail: it sums Stripe's PLN, within ~0.3% of NBP, enough to see the thresholds coming.
func printVATThresholds(all []entry, yearStart time.Time) {
	var euLastYear, euThisYear, plThisYear int64
	soldLastYear, firstSale := false, yearStart.AddDate(1, 0, 0)
	for _, e := range all {
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
	}
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
}

// monthTax is one month's income on the tax scale, in grosze. Stripe counts at the PLN it credited, which
// equals revenue at the NBP rate plus the exchange-rate differences.
type monthTax struct {
	stripe, patreon int64 // revenue
	fees, purchases int64 // Stripe and Patreon fees; invoices from koszty.csv
	imported        int64 // the part of purchases that is import of services (VAT-9M)
	income          int64
}

func monthTaxFor(month string, all []entry, stripeFees map[string]int64, costs []cost) monthTax {
	t := monthTax{fees: stripeFees[month]}
	for _, e := range all {
		if e.at.Format("2006-01") == month {
			t.stripe += e.stripePLN
		}
	}
	if p, ok := findPatreon(month); ok {
		t.patreon, t.fees = p.revenue, t.fees-p.costs
	}
	for _, c := range costs {
		if c.day.Format("2006-01") == month {
			t.purchases += c.pln
			if c.imported {
				t.imported += c.pln
			}
		}
	}
	t.income = t.stripe + t.patreon - t.fees - t.purchases
	return t
}

// pitScale is the tax-scale PIT on income since 1 January after the 3 600 zł reducing amount, in grosze;
// the base and the tax are rounded to full złoty.
func pitScale(income int64) int64 {
	base := (income + 50) / 100
	tax := (base*12 + 50) / 100
	if base > 120_000 {
		tax = 14_400 + ((base-120_000)*32+50)/100
	}
	return max(tax-3_600, 0) * 100
}

type cost struct {
	day      time.Time
	pln      int64 // grosze
	imported bool  // no VAT on the invoice: import of services, VAT-9M is due
}

// readCosts reads purchase invoices dated in [from, to) from ./koszty.csv, converting foreign ones at the NBP rate
// from the business day before the invoice date (art. 11a). No file means no invoices.
func readCosts(from, to time.Time) []cost {
	if _, err := os.Stat("koszty.csv"); errors.Is(err, fs.ErrNotExist) {
		return nil
	}
	var costs []cost
	for _, r := range readCSV("koszty.csv")[1:] { // data,opis,kwota,waluta,import
		d, err := time.ParseInLocation("2006-01-02", r[0], warsaw)
		if err != nil {
			log.Fatalf("koszty.csv: %v", err)
		}
		if d.Before(from) || !d.Before(to) {
			continue
		}
		c := cost{day: d, pln: cents(r[2]), imported: r[4] == "tak"}
		if cur := strings.ToUpper(r[3]); cur != "PLN" {
			rt, err := rateBefore(fetchRates(cur, d.AddDate(0, 0, -10), d), day(d))
			if err != nil {
				log.Fatal(err)
			}
			c.pln = toPLN(c.pln, cur, rt.Mid)
		}
		costs = append(costs, c)
	}
	return costs
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
func fetchEntries(key string, from, to time.Time) ([]entry, map[string]int64) {
	q := url.Values{"limit": {"100"}, "expand[]": {"data.source"},
		"created[gte]": {strconv.FormatInt(from.Unix(), 10)}, "created[lt]": {strconv.FormatInt(to.Unix(), 10)}}
	var entries []entry
	fees := map[string]int64{} // per RRRR-MM, including PayPal's fees taken through Stripe
	for {
		var page struct {
			Data    []balanceTx `json:"data"`
			HasMore bool        `json:"has_more"`
		}
		getJSON(stripeAPI+"balance_transactions?"+q.Encode(), key, &page)
		for _, bt := range page.Data {
			at := time.Unix(bt.Created, 0).In(warsaw)
			fees[at.Format("2006-01")] += bt.Fee
			if bt.ReportingCategory == "fee" { // Billing, Radar and other fees charged on their own
				fees[at.Format("2006-01")] -= bt.Amount
			}
			if bt.Source == nil {
				continue
			}
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
			return entries, fees
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

// fetchRates returns NBP table A mid rates for [from, to] in date order. NBP answers at most 367 days at once,
// so longer spans are fetched in halves.
func fetchRates(code string, from, to time.Time) []rate {
	if to.Sub(from) > 300*24*time.Hour {
		mid := from.Add(to.Sub(from) / 2)
		return append(fetchRates(code, from, mid), fetchRates(code, mid.AddDate(0, 0, 1), to)...)
	}
	var body struct{ Rates []rate }
	getJSON(fmt.Sprintf("%sexchangerates/rates/a/%s/%s/%s/?format=json", nbpAPI, strings.ToLower(code), day(from), day(to)), "", &body)
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
