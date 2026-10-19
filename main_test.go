package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestRatesAndDates(t *testing.T) {
	rates := []rate{{Date: "2026-09-25", Mid: 3.7}, {Date: "2026-09-28", Mid: 3.8}} // Friday, Monday
	for day, want := range map[string]float64{"2026-09-27": 3.7, "2026-09-28": 3.7, "2026-09-29": 3.8} {
		if r, err := rateBefore(rates, day); err != nil || r.Mid != want {
			t.Errorf("rateBefore(%s) = %v, %v; want %v", day, r.Mid, err, want)
		}
	}
	if _, err := rateBefore(rates, "2026-09-25"); err == nil {
		t.Error("rateBefore must fail without an earlier table")
	}
	for _, c := range []struct {
		minor    int64
		currency string
		mid      float64
		want     int64
	}{
		{1000, "USD", 3.8445, 3845},            // 3844,5 gr rounds up
		{-1000, "USD", 3.8445, -3845},          // refunds round the same way
		{1000, "JPY", 0.02464, 2464},           // ¥1000: Stripe sends whole yen
		{100_000_000, "IDR", 0.0002161, 21610}, // 1 mln IDR needs all 7 decimals of the rate
	} {
		if got := toPLN(c.minor, c.currency, c.mid); got != c.want {
			t.Errorf("toPLN(%d %s) = %d, want %d", c.minor, c.currency, got, c.want)
		}
	}
	for income, want := range map[int64]int64{30_000_00: 0, 40_000_00: 1_200_00, 130_000_00: 14_000_00, -500_00: 0} {
		if got := pitScale(income); got != want { // 12% minus 3 600 zł, 32% above 120 000 zł
			t.Errorf("pitScale(%d) = %d, want %d", income, got, want)
		}
	}
	jan1 := time.Date(2026, 1, 1, 0, 0, 0, 0, warsaw)
	if got, _ := exemptionLimitFor(false, time.Date(2026, 5, 30, 2, 39, 0, 0, warsaw), jan1); got != 14_202_739 { // 216 of 365 days
		t.Errorf("prorated limit = %d", got)
	}
	if got, _ := exemptionLimitFor(true, time.Date(2026, 5, 30, 2, 39, 0, 0, warsaw), jan1); got != exemptionLimit {
		t.Errorf("full limit = %d", got)
	}
	if got := day(time.Date(2026, 8, 31, 22, 30, 0, 0, time.UTC).In(warsaw)); got != "2026-09-01" { // already September in Warsaw
		t.Errorf("Warsaw day = %s", got)
	}
}

func TestFindPatreon(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	os.Mkdir(filepath.Join(home, "Downloads"), 0o755)
	csv := "Month – successful transactions,Currency,Total gross earnings,Patreon platform fees,Total payment processing fees," +
		"Merch costs (items + delivery),Refunds,Patreon adjustments,Recovered payments,Total net earnings\n" +
		"2026-09,PLN,299.34,-29.95,-28.78,0.00,-10.00,0.00,0.00,230.61\n"
	os.WriteFile(filepath.Join(home, "Downloads", "patreon_earnings_breakdown_by_month_2026-05-01_to_2026-10-19.csv"), []byte(csv), 0o644)

	if p, ok := findPatreon("2026-09"); !ok || p.revenue != 28934 || p.costs != -5873 || p.net != 23061 {
		t.Errorf("findPatreon = %+v, %v", p, ok)
	}
	if p, ok := findPatreon("2026-07"); ok || p.file == "" {
		t.Errorf("a month without a row must report the file but no data: %+v, %v", p, ok)
	}
}

func TestFetchEntries(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer rk_test" {
			t.Error("request without the API key")
		}
		switch r.URL.Path {
		case "/balance_transactions":
			fmt.Fprint(w, `{"has_more": false, "data": [
				{"id": "txn_1", "amount": 3843, "fee": 412, "currency": "pln", "created": 1790000000, "source": {"object": "charge", "id": "py_1",
					"amount": 999, "currency": "usd", "created": 1790000000, "payment_method_details": {"type": "paypal", "paypal": {"country": "DE"}}}},
				{"id": "txn_2", "amount": -3800, "currency": "pln", "created": 1792000000, "source": {"object": "refund", "id": "re_1",
					"amount": 999, "currency": "usd", "charge": "ch_1"}},
				{"id": "txn_3", "amount": -412, "currency": "pln", "created": 1790000000, "reporting_category": "fee", "source": null},
				{"id": "txn_4", "amount": -100000, "currency": "pln", "created": 1791000000, "source": {"object": "payout", "id": "po_1"}}]}`)
		case "/charges/ch_1":
			fmt.Fprint(w, `{"id": "ch_1", "created": 1789000000, "payment_method_details": {"type": "card", "card": {"country": "US"}}}`)
		default:
			t.Errorf("unexpected request %s", r.URL.Path)
		}
	}))
	defer srv.Close()
	stripeAPI = srv.URL + "/"

	got, fees := fetchEntries("rk_test", time.Unix(0, 0), time.Unix(2000000000, 0))
	if fees["2026-09"] != 824 { // the charge's fee plus a fee charged on its own
		t.Errorf("fees = %v", fees)
	}
	if len(got) != 2 {
		t.Fatalf("got %d entries, want the sale and the refund", len(got))
	}
	if s := got[0]; s.id != "py_1" || s.method != "paypal" || s.country != "DE" || s.amount != 999 || s.stripePLN != 3843 {
		t.Errorf("sale = %+v", s)
	}
	if r := got[1]; r.kind != "zwrot" || r.id != "ch_1" || r.country != "US" || r.amount != -999 || r.stripePLN != -3800 || r.sold.Unix() != 1789000000 {
		t.Errorf("refund = %+v", r)
	}
}
