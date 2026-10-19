# Automating the monthly Stripe → księgowa handover: build (Go) vs. no-code vs. existing tools (incl. NBP API)

Research date: 2026-10-19. Live API behaviour (NBP, Go module proxy, GitHub) was tested directly on that date; results marked "live test".

## Go implementation: stripe-go SDK (version, pagination, BalanceTransactions vs. Reporting API), restricted keys, multi-currency/refunds/disputes/payouts

### Takeaway
The current SDK is `github.com/stripe/stripe-go/v87` (v87.0.0, released 2026-10-01, pins API `2026-09-30.endive`, needs Go 1.24+). There are two workable data sources. (a) A Reporting API run of `balance_change_from_activity.itemized.3` (or `.7`) with `timezone=Europe/Warsaw` returns a ready CSV that already has `customer_facing_amount`/`customer_facing_currency` (the original currency of each charge, refund or dispute) and `payment_method_type`. This means less code, but it needs a live key with write access to report runs. (b) Listing `/v1/balance_transactions` with a `created` range and `expand[]=data.source` is fully read-only and sandbox-testable, but you resolve the original currency yourself.

### Cited Findings
- **Current version.** The latest stripe-go major is v87: `v87.0.0`, published 2026-10-01T00:45Z. The latest v86 is `v86.4.2` (2026-09-09). There is no v88 yet. — [Go module proxy, v87 @latest](https://proxy.golang.org/github.com/stripe/stripe-go/v87/@latest); [v86 @latest](https://proxy.golang.org/github.com/stripe/stripe-go/v86/@latest) (live test)
- The repo had 2,644 stars and was last pushed 2026-10-19 (GitHub API, live test). — [stripe/stripe-go](https://github.com/stripe/stripe-go)
- **v87.0.0 changes.** It "changes the pinned API version to `2026-09-30.endive`". Breaking changes: decimal API fields moved from `float64` to `decimal.Decimal` (adds `github.com/shopspring/decimal` as a public dependency), and support for Go 1.22/1.23 was dropped ("Go 1.24 is now the minimum supported Go version"). — [stripe-go CHANGELOG](https://github.com/stripe/stripe-go/blob/master/CHANGELOG.md)
- **Major-version cadence (first `.0.0` of each major):** v80 2024-10-01, v81 2024-10-29, v82 2025-04-01, v83 2025-09-30, v84 2025-11-18, v85 2026-03-26, v86 2026-06-06, v87 2026-10-01. That is 4 majors in the last 12 months (live test). — [Go module proxy, e.g. v85.0.0.info](https://proxy.golang.org/github.com/stripe/stripe-go/v85/@v/v85.0.0.info)
- **Pinned API versions across recent majors:** `basil` (2025-04…08), `clover` (2025-09-30…2026-02-25), `dahlia` (2026-03-25…2026-08-26), `endive` (2026-09-30). Minor SDK releases bump the pinned API version roughly monthly. — [stripe-go CHANGELOG](https://github.com/stripe/stripe-go/blob/master/CHANGELOG.md)
- **SDK policy.** "Each SDK version uses the API version that is current at the time of its release." Also: "Older major versions of the package continue to be available for use, but won't receive any additional updates."
  - Go policy: "We support Go 1.24+". Go 1.24 is marked deprecated (EOL Feb 2026), with Stripe support dropped March 2027. "we'll drop support for the oldest supported Go version with every major API release."
  - Source: [Stripe – How SDK versioning works](https://docs.stripe.com/sdks/versioning)
- **New client pattern.** Use `sc := stripe.NewClient(apiKey)`. `stripe.Client` was introduced in v82.1. The legacy `client.API` is deprecated, and the resource pattern (`customer.List`…) is "planned for deprecation in a future release". By default the SDK performs "up to two retries" and telemetry is on (disable with `EnableTelemetry: stripe.Bool(false)`). — [stripe-go README](https://github.com/stripe/stripe-go)
- **Pagination changed in v85.0.0.** "`List` and `Search` methods using `stripe.Client` now return a `struct` instead of `Seq2`… you will need to add an additional call to `.All(ctx)`", e.g. `for c, err := range sc.V1Customers.List(ctx, nil).All(ctx) {`. — [stripe-go CHANGELOG (85.0.0 section)](https://github.com/stripe/stripe-go/blob/master/CHANGELOG.md)
  - In v87, `V1List[T]` is a struct whose `All(ctx)` "will continue to fetch pages of items as needed" (auto-pagination). — [iter.go @v87.0.0](https://github.com/stripe/stripe-go/blob/v87.0.0/iter.go)
  - **Conflict:** the README at tag v87.0.0 still shows `for pi, err := range sc.V1PaymentIntents.List(context.TODO(), params) {` without `.All(ctx)`. That form does not match the v87 signature. — [README @v87.0.0](https://github.com/stripe/stripe-go/blob/v87.0.0/README.md) vs. [balancetransaction_service.go @v87.0.0](https://github.com/stripe/stripe-go/blob/v87.0.0/balancetransaction_service.go)
- **Exact v87 signatures:**
  - `func (c v1BalanceTransactionService) List(ctx context.Context, listParams *BalanceTransactionListParams) *V1List[*BalanceTransaction]`, exposed as `sc.V1BalanceTransactions`.
  - Report runs are on `sc.V1ReportingReportRuns`, with `Create(ctx, *ReportingReportRunCreateParams)`, `Retrieve` and `List`.
  - Files are on `sc.V1Files`.
  - Sources: [balancetransaction_service.go](https://github.com/stripe/stripe-go/blob/v87.0.0/balancetransaction_service.go), [reporting_reportrun_service.go](https://github.com/stripe/stripe-go/blob/v87.0.0/reporting_reportrun_service.go), [stripe_client.go](https://github.com/stripe/stripe-go/blob/v87.0.0/stripe_client.go)
- **`BalanceTransactionListParams` (v87)** has `Created *int64`, `CreatedRange *RangeQueryParams` (`form:"created"`), `Currency`, `Expand` (plus an `AddExpand()` helper), `Payout`, `Source` and `Type`. `RangeQueryParams` has `GreaterThan`/`GreaterThanOrEqual`/`LesserThan`/`LesserThanOrEqual`, which map to `gt/gte/lt/lte`. — [balancetransaction.go @v87.0.0](https://github.com/stripe/stripe-go/blob/v87.0.0/balancetransaction.go); [params.go @v87.0.0](https://github.com/stripe/stripe-go/blob/v87.0.0/params.go)
- **List endpoint behaviour.** `created` filters by "date interval". `limit` ranges from 1 to 100 (default 10). Results come back "with the most recent transactions appearing first". `payout` filter: "For automatic Stripe payouts only, only returns transactions that were paid out on the specified payout ID". `type` includes `charge`, `payment`, `refund`, `payment_refund`, `adjustment`, `payout`, `stripe_fee`, `stripe_fx_fee` and ~50 others. — [Stripe API – List balance transactions](https://docs.stripe.com/api/balance_transactions/list)
- **Balance-transaction object fields** (from the documented sample response): `amount`, `available_on`, `created`, `currency`, `description`, `exchange_rate`, `fee`, `fee_details`, `net`, `reporting_category`, `source`, `status`, `type`. — [Stripe API – List balance transactions](https://docs.stripe.com/api/balance_transactions/list)
- **Reporting API: creating a run.**
  - Create a run with `report_type` (e.g. `balance_change_from_activity.itemized.3`), `parameters[interval_start]` (inclusive) and `parameters[interval_end]` (exclusive) as Unix timestamps, plus optional `timezone`, `columns`, `currency` and `reporting_category`.
  - The `timezone` parameter "is optional and defaults to UTC if not supplied".
  - "A live-mode API key is required."
  - Sources: [Stripe – Reports API](https://docs.stripe.com/reports/api); [reporting_reportrun.go @v87.0.0](https://github.com/stripe/stripe-go/blob/v87.0.0/reporting_reportrun.go)
- **Reporting API: getting the result.**
  - A run starts as `status="pending"`. When it is `succeeded`, `result.url` points to `https://files.stripe.com/v1/files/<id>/contents`, which is fetched with the API key.
  - "Most runs complete within a few minutes". Webhooks `reporting.report_run.succeeded`/`failed` are sent.
  - "Stripe prepares data for your reports on a semi-daily basis". Runs must fall within the report type's `data_available_start`/`data_available_end`, otherwise you get HTTP 400.
  - Stripe "might throttle requests if too many reports are running simultaneously" (HTTP 429).
  - Sandboxes don't receive `reporting.report_type.updated` events.
  - Source: [Stripe – Reports API](https://docs.stripe.com/reports/api)
- **CSV vs. API amounts.** "The CSV reports format monetary amounts in major currency units as a decimal number… This differs from the Stripe API, where you specify amounts in the currency's minor unit." — [Stripe – Reports API](https://docs.stripe.com/reports/api)
- **Balance report types.**
  - Current types include `balance_change_from_activity.itemized.1/.2/.3/.6/.7`, `balance_change_from_activity.summary.1/.2`, `balance.summary.1` and `payouts.itemized.1–.5`.
  - `itemized.3` and `itemized.7` accept `timezone`; `itemized.6` does not.
  - Itemized reports also accept `decimal_separator`.
  - Source: [Stripe – Balance report type](https://docs.stripe.com/reports/report-types/balance)
- **Itemized columns:**
  - `balance_transaction_id`, `created` ("in the requested timezone, or UTC if not provided"), `created_utc`, `available_on` ("The date the balance transaction's net funds will become available in the Stripe balance"), `reporting_category`, `fee`, `charge_id`, `charge_created`, `payment_method_type`, `dispute_reason`, `automatic_payout_effective_at`, `refund_metadata[key]`.
  - `customer_facing_amount`: "For transactions associated with charges, refunds, or disputes, the amount of the original charge, refund, or dispute. If the customer was charged in a different currency than your account's default, this field will reflect the…" [original currency amount].
  - `customer_facing_currency` gives the ISO code for that amount.
  - Source: [Stripe – Balance report type](https://docs.stripe.com/reports/report-types/balance)
- **Reporting categories.** Values in the v87 enum include `charge`, `charge_failure`, `dispute`, `dispute_reversal`, `fee`, `advance`… The report docs describe `reporting_category` as "a new categorization of balance transactions, meant to improve on the current `type` field". — [balancetransaction.go @v87.0.0](https://github.com/stripe/stripe-go/blob/v87.0.0/balancetransaction.go); [Stripe – Balance report type](https://docs.stripe.com/reports/report-types/balance)
- **Restricted keys: what they are.**
  - RAKs (`rk_live_…`) let you set "**Read**, **Write**, or **None**" per resource. "All Stripe APIs support restricted API keys."
  - Mapping: GET is read; POST/DELETE is write. "Write permissions imply read permissions".
  - Missing permissions return an error "explaining which permissions to add".
  - Stripe recommends "one restricted key per service or use case".
  - Source: [Stripe – Restricted API keys](https://docs.stripe.com/keys/restricted-api-keys)
- **Restricted keys: handling.**
  - Keys are shown once ("You can't retrieve it later").
  - Rotation gives up to a 7-day grace period.
  - "Access policies" (IP/CIDR, or ASN/country/VPN/Tor blocking) replaced IP restrictions.
  - Storage advice: "Store sensitive keys in a secrets vault… If you can't use a secrets vault, use environment variables"; "Don't put keys in source code or configuration files checked into version control."
  - Source: [Stripe – API keys](https://docs.stripe.com/keys)
- **Settlement currency.** "Before incoming currencies settle in your account, Stripe automatically converts them into the default currency of your home country." Settlement in additional currencies can be enabled in Dashboard settings and needs a bank account in that currency for payouts. A balance in a currency with no bank account is converted on manual payout. Stripe fees in a currency without a bank account are converted "at the baseline exchange rate… doesn't incur any additional fees". — [Stripe – Settle in additional currencies](https://docs.stripe.com/payments/currencies/settlement-payouts)
- **Poland pricing.**
  - Standard EEA cards "1.5% + 1.00zł"; premium EEA "1.9% + 1.00zł"; UK "2.5% + 1.00zł"; international "3.25% + 1.00zł" plus "+ 2% if currency conversion is required".
  - Dispute received fee 90.00 zł, plus a 90.00 zł "dispute countered" fee that is refunded if the dispute is won.
  - PayPal fees are not listed on the page.
  - Source: [Stripe pricing (PL)](https://stripe.com/en-pl/pricing)
- **Overkill for this size.** Stripe Data Pipeline sends data to "Snowflake, Amazon Redshift, Databricks, or BigQuery". — [Stripe – Reports API](https://docs.stripe.com/reports/api)

### Inferences
- **Minimal Go path A (least code): Reporting API.** Do one `Create` on `balance_change_from_activity.itemized.3` (or `.7`) with `interval_start`/`interval_end` set to the Warsaw month bounds, `timezone=Europe/Warsaw`, and columns `balance_transaction_id, created, reporting_category, gross, fee, net, currency, customer_facing_amount, customer_facing_currency, payment_method_type, charge_id, description`. Then poll `Retrieve` until `succeeded`, GET `result.url`, and append NBP columns.
  - Advantages: no pagination, no per-source-type logic, the same CSV an accountant may already know from the Dashboard, and `payment_method_type` separates PayPal-via-Stripe from cards.
  - Drawbacks: needs a live key; creating a run is a POST, so by Stripe's GET=read / POST=write rule the key needs **Write** on report runs and likely Read on Files (exact permission names not confirmed); data lags semi-daily, so run it on the 2nd–3rd of the month, not at 00:05 on the 1st; and it is not testable in a sandbox with real data.
- **Minimal Go path B (strictly read-only): balance transactions.**
  ```go
  sc := stripe.NewClient(os.Getenv("STRIPE_RAK")) // rk_live_… with Read-only perms
  p := &stripe.BalanceTransactionListParams{CreatedRange: &stripe.RangeQueryParams{
      GreaterThanOrEqual: from.Unix(), LesserThan: to.Unix()}} // Warsaw month bounds
  p.Limit = stripe.Int64(100)
  p.AddExpand("data.source") // Charge / Refund / Dispute → original amount+currency
  for bt, err := range sc.V1BalanceTransactions.List(ctx, p).All(ctx) { /* … */ }
  ```
  Here `bt.Amount` and `bt.Currency` are in the settlement currency (PLN unless multi-currency settlement is on), i.e. already converted at Stripe's rate (incl. the 2% fee). The NBP conversion must use the expanded source's original amount and currency, which is what the report's `customer_facing_*` columns do for you.
- **Refunds, disputes, payouts.** Each is its own balance transaction, dated when it happens. A refund in November of an October sale shows up in November's list or report as a negative row with `reporting_category=refund` (dispute / dispute_reversal similarly). Payouts (`type=payout`) are transfers to the bank, not revenue, and can be excluded from the revenue sheet. The `payout` filter or the `payouts.itemized.*` report lets the accountant match bank-statement lines.
- **Upgrade discipline.** For a monthly batch job, pin one stripe-go major in `go.mod` and upgrade deliberately, about once a year, when the Go toolchain support window forces it. Never copy README snippets blindly: the v85 `.All(ctx)` change shows the README lagging the code. Use `decimal.Decimal` or integer minor units for money, not `float64`; the v87 move to shopspring/decimal reinforces this.

### Gaps
- Exact Dashboard permission names (e.g. "Report runs and report types", "Files") required for the Reporting API with a restricted key: the RAK docs fetched don't list them; only the GET=read / POST=write rule is documented.
- What differs between `balance_change_from_activity.itemized.3`, `.6` and `.7` (column sets/defaults) was not checked.
- Whether a PL-based Stripe account can enable EUR/USD settlement: the settlement doc is generic and doesn't list PL's currencies.
- The meaning of `exchange_rate` on the balance transaction object wasn't fetched verbatim.
- No documented end-of-life for old Stripe API versions was found; only that old SDK majors "won't receive any additional updates".

## NBP Web API (api.nbp.pl): endpoints, formats, weekends/holidays, previous-business-day lookup, limits, stability

### Takeaway
Use table A, single-currency range queries: `GET https://api.nbp.pl/api/exchangerates/rates/a/{code}/{start}/{end}/?format=json`. Make one call per currency per month, covering roughly 10 days before the month starts through month-end. Then pick, for each transaction, the last `effectiveDate` strictly before the transaction's Warsaw-local date. Never call per-date endpoints in a loop. Days without a table (weekends, holidays, including 24 Dec since 2025) just return 404 "Brak danych". Two live gotchas: Go's default User-Agent `Go-http-client/1.1` was answered with HTTP 403 on 2026-10-19, and the documented 93-day limit only applies to `tables/…` (single-currency `rates/…` allows 367 days).

### Cited Findings
- **URL templates (base `https://api.nbp.pl/api/exchangerates/`):** `rates/a/{code}/`, `rates/a/{code}/today/`, `rates/a/{code}/{date}/`, `rates/a/{code}/{startDate}/{endDate}/`, `rates/a/{code}/last/{topCount}/`. For whole tables: `tables/a/`, `tables/a/today/`, `tables/a/{date}/`, `tables/a/{startDate}/{endDate}/`, `tables/a/last/{topCount}/`. — [NBP Web API docs](https://api.nbp.pl/)
- **Documented rules:**
  - A range query "nie może obejmować przedziału dłuższego, niż 93 dni" (cannot cover more than 93 days).
  - 404 is returned for no data, 400 for malformed queries, and 400 "Przekroczony limit" when the data limit is exceeded.
  - Data is available "od 2 stycznia 2002 r." (from 2 January 2002).
  - JSON is the default; select it with `?format=json` or `Accept: application/json` (XML via `?format=xml`).
  - HTTPS-only from 1 August 2025.
  - The docs say nothing about rate limiting, quotas, terms of use, or corrections of published data.
  - Source: [NBP Web API docs](https://api.nbp.pl/)
- **Live test (2026-10-19): limits as actually enforced.**
  - `tables/a/2026-01-01/2026-04-05` returned `400 BadRequest - Przekroczony limit 93 dni / Limit of 93 days has been exceeded`. — [request](https://api.nbp.pl/api/exchangerates/tables/a/2026-01-01/2026-04-05/?format=json)
  - The single-currency `rates/a/eur/2026-01-01/2026-04-30` (120 days) returned **200** with 83 rates, and a 367-day rates query also returned 200. `rates/a/eur/2022-01-01/2025-12-31` returned `400 BadRequest - Przekroczony limit 367 dni / Limit of 367 days has been exceeded`. — [120-day request](https://api.nbp.pl/api/exchangerates/rates/a/eur/2026-01-01/2026-04-30/?format=json); [4-year request](https://api.nbp.pl/api/exchangerates/rates/a/eur/2022-01-01/2025-12-31/?format=json)
  - `last/300` returned `400 BadRequest - Przekroczony limit 255 wyników / Maximum size of 255 data series has been exceeded`; `last/255` returned 200. — [request](https://api.nbp.pl/api/exchangerates/rates/a/usd/last/300/?format=json)
- **Live test: response shape.** `{"table":"A","currency":"dolar amerykański","code":"USD","rates":[{"no":"202/A/NBP/2026","effectiveDate":"2026-10-16","mid":3.8827}]}`. The JSON keys are lowercase (the docs list them capitalised). — [request](https://api.nbp.pl/api/exchangerates/rates/a/usd/2026-10-16/?format=json)
- **Live test: weekends and holidays.**
  - Saturday 2026-10-17 → `404 NotFound - Not Found - Brak danych` (plain text, not JSON). — [request](https://api.nbp.pl/api/exchangerates/rates/a/usd/2026-10-17/?format=json)
  - 2025-11-11 (Independence Day) → 404. — [request](https://api.nbp.pl/api/exchangerates/rates/a/usd/2025-11-11/?format=json)
  - 2025-12-24 (Wigilia, a public holiday from 2025) → 404. — [request](https://api.nbp.pl/api/exchangerates/rates/a/usd/2025-12-24/?format=json)
  - The range 2025-12-19…31 jumps from table `248/A/NBP/2025` (2025-12-23) to `249/A/NBP/2025` (2025-12-29). — [request](https://api.nbp.pl/api/exchangerates/rates/a/usd/2025-12-19/2025-12-31/?format=json)
  - The 2026 range skips 2026-01-06 (Epiphany). — [request](https://api.nbp.pl/api/exchangerates/rates/a/eur/2026-01-01/2026-04-30/?format=json)
- **Live test: error ambiguity and quirks.**
  - An invalid currency code (`xyz`) returns the **same** `404 NotFound - Not Found - Brak danych` as a weekend. — [request](https://api.nbp.pl/api/exchangerates/rates/a/xyz/2026-10-16/)
  - A malformed date (`2026-13-01`) returns `404 NotFound`, with a UTF-8 BOM before the text. — [request](https://api.nbp.pl/api/exchangerates/rates/a/usd/2026-13-01/)
  - `HEAD` returns `405 Method Not Allowed` (`Allow: GET`).
  - 20 rapid sequential requests all returned 200/404 normally; no 429 or throttling was seen.
- **Live test: User-Agent block.**
  - The exact header `User-Agent: Go-http-client/1.1` got **403 Forbidden**.
  - These all got 200: `Go-http-client/2.0`, empty UA, `curl/8.7.1`, `python-requests/2.32.3`, `Mozilla/5.0`, and a custom `stripe-nbp-report/1.0`.
  - A Go 1.25.5 program using `http.Get` failed with `nbp 403 Forbidden` and worked after `req.Header.Set("User-Agent", "stripe-nbp-report/1.0")`.
  - The response sets a cookie (`ee3la5eizeiY4Eix=…`), consistent with a WAF/bot filter.
  - Source: [request](https://api.nbp.pl/api/exchangerates/rates/a/usd/2026-10-16/?format=json) (live test)
- **Live test: validated lookup logic (Go, stdlib only).** A range fetch for 2025-12-10…2026-01-31 plus "last table strictly before D" gave:
  - tx 2025-12-24 → NBP 2025-12-23 (`248/A/NBP/2025`, mid 3.5848)
  - tx 2025-12-29 (Mon) → 2025-12-23 (6 calendar days back)
  - tx 2026-01-02 → 2025-12-31 (`251/A/NBP/2025`, 3.6016)
  - tx 2026-01-07 → 2026-01-05 (`002/A/NBP/2026`, 3.6045)
  - The whole check, including Warsaw month bounds and a test `main`, was 52 non-blank lines. — [request used](https://api.nbp.pl/api/exchangerates/rates/a/usd/2025-12-10/2026-01-31/?format=json)
- **Publication times.** "Tabela A kursów średnich walut obcych aktualizowana jest w każdy dzień roboczy, między godziną 11:45 a 12:15" (every business day, between 11:45 and 12:15). Table B is published on Wednesdays (or the previous day if Wednesday is a day off). Table C is published 7:45–8:15. The page says nothing about corrections. — [NBP – terminy publikacji kursów walut](https://nbp.pl/statystyka-i-sprawozdawczosc/kursy/informacja-o-terminach-publikacji-kursow-walut/)
- **Existing tools using this logic.** A Polish single-file tool for exactly this rule "Steps back to the previous day… If there is no table that day (weekend / holiday → HTTP 404), it automatically steps back another day, up to 7 days". It shows the NBP table date used "so the booking is auditable". — [kaszanisko/nbp-kursy-walut](https://github.com/kaszanisko/nbp-kursy-walut)
- **Same rule in accounting guides.** Polish guides state that foreign-currency revenue is converted at the NBP average rate from the last business day before the revenue date (tax specifics are covered by other researchers). — [stripto.pl guide](https://stripto.pl/jak-rozliczyc-stripe); [fakturaxl.pl](https://www.fakturaxl.pl/faktura-ze-stripe-jak-rozliczyc)

### Inferences
- **Robust "last business day before D".** Don't hard-code holidays and don't walk back day by day with per-date calls. Fetch one range per currency covering at least `[monthStart − 10 days, monthEnd]`; the longest observed gap is 6 days (23 → 29 Dec), and 10 days is a safe margin. Binary-search for the last `effectiveDate < D`, and panic or fail if none is found. This uses one HTTP call per currency per month, stays within both the 93-day and 367-day limits, and makes holidays (including new ones like 24 Dec) automatic.
- `D` must be the **Europe/Warsaw** calendar date of the transaction, not the UTC date. Payments between 22:00/23:00 UTC and midnight belong to the next Warsaw day.
- **Validate currency codes** against `tables/a/` once, or treat a 404 on a *range* query as an error: a range that spans business days only returns 404 for a bad code.
- **Always set a custom User-Agent in Go.** The 403 may be a WAF rule that changes over time, so fail loudly on non-200 rather than silently skipping a currency.
- **Store NBP `no` (table number) and `effectiveDate` next to each converted row.** Accountants usually cite them, and it makes the sheet auditable.
- **Historical data.** No statement on corrections was found. Saving the fetched rates with the report each month freezes what was used. Re-fetching later should return the same values, but that is not a documented guarantee.
- The `przemekperon/nbp-exchange-rates-golang` library is not needed: the stdlib code is about 30 lines. It is also a 1-star repo from a single day; see the OSS section.

### Gaps
- No official NBP statement on rate limits, acceptable use, or the User-Agent filtering. The 403 is an observation from one day and one network.
- Why the docs' 93-day rule isn't enforced for `rates/` (only for `tables/`) is undocumented and could change. Design for ≤93 days.
- No source on whether NBP ever revises published table A values.

## Delivery and scheduling: email, shared folders, repos; GitHub Actions cron, launchd/cron, Cloudflare Workers; secret storage

### Takeaway
The lazy, reliable combination is to run the binary monthly on your own Mac via `launchd`, with the key in macOS Keychain or an env file. launchd catches up after sleep, while cron silently skips. Write the CSV/XLSX into a Google Drive or Dropbox folder already shared with the księgowa, which needs zero delivery code. GitHub Actions cron in a private repo is the server-side alternative, but it can be delayed or dropped at high load, and repo secrets hold a live Stripe key. Email from Go is possible, but `net/smtp` is frozen, Gmail app passwords are discouraged, and a transactional email API (e.g. Resend's free tier) is cleaner.

### Cited Findings
- **`net/smtp` (Go stdlib).** "The smtp package is frozen and is not accepting new features. Some external packages provide more functionality." It offers `SendMail`, `PlainAuth` and STARTTLS (RFC 3207), with no MIME/attachment helpers. Checked with `go doc` on Go 1.25.5. — [pkg.go.dev/net/smtp](https://pkg.go.dev/net/smtp)
- **Gmail app passwords:**
  - "App passwords can only be used with accounts that have 2-Step Verification turned on."
  - Not available if 2SV is "set up only for security keys", if "You're logged into a work, school, or another organization account", or with "Advanced Protection".
  - "we revoke your app passwords when you change your Google Account password"
  - "App passwords aren't recommended and are unnecessary in most cases."
  - Source: [Google Account Help – Sign in with app passwords](https://support.google.com/accounts/answer/185833?hl=en)
- **Resend.** Free plan: 3,000 emails/month, "limited to 100 emails per day", 3 domains. Pro: $20/mo for 50,000 emails. The page has no date marker. — [Resend pricing](https://resend.com/pricing)
- **Google Apps Script quotas.**
  - Consumer accounts: URL Fetch 20,000/day, email recipients 100/day, triggers total runtime 90 min/day.
  - Workspace accounts: URL Fetch 100,000/day, 1,500 recipients/day, 6 h/day of trigger runtime.
  - Both: script runtime 6 min/execution and 20 triggers/user/script.
  - Limits are "subject to elimination, reduction, or change at any time, without notice".
  - Source: [Apps Script quotas](https://developers.google.com/apps-script/guides/services/quotas)
- **GitHub Actions `schedule`:**
  - "can be delayed during periods of high loads… High load times include the start of every hour. If the load is sufficiently high enough, some queued jobs may be dropped."
  - "schedule your workflow to run at a different time of the hour"
  - Minimum interval is 5 minutes.
  - "By default, scheduled workflows run in UTC. You can optionally specify a timezone using an IANA timezone string"
  - Runs only on the default branch.
  - "In a public repository, scheduled workflows are automatically disabled when no repository activity has occurred in 60 days."
  - Source: [GitHub Docs – Events that trigger workflows](https://docs.github.com/en/actions/reference/workflows-and-actions/events-that-trigger-workflows)
- **GitHub Actions secrets.**
  - Encrypted with libsodium sealed boxes, and only exposed when "you explicitly include the secret in a workflow".
  - Log redaction is not guaranteed ("the runner can only redact secrets used within the current job").
  - Environment secrets can require reviewer approval.
  - Source: [GitHub Docs – Secrets](https://docs.github.com/en/actions/concepts/security/secrets)
- **macOS launchd vs. cron.** "Unlike cron which skips job invocations when the computer is asleep, launchd will start the job the next time the computer wakes up. If multiple intervals transpire before the computer is woken, those events will be coalesced into one event upon wake from sleep." The macOS `crontab` man page says cron/crontab "are officially supported under Darwin, their functionality has been absorbed into launchd(8)". — [launchd.plist(5) man page, macOS (read locally via `man launchd.plist`)](x-man-page://5/launchd.plist)
- **Cloudflare Workers Cron Triggers.**
  - "Cron Triggers execute on UTC time"
  - Config changes take "up to 15 minutes to propagate"
  - Scheduled Workers "will run on underutilized machines"
  - Examples are only in JavaScript/TypeScript/Python; Go is not mentioned.
  - Source: [Cloudflare Docs – Cron Triggers](https://developers.cloudflare.com/workers/configuration/cron-triggers/)
- **Stripe key storage.** Stripe advises a secrets vault, or env vars if none is available, never in source control. Access policies can restrict a key to known IPs or ASNs/countries. — [Stripe – API keys](https://docs.stripe.com/keys)
- **Existing project example.** The `stripe-ksef-bridge` project injects secrets via 1Password CLI: "Sekrety przez 1Password (`op`), nie na dysku… `op run --env-file=.env -- npm run dev`". — [Czarekkk/stripe-ksef-bridge](https://github.com/Czarekkk/stripe-ksef-bridge)

### Inferences
- **Simplest delivery: a shared synced folder.**
  - The program writes `2026-10_stripe_PLN.csv` (plus a summary) into `~/Google Drive/Księgowa/` or Dropbox. The desktop sync client uploads it, and the księgowa sees it in the shared folder.
  - No SMTP, no OAuth, no attachment MIME code.
  - Add a one-line macOS notification or a calendar reminder to eyeball totals before the accountant's deadline.
- **Email, if the accountant insists.** Use Resend or a similar API: one HTTPS POST with a base64 attachment, and a free tier far above 1 email/month. The alternative is `net/smtp` + Gmail app password + hand-built multipart MIME. It works, but the password is revoked on any Google password change and is unavailable on many Workspace accounts.
- **Scheduling.**
  - On a laptop: launchd `StartCalendarInterval` (Day=3, Hour=10), because of the Stripe report data lag and NBP's 11:45–12:15 publication (not needed if only "before-D" tables are used).
  - GitHub Actions: put a `timezone: Europe/Warsaw` cron at an odd minute (e.g. `17 9 3 * *`) in a **private** repo, plus `workflow_dispatch` for manual re-runs.
  - Cloudflare Workers is a poor fit for a Go binary (JS/TS/Python examples only) and is overkill for one run a month.
- **Secret exposure.** With GitHub Actions, the Stripe RAK lives in GitHub's secret store. Use a read-only RAK (path B) and an access policy if possible. GitHub-hosted runner IPs vary, so an IP policy is impractical; an ASN/country policy might work (not verified).
- **No-server alternative.** Google Apps Script with a monthly time-driven trigger: `UrlFetchApp` to Stripe and NBP, write to a Google Sheet shared with the księgowa, `MailApp` notification. The free quotas easily cover 1 run/month with a few dozen API pages. The key goes in Script Properties. Trade-off: JavaScript instead of Go, and a 6-min execution cap.

### Gaps
- GitHub Actions free-minute allowance for private repos and Cloudflare Workers free-plan cron limits were not fetched. Both are irrelevant at 1 run/month.
- Whether NBP's WAF blocks Apps Script's `UrlFetchApp` User-Agent was not tested.
- Whether Stripe access policies can be practically scoped to GitHub-hosted runners was not verified.

## Pitfalls, and an honest build/maintenance estimate for a ~150–300 line Go program

### Takeaway
The code is small. The traps are semantic: Warsaw vs. UTC month bounds; using `created` rather than `available_on` or payout dates; refunds and disputes landing in later months; Stripe's own FX (with its 2% fee) vs. the NBP rate; the original-currency amount vs. the settlement amount; and NBP quirks (403 for Go's default UA, 404 ambiguity). Realistic build time is about 1–2 focused days plus 1–2 monthly cycles of reconciliation with the księgowa. Maintenance is about 2–8 h/year if the SDK is pinned.

### Cited Findings
- **Warsaw month bounds (Go `time` + tzdata, live test).**
  - 2026-10-01 00:00 CEST = `2026-09-30T22:00:00Z` (unix 1790805600).
  - 2026-11-01 00:00 CET = `2026-10-31T23:00:00Z` (unix 1793487600).
  - The offset changes between summer and winter months.
  - Source: [pkg.go.dev/time/tzdata](https://pkg.go.dev/time/tzdata) (own test)
- **Embedded tzdata.** `time/tzdata` "provides an embedded copy of the timezone database… if the time package cannot find tzdata files on the system, it will use this embedded information", adding about 450 KB. This matters for scratch/distroless containers and some CI images. — [pkg.go.dev/time/tzdata](https://pkg.go.dev/time/tzdata)
- **Report timezone default.** The Reporting API `timezone` "defaults to UTC if not supplied". — [Stripe – Reports API](https://docs.stripe.com/reports/api)
- **`created` vs. `available_on`.** `available_on` is when "net funds will become available in the Stripe balance", and `created` is "Time at which the balance transaction was created". For refunds and disputes, `charge_created` gives the original charge time. — [Stripe – Balance report type](https://docs.stripe.com/reports/report-types/balance)
- **Payouts are not revenue.** Polish accounting and vendor guides state that a Stripe payout is not revenue: it is net of fees and refunds and bundles many transactions. Revenue is the full amount paid by the customer, and the Stripe fee is a separate cost. — [stripto.pl](https://stripto.pl/jak-rozliczyc-stripe); [generaloffice.pl](https://generaloffice.pl/rozliczenia-stripe/)
- **Stripe FX.** Stripe charges "+ 2% if currency conversion is required" on international cards for PL accounts. Charges in non-settlement currencies are auto-converted to the home default currency. — [Stripe pricing (PL)](https://stripe.com/en-pl/pricing); [Stripe – Settle in additional currencies](https://docs.stripe.com/payments/currencies/settlement-payouts)
- **Stripe's rate has no tax meaning.** Polish guides say Stripe's conversion rate is not used for tax; NBP is used instead. — [stripto.pl](https://stripto.pl/jak-rozliczyc-stripe)
- **Original-currency amounts.** `customer_facing_amount`/`customer_facing_currency` carry the original charge, refund or dispute currency amount. — [Stripe – Balance report type](https://docs.stripe.com/reports/report-types/balance)
- **API versioning and SDK churn.**
  - The SDK pins the API version, and each major can contain breaking changes: v85 changed `List` to need `.All(ctx)`, and v87 changed decimal types and dropped Go 1.22/1.23.
  - Four majors shipped between Nov 2025 and Oct 2026.
  - "Older major versions… continue to be available for use, but won't receive any additional updates."
  - Sources: [stripe-go CHANGELOG](https://github.com/stripe/stripe-go/blob/master/CHANGELOG.md); [Go module proxy](https://proxy.golang.org/github.com/stripe/stripe-go/v87/@latest); [Stripe – SDK versioning](https://docs.stripe.com/sdks/versioning)
- **Report data lag.** Report data is prepared "on a semi-daily basis". — [Stripe – Reports API](https://docs.stripe.com/reports/api)
- **NBP quirks** (live tests above): 403 for `Go-http-client/1.1`, identical 404 text for holidays and bad currency codes, 93/367-day limits, HTTPS-only since 2025-08-01. — [NBP Web API docs](https://api.nbp.pl/)

### Inferences
- **Line-count sketch for path A or B, stdlib + stripe-go only:**

  | Part | Lines |
  | --- | --- |
  | Config/env + Warsaw month bounds | ~20 |
  | Stripe fetch (report run + poll + download, or BT list with expand) | ~40–60 |
  | CSV parse/emit | ~40 |
  | NBP range fetch + lookup (measured: about 30 lines of logic, 52 lines with a test harness) | ~35 |
  | PLN conversion with integer grosze/decimal and per-row rounding | ~20 |
  | Summary totals per currency/category | ~20 |
  | Delivery (0 for a synced folder; ~30 for Resend; ~50 for SMTP+MIME) | 0–50 |

  Total ≈ 175–245 lines, within the 150–300 estimate.
- **Build effort.** About 1–2 focused days for someone fluent in Go: day 1 for data plumbing, day 2 for edge cases and a test with real data. Add 1–2 parallel monthly runs where the output is compared with the Dashboard balance report and the accountant's own numbers before trusting it. The expensive part is agreeing the column semantics with the księgowa (which date drives the rate, how refunds are shown, whether fees are listed), not the code.
- **Yearly maintenance.**
  - About 2–8 h/year: one deliberate stripe-go bump per year (Go toolchain support drives it; v87 needs Go 1.24+, and Stripe drops Go 1.24 in March 2027).
  - Occasional fixes when NBP or the WAF changes behaviour (e.g. the UA block).
  - Re-validating whenever business changes: new currency, multi-currency settlement turned on, new payment method, VAT status change.
  - The silent-failure risk (job didn't run, partial data) is the main operational cost. A run log or "last success" file in the shared folder, plus a calendar reminder, mitigates it cheaply.
- **Pitfall checklist for the report writer:**
  1. Month bounds in Europe/Warsaw, with `interval_end` or `lt` exclusive.
  2. Use `created`, never `available_on` or payout dates, for revenue rows.
  3. Refunds and disputes of earlier sales appear as negative rows in the month they occur; keep `charge_id` and `charge_created` so the accountant can link them.
  4. Convert the **original** currency amount (customer_facing) at NBP; keep Stripe's PLN amount and fee separately. The difference partly reflects Stripe's 2% FX fee, and its tax treatment is out of scope here.
  5. Money in minor units or decimal; the CSV is in major units.
  6. Never compute NBP on the UTC date.
  7. Custom UA for NBP; fail on any non-200.
  8. Pin stripe-go; don't trust README snippets across majors.
  9. Run on the 2nd–3rd of the month, after Stripe's report data is complete.
  10. Fixed timezone in the scheduler.

### Gaps
- No empirical data was found on how long similar Go or Python Stripe→NBP scripts took others to build, or what they cost to maintain. The estimate above is reasoned from API surface and the measured NBP part, not from case studies.
- How Stripe dates PayPal-via-Stripe transactions (`created` vs. capture/settlement) was not checked; Stripe features are covered by another researcher.

## No-code/low-code alternatives: Zapier, Make, n8n (templates, Polish invoicing tools, pricing, reliability)

### Takeaway
Zapier, Make and n8n are event-driven (one run per Stripe payment). They suit "create an invoice in inFakt/Fakturownia per payment", which also covers the NBP conversion because the invoicing tools do it. They are a poor fit for "monthly summary with a per-transaction previous-business-day NBP rate", which needs loops, 404 handling and a date search. For invoicing, the native integrations are cheaper and simpler than any iPaaS: Fakturownia's built-in Stripe app, or Stripto (from 19 zł/month) for Fakturownia/inFakt/wFirma.

### Cited Findings
- **Zapier pricing (page marker `last_updated_at: 2026-10-19`).**
  - Free: $0, 100 tasks/mo, two-step Zaps only, 15-min polling, no premium apps.
  - Professional: $19.99/mo billed annually or $29.99/mo billed monthly for 750 tasks; 2-min polling.
  - Team: from $69/mo annual for 2,000 tasks.
  - Overage is 1.25× (annual) or 2.5× (monthly) the base rate.
  - Source: [Zapier pricing](https://zapier.com/pricing)
- **Zapier inFakt ↔ Stripe.**
  - inFakt actions: "Create Client", "Create Invoice", "Create Invoice (Async)", "Mark Invoice as Paid", "Send Invoice via Email", "Send Invoice to KSeF".
  - inFakt triggers: "New Invoice", "New Invoice (With Filters)".
  - Stripe triggers: "New Charge", "New Payment", "New Invoice", "New Refund", "Failed Payment"… (16 events).
  - No ready-made Stripe→inFakt Zap template was shown.
  - Source: [Zapier – inFakt + Stripe](https://zapier.com/apps/infakt/integrations/stripe)
- **Fakturownia on Zapier.** Search did not surface a Zapier page for Fakturownia. — [search results incl. Zapier inFakt page](https://zapier.com/apps/infakt/integrations/stripe)
- **Fakturownia native Stripe app.**
  - It "automatycznie wystawiaj[e] dokumenty sprzedażowe do każdej płatności przyjętej przez Stripe" (automatically issues a sales document for every payment Stripe accepts).
  - Document types: Faktura VAT, Proforma, Paragon, or mixed (invoices for businesses, receipts for consumers).
  - Setup takes Stripe's Publishable **and Secret** key, and "Fakturownia automatycznie utworzy webhook w Stripe" (creates the Stripe webhook automatically).
  - It "jest w pełni kompatybilna z Krajowym Systemem e-Faktur (KSeF)" (fully compatible with KSeF), given an active KSeF integration.
  - Buyer NIP/address requires enabling data collection in Stripe.
  - The page doesn't say how foreign currency or NBP is handled, or whether past payments are back-filled.
  - Source: [Fakturownia – Stripe app](https://fakturownia.pl/apps/stripe-app)
- **Stripto (commercial, PL).**
  - "Stripto łączy Stripe z Fakturownią, inFakt lub wFirma i przy każdej płatności automatycznie tworzy fakturę VAT" (connects Stripe with Fakturownia, inFakt or wFirma and creates a VAT invoice for each payment). Covers subscriptions and renewals and multiple Stripe accounts, with Stripe connected via OAuth.
  - "Plany zaczynają się od 19 zł miesięcznie za 25 faktur" (plans start at 19 zł/month for 25 invoices); "Pierwsze 5 transakcji jest za darmo" (the first 5 transactions are free).
  - The page attributes NBP conversion to Fakturownia/inFakt and says nothing about a monthly zestawienie for the accountant.
  - Sources: [stripto.pl](https://stripto.pl/jak-rozliczyc-stripe); [Stripto – inFakt](https://stripto.pl/stripe-infakt); [Stripto – wFirma](https://stripto.pl/stripe-wfirma); [Stripto – Fakturownia](https://stripto.pl/stripe-fakturownia)
- **Make pricing.**
  - "As of August 27, credits are the billing unit in Make" (year not shown). Free: 1,000 credits/month, 1 data store. Most actions use 1 credit. — [Make pricing](https://www.make.com/en/pricing)
  - Core ≈ $9/mo annual or $10.59/mo monthly for 10k credits; Pro $16, Teams $29 (third-party summary, Sept 2026). — [jetadmin.io – Make pricing 2026](https://www.jetadmin.io/blog/make-pricing/)
  - Zapier's blog quotes higher figures (Core $12 annual), so the sources conflict. — [zapier.com blog](https://zapier.com/blog/make-com-pricing/)
- **Make templates.** Make has a Stripe app (41 modules per a third-party guide) and a published Stripe ↔ Fakturoid (Czech) template. No inFakt/Fakturownia/wFirma module or template was confirmed. — [hackceleration – Make Stripe](https://hackceleration.com/resources/integrations/make/stripe); [Make – Stripe + Fakturoid](https://www.make.com/en/integrations/stripe/fakturoid)
- **n8n pricing** (EUR, billed annually):
  - Starter €20/mo, 2.5K executions.
  - Pro €50/mo, 10K executions.
  - Business €667/mo, 40K executions, self-hosted.
  - Community Edition "available on GitHub" (self-hosted).
  - Pricing is "based on monthly workflow executions, regardless of complexity".
  - Source: [n8n pricing](https://n8n.io/pricing/)
- **n8n Stripe node.** 5 triggers and 20 actions (third-party guide). — [eesel.ai](https://www.eesel.ai/blog/stripe-integrations-with-n8n)

### Inferences
- **Monthly NBP summary via iPaaS.** Each Stripe payment would trigger: an NBP HTTP call for D−1, retries walking back on 404 (loops are awkward in Zapier, possible in Make/n8n), and a Google Sheets append.
  - That is about 3–6 tasks or credits per payment. Hundreds to low thousands of payments/month blows past Zapier Free (100 tasks), puts you on Zapier Professional ($20–30+/mo), or uses up most or all of Make Free's 1,000 credits.
  - Refunds and disputes need separate Zaps.
  - Webhook or polling misses mean silently missing rows, with no "re-run October" button.
  - More moving parts and more cost than a 200-line batch job, and less auditable.
- **Where iPaaS or native integrations win:** per-payment invoice creation in Polish invoicing software (incl. KSeF), where the invoicing tool applies the NBP rate itself.
  - Fakturownia's native app costs nothing extra beyond the Fakturownia plan. Caveat: it asks for the Stripe **secret** key, not a restricted key.
  - Stripto adds inFakt/wFirma for 19 zł+/mo.
  - Once invoices exist in the accounting tool, the księgowa may need no separate Stripe summary beyond the fee invoice and payout reconciliation. That is a question for the tax/invoicing researchers.
- **n8n self-hosted** is "free" but adds a server to maintain: more upkeep than a launchd job. It is overkill here.

### Gaps
- Whether Stripe is a Premium app on Zapier (which would exclude the Free plan) could not be confirmed on the pages fetched.
- No official Make page with the Core price was retrieved.
- No wFirma or Fakturownia Zapier/Make connectors were confirmed.
- Fakturownia's handling of non-PLN Stripe payments (NBP date rule) is not documented on its Stripe app page.

## Existing open-source projects and small tools (Stripe → Polish accounting, NBP, KSeF, Fakturownia/inFakt/wFirma, JPK)

### Takeaway
No open-source project builds a "Stripe → NBP-converted monthly zestawienie for the księgowa"; GitHub search returned zero repos for "stripe nbp", "stripe jpk", "stripe ksiegowosc", "stripe kpir", "stripe ryczalt" and "stripe ifirma". What exists is young and tiny. There are Stripe→KSeF invoice bridges (best: `frogermcs/stripe-ksef`, 16★, created Aug 2026, EU B2B only), a Go Stripe→wFirma bridge tied to OpenCart, a 1★ Go NBP client, and a single-HTML NBP lookup tool. None is mature enough to depend on. They are useful as reference code; for invoicing, commercial integrations (Fakturownia native, Stripto, ksiegai) are the "buy" options.

### Cited Findings
- **[frogermcs/stripe-ksef](https://github.com/frogermcs/stripe-ksef)** — ★16, TypeScript, created 2026-08-22, last push 2026-08-24 (GitHub API).
  - Syncs "finalized Stripe EU B2B invoices with the Polish National e-Invoicing System (KSeF) using the official FA(3) XML schema".
  - Runs as a CLI or a Next.js API route on Vercel cron; filters for "EU B2B customers"; stores KSeF status in Stripe metadata.
  - The README says it submits to the "KSeF Sandbox".
- **[Czarekkk/stripe-ksef-bridge](https://github.com/Czarekkk/stripe-ksef-bridge)** — ★1, TypeScript, created 2026-06-14, last push 2026-08-14.
  - Pipeline: webhook (`checkout.session.completed`, `invoice.paid`) → FA(3) → KSeF; "Only domestic sales (PLN + Polish VAT id) are auto-issued — anything else (e.g. a foreign buyer…) is flagged for manual issuance".
  - "zweryfikowane na środowisku TESTOWYM KSeF" (verified on the KSeF test environment) on 2026-06-13.
  - Uses a restricted Stripe key, libSQL and MailerSend; "Zastępuje ręczne wystawianie w SaldeoSMART" (replaces manual issuance in SaldeoSMART).
- **[ruslan-hut/wfsync](https://github.com/ruslan-hut/wfsync)** — ★1, Go 1.24, created 2025-06-06, last push 2026-10-18.
  - OpenCart ↔ Stripe ↔ wFirma bridge with "Automatic invoice generation on successful payment via webhooks".
  - Needs MongoDB + MySQL and a Telegram bot. It is e-commerce specific.
- **[DamianS21/stripe-infakt](https://github.com/DamianS21/stripe-infakt)** — ★0, Python, created 2025-04-26, last push 2026-01-31. "Integration of Infakt… with Stripe"; no README content.
- **[rskonieczka/stripe_wfirma_ksef](https://github.com/rskonieczka/stripe_wfirma_ksef)** (★0, PHP, a single day 2026-03-24) and **[polindvichka/stripe_wfirma_integration](https://github.com/polindvichka/stripe_wfirma_integration)** (★0, JS, last push 2025-05-20) are tiny or abandoned. **[nocodeboring/stripe-ksef](https://github.com/nocodeboring/stripe-ksef)** (★0, created and pushed 2026-04-23) has no description or language.
- **[przemekperon/nbp-exchange-rates-golang](https://github.com/przemekperon/nbp-exchange-rates-golang)** — ★1, created and last pushed 2026-05-08.
  - Zero-dependency Go client for tables A/B/C and gold.
  - "Distinguishes 'no data' (weekend/holiday) from generic 404 via `ErrNoData`"; "Built-in retries for HTTP 429… and 5xx"; Go 1.24+.
- **[kaszanisko/nbp-kursy-walut](https://github.com/kaszanisko/nbp-kursy-walut)** — ★0, created 2026-06-08.
  - Single HTML file; implements the previous-business-day walk-back (up to 7 days); EUR/USD/KRW; Comarch format (4 dp, comma).
  - Live demo on GitHub Pages.
- **Searches with no matches.** GitHub repo searches on 2026-10-19 for "stripe nbp", "stripe jpk", "stripe ksiegowosc", "stripe ksiegowa", "stripe raport ksiegowy", "stripe kpir", "stripe ewidencja przychodow", "stripe ryczalt", "stripe ifirma", "stripe pln nbp", "paypal nbp" and "stripe accounting poland" returned **no repositories**. "nbp kurs" returned only student or recruitment exercises. — [GitHub search](https://github.com/search?q=stripe+nbp&type=repositories)
- **Commercial: ksiegai.pl** (Tovernet Sp. z o.o., software, not an accounting office).
  - "importuje sprzedaż ze Stripe, rozbija prowizje, wylicza VAT od importu usług Stripe" (imports Stripe sales, splits out fees, calculates import-of-services VAT on Stripe fees).
  - Matches payouts to the bank (±4 days); entries need approval ("Każde księgowanie nadal wymaga Twojej zgody").
  - Supports VAT-9M / JPK_V7. Pricing is not listed, and NBP conversion is not mentioned.
  - Source: [ksiegai.pl – Stripe dla SaaS](https://www.ksiegai.pl/stripe-dla-saas/)
- **Commercial: Stripto** — Stripe → Fakturownia/inFakt/wFirma invoices, from 19 zł/mo. — [stripto.pl](https://stripto.pl/jak-rozliczyc-stripe)
- **Commercial: other.** CSV2Invoice (PeerPush, Mar 2026) turns Stripe/ThriveCart CSV exports into invoices for EU SaaS (search-result summary only; page not opened). — [peerpush.com/p/csv2invoice](https://peerpush.com/p/csv2invoice)

### Inferences
- **Fit for "monthly summary for the księgowa":** none of these does it.
  - The KSeF bridges solve a different problem (issuing e-invoices). They are 2–4 months old, single-maintainer, have ≤16★, and one only reached the KSeF test environment. Adopting them means owning them.
  - `wfsync` is OpenCart-specific and infrastructure-heavy.
  - The Go NBP library adds a dependency for about 30 lines of stdlib code.
  - `kaszanisko/nbp-kursy-walut` is good proof of the walk-back rule and handy for manual spot checks, but it is not automation.
- **Practical "buy" shortlist (handover side):**
  1. If the user already uses or would adopt Fakturownia: its native Stripe app.
  2. inFakt/wFirma: Stripto (19 zł+/mo) or Zapier's inFakt actions.
  3. Full Stripe-aware bookkeeping: ksiegai.pl (price unknown).
  Each moves the NBP conversion into accounting software the księgowa already trusts.
- **Build remains justified** only if the user wants a summary sheet (not per-payment invoices), e.g. B2C SaaS where the accountant books revenue from an ewidencja/zestawienie. In that case a ~200-line Go job, or even a monthly manual CSV + NBP column, beats adopting any of these repos.

### Gaps
- npm/PyPI were not searched directly (only GitHub and the Go proxy). There may be small npm/PyPI NBP clients, but nothing Stripe-specific surfaced in web search.
- Stripto/ksiegai pricing tiers above the entry level, and real user reviews, were not found.
- No information on whether Fakturownia/inFakt, when fed by Stripe, apply the NBP rate per payment date as the tax rules require.

## How Polish indie hackers / small SaaS founders handle the monthly handover

### Takeaway
First-hand Polish founder accounts are scarce. Reddit, wykop and 4programmers searches returned nothing relevant; the discourse is dominated by accounting-firm and vendor blogs. The observable pattern is:
- Founders with B2B/KSeF needs build or buy per-payment invoice bridges (two founder-built Stripe→KSeF repos in 2026, one explicitly replacing manual SaldeoSMART invoicing).
- Others rely on the invoicing tool's integration (Fakturownia/Stripto).
- Manual CSV exports are the default pain point: vendors claim "hours" for dozens of transactions (marketing, unverified).

### Cited Findings
- **Vendor time claims.** Manual handling "przy kilkudziesięciu zaczyna zajmować godziny" (with dozens of transactions it starts taking hours); with hundreds it is "po prostu nierealne" (simply unrealistic). Self-reported marketing. — [stripto.pl](https://stripto.pl/jak-rozliczyc-stripe)
- **A described monthly process (ksiegai.pl).** Connect Stripe → import payments, fees and payouts → the system creates a Stripe period settlement → import the bank statement → match the payout → review gaps "including the Stripe balance and the Stripe invoice". — [ksiegai.pl](https://www.ksiegai.pl/stripe-dla-saas/)
- **Founder motivation for building.** One founder-built bridge replaces manual issuance: "Zastępuje ręczne wystawianie w SaldeoSMART" (replaces manual issuance in SaldeoSMART). Another README opens by naming the lack of built-in Stripe invoice synchronisation for the Polish market as the problem ("napotykamy na problem braku wbudowanej synchronizacji faktur…", i.e. "we run into the problem of no built-in invoice synchronisation"). — [Czarekkk/stripe-ksef-bridge](https://github.com/Czarekkk/stripe-ksef-bridge); [frogermcs/stripe-ksef](https://github.com/frogermcs/stripe-ksef)
- **International anecdote.** An Indie Hackers founder built a Stripe → Google Sheets sync to escape the export–clean–import routine (search-result summary; post not opened). — [Indie Hackers – "wasting time exporting CSVs from Stripe"](https://www.indiehackers.com/post/wasting-time-exporting-csvs-from-stripe-8e6e119dc1)
- **Accounting guides (rules context, covered by others).** Revenue is the full customer payment at the NBP average rate from the last business day before the revenue date. The Stripe fee is a cost (import of services, reverse-charge VAT even for VAT-exempt taxpayers). Payouts are not revenue. — [generaloffice.pl](https://generaloffice.pl/rozliczenia-stripe/); [fakturaxl.pl](https://www.fakturaxl.pl/faktura-ze-stripe-jak-rozliczyc); [ifirma.pl](https://www.ifirma.pl/blog/aktualnosci/rozliczenie-faktury-od-stripe/)
- **Open question in practice.** Whether each transaction must be converted separately or an aggregate may be used appears as an open question in a tax-interpretation search result; the answer was not retrieved. — [podatki.biz](https://www.podatki.biz/artykuly/przychod-w-walucie-obcej-a-pit_14_20894.htm)

### Inferences
- With no reliable first-hand data, the honest framing for the user is this. The **manual** route (download the Stripe itemized balance CSV, add NBP columns with a small script or the kaszanisko tool, drop it in a shared folder) is probably 15–60 min/month at hundreds of transactions if NBP is scripted, and much longer if rates are looked up by hand. This is an estimate, not sourced.
- Automating the NBP column is the highest-leverage step. Automating delivery and scheduling saves little at one run a month.
- Agree the format with the księgowa first; many accountants prefer the invoicing tool's own records (e.g. Fakturownia/inFakt with KSeF) over a custom sheet.

### Gaps
- No Reddit, wykop, 4programmers, X or LinkedIn threads from Polish SaaS founders describing their monthly Stripe→księgowa routine, time spent, or regrets were found with the searches run. This remains the weakest-sourced area.
- No measured data on manual handover time; vendor claims only.
