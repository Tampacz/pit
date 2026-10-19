# Stripe native reporting/export for a monthly accounting handover + PayPal on Stripe (Polish JDG, SaaS)

*All sources retrieved 2026-10-19 (docs.stripe.com Markdown, stripe.com/en-pl pricing pages, support.stripe.com, paypal.com/pl). Prices are the Polish (PLN) price lists shown on stripe.com/en-pl that day.*

## 1. Dashboard financial reports for a monthly handover (reports, columns, time zone, formats, scheduling, accountant access)

### Takeaway
Stripe's free Dashboard reports already cover a monthly handover. The **Balance** report works like a bank statement: a summary, an itemized CSV of every balance change, and a payouts list. The **Payout reconciliation** report matches bank deposits to transactions, and the **Fees** report itemizes every Stripe fee with the tax on it. All three default to the previous month, can use the account's time zone, export to CSV (Balance also to PDF), and can be scheduled daily, weekly or monthly. The schedule only sends an email saying the report is ready to a Dashboard user. The simplest native route is to invite the accountant as a **View Only** (or **Accountant**) team member so they download the reports themselves. One limit: no Balance or payout report has a VAT-amount or customer-tax-ID column.

### Cited Findings
**Which reports exist (free)**
- Stripe lists two "free reports": **Balance** ("Reconcile your Stripe balance like a bank account and download a list of payouts") and **Payout reconciliation** ("Break down the individual transactions included in each payout"). The advanced paid tools are Revenue Recognition and Bank reconciliation. — [Stripe reporting](https://docs.stripe.com/reports)
- Report-selection guide: Balance is for "Download monthly transaction history / View monthly totals by transaction category / Reconcile your Stripe balance like a bank account / Download a list of your payouts". Payout reconciliation is for "Break down the individual transactions included in each payout". The **Fees** report is for "View all fees charged by Stripe, networks, and external partners". — [Select a report](https://docs.stripe.com/reports/select-a-report)
- Each report uses a different date to decide what falls in the period:
  - Balance uses the balance transaction `created` date. Automatic payouts are the exception and use `available_on`.
  - Payout reconciliation uses `automatic_payout_effective_at`.
  - Fees uses `balance_transaction_created` or `incurred_at`.
  - [Select a report](https://docs.stripe.com/reports/select-a-report)
- The Balance summary report "works like a bank statement, helping you reconcile your Stripe balance at the end of each month… export an itemized CSV of your complete transaction history, including any associated custom metadata. Transactions appear in your settlement currency after any required foreign currency conversion." Its sections are **Balance summary** (starting and ending balance), **Balance change from activity** (all transactions except payouts: charges, refunds, disputes, adjustments, fees, by reporting category) and **Payouts**. — [Balance summary report](https://docs.stripe.com/reports/balance)
- The Payout reconciliation report needs **automatic payouts**. It groups a payout by its *estimated arrival date*, not the date the bank posts it. Its sections are Balance summary, Payout reconciliation, Failed payouts and Ending balance reconciliation (transactions not yet paid out by the end date). Instant payouts can't be reconciled to transactions. — [Payout reconciliation report](https://docs.stripe.com/reports/payout-reconciliation)
- Fees report (Dashboard: Reports > All fees): shows "the fees that Stripe, card networks, and other partners deduct from your Stripe balances". Details:
  - Summary by product, or itemized per transaction.
  - "Fee data becomes available within 96 hours after a fee affects your balance."
  - It excludes Terminal device purchases, Capital, Atlas, "and fees that Stripe invoices after you use a service".
  - The `amount` column is the fee before tax and `tax` is the tax on the fee.
  - It is "not a PDF invoice or receipt". Stripe's tax invoices are on the Documents/Invoice history page.
  - [Fees report](https://docs.stripe.com/reports/all-fees)

**Exact report IDs and columns (the Dashboard exports and the API share the same schemas)**
- **Balance summary**: `balance.summary.2` per the Balance page. The report-types page lists `balance.summary.1`, a doc inconsistency. Columns are category, currency, description, net_amount. — [Balance summary report](https://docs.stripe.com/reports/balance); [Balance report type](https://docs.stripe.com/reports/report-types/balance)
- **Balance change from activity summary** `balance_change_from_activity.summary.2`: count, currency, fee, gross, net, reporting_category. — [Balance summary report](https://docs.stripe.com/reports/balance)
- **Itemized balance change from activity** `balance_change_from_activity.itemized.7` (newest version; v1, v2, v3 and v6 are still listed). — [Balance summary report](https://docs.stripe.com/reports/balance); [Balance report type](https://docs.stripe.com/reports/report-types/balance)
  - Default columns: `balance_transaction_id, created, available_on, currency, gross, fee, net, reporting_category, description`.
  - Customer: `customer_id, customer_name, customer_email, customer_description`.
  - Customer address: `customer_address_line1/line2/city/state/postal_code/country`.
  - Shipping address: `customer_shipping_address_*` and `shipping_address_*`.
  - Card: `card_brand, card_country, card_funding, card_network, card_address_*`.
  - Payment and charge: `charge_id, charge_created, payment_intent_id, payment_method_type, is_link`.
  - Billing: `invoice_id, invoice_number, subscription_id`.
  - Original currency: `customer_facing_amount, customer_facing_currency` (the amount "as seen by the customer" when charged in a currency other than the account default).
  - Payout: `automatic_payout_id, automatic_payout_effective_at`, `trace_id`.
  - Other: `dispute_reason, statement_descriptor, source_id`.
  - Metadata: `payment_metadata[key], refund_metadata[key], transfer_metadata[key]`.
  - Withheld tax: `withheld_tax`, `fee_net_of_withheld_tax`.
  - Every timestamp also has a `*_utc` twin.
  - **No column for the VAT/tax amount charged to the customer, the customer tax ID, or an exchange rate.**
- **Itemized payouts** `payouts.itemized.5`: balance_transaction_id, payout_id, effective_at, payout_expected_arrival_date, payout_status, gross/fee/net, currency, reporting_category, description, plus optional payout_destination_id, trace_id and similar. **Payouts summary** is `payouts.summary.2`. — [Balance summary report](https://docs.stripe.com/reports/balance)
- Payout reconciliation report types — [Payout reconciliation report](https://docs.stripe.com/reports/payout-reconciliation); [Payout reconciliation report type](https://docs.stripe.com/reports/report-types/payout-reconciliation):
  - `payout_reconciliation.summary.2`
  - `payout_reconciliation.itemized.7`
  - `payout_reconciliation.by_id.itemized.4` (a single payout; required parameter `payout`)
  - `failed_payouts.summary.1` and `failed_payouts.itemized.2`
  - `ending_balance_reconciliation.summary.1` and `ending_balance_reconciliation.itemized.4` (require only `interval_end`)
- The **`exchange_rate`** column ("captures the foreign exchange rate applied to a balance transaction") is offered in `payout_reconciliation.itemized.6` and `.7`. It is not in the Balance itemized report. — [Payout reconciliation report type](https://docs.stripe.com/reports/report-types/payout-reconciliation)
- Itemized exports downloaded from the Dashboard (Balance and Payout reconciliation) add **Product IDs**, **Product Names** and **Country**. For Country: "For card payments, this is the card's country of issuance or billing address country… **For multi-region payment methods like SEPA or PayPal, it defaults to your account country.**" — [Balance summary report](https://docs.stripe.com/reports/balance); [Payout reconciliation report](https://docs.stripe.com/reports/payout-reconciliation)
- Fees itemized reports are `all_fees.balance_transaction_created.itemized.2` and `all_fees.incurred_at.itemized.2`. The summaries are `all_fees.balance_transaction_created.summary.2` and `all_fees.incurred_at.summary.2`. — [Fees report](https://docs.stripe.com/reports/all-fees); [Fees report type](https://docs.stripe.com/reports/report-types/fees)
  - Itemized columns: `amount, tax, currency, fee_description, feature_name, feature_description, product, suite, incurred_at, incurred_by, incurred_by_type (charge/refund/payout/transfer/invoice), balance_transaction_id, balance_transaction_created, fee_transaction_id, pricing_tier, settled_at, settled_via (credit/balance/v2_balance/invoice/multiple)`.
- Reporting categories merge balance transaction types. `charge` and `payment` (non-card payment methods such as PayPal) both become `charge`. `refund` and `payment_refund` become `refund`. `stripe_fee` becomes `fee`. — [Reporting categories](https://docs.stripe.com/reports/reporting-categories)
- Metadata can be added as custom report columns. — [Create custom columns using metadata](https://docs.stripe.com/reports/metadata)

**Time zone, month boundaries, data availability**
- Reports open on the **prior month** by default, and selected dates are inclusive (00:00 to 23:59 in the selected time zone). Reports can use "either your Stripe account's time zone or Coordinated Universal time (UTC)". The choice affects both the filtering and how dates are shown. — [How report configuration works](https://docs.stripe.com/reports/options)
- Financial reports default to the account's time zone, and UTC can be chosen per session. When a time zone is chosen, the old UTC datetime columns get a time-zone-adjusted duplicate (the UTC columns end in `_utc`). If earlier periods were downloaded in UTC, download the current period in both, or transactions near the period boundary can appear duplicated or missing. — [Time zone customization: changes to reports](https://support.stripe.com/questions/time-zone-customization-changes-to-reports)
- Data availability:
  - Balance and Payout reconciliation data arrive "within 12 hours"; a day's data is normally ready by 12:00 the next day.
  - Dashboard reports include only complete days.
  - The first automatic payout after a manual plan, or one more than 180 days after the previous automatic payout, can take longer.
  - [How report configuration works](https://docs.stripe.com/reports/options); [Balance summary report](https://docs.stripe.com/reports/balance)
- Reports are per **settlement currency**. Accounts with several settlement currencies (e.g. PLN and EUR) view each currency separately with a currency selector. — [How report configuration works](https://docs.stripe.com/reports/options)

**Download formats**
- Balance report — [Balance summary report](https://docs.stripe.com/reports/balance):
  - **PDF** or **CSV**. The summary PDF covers the Balance summary, Balance change from activity and Payouts sections.
  - The itemized PDF includes "the first 5,000 rows". Export CSV for all rows.
  - Each CSV holds only the section it was exported from.
- Payout reconciliation: CSV for all accounts. PDF (Summary/Itemized, up to 5,000 rows) appears only "if the download dialog for your account includes a file format selector". — [Payout reconciliation report](https://docs.stripe.com/reports/payout-reconciliation)
- No XLSX option is documented for financial reports. Stripe mentions CSV/PDF/Excel downloads only for **Revenue Recognition**. — [Stripe reporting](https://docs.stripe.com/reports)

**Scheduling and delivery**
- Scheduled reports, step by step:
  1. Open Reports hub > report > **Schedule…**.
  2. Choose daily, weekly or monthly, plus the columns and currency.
  3. "Download the report after you receive an email notification that it's ready."
  4. To stop, set the schedule to `None`.
  - [How report configuration works](https://docs.stripe.com/reports/options)

**Giving the accountant their own Dashboard access**
- **View Only** role: "for people who need to view payments, balance, and connected accounts, but can't edit any of them. This role can also export data and download reports." [User roles](https://docs.stripe.com/get-started/account/teams/roles)
  - Can: "Create, view, and download financial reports", "Bulk exports of payments, customers, and products", view payouts, balance, customers, disputes and events.
  - Can't: refund, edit customers or products, pay out, edit bank details, manage API keys or team members.
- **Accountant** role: "for accountants who need to access Stripe Revenue Recognition and its features for monthly bookkeeping". It can also "Create, view, and download financial reports" and view payments, balance, invoices, customers and payouts, but can't edit them, change pricing plans or sign up for or cancel Revenue Recognition. — [User roles](https://docs.stripe.com/get-started/account/teams/roles)
- **Analyst** role: can also refund payments, pay out balance and edit customers and products. It is broader than an external accountant needs. — [User roles](https://docs.stripe.com/get-started/account/teams/roles)

### Inferences
- A complete native monthly pack would be:
  1. The Balance summary PDF.
  2. The itemized balance-change CSV, with extra columns `customer_name, customer_email, customer_address_country, invoice_number, payment_method_type, customer_facing_amount/currency, charge_id`.
  3. The itemized payouts CSV, for matching bank deposits.
  4. The Fees itemized CSV with the `tax` column.
  5. Stripe's monthly fee tax invoice (see section 5).
  All of these are free.
- **Don't use the Dashboard "Country" column as the customer's country for PayPal payments**: it shows PL (the account country). For VAT place-of-supply evidence, use `customer_address_country` (only filled if Checkout/Billing collects an address) or PayPal payer data from the API or Sigma (section 6).
- The Polish accountant will want **Europe/Warsaw** month boundaries. Check that the account time zone is Europe/Warsaw (it defaults to where the business is registered) and stay consistent month to month, because switching between UTC and local time creates boundary gaps or duplicates.
- Stripe's scheduled reports don't email the file to an external address; they only notify a Dashboard user. A **View Only** login for the accountant, with their own monthly schedule, needs no code. A script is only worth it if the accountant must receive files without logging in, or wants Stripe, PayPal-fee and Patreon data merged into one file.
- Several months of CSV per currency is easier for an accountant to handle than PDFs: the itemized PDF is capped at 5,000 rows, which is enough at "hundreds to low thousands" of transactions per month.

### Gaps
- The docs don't say whether a scheduled-report email can go to non-team addresses, or whether it attaches the file (the wording implies a download link that needs a login).
- I couldn't confirm whether the View Only or Accountant role can open Stripe's fee tax invoices (Settings > Plans and fees / Documents).
- The Dashboard menu path for the account time-zone setting isn't documented in the pages I read. The support article only says it is changed in account settings.
- The Balance summary ID is inconsistent in Stripe's docs (`balance.summary.1` vs `balance.summary.2`).

## 2. Stripe Reporting API (report types, creating a run for a date range, downloading, permissions, lag)

### Takeaway
Every Dashboard financial report download is also available through the free Reporting API. The steps are:
1. POST `/v1/reporting/report_runs` with `report_type`, `parameters[interval_start]` (inclusive) and `parameters[interval_end]` (exclusive) as Unix timestamps, an optional IANA `timezone`, `columns[]`, `currency` and `decimal_separator`.
2. Poll the run, or wait for the `reporting.report_run.succeeded` webhook.
3. Download `result.url` (files.stripe.com) with the API key.

A restricted key needs **Financial Reports: Read** (`report_runs_and_report_types_read`). Tax report types also need **Tax Transaction Reports**.

### Cited Findings
- "All of the available downloads" for the Balance, Payout reconciliation, Fees, Tax, Connect and Commerce reports are available from the API. CSV amounts are in **major units** (`10.00`), while the API uses minor units (`1000`). — [How the Reports API works](https://docs.stripe.com/reports/api)
- Run parameters — [How the Reports API works](https://docs.stripe.com/reports/api):
  - Nearly every report type needs `interval_start` (inclusive) and `interval_end` (exclusive) as Unix timestamps, and both must lie within the type's `data_available_start`/`data_available_end`, or the API returns a 400.
  - `timezone` (an IANA name) is optional and "defaults to UTC if not supplied". It only exists on report types that have a timezone parameter.
  - `currency` and `report_category` filter rows. `columns` customizes the selection and order of columns.
- Optional parameters on `balance_change_from_activity.itemized.7` are `timezone, currency, reporting_category, decimal_separator, columns`. `payouts.itemized.5` takes the same set. — [Balance report type](https://docs.stripe.com/reports/report-types/balance)
- Example call: `curl https://api.stripe.com/v1/reporting/report_runs -d report_type=balance_change_from_activity.itemized.3 -d parameters[interval_start]=… -d parameters[interval_end]=… -d parameters[timezone]=America/Los_Angeles -d parameters[columns][]=created …`. It returns `status=pending`, then `succeeded` with `result: {object: file, url: https://files.stripe.com/v1/files/file_…/contents, type: csv}`, which you fetch with the API key. — [How the Reports API works](https://docs.stripe.com/reports/api)
- "Certain report types require a live-mode API key." — [Create a Report Run (API reference)](https://docs.stripe.com/api/reporting/report_run/create)
- Notifications and availability — [How the Reports API works](https://docs.stripe.com/reports/api):
  - Stripe sends `reporting.report_type.updated` twice a day (00:00 and 12:00 UTC). A webhook endpoint must subscribe to it explicitly; "all events" endpoints don't receive it, and sandboxes don't get it.
  - When a run finishes, Stripe sends `reporting.report_run.succeeded` or `.failed`.
  - "Most runs complete within a few minutes."
  - Too many concurrent runs can be throttled with HTTP 429.
  - Runs started from the Dashboard don't trigger these webhooks.
- Restricted-key permission: "Reporting | Financial Reports | `report_runs_and_report_types_read`… Read access lets you list report types and retrieve generated reports." — [Permissions reference](https://docs.stripe.com/stripe-apps/reference/permissions); Reports API v2 also says "grant **Read** access to **Financial Reports**" — [Reports API v2](https://docs.stripe.com/reports/v2-api)
- Tax report types: `tax.itemized_export.1`, `tax.summarized_export.1` and `tax.transactions.itemized.2`. A restricted key "must have the **Tax Transaction Reports** permission". — [Tax report type](https://docs.stripe.com/reports/report-types/tax)
- Newer APIs, all in preview:
  - **Reports API v2** is "currently in private preview" (header `Stripe-Version: 2025-07-30.preview`) and works with Organization keys and sandboxes. — [Reports API v2](https://docs.stripe.com/reports/v2-api)
  - The **Query Run API** runs Sigma SQL programmatically and "requires an active Sigma subscription". — [Query Run API](https://docs.stripe.com/reports/query-runs)
  - A new **Data API** (preview) "consolidates" the Reports API v2 and Query Run API, with a Sigma requirement for query runs. — [Access Stripe data with the API](https://docs.stripe.com/data/api)

### Inferences
- A minimal Go script would create three runs (`balance_change_from_activity.itemized.7`, `payouts.itemized.5`, `all_fees.balance_transaction_created.itemized.2`) with `timezone=Europe/Warsaw`. For September 2026 the interval would be `interval_start` = 2026-09-01 00:00 Warsaw (2026-08-31 22:00 UTC) and `interval_end` = 2026-10-01 00:00 Warsaw. The script then polls, downloads the CSVs and emails them. It's a short program, and no paid product is needed.
- Run it no earlier than about the 2nd of the month for Balance and Payout data (12-hour lag), and after about the 5th if the Fees report is included (96-hour lag).
- `decimal_separator` (e.g. a comma) can produce Polish-friendly CSV numbers without post-processing.
- The API adds no data that the Dashboard export lacks (same schemas). The only gain is automation and delivery.

### Gaps
- It isn't documented whether a restricted key also needs **Files: Read** (`file_read`) to download `result.url`. Grant it if the download returns 403. (Not verified.)
- I didn't find documented rate or volume limits beyond the 429 on concurrent runs.

## 3. Stripe Sigma (scheduled queries, pricing), Data Pipeline, Revenue Recognition and accounting integrations

### Takeaway
**Sigma** is Stripe's SQL tool for custom reports. It is the cheapest native way to get a recurring, custom CSV emailed on a schedule, and it costs from **37–56 zł/month** for up to 250 charges. Its scheduled queries email team members a download link, or fire a webhook. Note that its templates use UTC month boundaries. **Data Pipeline** (from 187–243 zł/month, warehouse sync) and **Revenue Recognition** (from 95 zł/month, accrual accounting under ASC 606/IFRS 15, two accounting standards) are overkill for a cash-basis monthly handover. The Xero and QuickBooks integrations are third-party Stripe Apps; Stripe itself doesn't support third-party reconciliation of its reports.

### Cited Findings
**Sigma: what it does**
- Sigma is a "SQL environment in the Stripe Dashboard… to create custom financial reports for accounting and reconciliation". It is read-only. The templates map to report types: "Itemized balance change from activity" → `balance_change_from_activity.itemized.3`, "Itemized payouts" → `payouts.itemized.3`, "Detailed payout reconciliation" → `payout_reconciliation.itemized.5`, "Itemized ending balance reconciliation" → `ending_balance_reconciliation.itemized.4`. — [How Sigma works](https://docs.stripe.com/data/how-sigma-works)
- Differences from the Dashboard reports — [How Sigma works](https://docs.stripe.com/data/how-sigma-works):
  - "Sigma filters templates by the UTC time zone."
  - A Sigma template range "January 13 to January 14" covers Jan 13 00:00–23:59 only (the end date is exclusive).
  - Templates return all currencies.
  - Templates don't include metadata unless you use the "Metadata to column" template.
  - Sigma is free in sandboxes.
- Fee-level data is in the Sigma `itemized_fees` table. — [Fees report](https://docs.stripe.com/reports/all-fees)

**Sigma: scheduled queries**
- Scheduled queries run "daily, weekly, or monthly". Monthly results "are usually available by 2pm UTC on the 1st of the month". Delivery — [Schedule queries](https://docs.stripe.com/data/schedule-queries):
  - The creator is subscribed by default.
  - "To notify other team members as well, enter their email addresses."
  - The email includes "a link to download the results in CSV format".
  - Subscribers can unsubscribe.
  - Alternatively, the webhook `sigma.scheduled_query_run.created` carries `file.url`, downloadable "using your live secret API key".

**Sigma pricing (Poland)**
- 30-day free trial for new users. "Charges" counts successful charges "both on Stripe and through third-party payments processors". Every plan includes "Schedule custom reports for recurring email delivery". — [Stripe Sigma pricing (PL)](https://stripe.com/en-pl/sigma/pricing)
- The price slider has five tiers. The smallest tier has both a monthly and an annual option; the rest are annual — [Stripe Sigma pricing (PL)](https://stripe.com/en-pl/sigma/pricing):

  | Monthly charges | Plan | Price | Each additional charge |
  | --- | --- | --- | --- |
  | Up to 250 | Monthly | 56 zł/month | 0.224 zł |
  | Up to 250 | Annual | 37 zł/month | 0.149 zł |
  | Up to 2,500 | Annual | 224 zł/month | 0.093 zł |
  | Up to 10,000 | Annual | 840 zł/month | 0.093 zł |
  | Up to 25,000 | Annual | 1,680 zł/month | 0.075 zł |
  | 25,000+ | Annual | 1,680 zł/month, or custom | 0.075 zł |
- Cancel under Settings > Sigma; access continues until the end of the billing cycle. — [How Sigma works](https://docs.stripe.com/data/how-sigma-works)

**Data Pipeline**
- Data Pipeline sends all Stripe data to Snowflake, Redshift, Databricks or BigQuery, or as Parquet files to S3, GCS or Azure. There is a full refresh "every 3 hours". — [How Data Pipeline works](https://docs.stripe.com/data/access-data-in-warehouse)
- Pricing (Poland) — [Data Pipeline pricing (PL)](https://stripe.com/en-pl/data-pipeline/pricing):
  - **Warehouse/cloud-storage sync** includes "complimentary access to Stripe Sigma".
    - Monthly plan: 243 zł/month for 1,000 charges, plus 0.261 zł per additional charge.
    - Annual plans: 187 zł/month for 1,000 charges (+0.224 zł); 280 zł/month for 2,500 (+0.112 zł); 1,045 zł/month for 10,000.
  - **Postgres real-time sync** (public preview): 0.009 zł per row synced with 5,000 free rows/month, or subscriptions from 370 zł/month for 50,000 rows.

**Revenue Recognition**
- Revenue Recognition automates accrual accounting (ASC 606 / IFRS 15). Reports can be downloaded as CSV, PDF or Excel. — [Stripe reporting](https://docs.stripe.com/reports); [Revenue Recognition](https://docs.stripe.com/revenue-recognition)
- PL pricing: 95 zł/month on a monthly plan for up to 37,500 zł of monthly volume, with 0.25% on additional volume. Annual plans start at 730 zł/month. 30-day free trial. Features: journal entries, balance sheet, income statement, chart of accounts mapping. — [Revenue Recognition pricing (PL)](https://stripe.com/en-pl/revenue-recognition/pricing)

**Other tools and integrations**
- Bank reconciliation "is currently only available for users with direct US-based Stripe accounts". — [Bank reconciliation](https://docs.stripe.com/bank-reconciliation)
- Accounting integrations — [Integrate an accounting tool](https://docs.stripe.com/accounting-integrations); [Accounting integrations (Stripe Apps)](https://docs.stripe.com/stripe-apps/embedded-apps/accounting-integrations); [Select a report](https://docs.stripe.com/reports/select-a-report):
  - QuickBooks: an IIF export for QuickBooks Desktop. For QuickBooks Online, a third-party integration is needed.
  - Xero and QuickBooks Online apps exist, "built by 3rd party developers".
  - The Stripe App Marketplace has an accounting category.
  - "Stripe doesn't support third-party accounting tools reconciliation for Stripe reports."

### Inferences
- For a JDG with hundreds to low thousands of charges, Sigma costs roughly 56–224 zł/month. A monthly scheduled query could emit a single "accountant CSV": PL-time month, gross/fee/net, payment method, customer country and invoice number, joined to tax data. Recipients would need to be team members, so a free Dashboard login plus the free Balance report export covers most of the same need at no cost.
- If you use Sigma, convert timestamps to `Europe/Warsaw` in SQL. The templates are UTC, so the last 1–2 hours of each Polish month would otherwise land in the wrong month.
- Data Pipeline and Revenue Recognition solve problems (a data warehouse, IFRS/ASC accrual) that a Polish JDG handing data to an external accountant doesn't have. Both are overkill.

### Gaps
- I couldn't confirm whether Sigma email recipients can be addresses outside the team, or whether the download link needs a Dashboard login.
- The Sigma tier-to-price mapping is inferred from the order of the PL pricing slider (5 tiers, 6 price cards). Confirm in Dashboard > Your plans.
- Polish accounting apps on the Stripe App Marketplace are not covered (out of scope: another researcher covers invoicing software).

## 4. Stripe Tax: EU OSS and Polish VAT support, reports/exports, 2026 pricing, worth it for a small EU B2C SaaS?

### Takeaway
Stripe Tax supports Poland (VAT, all product tax codes) and the EU Union One Stop Shop (OSS) scheme: it tracks the €10,000 EU-wide threshold and produces **itemized and summarized exports**, which Stripe says suit "country-level filings and VAT OSS". It doesn't file a Polish Union-OSS return itself. Its own automated EU filing only covers non-Union OSS in Ireland; partners such as Taxually and Marosa file EU returns. **Tax Basic** has no fixed fee: 0.5% per transaction in locations where you're registered (no-code Billing/Checkout) or 2 zł per API transaction.

### Cited Findings
- Poland: tax type VAT, supported product types "All PTCs", business and customer location both supported. The PL page links the MF (Polish Ministry of Finance) pages on VAT registration and OSS. Registration threshold "1 transaction". — [Stripe Tax – Poland](https://docs.stripe.com/tax/supported-countries/european-union/poland)
- EU rules — [Stripe Tax – European Union](https://docs.stripe.com/tax/supported-countries/european-union):
  - EU-based sellers whose cross-border EU sales stay under **10,000 EUR** a year can choose to have Stripe track the EU-wide threshold instead of individual country thresholds. Above it, you must collect VAT in customers' countries, either by registering for **Union OSS** in your home country or by registering per country.
  - Digital services are "Generally taxable in the customer's country".
  - Reverse charge applies to cross-border B2B sales where the customer has a business tax ID. "Stripe supports reverse charge only for cross-border sales, not for sales within the same country."
- Exports are in the Dashboard under Tax > **Locations**:
  - **Itemized exports** have line-item detail, including non-taxable transactions, and filters by jurisdiction. For OSS: "you can download an itemized export of all your EU transactions… can assist in preparing your VAT OSS return… Make sure to filter out domestic transactions".
  - **Summarized exports** aggregate per jurisdiction: "Use this export for country-level filings and VAT OSS".
  - Both show integration-currency and **filing-currency** amounts, with a `filing_exchange_rate`.
  - [Tax reporting](https://docs.stripe.com/tax/reports)
- API/CSV schema for `tax.itemized_export.1` — [Tax report type](https://docs.stripe.com/reports/report-types/tax):
  - Optional parameters: `country, taxable, is_eu_vat_moss, tax_type`, among others.
  - Columns: `customer_tax_id, customer_name, customer_email, invoice_number, destination_resolved_address_country, tax_rate, tax_amount, taxable_amount, taxability_reason (incl. reverse_charge), filing_currency, filing_exchange_rate, filing_tax_amount, transaction_date, type (checkout/invoice/payment_intent/credit_note/refund)`.
- `tax.summarized_export.1` columns: `country_code, jurisdiction_name, filing_currency, filing_total_sales, filing_total_tax_collected, filing_tax_payable, total_sales, total_tax_collected, total_sales_refunded`, and others. — [Tax report type](https://docs.stripe.com/reports/report-types/tax)
- Filing: "Stripe Tax offers automated filing in the US… and automated filing for selected registrations in Ireland and Canada" (Ireland means non-Union OSS registrations only). Partners: **Taxually** (filing credits included in Tax Complete) and **Marosa** ("VAT filing… Users selling in the EU"). Global filing is a private-preview waitlist. — [File and remit](https://docs.stripe.com/tax/filing)
- PL pricing — [Stripe Tax pricing (PL)](https://stripe.com/en-pl/tax/pricing):
  - **Tax Basic** ("Pay for what you use, with no recurring fees"):
    - **0.5% per transaction** where you're registered to collect, for no-code Billing/Checkout/Invoicing/Payment Links.
    - **2 zł per transaction** through the API, with 10 calculation calls included and 0.20 zł per extra call.
  - **Tax Complete** (1-year contract, with registrations and filings) is 360 / 1,700 / 4,000 / 6,000 zł per month:
    - 360 zł: 2 registrations/year, 200 transactions/month, 4 filings/year.
    - 1,700 zł: 4 registrations, 1,000 transactions, 12 filings.
  - Both plans include threshold monitoring and alerts.
  - Stripe registers for you only in Canada (federal simplified), the EU **non-Union** OSS and the US.

### Inferences
- **Once the business is OSS-registered, Stripe Tax is likely worth enabling.** For a Polish JDG selling B2C SaaS across the EU, it determines the customer's country and VAT rate, validates VAT IDs for reverse charge, stores evidence and gives the accountant a per-country OSS export with PLN/EUR filing amounts. At 0.5% of registered-jurisdiction transactions, that is far cheaper than doing it by hand. Below the threshold, or while VAT-exempt (Polish law, out of scope), monitoring alone may suffice. Under Tax Basic, transactions outside registered locations appear not to be charged.
- Without Stripe Tax, the Balance and Payout reports have **no VAT columns**. The VAT breakdown would then have to come from invoice data, i.e. Billing invoices, Sigma or a custom script.
- The accountant would still file the Polish OSS/JPK returns manually, or through Taxually/Marosa, using Stripe's exports.

### Gaps
- The docs don't say whether Stripe's PLN `filing_exchange_rate` source matches the Polish OSS conversion rule (ECB rate on the last day of the quarter). The accountant should check (Polish tax law is another researcher's area).
- Tax Basic's handling of transactions outside registered jurisdictions (no charge) is inferred from the pricing wording ("where you're registered"), not stated explicitly.

## 5. Stripe's own fee invoices (tax/VAT invoices) for EU/Polish accounts

### Takeaway
Since **1 July 2025** Stripe bills EU accounts from **Stripe Payments Europe Limited (Ireland)**. It issues a **monthly PDF tax invoice** covering Stripe's fees, available by the **10th of the following month**. No VAT is charged to Polish accounts ("intracommunity supply of services"); the customer self-assesses VAT under reverse charge. The invoice is downloaded in the Dashboard. I found no API or auto-email option. **PayPal's fees are not on it**; they come from PayPal.

### Cited Findings
- "Starting July 1, 2025, Stripe will switch to billing its services from Ireland." The issuer is Stripe Payments Europe Limited, Ireland. — [What tax invoices… after July 1, 2025](https://support.stripe.com/questions/what-tax-invoices-or-tax-documents-will-i-receive-for-fees-and-services-after-july-1-2025)
  - "VAT is not charged, being an intracommunity supply of services". You "may need to self-assess any VAT due using the reverse charge mechanism".
  - The document is a "monthly" "Official Stripe Payments Europe Limited - Ireland PDF tax invoice", and it "is **NOT** an e-invoice".
  - Location given: dashboard.stripe.com/settings/documents.
- Stripe supplies payment processing "and other services (Billing, Radar, and Tax)" from its Irish establishment. It does not charge VAT on fees to European account holders except those in Ireland, Norway, Switzerland/Liechtenstein and the UAE. — [Global taxation of Stripe fees](https://support.stripe.com/questions/global-taxation-of-stripe-fees)
- Download path: "**Settings > Plans and fees > Invoice history**". — [Download tax invoices for Stripe fees](https://support.stripe.com/questions/download-tax-invoices-for-stripe-fees)
  - Issued monthly, "available by the 10th of the following month". They cover "Stripe fees and the taxes charged on those fees".
  - None is issued if business information is missing or "No taxes were applied to your activity that month".
  - Supported locations include "Europe".
- The Fees report doc also points to "the Documents page" at `dashboard.stripe.com/settings/plans-and-fees/invoice-history`. — [Fees report](https://docs.stripe.com/reports/all-fees)
- "PayPal fees aren't included in your tax invoice from Stripe. You can access these documents from your PayPal dashboard." — [PayPal payments](https://docs.stripe.com/payments/paypal)
- Keep the business address and indirect tax number (Tax Details in Business Settings) current, because the invoices use them. — [Download tax invoices for Stripe fees](https://support.stripe.com/questions/download-tax-invoices-for-stripe-fees); [Global taxation of Stripe fees](https://support.stripe.com/questions/global-taxation-of-stripe-fees)
- The Fees report (CSV) itemizes each fee with its `tax` and is the supporting detail behind the invoice. — [Fees report](https://docs.stripe.com/reports/all-fees)

### Inferences
- The accountant needs the monthly Stripe Payments Europe PDF for each month, plus PayPal's own fee documents.
- **"No taxes were applied" → no invoice.** Since Stripe charges Polish accounts 0% VAT under reverse charge, the condition may be read literally. Check that the Polish VAT ID (NIP/EU VAT) is in Tax Details, and that an invoice actually appears each month.
- Billing, Tax and Radar fees are probably on the same monthly invoice, since they are supplied by the same Irish establishment. That is not explicitly confirmed.

### Gaps
- The Dashboard location is inconsistent across Stripe pages: "Settings > Plans and fees > Invoice history" vs "Settings > Documents" (`/settings/documents`). It probably moved; check both.
- I found no API endpoint or automatic email for Stripe's fee invoices. A script likely can't fetch them (not confirmed either way).
- The pages don't confirm whether Sigma or other subscription fees and Billing usage fees appear on the same invoice, or whether a Poland-specific format exists.

## 6. PayPal on Stripe (availability, onboarding, settlement, fees, reporting/API fields, refunds/disputes, PayPal-side documents)

### Takeaway
The user's description matches **"PayPal processed through Stripe"**, which is available to Polish Stripe accounts, with the **"Add PayPal funds to your Stripe balance"** settlement preference: payments "are immediately transferred from PayPal to your Stripe balance, without the need for you to take any action". Setup requires connecting the merchant's own PayPal account. Fees are charged twice:
- **Stripe**: 0.2% + 0.50 zł per transaction, +2% if currency conversion is needed.
- **PayPal**: PayPal's own fee under the merchant's PayPal account, 2.90% + 1.35 PLN for domestic commercial transactions.

Stripe deducts both, and they show as "fees" in Stripe reports. **Only Stripe's part is on Stripe's tax invoice**; PayPal's fee documents come from the PayPal dashboard.

### Cited Findings
**Availability and onboarding**
- The page covers "PayPal processed through Stripe, available to businesses in supported European countries". **PL is listed.** It works with Checkout, Payment Links, Elements, **Subscriptions** and **Invoicing**. Recurring payments: "Yes, but might require approval". Refunds within 180 days; disputes supported. — [PayPal payments](https://docs.stripe.com/payments/paypal)
- Presentment currencies: EUR, GBP, USD, CHF, CZK, DKK, NOK, **PLN**, SEK, AUD, CAD, HKD, NZD, SGD. Payouts follow "Standard payout timing". — [PayPal payments](https://docs.stripe.com/payments/paypal)
- Activation — [Activate PayPal payments](https://docs.stripe.com/payments/paypal/activate):
  - Availability: "EU (except for Hungary), Liechtenstein, Norway, the United Kingdom, and Switzerland".
  - Payment methods settings > PayPal > **Turn on** > choose a **Settlement preference** > **Continue to PayPal** to "connect your Stripe and PayPal accounts. You can connect an existing PayPal account or create a new one".
  - The connection can sit in a pending state until the PayPal email is confirmed or PayPal finishes its checks.
  - Switching PayPal accounts disables recurring payments until customers give new mandates.

**Settlement**
- "If you settle funds from the PayPal payments on Stripe… The funds from the payments you receive are **immediately transferred from PayPal to your Stripe balance**." Stripe settlement gives "automatic transactions reconciliation" and a Dashboard that works as for other methods. — [Choose settlement preference](https://docs.stripe.com/payments/paypal/choose-settlement-preference)
- With PayPal settlement, the Stripe balance transaction has a zero amount and Net is minus Stripe's fee. You must reconcile by hand using PayPal's settlement report. — [Choose settlement preference](https://docs.stripe.com/payments/paypal/choose-settlement-preference)
- A currency conversion happens when presentment currency ≠ settlement currency. To avoid it, add a settlement currency for each presented currency. — [Choose settlement preference](https://docs.stripe.com/payments/paypal/choose-settlement-preference)
- On Stripe settlement, "the sales amount minus PayPal fees is added to your Stripe account automatically". Changing the preference needs a support ticket (usually up to 72 hours), and funds can't be moved between the Stripe and PayPal balances. — [PayPal settlement choice (support)](https://support.stripe.com/questions/paypal-settlement-choice)
- Multi-currency settlement is available in the EU and supports up to 18 currencies. Each settlement currency needs its own bank account in that currency. — [Multi-currency settlement](https://docs.stripe.com/payouts/multi-currency-settlement)
- The PL price list shows "Multi-currency settlement… 1% of pay out volume or a minimum fee". — [Stripe pricing (PL)](https://stripe.com/en-pl/pricing)

**Fees**
- Stripe's PL pricing for PayPal: "**0.2% + 0.50zł + PayPal fees** per successful charge", "+ 2% if currency conversion is required". — [Local payment methods pricing (PL)](https://stripe.com/en-pl/pricing/local-payment-methods)
- For comparison: standard EEA cards 1.5% + 1.00 zł; premium EEA cards 1.9% + 1.00 zł; UK cards 2.5% + 1.00 zł; international cards 3.25% + 1.00 zł (+2% with conversion); BLIK 1.6% + 1.00 zł; card dispute fee 90 zł. — [Stripe pricing (PL)](https://stripe.com/en-pl/pricing)
- On refunds, Stripe does not return "the payment processing… and currency conversion fees from the original transaction". — [Stripe pricing (PL)](https://stripe.com/en-pl/pricing)
- PayPal's PL business fees (page last updated 15 July 2026) — [PayPal business fees (PL)](https://www.paypal.com/pl/business/paypal-business-fees):
  - Domestic commercial transactions: **2.90% + fixed fee**. The fixed fee is 1.35 PLN per transaction in PLN and 0.35 EUR in EUR.
  - Extra for international senders: 0% from the EEA, 1.29% from the UK, 1.99% from other markets.
  - Currency conversion: 3.00% above the base rate.
  - Dispute fee: 60 PLN standard, 120 PLN high-volume.
- PayPal fees appear in the "**Less fees** sections of Balance and Balance Activity reports". They aren't on Stripe's tax invoice: "You can access these documents from your PayPal dashboard." — [PayPal payments](https://docs.stripe.com/payments/paypal)
- The BalanceTransaction `fee_details[].type` is one of `application_fee`, **`payment_method_passthrough_fee`**, `stripe_fee`, `tax` or `withheld_tax`. — [BalanceTransaction object](https://docs.stripe.com/api/balance_transactions/object)

**How PayPal payments appear in Stripe data**
- Charge example: id `py_…`, `payment_method_details: {type: "paypal", paypal: {transaction_id (PayPal capture id), payer_id, payer_email, payer_name}}`. `billing_details` includes address and country. `payment_method_options.paypal.reference` (your order/invoice ID) appears as "Invoice ID" in PayPal's settlement report, and `transaction_id` appears as "Transaction ID". — [PayPal payout reconciliation](https://docs.stripe.com/payments/paypal/payout-reconciliation)
- The PaymentMethod object has `paypal.country` ("buyer's country… provided by PayPal"), `paypal.payer_email` and `paypal.payer_id`. — [PaymentMethod object](https://docs.stripe.com/api/payment_methods/object)
- In report exports, the `payment_method_type` column gives "The type of payment method used in the related payment". The Dashboard "Country" column defaults to the account country for PayPal. — [Balance summary report](https://docs.stripe.com/reports/balance)
- Non-card payments have balance transaction type `payment` / `payment_refund`, reported under the `charge` / `refund` reporting categories. — [Reporting categories](https://docs.stripe.com/reports/reporting-categories)

**Refunds and disputes**
- On Stripe settlement, refunds and disputes are paid from the Stripe balance. If PayPal charges a fee when a dispute closes, "it's withdrawn from your Stripe balance through a Balance Transaction of type `adjustment`". — [Choose settlement preference](https://docs.stripe.com/payments/paypal/choose-settlement-preference)
- Dispute process — [Disputed PayPal payments](https://docs.stripe.com/payments/paypal/disputed-payments):
  - Buyers can dispute within 180 days.
  - Evidence is submitted in Stripe within 2–19 days, and PayPal aims to decide within 30 days.
  - Talking to the buyer and filing appeals must be done in PayPal.
  - "PayPal might charge fees for disputes… Stripe doesn't charge any additional fees for PayPal disputes."
  - One PayPal case can show up as several Stripe disputes.

**PayPal-side reports**
- PayPal's Settlement Report is in the PayPal business account under Activity > All Reports > Transactions > Settlement, or by sFTP on request. — [PayPal payout reconciliation](https://docs.stripe.com/payments/paypal/payout-reconciliation)

### Inferences
- **The user's setup is "PayPal processed through Stripe" with Stripe settlement.** For accounting, each PayPal sale is an ordinary Stripe transaction in the Balance and Payout reports (`payment_method_type` = paypal, charge id `py_…`). Its `fee` is Stripe's 0.2% + 0.50 zł plus PayPal's fee, probably shown as a `payment_method_passthrough_fee` line. The money then reaches the bank in the normal Stripe payout.
- The accountant still needs **PayPal's own fee documents/statements** (from paypal.com) to support the PayPal part of the fees, because Stripe's invoice excludes it.
- The PayPal account probably also shows each transaction and its automatic transfer to Stripe, since payments are processed in the merchant's PayPal account and then swept. This is not documented by Stripe.
- To avoid Stripe's 2% conversion on EUR/USD PayPal payments, present PLN only, or add EUR/USD settlement currencies with matching bank accounts. Weigh the latter against the 1% multi-currency settlement fee now shown on the PL price list.
- For VAT and customer-country evidence on PayPal sales, use `paypal.country` or `billing_details.address.country` from the API or Sigma, not the Dashboard export's "Country" column.

### Gaps
- Stripe doesn't document how long the PayPal-to-Stripe transfer takes beyond "immediately", or whether funds sit in "pending" under the normal payout schedule. The doc says only "Standard payout timing applies".
- I couldn't confirm whether the "PayPal fees" in Stripe's price are PayPal's standard PL schedule (2.90% + 1.35 PLN) or a partner-specific rate. It depends on the merchant's PayPal account.
- I couldn't confirm whether PayPal returns its fee on refunds, or what form PayPal's fee invoice takes for Polish business accounts (monthly invoice vs. statements).
- The `payment_method_passthrough_fee` → PayPal mapping is inferred; I found no explicit Stripe statement.

## 7. Customer-facing documents from Stripe Billing/Invoicing (receipts, invoice PDFs, numbering, customer tax IDs)

### Takeaway
Stripe creates receipts for every successful payment and refund and can email them automatically. Billing and Invoicing produce invoice PDFs and hosted invoice pages, with sequential numbering across the whole account by default for EU accounts. Customer tax IDs can be stored, including `eu_vat` (checked against VIES, the EU VAT number database) and `pl_nip` (not validated). Subscriptions cost 0.7% of Billing volume; one-off invoices cost 0.4–0.5% per paid invoice. Whether these documents satisfy Polish invoicing and KSeF (the national e-invoicing system) rules is out of scope here.

### Cited Findings
- "Stripe creates receipts for all successful payments and refunds, including invoice payments and recurring payments as part of a subscription." — [Receipts and paid invoices](https://docs.stripe.com/receipts)
  - Automatic email receipts are configured under Settings > Business > Customer emails. Manual receipts can be sent from Transactions > Payments.
  - After an invoice is paid, the customer gets an email summary "includ[ing] links to download PDFs of both the invoice and payment receipt".
- Invoice lifecycle: draft → finalized (`open`) → paid. Most details can't be changed after finalization. Customer emails can be sent at each stage. — [Invoicing overview](https://docs.stripe.com/invoicing/overview)
- Invoice numbering: either per customer (customer prefix + sequence) or "Sequentially across your account (account-level)". "European Union member countries and the United Kingdom typically require account level sequencing". Changes apply only to future invoices. — [Customize invoices](https://docs.stripe.com/invoicing/customize)
- Customer tax IDs (Poland): `eu_vat` "European VAT number" (PL1234567890), validated: Yes, automatically against VIES. `pl_nip` "Polish NIP number" (1234567890), validated: No. — [Customer tax IDs](https://docs.stripe.com/billing/taxes/tax-ids)
- Invoice numbers and IDs link to payments in reports: `invoice_id`/`invoice_number` in the Balance itemized report; `invoice_number`/`customer_tax_id` in Tax exports. — [Balance summary report](https://docs.stripe.com/reports/balance); [Tax report type](https://docs.stripe.com/reports/report-types/tax)
- PL pricing — [Billing pricing (PL)](https://stripe.com/en-pl/billing/pricing); [Invoicing pricing (PL)](https://stripe.com/en-pl/invoicing/pricing); [Stripe pricing (PL)](https://stripe.com/en-pl/pricing):
  - **Billing** pay-as-you-go: 0.7% of Billing volume ("Excludes one-off invoices").
  - **Invoicing** Starter: 0.4% per paid invoice. Plus: 0.5%.
  - "Post-payment invoices" (an invoice generated automatically after a one-off purchase): 0.4% of the transaction total, capped at US$2.00 per invoice.

### Inferences
- If the SaaS uses Stripe Billing subscriptions, every payment already has an invoice number. The `invoice_number` column in the Balance itemized export lets the accountant link each payment to the customer-facing document.
- For EU B2B buyers, collecting `eu_vat` (VIES-validated) is what makes Stripe Tax apply reverse charge.

### Gaps
- I didn't verify which columns the Dashboard's Invoices list CSV export contains (e.g. per-invoice tax and customer tax ID). It could be an alternative source for a VAT breakdown without Stripe Tax.
- Polish legal adequacy of Stripe invoices and receipts, and KSeF, are deliberately not covered here (another researcher's area).
