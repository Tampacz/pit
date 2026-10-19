# Polish income-tax and VAT record-keeping for a JDG selling SaaS via Stripe (incl. PayPal-on-Stripe): what the accountant needs each month (law as of 19 Oct 2026)

Legend: **[LAW]** = statute/regulation text, read directly from the official Dz.U. PDFs. **[OFFICIAL]** = MF answer or KIS individual interpretation (an interpretation binds only its applicant). **[PRACTICE]** = accounting offices, invoicing/bookkeeping vendors, tax portals. **[NEW 2025] / [NEW 2026]** = changed in 2025–2026. **[2027]** = scheduled or pending. Consolidated texts used: PIT act t.j. Dz.U. 2025 poz. 163 (stan na 10 Jan 2025), ryczałt act t.j. Dz.U. 2025 poz. 843 (13 Jun 2025), VAT act t.j. Dz.U. 2025 poz. 775 (21 May 2025). Later amendments, such as the 240,000 zł VAT limit, are not in those texts and are cited separately.

## 1. Income-tax records: ewidencja przychodów (ryczałt) vs. KPiR (skala/liniowy). Aggregated entries, source documents, invoice obligations, 2025–2026 changes

### Takeaway
Since 1 Jan 2026 two new regulations apply: Dz.U. 2025 poz. 1294 for ryczałt and poz. 1299 for KPiR. Under both, non-invoiced (B2C) sales may still be recorded as **one amount per day** on an end-of-day **dowód wewnętrzny**. Under ryczałt that amount is split by ryczałt rate. Under KPiR a daily **zestawienie sprzedaży** is an alternative, with a single monthly KPiR line. A **monthly** aggregate with no daily record is not provided for. **Every invoice must now be entered individually**: daily or monthly invoice summaries were abolished on 1 Jan 2026, and MF confirmed on 3 Jul 2026 that this will not be reversed. Invoices are mandatory for B2B and issued to consumers only on request. Electronic books with annual JPK_PKPIR/JPK_EWP filing are being phased in for 2026 and 2027.

### Cited Findings
**Ryczałt (ryczałt od przychodów ewidencjonowanych)**
- [LAW] Art. 15 ust. 1 ustawy o zryczałtowanym podatku dochodowym: ryczałt taxpayers must keep purchase evidence, a wykaz środków trwałych/WNiP and, **separately for each tax year**, an ewidencja przychodów. Under art. 15 ust. 4 the ewidencja and its supporting evidence (dowody) are kept at the place of business or at the biuro rachunkowe that keeps the records. — [Ustawa o zryczałtowanym podatku dochodowym, t.j. Dz.U. 2025 poz. 843](https://eli.gov.pl/eli/DU/2025/843/ogl/pol/pdf)
- [LAW][NEW 2026] Rozporządzenie Ministra Finansów i Gospodarki z 6 września 2025 r. w sprawie prowadzenia ewidencji przychodów i wykazu środków trwałych oraz WNiP (Dz.U. 2025 poz. 1294) is in force from **1 Jan 2026** (§ 12). It repeals the 2021 regulation, Dz.U. 2021 poz. 2414 (§ 11). — [Dz.U. 2025 poz. 1294](https://eli.gov.pl/eli/DU/2025/1294/ogl/pol/pdf)
- [LAW] § 4 ust. 1 (poz. 1294): entries are based on invoices where sales are invoiced. For sales without invoices they are based on "wystawionego na koniec dnia dowodu wewnętrznego, w którym w jednej kwocie jest wykazana wysokość przychodów za dany dzień z podziałem na przychody objęte poszczególnymi stawkami ryczałtu". Monthly entries from monthly fiscal reports (§ 4 ust. 3) are allowed **only** for cash-register users. — [Dz.U. 2025 poz. 1294](https://eli.gov.pl/eli/DU/2025/1294/ogl/pol/pdf)
- [LAW] § 6 (poz. 1294): entries are made "najpóźniej do 20. dnia każdego miesiąca za miesiąc poprzedni". — [Dz.U. 2025 poz. 1294](https://eli.gov.pl/eli/DU/2025/1294/ogl/pol/pdf)
- [LAW] Template columns of the 2026 ewidencja (annex to poz. 1294):
  - Lp.; data wpisu; data uzyskania przychodu.
  - **numer identyfikujący fakturę wystawioną przy użyciu KSeF**.
  - numer dowodu księgowego.
  - **identyfikator podatkowy kontrahenta**, which includes the country code and is *not* filled for entries based on a dowód wewnętrzny or fiscal reports.
  - Revenue split by rate: 17 / 15 / 14 / 12.5 / 12 / 10 / 8.5 / 5.5 / 3 %; ogółem; uwagi. Taxpayers paying ryczałt quarterly may enter the date the revenue was received in "Uwagi".
  - Source: [Dz.U. 2025 poz. 1294, załącznik](https://eli.gov.pl/eli/DU/2025/1294/ogl/pol/pdf)
- [LAW] § 3 ust. 4 pkt 1 and 3 (poz. 1294): the ewidencja is still deemed rzetelna and niewadliwa if unrecorded or misrecorded revenue does not exceed 0.5% of the year's revenue, or if errors are corrected before an audit starts. — [Dz.U. 2025 poz. 1294](https://eli.gov.pl/eli/DU/2025/1294/ogl/pol/pdf)
- [LAW] Ryczałt rates matter for the split. Art. 12 ust. 1 pkt 2b sets **12%** for services "związanych z oprogramowaniem (PKWiU ex 62.01.1)", "Oryginały oprogramowania" (62.01.2), software consulting (ex 62.02), installation (ex 62.09.20.0) and network management (62.03.1). The same 12% covers publishing software packages (58.29.1, 58.29.2, ex 58.29.3 "z wyłączeniem pobierania oprogramowania w trybie on-line"). The **15%** list includes "przetwarzania danych (PKWiU ex 63.11.1)". The SaaS classification therefore drives which column revenue goes into. — [ryczałt act t.j. Dz.U. 2025 poz. 843](https://eli.gov.pl/eli/DU/2025/843/ogl/pol/pdf)

**KPiR (skala podatkowa / podatek liniowy)**
- [LAW][NEW 2026] Rozporządzenie MFiG z 6 września 2025 r. w sprawie prowadzenia podatkowej księgi przychodów i rozchodów (Dz.U. 2025 poz. 1299) is in force from **1 Jan 2026** (§ 25). It replaces Dz.U. 2019 poz. 2544. — [Dz.U. 2025 poz. 1299](https://eli.gov.pl/eli/DU/2025/1299/ogl/pol/pdf)
- [LAW] § 6 ust. 1 (poz. 1299): entries are made "w języku polskim i w walucie polskiej". — [Dz.U. 2025 poz. 1299](https://eli.gov.pl/eli/DU/2025/1299/ogl/pol/pdf)
- [LAW] § 11 ust. 1 (poz. 1299): sales revenue entries are based on (1) issued invoices, (2) fiscal reports, or (3) "wystawionego na koniec dnia dowodu wewnętrznego, w którym w jednej kwocie jest wykazana wysokość tych przychodów za dany dzień, albo prowadzonego zestawienia sprzedaży (ewidencji sprzedaży) – w przypadku przychodów, które nie są dokumentowane fakturami lub raportami fiskalnymi". — [Dz.U. 2025 poz. 1299](https://eli.gov.pl/eli/DU/2025/1299/ogl/pol/pdf)
- [LAW] § 11 ust. 2–3 (poz. 1299):
  - Invoice and fiscal-report entries: by the 20th of the following month.
  - Dowód wewnętrzny entries: "raz dziennie po zakończeniu dnia, nie później niż przed rozpoczęciem działalności w dniu następnym".
  - Zestawienie sprzedaży route: one KPiR entry "w jednej pozycji na koniec każdego miesiąca", by the 20th of the following month.
  - Source: [Dz.U. 2025 poz. 1299](https://eli.gov.pl/eli/DU/2025/1299/ogl/pol/pdf)
- [LAW] § 17–18 (poz. 1299): the zestawienie sprzedaży contains at least the Lp., the date of revenue not documented by invoice, and the amount. Entries are made **once a day** after the day ends, before business starts the next day. — [Dz.U. 2025 poz. 1299](https://eli.gov.pl/eli/DU/2025/1299/ogl/pol/pdf)
- [LAW] KPiR template notes (poz. 1299):
  - Column 4 holds the KSeF number for invoices issued in KSeF.
  - Columns 5–7 (counterparty tax ID incl. country code, or name and address) are filled for invoice- or receipt-documented sales and *not* for entries based on a dowód wewnętrzny or fiscal reports.
  - Source: [Dz.U. 2025 poz. 1299](https://eli.gov.pl/eli/DU/2025/1299/ogl/pol/pdf)

**Invoice-by-invoice entry since 2026 (both regimes)**
- [OFFICIAL][NEW 2026] MF answer to **interpelacja nr 17841**, dated 3 Jul 2026 (znak DD2.054.8.2026, signed by Podsekretarz Stanu J. Neneman):
  - Until 31 Dec 2025, entries could be made from a daily zestawienie sprzedaży/faktur (old § 5 of Dz.U. 2021 poz. 2414 and § 18 ust. 3 of Dz.U. 2019 poz. 2544).
  - The two new regulations "zrezygnowano z przepisów umożliwiających dokonywania zapisów … na podstawie dziennego zestawienia sprzedaży/faktur". Now "zapisy poszczególnych faktur są dokonywane bezpośrednio w księgach podatkowych".
  - The stated reason is JPK_PD analysis of individual documents. The rule expressly covers OSS/WSTO and dropshipping sellers issuing large volumes of invoices.
  - MF "nie planuje podjęcia działań w kierunku przywrócenia" the summary-based entries.
  - Source: [MF answer to interpelacja 17841](https://api.sejm.gov.pl/sejm/term10/interpellations/attachment/ATTDVRCQQ/i17841-o1.pdf)
- [PRACTICE] Portal confirmation: since January 2026 businesses "nie mogą ewidencjonować zbiorczych zestawień faktur" in KPiR or the ewidencja przychodów, and "Każdą fakturę należy wpisywać odrębnie". — [podatki.biz, PKPiR 2026 – MF: nie będzie zmian](https://www.podatki.biz/artykuly/pkpir-2026-zapisy-w-ksiegach-na-podstawie-zestawien-sprzedazy-faktur-mf-nie-bedzie-zmian_23_62822.htm)

**Electronic books and JPK_PKPIR / JPK_EWP**
- [LAW][NEW 2026 / 2027] Art. 24a ust. 1e PIT, in force 1 Jan 2026: KPiR (and the ewidencja środków trwałych) must be kept "przy użyciu programów komputerowych" and sent after the tax year, by the annual-return deadline, in the JPK structure (art. 193a Ordynacja podatkowa). Art. 15 ust. 12 of the ryczałt act imposes the same obligation for the ewidencja przychodów. Transitional art. 66 provisions take effect partly on 1 Jan 2026 and partly on 1 Jan 2027. — [PIT act t.j. Dz.U. 2025 poz. 163](https://eli.gov.pl/eli/DU/2025/163/ogl/pol/pdf); [ryczałt act t.j.](https://eli.gov.pl/eli/DU/2025/843/ogl/pol/pdf)
- [PRACTICE] Phasing as described by portals:
  - Monthly JPK_V7M filers: electronic books from 1 Jan 2026; first files for 2026 due by 30 Apr 2027.
  - Quarterly filers and **VAT-exempt** taxpayers: from 1 Jan 2027, with first sending for 2027 in 2028.
  - Sources partly disagree on the second group's sending date.
  - Sources: [Poradnik Przedsiębiorcy](https://poradnikprzedsiebiorcy.pl/-nowe-pliki-jpk-kogo-dotycza-i-jak-sie-przygotowac); [infor.pl](https://ksiegowosc.infor.pl/wiadomosci/7565795,elektroniczne-ksiegi-podatkowe-od-2026-r-kogo-obejmuje-obowiazek-i-kiedy-przeslac-jpk-koniec-z-papierowa-ksiegowoscia-tysiace-firm-juz-musi-prowadzic-elektroniczne-ksiegi-sprawdz-czy-ty-tez-ewidencja-podatkowa-ksiega-przychodow-i-rozchodow-kpir-pelna-ksiegowosc.html); [wFirma help JPK_EWP](https://pomoc.wfirma.pl/-jednolity-plik-kontrolny-dla-ewidencji-przychodow-jpk-ewp); [ifirma](https://www.ifirma.pl/blog/jpk_pkpir-co-to-jest-i-od-kiedy-obowiazuje/)
- [PRACTICE][NEW 2026] Former monthly VAT filers who return to the VAT exemption from 1 Jan 2026 thereby postpone the electronic PKPiR / JPK_PKPiR obligation. — [podatki.biz](https://www.podatki.biz/artykuly/zmiany-2026-nowy-limit-zwolnienia-z-vat-i-kasowego-pit_4_60826.htm)

**When an invoice must be issued (ustawa o VAT)**
- [LAW] Art. 106b ust. 1 pkt 1: an invoice is mandatory for sales to another podatnik / VAT taxpayer (incl. foreign) or a legal person. Ust. 1 pkt 4 also requires one for a prepayment, except that under ust. 1a none is needed if the payment and the service fall in the same month. — [VAT act t.j. Dz.U. 2025 poz. 775](https://eli.gov.pl/eli/DU/2025/775/ogl/pol/pdf)
- [LAW] Art. 106b ust. 3: for consumers, an invoice is issued only "na żądanie nabywcy", if the request is made "w terminie 3 miesięcy, licząc od końca miesiąca, w którym … wykonano usługę bądź otrzymano całość lub część zapłaty". — [VAT act t.j.](https://eli.gov.pl/eli/DU/2025/775/ogl/pol/pdf)
- [LAW] Art. 106b ust. 2: no invoice obligation for sales exempt under art. 43 ust. 1, **art. 113 ust. 1 i 9** or **art. 113a ust. 1**. The reference to the EU SME scheme was added in 2025. — [VAT act t.j.](https://eli.gov.pl/eli/DU/2025/775/ogl/pol/pdf)
- [LAW] Art. 106a pkt 1 lit. b: Polish invoicing rules do not apply to special-procedure (OSS/IOSS) transactions only where the member state of identification is *not* Poland. For a Polish OSS registrant they therefore apply. — [VAT act t.j.](https://eli.gov.pl/eli/DU/2025/775/ogl/pol/pdf)

### Inferences
- **Legal minimum granularity for B2C Stripe revenue without invoices is per day, not per month.** A monthly handover must therefore include (or be built from) a **day-by-day table**: Polish calendar date, PLN amount and, for ryczałt, the split by ryczałt rate. A single monthly total satisfies neither § 4 ust. 1 of poz. 1294 nor § 11 and § 17–18 of poz. 1299.
- The KPiR daily dowód wewnętrzny or zestawienie must formally be made **before business starts the next day**. A monthly email cannot satisfy that timing on its own. An automated daily job (e.g., the planned Go script run daily) that generates and archives the daily document is the natural fit. The monthly package then forwards the archived dailies plus the per-transaction detail.
  - Ryczałt is looser: the dowód is dated end-of-day, but entries can be made up to the 20th of the following month.
- **Avoid unnecessary invoices.** If an invoice is issued for every subscription charge, every one must be entered individually from 2026 (e.g., if Stripe-generated invoice PDFs are treated as VAT invoices). Not issuing B2C invoices unless requested (art. 106b ust. 3) keeps B2C revenue in the daily aggregates. B2B invoices always need individual entry with the buyer's tax ID and country code (and the KSeF number where one exists).
- **Rate mapping is a per-product decision.** For ryczałt, each product or plan must map to a ryczałt rate column (12% vs 15% vs 8.5% depends on PKWiU classification), so the export should carry a product/plan identifier.

### Gaps
- No official statement was found on whether a payment-processor report (Stripe export) may itself serve as the dowód wewnętrzny or zestawienie sprzedaży. The regulations only prescribe minimum content and timing. Confirm with the accountant.
- The correct ryczałt rate (and PKWiU) for this specific SaaS was not researched. This is an open point for the accountant, and possibly for an interpretation.
- I did not verify whether Stripe Billing's auto-generated "invoices" meet art. 106e and should be treated as VAT invoices (KSeF aspects belong to the other researcher).

## 2. Revenue-recognition date for SaaS subscriptions and one-off payments (PIT and ryczałt). Practice, refunds, chargebacks, corrections

### Takeaway
The statutory rule is the date the service is performed, but no later than the invoice date or payment date (art. 14 ust. 1c PIT). For services agreed to be settled in billing periods, it is the **last day of the billing period** (ust. 1e). Prepayments for services to be performed in later reporting periods are not revenue (ust. 3 pkt 1). Ryczałt uses the same art. 14 rules. Accountants and vendors in practice usually book Stripe revenue at the **charge/payment date**, gross of fees. No PIT interpretation specific to SaaS subscriptions was found, so the method must be agreed with the accountant. Corrections (refunds) go into the period of the correcting invoice or other correcting document, not back to the original month.

### Cited Findings
- [LAW] Art. 14 ust. 1 PIT: revenue = "kwoty należne, choćby nie zostały faktycznie otrzymane, po wyłączeniu wartości zwróconych towarów, udzielonych bonifikat i skont". For VAT payers, revenue is reduced by output VAT. — [PIT act t.j. Dz.U. 2025 poz. 163](https://eli.gov.pl/eli/DU/2025/163/ogl/pol/pdf)
- [LAW] Art. 14 ust. 1c PIT: revenue date = "dzień … wykonania usługi, albo częściowego wykonania usługi, nie później niż dzień: 1) wystawienia faktury albo 2) uregulowania należności" (subject to ust. 1e, 1h–1j, 1n–1p). — [PIT act t.j.](https://eli.gov.pl/eli/DU/2025/163/ogl/pol/pdf)
- [LAW] Art. 14 ust. 1e PIT: "Jeżeli strony ustalą, że usługa jest rozliczana w okresach rozliczeniowych, za datę powstania przychodu uznaje się ostatni dzień okresu rozliczeniowego określonego w umowie lub na wystawionej fakturze, nie rzadziej niż raz w roku." — [PIT act t.j.](https://eli.gov.pl/eli/DU/2025/163/ogl/pol/pdf)
- [LAW] Art. 14 ust. 1i PIT: where ust. 1c, 1e and 1h do not apply, revenue arises on the day payment is received. — [PIT act t.j.](https://eli.gov.pl/eli/DU/2025/163/ogl/pol/pdf)
- [LAW] Art. 14 ust. 3 pkt 1 PIT: not revenue are "pobranych wpłat lub zarachowanych należności na poczet dostaw towarów i usług, które zostaną wykonane w następnych okresach sprawozdawczych". — [PIT act t.j.](https://eli.gov.pl/eli/DU/2025/163/ogl/pol/pdf)
- [LAW] Art. 14 ust. 1j PIT: an option to treat such prepayments as revenue on receipt exists **only** for sales registered on a cash register, chosen for the whole year. A kasa-exempt online seller cannot use it. — [PIT act t.j.](https://eli.gov.pl/eli/DU/2025/163/ogl/pol/pdf)
- [LAW] Art. 14 ust. 1m–1n PIT: a correction not caused by a calculation error or obvious mistake is made by adjusting revenue "w okresie rozliczeniowym, w którym została wystawiona faktura korygująca lub, w przypadku braku faktury, inny dokument potwierdzający przyczyny korekty". If revenue in that period is insufficient, costs are increased. — [PIT act t.j.](https://eli.gov.pl/eli/DU/2025/163/ogl/pol/pdf)
- [LAW] Ryczałt follows PIT here:
  - Art. 6 ust. 1 of the ryczałt act taxes revenues "o których mowa w art. 7a ust. 4 lub art. 14 ustawy o podatku dochodowym", so the art. 14 date rules apply.
  - Art. 6 ust. 1e: if a ryczałt taxpayer has too little revenue in the period of an art. 14 ust. 1m reduction, the reduction carries to subsequent periods.
  - Source: [ryczałt act t.j.](https://eli.gov.pl/eli/DU/2025/843/ogl/pol/pdf)
- [LAW][NEW 2025] **Kasowy PIT**, art. 14c PIT (from 2025):
  - Optional. Conditions: prior-year revenue ≤ 1,000,000 zł and no full books; written election by 20 Feb (or by the 20th of the month after starting).
  - Revenue date = day of payment, but no later than 2 years after the invoice.
  - Applies **only** to revenue from transactions with a przedsiębiorca documented by invoices (art. 14c ust. 4). It is therefore irrelevant for B2C Stripe sales.
  - Requires an ewidencja of invoices: art. 24a ust. 1da PIT; art. 15 ust. 11a ryczałt act.
  - Source: [PIT act t.j.](https://eli.gov.pl/eli/DU/2025/163/ogl/pol/pdf). A podatki.biz headline mentions a "nowy limit kasowego PIT" for 2026, but the article body did not state it: [podatki.biz](https://www.podatki.biz/artykuly/zmiany-2026-nowy-limit-zwolnienia-z-vat-i-kasowego-pit_4_60826.htm)
- [LAW][VAT, for comparison] Art. 19a ust. 3: services with successive payment or settlement dates are performed at the end of each period. Art. 19a ust. 8: a prepayment creates the VAT tax point on receipt. — [VAT act t.j.](https://eli.gov.pl/eli/DU/2025/775/ogl/pol/pdf)
- [OFFICIAL] KIS, 20 Jan 2025, **0114-KDIP4-1.4012.551.2024.2.RMA** (booking-platform SaaS subscription, automatic card charges):
  - The service is an e-service (art. 2 pkt 26).
  - Monthly subscription paid in advance: VAT tax point on receipt of payment (art. 19a ust. 8), or at the end of the billing period if unpaid (art. 19a ust. 3).
  - Fixed-term subscription where access did not depend on payment: art. 19a ust. 1, i.e. on performance; the applicant's position was incorrect.
  - Source: [Inforlex – 0114-KDIP4-1.4012.551.2024.2.RMA](https://www.inforlex.pl/dok/tresc,FOB0000000000006839986,Interpretacja-indywidualna-z-dnia-20-stycznia-2025-r-Dyrektor-Krajowej-Informacji-Skarbowej-sygn-0114-KDIP4-1-4012-551-2024-2-RMA.html)
- [PRACTICE] Stripto (vendor, 2026): revenue is the full amount the customer paid, not the net payout. It arises when Stripe receives the payment (the card is charged), not when funds reach the bank, and each renewal is new revenue on its renewal date. — [stripto.pl](https://stripto.pl/jak-rozliczyc-stripe)
- [PRACTICE] Biuro rachunkowe General Office (2026):
  - "Przychodem jest kwota, którą zapłacił klient, z datą sprzedaży, a nie przelew ze Stripe".
  - "Zwrot zmniejsza przychód w okresie, w którym wystawiasz fakturę korygującą".
  - "Chargeback … nie zawsze jest zwrotem".
  - "Opłatę za spór traktujemy jak prowizję".
  - Source: [generaloffice.pl](https://generaloffice.pl/rozliczenia-stripe/)
- [PRACTICE] KsięgaI (vendor): "wypłata Stripe nie jest przychodem". — [ksiegai.pl](https://www.ksiegai.pl/stripe-dla-saas/)
- [LAW][VAT corrections] Art. 29a ust. 13 has two parallel versions in the t.j.:
  - Pre-KSeF: reduction in the period the correcting invoice is issued, provided documentation of the agreed conditions exists.
  - KSeF version: reduction in the period the structured correcting invoice is issued; for a non-structured correction, when the buyer's confirmation of receipt is obtained (ust. 13a).
  - Source: [VAT act t.j.](https://eli.gov.pl/eli/DU/2025/775/ogl/pol/pdf)

### Inferences
- For a monthly subscription renewed mid-month and paid upfront by card, three readings compete:
  - (a) **Charge date**: art. 14 ust. 1c pkt 2, "nie później niż … uregulowania należności". Common practice.
  - (b) **Last day of the billing period**: art. 14 ust. 1e, if billing periods are agreed in the terms.
  - (c) **Exclusion as a prepayment for a later period**: art. 14 ust. 3 pkt 1.
  - The **handover data must therefore carry both the payment timestamp and the service period (start/end)** of each charge, e.g. the Stripe invoice line period. The accountant can then apply the chosen method consistently and use the matching NBP rate date.
- **Annual prepaid plans** are where the methods diverge most. Reading (b) could push revenue to the last day of the annual period, possibly into the next tax year, while VAT is due on receipt under art. 19a ust. 8. Agree this explicitly with the accountant.
- **Refunds of non-invoiced B2C sales**: under art. 14 ust. 1m the "inny dokument potwierdzający przyczyny korekty" can be the refund record. The export needs refund ID, date, amount, reason and a link to the original charge, booked in the month of the refund, not back-dated.
- **Disputes and chargebacks** need lifecycle fields: opened date, amount withdrawn, dispute fee, outcome (won/lost), reinstatement date. Whether a lost dispute reduces revenue is an accountant decision.
- Stripe timestamps are UTC. Day bucketing (daily dowód, revenue date, NBP "previous business day") must use **Europe/Warsaw** local dates.

### Gaps
- No KIS **PIT** interpretation specific to SaaS/subscription revenue dates (1c vs 1e vs 3 pkt 1) was found. This is an open point for the accountant.
- No explicit provision was found on how to evidence and time **VAT** corrections for B2C sales never documented by an invoice (art. 29a ust. 13 is invoice-based). Practice must be agreed with the accountant.
- The 2026 kasowy PIT limit was not verified (the 2025 text says 1,000,000 zł). It is irrelevant for B2C.

## 3. Currency conversion and exchange-rate differences (PIT/ryczałt, VAT, OSS)

### Takeaway
- **Income tax**: foreign-currency revenue is converted at the NBP average rate from the last business day before the revenue date (art. 11a ust. 1 PIT). Exchange differences (art. 24c) apply under KPiR **and also under ryczałt**: art. 6 ust. 1c of the ryczałt act applies art. 24c, and there negative differences reduce revenue. So the assumption that they are "not relevant for ryczałt" is wrong.
- **VAT**: conversion uses the NBP (or optionally ECB) rate from the day before the tax point (art. 31a).
- **OSS**: returns are in EUR at the ECB rate on the last day of the quarter (art. 130c ust. 6–7).

### Cited Findings
- [LAW] Art. 11a ust. 1 PIT: "Przychody w walutach obcych przelicza się na złote według kursu średniego walut obcych ogłaszanego przez Narodowy Bank Polski z ostatniego dnia roboczego poprzedzającego dzień uzyskania przychodu." Ust. 2 applies the same for costs, keyed to the day the cost is incurred. — [PIT act t.j.](https://eli.gov.pl/eli/DU/2025/163/ogl/pol/pdf)
- [LAW] Art. 24c ust. 1–2 PIT: exchange differences increase revenue (positive) or costs (negative). A positive difference arises when the revenue due, converted at the NBP rate, is lower than its value on the day of receipt "przeliczonej według faktycznie zastosowanego kursu waluty z tego dnia" (pkt 1). Pkt 3 covers foreign-currency funds whose value on inflow is lower than on outflow. — [PIT act t.j.](https://eli.gov.pl/eli/DU/2025/163/ogl/pol/pdf)
- [LAW] **Ryczałt**, art. 6 ust. 1c: "Do przychodów, o których mowa w ust. 1, z tytułu różnic kursowych stosuje się odpowiednio art. 24c ustawy o podatku dochodowym, z tym że ujemne różnice kursowe … zmniejszają przychody uzyskane w roku podatkowym, w którym powstały te różnice." — [ryczałt act t.j.](https://eli.gov.pl/eli/DU/2025/843/ogl/pol/pdf)
- [LAW] The ryczałt act has no separate conversion rule for foreign-currency revenue; the only "kurs" provision concerns EUR thresholds. I searched the t.j. text. — [ryczałt act t.j.](https://eli.gov.pl/eli/DU/2025/843/ogl/pol/pdf)
- [PRACTICE] Under ryczałt only sales-related exchange differences are shown, and both positive and negative adjust revenue. For a PLN bank account the bank's actually applied rate is used; for a foreign-currency account, the NBP rate of the day before inflow. Sources disagree on whether the difference arises on the day funds reach the bank. — [Poradnik Przedsiębiorcy, ryczałt a różnice kursowe](https://poradnikprzedsiebiorcy.pl/-ryczalt-a-roznice-kursowe)
- [LAW] KPiR entries are made in PLN (§ 6 ust. 1 of Dz.U. 2025 poz. 1299). — [Dz.U. 2025 poz. 1299](https://eli.gov.pl/eli/DU/2025/1299/ogl/pol/pdf)
- [LAW] VAT, art. 31a ust. 1: NBP average rate from the last business day before the tax point. Alternatively, the latest ECB rate published for the day before the tax point; non-EUR currencies are then converted via their EUR rate.
  - [NEW 2026] Ust. 1a (KSeF era): where the tax point is the invoice date, the NBP rate from the day before the supply date (106e ust. 1 pkt 1) may be used, if the structured invoice is issued by the next day.
  - Source: [VAT act t.j.](https://eli.gov.pl/eli/DU/2025/775/ogl/pol/pdf)
- [LAW] OSS, art. 130c ust. 6–7: "Kwoty w deklaracji VAT wyrażane są w euro". Payments in other currencies are converted at the ECB rate "opublikowany … w ostatnim dniu danego okresu rozliczeniowego" (or the next day's rate if none was published). — [VAT act t.j.](https://eli.gov.pl/eli/DU/2025/775/ogl/pol/pdf)
- [PRACTICE] Stripto: use the NBP rate from the last business day before the payment day, not Stripe's rate or the payout-day rate. — [stripto.pl](https://stripto.pl/jak-rozliczyc-stripe)
- [PRACTICE] General Office: "Jeśli trzymasz saldo w euro albo dolarach i wypłacasz je później, powstają różnice kursowe." — [generaloffice.pl](https://generaloffice.pl/rozliczenia-stripe/)
- [Stripe] PayPal-on-Stripe presentment currencies include EUR, GBP, USD, CHF, CZK, DKK, NOK, **PLN**, SEK and others. Funds settle per the chosen settlement preference (Stripe balance or PayPal balance). — [Stripe docs – PayPal](https://docs.stripe.com/payments/paypal)

### Inferences
- **Settlement currency determines where FX differences arise.**
  - If the Stripe account settles foreign charges into a **PLN** balance, Stripe converts at charge time at its own rate. That rate is arguably the "faktycznie zastosowany kurs", so each foreign-currency charge produces a small per-transaction exchange difference against NBP(D-1). This applies under KPiR and ryczałt alike.
  - If EUR is held and paid out to an EUR account and converted later, differences arise at payout/conversion (art. 24c ust. 2 pkt 3 logic).
- **Required export fields:** presentment currency and amount; settlement currency and amount; Stripe exchange rate; the NBP table number, date and rate used; the PLN value; and the computed difference (or enough data for the accountant to compute it).
- The NBP rate date depends on the revenue date chosen in section 2: charge date vs. billing-period end. The script should not hard-code one rule without the accountant's sign-off.
- **Separate FX regimes:** OSS figures need their own EUR conversion (ECB rate at quarter-end). The data for VAT (OSS), VAT (domestic JPK) and income tax may legitimately carry three different PLN/EUR figures for the same charge.

### Gaps
- No official source was found on whether the "dzień otrzymania" for art. 24c is the day the charge is credited to the Stripe balance or the day the payout reaches the bank. Portal sources disagree.
- I did not find a KIS interpretation confirming that ryczałt revenue conversion follows art. 11a PIT. The ryczałt act is silent and practice applies art. 11a; confirm with the accountant.
- The currency in which Stripe's monthly fee invoice is denominated for a PLN-settling Polish account was not verified.

## 4. VAT on electronically supplied services. B2C PL / B2C EU (EUR 10,000, OSS) / B2B EU (reverse charge, VAT-UE) / non-EU (NP); customer-location evidence; JPK_V7 markings; subject exemption (240,000 zł from 2026)

### Takeaway
- **B2C**: SaaS is an e-service (art. 2 pkt 26 ↔ art. 7 of Reg. 282/2011).
  - Polish consumers: Polish VAT (23%).
  - Other EU consumers: Polish VAT until EU-wide cross-border B2C e-services exceed EUR 10,000 (42,000 zł) in the current or previous year; after that, VAT of the consumer's state, normally via OSS (quarterly EUR return, 10-year records).
- **B2B EU**: reverse charge (art. 28b), invoice marked "odwrotne obciążenie", monthly VAT-UE informacja podsumowująca, and VAT-UE registration even for VAT-exempt sellers.
- **Non-EU customers**: outside Polish VAT (NP).
- **Evidence**: two non-contradictory items of customer-location evidence, or one item under EUR 100,000.
- **JPK_V7**: non-invoiced B2C sales outside a cash register are shown in summary as "WEW" (practice). OSS sales appear to be excluded from JPK_V7 (vendor practice; no official text found).
- **Subject exemption**: 240,000 zł from 1 Jan 2026. Art. 113 ust. 13 does not exclude SaaS (only legal, advisory, jewellery and debt-collection services, plus certain goods).

### Cited Findings
- [LAW] Art. 2 pkt 26 VAT: "usługach elektronicznych – … usługi świadczone drogą elektroniczną, o których mowa w art. 7 rozporządzenia 282/2011". — [VAT act t.j.](https://eli.gov.pl/eli/DU/2025/775/ogl/pol/pdf)
- [OFFICIAL] KIS treated these as e-services:
  - Hosting, cloud, VPN, SSL and website builders sold online and automated (26 Mar 2025, 0113-KDIPT1-3.4012.17.2025.2.ALN): [Inforlex](https://www.inforlex.pl/dok/tresc,FOB0000000000006902420,Interpretacja-indywidualna-z-dnia-26-marca-2025-r-Dyrektor-Krajowej-Informacji-Skarbowej-sygn-0113-KDIPT1-3-4012-17-2025-2-ALN.html)
  - A SaaS subscription platform (20 Jan 2025, 0114-KDIP4-1.4012.551.2024.2.RMA): [Inforlex](https://www.inforlex.pl/dok/tresc,FOB0000000000006839986,Interpretacja-indywidualna-z-dnia-20-stycznia-2025-r-Dyrektor-Krajowej-Informacji-Skarbowej-sygn-0114-KDIP4-1-4012-551-2024-2-RMA.html)
- [LAW] Art. 28k VAT:
  - Ust. 1: for B2C e-services, place of supply = the customer's seat, permanent address or usual residence.
  - Ust. 2: the exception applies if the supplier is established in one member state only and total B2C TBE services plus WSTO to other EU states "nie przekroczyła w trakcie roku podatkowego ani w poprzednim roku podatkowym kwoty 10 000 euro". For a Poland-only supplier, "wyrażona w złotych równowartość kwoty w euro wynosi 42 000 zł".
  - Ust. 3: once exceeded, ust. 1 applies "począwszy od usługi, w związku ze świadczeniem której przekroczono tę kwotę".
  - Ust. 4: the supplier may opt for customer-location taxation.
  - Source: [VAT act t.j.](https://eli.gov.pl/eli/DU/2025/775/ogl/pol/pdf)
- [LAW] B2B with EU businesses:
  - Art. 28b ust. 1: B2B place of supply = the recipient's seat.
  - Art. 106e ust. 1 pkt 18: the invoice must contain "odwrotne obciążenie".
  - Art. 100 ust. 1 pkt 4 and ust. 3: VAT-UE informacja podsumowująca for art. 28b services to EU-identified taxable persons where the recipient pays the tax, filed monthly "w terminie do 25. dnia miesiąca następującego".
  - Art. 97 ust. 3 pkt 2: podatnicy other than active VAT payers (i.e. incl. VAT-exempt) must register as VAT-UE before providing such services.
  - Source: [VAT act t.j.](https://eli.gov.pl/eli/DU/2025/775/ogl/pol/pdf)
- [LAW] OSS (procedura unijna), art. 130c–130d VAT:
  - Electronic **quarterly** returns, due by the end of the month after the quarter, filed "niezależnie od tego, czy … miały miejsce" (i.e. nil returns too).
  - Content per member state of consumption: net value of services, VAT by rate, and total VAT.
  - Amounts and payment in EUR (ust. 6–8).
  - Under art. 130d, electronic records of OSS transactions per **art. 63c of Reg. 282/2011**, kept "przez okres 10 lat".
  - Source: [VAT act t.j.](https://eli.gov.pl/eli/DU/2025/775/ogl/pol/pdf)
- [LAW] **Evidence of customer location**, Reg. 282/2011 art. 24b:
  - Point (d): the customer is presumed to be located "at the place identified as such by the supplier on the basis of two items of non-contradictory evidence" listed in art. 24f.
  - Second paragraph: where cross-border TBE supplies by a supplier established in one member state "does not exceed EUR 100 000" in the current and preceding calendar year, **one** item of evidence suffices. It must come from a person involved in the supply other than the supplier or customer (art. 24f points (a)–(e)).
  - Exceeding the threshold ends the simplification from that point.
  - Introduced by Implementing Reg. (EU) 2017/2459.
  - Sources: [Reg. 282/2011 art. 24b (legislation.gov.uk copy)](https://www.legislation.gov.uk/eur/2011/282/article/24b); [EUR-Lex – Reg. 2017/2459](https://eur-lex.europa.eu/eli/reg_impl/2017/2459/oj/eng)
- [LAW] Art. 2 pkt 22 VAT: "sprzedaży" = "odpłatną dostawę towarów i odpłatne świadczenie usług na terytorium kraju, eksport towarów oraz wewnątrzwspólnotową dostawę towarów". — [VAT act t.j.](https://eli.gov.pl/eli/DU/2025/775/ogl/pol/pdf)
- [LAW][NEW 2026] **Zwolnienie podmiotowe, art. 113 ust. 1**:
  - The May 2025 t.j. still states 200,000 zł; the limit was raised to **240,000 zł from 1 Jan 2026** (net of VAT).
  - Per MF's position on transitional rules: taxpayers whose 2025 sales exceeded 200,000 zł but not 240,000 zł may return to the exemption in any month of 2026 without the one-year waiting period, by filing an updated VAT-R.
  - One portal (amavat.pl) still states 200,000 zł for 2026, which conflicts with the others and appears outdated.
  - Sources: [VAT act t.j. (old 200k text)](https://eli.gov.pl/eli/DU/2025/775/ogl/pol/pdf); [podatki.biz – MF o przepisach przejściowych](https://www.podatki.biz/artykuly/nowy-limit-zwolnienia-podmiotowego-z-vat-w-2026-r-mf-o-przepisach-przejsciowych_4_60302.htm); [podatki.biz – Zmiany 2026](https://www.podatki.biz/artykuly/zmiany-2026-nowy-limit-zwolnienia-z-vat-i-kasowego-pit_4_60826.htm); contradicted by [amavat.pl](https://amavat.pl/limit-zwolnienia-z-vat-w-2026-co-sie-zmienia/)
- [LAW] **Art. 113 ust. 13** (exclusions from the exemption):
  - Pkt 1: goods in załącznik nr 12.
  - Pkt 2: services "prawnicze", "w zakresie doradztwa" (except agricultural advisory), "jubilerskie", "ściągania długów, w tym factoringu".
  - Pkt 3: repealed.
  - Nothing targets e-services or SaaS.
  - Source: [VAT act t.j.](https://eli.gov.pl/eli/DU/2025/775/ogl/pol/pdf)
- [LAW][NEW 2025] Art. 113a (EU SME scheme): exempts sales in Poland by SMEs established in other member states if their EU-wide annual turnover is ≤ EUR 100,000 and the art. 113 ust. 1 limit is met. The 2025 scheme is also referenced in art. 106b ust. 2 and art. 109 ust. 1. — [VAT act t.j.](https://eli.gov.pl/eli/DU/2025/775/ogl/pol/pdf)
- [LAW] Records:
  - Art. 109 ust. 1: VAT-exempt (art. 113 / 113a) taxpayers must keep "ewidencję sprzedaży za dany dzień, nie później jednak niż przed dokonaniem sprzedaży w dniu następnym". This is what monitors the 240k limit.
  - Art. 109 ust. 3: active payers keep a record with data allowing correct settlement and the VAT-UE informacja, incl. sale type, tax base and VAT by rate, and corrections.
  - Source: [VAT act t.j.](https://eli.gov.pl/eli/DU/2025/775/ogl/pol/pdf)
- [PRACTICE] JPK_V7, non-invoiced B2C sales exempt from the cash register: recorded on the basis of an internal document with the **"WEW"** document type, generally as one summary entry. RO, WEW and FP markings cannot be combined on one entry. — [Poradnik Przedsiębiorcy – zwolnienie z kasy a JPK](https://poradnikprzedsiebiorcy.pl/-ewidencjonowanie-sprzedazy-przy-zwolnieniu-z-kasy-fiskalnej); [jpk.info.pl – sprzedaż poza kasą (one vs many WEW)](https://jpk.info.pl/aktualnosci/sprzedaz-poza-kasa-fiskalna/)
- [PRACTICE] JPK_V7 procedure markings: **"WSTO_EE"** (from Jan 2022; "EE" was used Jul–Dec 2021, and EE covered only art. 28k services before July 2021) applies to WSTO and to art. 28k telecom/broadcast/e-services for non-taxable customers in other EU states. — [jpk.info.pl – nowe oznaczenia procedur](https://jpk.info.pl/aktualnosci/nowe-oznaczenia-procedur-podatkowych-jpk-v7/)
- [PRACTICE, vendor doc] Comarch Optima: documents marked "Procedura OSS" are excluded from the JPK_V7 file. — [Comarch Optima help](https://pomoc.comarch.pl/optima/pl/2026_5/?p=30357)
- [PRACTICE, non-official] Field K_11 (supplies or services outside Poland) holds "eksport usług"-type documents rather than OSS transactions. — [fakturowo.pl procedures PDF](https://www.fakturowo.pl/pomoc/procedury-jpk-v7.pdf)

### Inferences
- **Per-transaction VAT category needed.** Every Stripe charge should carry one of:
  1. PL B2C taxed in PL;
  2. EU B2C taxed in PL (under the EUR 10k / 42,000 zł threshold; "WSTO_EE" marking for an active payer);
  3. EU B2C under OSS (country and local rate);
  4. EU B2B reverse charge (VIES-validated VAT ID; VAT-UE);
  5. non-EU NP;
  6. PL sales exempt under art. 113.
- **Running thresholds the script should track:**
  - 240,000 zł (art. 113, Polish "sprzedaż" only).
  - 42,000 zł / EUR 10,000 (art. 28k): EU cross-border B2C. Once crossed, the very next supply changes its place of supply.
  - EUR 100,000 (art. 24b, evidence simplification).
  - 20,000 zł for the cash register (section 5).
- **Location evidence fields per B2C charge:** billing-address country, card issuing country (BIN), IP-geolocated country (if captured), and possibly PayPal payer country. Below EUR 100k a single third-party item, e.g. the card/bank country from the payment provider, is enough. Above it, two non-contradictory items, with OSS records kept 10 years.
- **Services supplied outside Poland** (OSS B2C, EU B2B, non-EU) fall outside "sprzedaż" (art. 2 pkt 22). This matters for the art. 113 limit and the cash-register rules, but all such revenue remains **taxable income** for PIT/ryczałt.
- **A VAT-exempt seller still has VAT-UE duties** if it sells B2B to EU businesses: VAT-UE registration plus monthly VAT-UE informacja (art. 97 ust. 3 pkt 2, art. 100 ust. 1 pkt 4).

### Gaps
- No **official** MF/KAS text was found stating that OSS-settled sales are excluded from JPK_V7; only vendor documentation. Confirm with the accountant.
- I did not verify K_12 ("w tym usługi z art. 100 ust. 1 pkt 4") in the current JPK_V7 schema, nor the exact paragraph of the JPK regulation defining WEW and WSTO_EE.
- Whether a VAT-exempt JDG can register for OSS (rather than becoming VAT-active) was not confirmed.
- Whether Polish SMEs can use the 2025 cross-border SME scheme in other member states as an alternative to OSS was not researched.
- Whether PIT/ryczałt revenue from OSS sales is net of the **foreign** VAT: the PIT act's art. 14 ust. 1 speaks of "należny podatek od towarów i usług", and I found no explicit OSS rule. This is an open point.

## 5. Cash register (kasa rejestrująca) exemption for SaaS paid via Stripe / PayPal-on-Stripe. Applicable positions, conditions, interpretations, implications for the export

### Takeaway
The current regulation is Dz.U. 2024 poz. 1902 (from 1 Jan 2025, amended 27 Mar 2026 in unrelated areas). Its exemptions last no longer than **31 Dec 2027**. A pure SaaS cannot use poz. 27 (e-services exempt only if the taxpayer also provides telecom services). It can rely on:
- the **20,000 zł** B2C sales limit (§ 3 ust. 1 pkt 1, or pro rata for new businesses); and/or
- **poz. 42**: services to consumers, fully paid "za pośrednictwem poczty, banku lub SKOK", where "z ewidencji i dowodów dokumentujących zapłatę jednoznacznie wynika, jakiej konkretnie czynności dotyczyła".

The regulation does not mention payment institutions. KIS has accepted payment intermediaries (PayU, PayPal, Przelewy24; and in 2025 unnamed "pośrednicy płatności" for online e-services) when the money ends up in the taxpayer's bank account and each payment is traceable. No interpretation naming Stripe was found.

### Cited Findings
- [LAW] Art. 111 ust. 1 VAT: "Podatnicy dokonujący sprzedaży na rzecz osób fizycznych nieprowadzących działalności gospodarczej oraz rolników ryczałtowych są obowiązani prowadzić ewidencję sprzedaży przy zastosowaniu kas rejestrujących." — [VAT act t.j.](https://eli.gov.pl/eli/DU/2025/775/ogl/pol/pdf)
- [LAW] **Rozporządzenie MF z 17 grudnia 2024 r. w sprawie zwolnień z obowiązku prowadzenia ewidencji sprzedaży przy zastosowaniu kas rejestrujących** (Dz.U. 2024 poz. 1902), in force 1 Jan 2025.
  - Amended by Rozporządzenie MFiG z 27 marca 2026 r. (Dz.U. 2026 poz. 420), in force 31 Mar 2026. The amendment concerns parking services and non-powered self-service devices and is not relevant to SaaS.
  - Sources: [ISAP – Dz.U. 2024 poz. 1902](https://isap.sejm.gov.pl/isap.nsf/DocDetails.xsp?id=WDU20240001902); [ISAP – Dz.U. 2026 poz. 420](https://isap.sejm.gov.pl/isap.nsf/DocDetails.xsp?id=WDU20260000420); [Forsal – projekt zmian 2026](https://forsal.pl/gospodarka/aktualnosci/artykuly/10640481,nowe-zwolnienia-z-kas-fiskalnych-2026-projekt-rozporzadzenia.html)
- [LAW] Conditions in Dz.U. 2024 poz. 1902:
  - § 2 ust. 1: annex exemptions apply per tax year, "nie dłużej jednak niż do dnia 31 grudnia 2027 r.".
  - § 3 ust. 1 pkt 1–2: prior-year sales to consumers and ryczałt farmers "nie przekroczyła kwoty 20 000 zł", pro rata for businesses started after 31 Dec 2024.
  - § 3 ust. 1 pkt 3–4: the share of annex-exempt sales in consumer sales exceeds 80%.
  - § 3 ust. 3: poz. 40 and 50 sales and VAT are ignored in the calculation.
  - § 5: the exemption lapses two months after the trigger.
  - Source: [Dz.U. 2024 poz. 1902 text](https://api.sejm.gov.pl/eli/acts/DU/2024/1902/text.html)
- [LAW] **Annex poz. 42**: "Świadczenie usług na rzecz osób fizycznych nieprowadzących działalności gospodarczej oraz rolników ryczałtowych, jeżeli świadczący usługę otrzyma w całości zapłatę za wykonaną czynność za pośrednictwem poczty, banku lub spółdzielczej kasy oszczędnościowo-kredytowej (odpowiednio na rachunek bankowy podatnika lub na rachunek podatnika w spółdzielczej kasie oszczędnościowo-kredytowej, której jest członkiem), a z ewidencji i dowodów dokumentujących zapłatę jednoznacznie wynika, jakiej konkretnie czynności dotyczyła". — [Dz.U. 2024 poz. 1902 text](https://api.sejm.gov.pl/eli/acts/DU/2024/1902/text.html)
- [LAW] **Annex poz. 27**: "Usługi elektroniczne, o których mowa w art. 2 pkt 26 ustawy, jeżeli podatnik świadczy usługi telekomunikacyjne". This is not available to a SaaS that provides no telecom services. — [Dz.U. 2024 poz. 1902 text](https://api.sejm.gov.pl/eli/acts/DU/2024/1902/text.html)
- [LAW] The regulation defines neither "bank" nor "rachunek" and does not mention payment or e-money institutions. The § 4 ust. 1 pkt 2 exclusion list (e.g. passenger transport, vehicle repair, medical, legal, tax advisory, hairdressing, parking, device-operated services) does not name software or SaaS. — [Dz.U. 2024 poz. 1902 text](https://api.sejm.gov.pl/eli/acts/DU/2024/1902/text.html)
- [OFFICIAL] KIS, 26 Mar 2025, **0113-KDIPT1-3.4012.17.2025.2.ALN**:
  - Facts: fully automated online e-services for consumers, prepaid only electronically "za pośrednictwem pośredników płatności", full payment landing in the bank account (online payment site or per-customer virtual accounts), records identifying the customer and the transaction.
  - Held: exemption under § 3 ust. 1 pkt 1 until the 20,000 zł limit is exceeded, **and** under § 2 ust. 1 with **poz. 42** until 31 Dec 2027.
  - "Niespełnienie jednego z warunków oznacza brak możliwości zastosowania zwolnienia."
  - KIS did not separately analyse whether a payment operator is a "bank".
  - Source: [Inforlex](https://www.inforlex.pl/dok/tresc,FOB0000000000006902420,Interpretacja-indywidualna-z-dnia-26-marca-2025-r-Dyrektor-Krajowej-Informacji-Skarbowej-sygn-0113-KDIPT1-3-4012-17-2025-2-ALN.html)
- [OFFICIAL] KIS, 26 May 2025, **0111-KDIB3-2.4012.205.2025.2.AK**: a tutor above 20,000 zł may use the exemption because all payments went to the bank account and the records showed which transaction each concerned. — [Inforlex](https://www.inforlex.pl/dok/tresc,FOB0000000000006964208,Interpretacja-indywidualna-z-dnia-26-maja-2025-r-Dyrektor-Krajowej-Informacji-Skarbowej-sygn-0111-KDIB3-2-4012-205-2025-2-AK.html)
- [OFFICIAL, older; position numbering of earlier regulations]
  - Dyrektor Izby Skarbowej w Warszawie, 26 Feb 2015: payment via **PayU** satisfies "za pośrednictwem banku" because funds pass from the buyer's account through the intermediary to the seller's account.
  - Interpretation of 19 Apr 2016 took a similar view for **PayPal / Przelewy24**.
  - The tax authority distinguished PayU/PayPal from credit-card payment.
  - Sources: [eGospodarka – "tak dla PayU, nie dla karty kredytowej"](https://www.podatki.egospodarka.pl/132471,Zwolnienie-z-kasy-tak-dla-platnosci-PayU-nie-dla-karty-kredytowej,1,68,1.html); [Poradnik Przedsiębiorcy – PayU](https://poradnikprzedsiebiorcy.pl/-czy-platnosc-przez-payu-zwalnia-z-kasy-fiskalnej); [PayPal a kasa fiskalna](https://polskiekasyfiskalneonline.pl/paypal-a-kasa-fiskalna-czy-mozliwe-jest-zwolnienie/)
- [OFFICIAL, summary only] KIS, 10 Feb 2026, 0111-KDIB3-2.4012.924.2025.2.KK: card payments via a terminal; reported to cite a court ruling that a cash register is not required if the conditions are met. I saw only a search abstract. — [Inforlex](https://www.inforlex.pl/dok/tresc,FOB0000000000007520205,Interpretacja-indywidualna-z-dnia-10-lutego-2026-r-Dyrektor-Krajowej-Informacji-Skarbowej-sygn-0111-KDIB3-2-4012-924-2025-2-KK.html)
- [PRACTICE] General Office: services paid fully by transfer "mogą być zwolnione z kasy do końca 2027 r.", but "Stripe nie jest bankiem i wypłaca kwoty zbiorczo, dlatego urzędy oceniają to różnie". — [generaloffice.pl](https://generaloffice.pl/rozliczenia-stripe/)
- [PRACTICE] Stripto: B2C online sales paid to a bank account with an identifiable buyer are "zazwyczaj" exempt; confirm with the accountant. — [stripto.pl](https://stripto.pl/jak-rozliczyc-stripe)
- [Stripe] PayPal-on-Stripe settles either to the Stripe balance or to the PayPal balance, depending on the chosen **settlement preference**. Refunds are funded from the corresponding balance. — [Stripe docs – PayPal](https://docs.stripe.com/payments/paypal)

### Inferences
- **The poz. 42 audit trail.** To rely on poz. 42 with Stripe, the taxpayer must show for each B2C payment:
  - (a) that it was paid **in full** electronically;
  - (b) that the money reached the **taxpayer's bank account**, via the Stripe payout;
  - (c) **which specific service** it paid for.
- Stripe pays out net, aggregated amounts, so the chain charge → balance transaction → payout → bank-statement line has to be reconstructable. **Required retained fields per transaction:**
  - Stripe charge/payment ID and balance-transaction ID;
  - date/time;
  - gross amount and currency;
  - customer identifier (customer ID, email/name);
  - the product/plan and service period, or the invoice/subscription ID;
  - payment method (card / PayPal);
  - fee and net amount;
  - **payout ID with payout date and amount**, matched to the bank statement;
  - refunds and disputes linked to the original charge.
- The itemized payout reconciliation (charges per payout) plus bank statements are the "dowody dokumentujące zapłatę". A monthly aggregate alone would not show "jakiej konkretnie czynności dotyczyła".
- For PayPal-on-Stripe, choosing **"settle to Stripe"** keeps one evidentiary chain (Stripe payout → bank). Settling to the PayPal balance adds a second intermediary hop that must also be documented.
- **What counts toward the 20,000 zł limit:** B2C sales taxed in Poland, including EU consumers below the EUR 10k threshold. B2C supplies whose place of supply is abroad (OSS, non-EU) are not "sprzedaż" under art. 2 pkt 22. The limit should be tracked as a running total even if poz. 42 is relied on, as a fallback.
- **VAT-exempt sellers are not excused from the cash register.** Art. 111 ust. 1 addresses "podatnicy" without excluding art. 113-exempt ones, so a VAT-exempt JDG still needs a kasa exemption. This is my reading of the text; confirm with the accountant.
- **Scheduled cliff:** the annex exemptions expire **31 Dec 2027**. Watch for a successor regulation in 2027.

### Gaps
- No KIS interpretation **naming Stripe** (or PayPal-on-Stripe) was found. Whether card payments acquired by Stripe count as "za pośrednictwem banku" remains uncertain: older practice distinguished intermediaries from card payments, while newer abstracts suggest card payments may qualify. The user may want an own interpretation or the accountant's documented position.
- The full texts of the 2016 PayPal/Przelewy24 and the 10 Feb 2026 interpretations were not read.

## 6. Stripe's own fees (and PayPal fees): VAT (import of services), VAT-UE registration / VAT-9M for VAT-exempt JDGs, deductibility

### Takeaway
Stripe supplies payment processing **and** other services (Billing, Radar, Tax) to EU account holders from Ireland, **without VAT**, with tax invoices issued monthly (by the 10th). The Polish JDG must self-account (import usług, art. 28b with art. 17 ust. 1 pkt 4).
- **Pure payment processing** is arguably exempt under art. 43 ust. 1 pkt 40. Non-payment products (Billing, Tax, Radar) are not payment services and would be 23%. Polish practice is split, and no Stripe-specific KIS ruling was found.
- **Active VAT payers**: neutral via JPK_V7M (output equals input).
- **VAT-exempt JDG**:
  - VAT-UE registration before the first purchase (art. 97 ust. 3 pkt 1);
  - VAT-9M by the 25th of the following month (art. 99 ust. 9), paying 23% on taxable fees;
  - under current practice, also reporting the base of exempt imports;
  - a pending 2026/2027 VAT package would end reporting of **exempt** imports, but it was not enacted as of late Sep 2026.
- **PayPal fees** in PayPal-on-Stripe are **not** on Stripe's tax invoice; the documents come from the PayPal dashboard.
- **Deductibility**: fees are deductible costs under KPiR and irrelevant under ryczałt (revenue stays gross).

### Cited Findings
- [Stripe] "Stripe supplies payment processing and other services (such as Billing, Radar, and Tax) from our establishment in Ireland" and "does not charge VAT on Stripe fees" except for account holders in Ireland, Norway, Switzerland/Liechtenstein and the UAE. Users who may need to self-account should consult a tax advisor. Tax invoices are in Dashboard → Documents, "issued by the 10th of the following month". — [Stripe Support – Global taxation of Stripe fees](https://support.stripe.com/questions/global-taxation-of-stripe-fees)
- [Stripe] PayPal on Stripe: "PayPal fees are included in the Less fees sections of Balance and Balance Activity reports. PayPal fees aren't included in your tax invoice from Stripe. You can access these documents from your PayPal dashboard." — [Stripe docs – PayPal](https://docs.stripe.com/payments/paypal)
- [LAW] VAT on imported services:
  - Art. 17 ust. 1 pkt 4: buyers of services from suppliers without a Polish establishment are the taxpayers (for art. 28b services, a podatnik under art. 15).
  - Art. 97 ust. 3 pkt 1: podatnicy other than active VAT payers (incl. exempt) must register as VAT-UE before purchasing art. 28b services "jeżeli usługi te stanowiłyby u nich import usług".
  - Art. 99 ust. 9: where art. 17 ust. 1 pkt 4 taxpayers have no other return obligation, a return is filed "do 25. dnia miesiąca następującego po miesiącu, w którym powstał obowiązek podatkowy". This is the basis for VAT-9M.
  - Source: [VAT act t.j.](https://eli.gov.pl/eli/DU/2025/775/ogl/pol/pdf)
- [LAW] Art. 43 ust. 1 pkt 40: exempt are "usługi w zakresie depozytów środków pieniężnych, prowadzenia rachunków pieniężnych, wszelkiego rodzaju transakcji płatniczych, przekazów i transferów pieniężnych, długów, czeków i weksli oraz usługi pośrednictwa w świadczeniu tych usług". — [VAT act t.j.](https://eli.gov.pl/eli/DU/2025/775/ogl/pol/pdf)
- [LAW] VAT-UE informacja podsumowująca covers services **supplied** under art. 28b (art. 100 ust. 1 pkt 4), not services purchased. A vendor claim that buying Stripe services triggers a monthly VAT-UE return is therefore inconsistent with art. 100 ust. 1. — [VAT act t.j.](https://eli.gov.pl/eli/DU/2025/775/ogl/pol/pdf); contradicting [stripto.pl](https://stripto.pl/jak-rozliczyc-stripe)
- [PRACTICE, conflicting]
  - ifirma (13 Oct 2021): payment-intermediation services are VAT-exempt and entered with "zw". VAT-exempt taxpayers must still register VAT-UE; VAT from VAT-9M is entered in KPiR when the declaration is filed (cost only once paid). The article names the issuer "Stripe Payments International", which is outdated vs. Stripe's statement that EU services come from Ireland. — [ifirma.pl](https://www.ifirma.pl/blog/aktualnosci/rozliczenie-faktury-od-stripe/)
  - General Office (2026): "Opłaty Stripe zwykle traktuje się jako opodatkowane", so the VAT-exempt JDG pays 23% via VAT-9M; "Inaczej bywa przy PayPal, którego opłaty są zwykle zwolnione z VAT". VAT-R/VAT-UE registration "przed pierwszym zakupem usługi". — [generaloffice.pl](https://generaloffice.pl/rozliczenia-stripe/)
  - Stripto: reverse-charge 23% on the fee, with a monthly consolidated invoice from Stripe Payments Europe Ltd. — [stripto.pl](https://stripto.pl/jak-rozliczyc-stripe)
  - KsięgaI demo: VAT-RC 23% on fees; VAT-9M vs JPK_V7 depending on status. — [ksiegai.pl](https://www.ksiegai.pl/stripe-dla-saas/)
- [OFFICIAL via secondary] NSA **I FSK 1928/23**: a Polish buyer obliged to settle an import of exempt financial services cannot unilaterally opt to tax it. — [MDDP](https://www.mddp.pl/import-uslug-finansowych-nabywca-nie-moze-sam-wybrac-opodatkowania-vat/)
- [OFFICIAL via secondary] Current law on **exempt** imports. Article by Akademia LTCA, 22 Jun 2026, citing:
  - DKIS 05.09.2025 **0113-KDIPT1-2.4012.569.2025.2.SM**: base reported in VAT-9M; VAT-UE registration required.
  - DKIS 29.04.2026 **0114-KDIP1-2.125.2026.3.JO**: VAT-9M shows only the base, no output tax, for exempt imports.
  - Source: [Akademia LTCA](https://akademialtca.pl/blog/import-uslug-zwolnionych-czy-wykazywac-w-jpk-vat-9m)
- [2027 / PENDING] Proposed new **art. 17 ust. 1j** (government bill UD314, druk 2677): importers of exempt services would no longer be taxpayers for them, so no VAT-UE registration and no JPK/VAT-9M reporting. Status by source:
  - Akademia LTCA (June 2026): from 1 Oct 2026.
  - ifirma (29 Sep 2026): the package "jest jeszcze na etapie prac legislacyjnych"; no longer reporting the base of exempt imports in "JPK_VAT, VAT-9M … VAT-8"; most changes from 1 Jan 2027.
  - Crowe: described as planned, covering both art. 43 ust. 1 and SME exemptions.
  - Sources: [Akademia LTCA](https://akademialtca.pl/blog/import-uslug-zwolnionych-czy-wykazywac-w-jpk-vat-9m); [ifirma – zmiany VAT X 2026 / I 2027](https://www.ifirma.pl/blog/zmiany-w-podatku-vat-od-pazdziernika-2026-r-z-uwzglednieniem-zmian-od-stycznia-2027-r/); [Crowe](https://www.crowe.com/pl/insights/zmiany-w-vat-2026-czyli-planowana-nowelizacja-przepisow)
- [PRACTICE] Revenue is the gross customer payment and the Stripe fee is a separate cost; the payout is not revenue. — [stripto.pl](https://stripto.pl/jak-rozliczyc-stripe); [ksiegai.pl](https://www.ksiegai.pl/stripe-dla-saas/); [generaloffice.pl](https://generaloffice.pl/rozliczenia-stripe/)

### Inferences
- **Split the fee invoice by Stripe product.** A SaaS using Stripe **Billing** (subscriptions) and possibly **Tax/Radar** will likely have fees that are not payment services (Billing, Tax, Radar) alongside processing fees, and each line needs a VAT classification. For a **VAT-exempt** JDG, the non-payment lines mean real 23% VAT payable through VAT-9M every month, which is a hidden cost of Stripe Billing. Processing fees treated as exempt mean base-only VAT-9M reporting until the art. 17 ust. 1j change takes effect, if it does. The accountant needs the **Stripe tax invoice itself** plus the per-product fee breakdown.
- **Ryczałt:** fees reduce nothing because ryczałt is levied on revenue (art. 6 ust. 1, art. 12). Revenue must be reported **gross**, and payouts must never be used as revenue. Under KPiR the fees (and VAT-9M VAT for exempt payers) are costs.
- **PayPal-on-Stripe:** PayPal's fee documents must be downloaded from the PayPal dashboard monthly. Their VAT treatment (exempt payment service vs taxable) must be classified separately; the issuing PayPal entity was not verified.

### Gaps
- No KIS interpretation specifically on **Stripe** (processing vs Billing/Tax/Radar fees) was found. Search abstracts mentioned a 2023 interpretation treating payment-operator commission as exempt and a 6 Jul 2026 interpretation treating a platform commission as a taxable import, but I could not open or verify those documents.
- The exact line items, currency and issuing entity name on a Polish account's Stripe tax invoice were not verified (Stripe states the "establishment in Ireland" supplies the services). The PayPal fee-document issuer and frequency were also not verified.
- Whether the art. 17 ust. 1j / exempt-import deregulation has been **enacted and published** (and from when) remains unconfirmed as of 19 Oct 2026. Re-check ISAP.

## 7. What Polish accounting offices typically request monthly from SaaS/e-commerce clients paid via Stripe/PayPal/PayU/Przelewy24, and what the monthly handover should contain

### Takeaway
Accounting offices ask monthly for:
- a **per-transaction sales listing** (Stripe Balance/transactions with the amount paid and date);
- **payout reports**;
- **Stripe fee invoices** (plus PayPal fee documents for PayPal-on-Stripe);
- **refund/chargeback** information;
- **bank statements** (incl. FX accounts).

A bank statement alone is insufficient because payouts are net and aggregated. The law requires per-day aggregates for non-invoiced revenue and individual entries for invoices. The cash-register exemption and VAT/OSS require per-transaction traceability and location evidence. So the handover should be **per-transaction detail plus derived daily and monthly summaries**, not a summary only.

### Cited Findings
- [PRACTICE] General Office (biuro rachunkowe, 2026):
  - "Co miesiąc prosimy o: zestawienie transakcji (Balance) z kwotą zapłaconą przez klienta i datą", "zestawienie wypłat (Payouts), faktury Stripe za prowizje", refund and chargeback info, and bank statements incl. foreign-currency accounts.
  - "Stripe wypłaca jedną kwotę za setki transakcji", so the bank statement alone is not enough.
  - Revenue is booked "na podstawie faktur, raportów z kasy fiskalnej albo dziennego zestawienia sprzedaży"; "zestawienie sprzedaży prowadzisz codziennie, nie raz w miesiącu".
  - Source: [generaloffice.pl](https://generaloffice.pl/rozliczenia-stripe/)
- [PRACTICE] KsięgaI (vendor): imports sales, fees and payouts; "Zbiorcze rozliczenie prowizji za okres"; matches payouts to bank lines by amount and date ±4 days. — [ksiegai.pl](https://www.ksiegai.pl/stripe-dla-saas/)
- [PRACTICE] Stripto (vendor): fee invoice from Settings → Documents; revenue gross at charge date; NBP D-1 rate; B2B invoices via KSeF. — [stripto.pl](https://stripto.pl/jak-rozliczyc-stripe)
- [LAW] Statutory timing anchors for the handover:
  - Ewidencja/KPiR invoice entries by the 20th (§ 6 poz. 1294; § 11 ust. 2 poz. 1299).
  - KPiR daily dowód/zestawienie entries before the next business day (§ 11 ust. 3, § 18 poz. 1299).
  - VAT-UE informacja by the 25th (art. 100 ust. 3).
  - VAT-9M by the 25th (art. 99 ust. 9).
  - OSS quarterly by the end of the following month (art. 130c ust. 2).
  - Daily VAT-exempt sales record (art. 109 ust. 1).
  - Sources: [Dz.U. 2025 poz. 1294](https://eli.gov.pl/eli/DU/2025/1294/ogl/pol/pdf); [Dz.U. 2025 poz. 1299](https://eli.gov.pl/eli/DU/2025/1299/ogl/pol/pdf); [VAT act t.j.](https://eli.gov.pl/eli/DU/2025/775/ogl/pol/pdf)
- [OFFICIAL] Invoices must be entered one by one in KPiR/ewidencja from 2026 (no zestawienie). — [MF answer, interpelacja 17841](https://api.sejm.gov.pl/sejm/term10/interpellations/attachment/ATTDVRCQQ/i17841-o1.pdf)
- [Stripe] PayPal fees appear in Balance reports ("Less fees") but not on Stripe's tax invoice. — [Stripe docs – PayPal](https://docs.stripe.com/payments/paypal)

### Inferences
Below is the proposed handover specification, synthesized from the law and practice above.

**A. Per-transaction file (CSV), all branches.** One row per charge, refund, dispute event or fee adjustment, covering Stripe and PayPal-on-Stripe:
- *Identifiers:* Stripe object IDs (payment/charge, balance transaction, invoice number if any, subscription, customer); payout ID and payout date.
- *Timing:* timestamp in UTC **and** Europe/Warsaw date; service period start/end.
- *Product:* product/plan, for the ryczałt rate and PKWiU mapping.
- *Amounts and currencies:* presentment currency and gross amount; tax collected (if any); fee (Stripe vs PayPal); net; settlement currency and Stripe FX rate; NBP table, date and rate used; PLN amounts.
- *Customer and VAT data:* customer type (B2B/B2C); customer VAT ID and VIES check result; location-evidence fields (billing country, card country, IP country, PayPal payer country); assigned VAT category (see section 4).
- *Payment:* payment method.
- *Links:* refund/dispute linkage to the original charge.

This file underpins the poz. 42 cash-register exemption, OSS records (10-year retention) and the art. 14 ust. 1m corrections.

**B. Daily aggregate (dowód wewnętrzny / zestawienie sprzedaży)** for revenue not covered by invoices, numbered, per Warsaw date:
- Ryczałt: PLN totals split by ryczałt rate (§ 4 ust. 1 poz. 1294).
- KPiR: daily PLN total (§ 11, § 17–18 poz. 1299).
- VAT-active: additionally split by VAT category and rate for the JPK_V7 "WEW" summary.
- Ideally generated and archived **daily**, then bundled monthly.

**C. Invoices issued** (B2B always; B2C only on request): an individual list with number, date, KSeF number (if any), buyer tax ID with country code, net/VAT/gross, currency, and rate/NP/reverse-charge flag. Include correcting invoices.

**D. Monthly VAT-specific summaries:**
- EU B2B per customer VAT ID (VAT-UE);
- EU B2C per country under OSS, quarterly in EUR at the quarter-end ECB rate;
- EU B2C under the threshold taxed in PL ("WSTO_EE");
- non-EU NP total;
- PL B2C total.

**E. Running thresholds:**
- 240,000 zł (art. 113);
- 42,000 zł / EUR 10,000 (art. 28k);
- 20,000 zł B2C (cash register, § 3);
- EUR 100,000 (evidence, art. 24b).

**F. Supporting documents:** Stripe monthly tax invoice (fees); PayPal fee documents; Stripe payout reconciliation (charges per payout); bank statements (PLN and FX); dispute evidence/outcomes.

**G. Branch differences (who uses what):**

| Branch | What the accountant uses |
|---|---|
| Ryczałt | B (rate split), C, F, and the exchange-difference data (art. 6 ust. 1c) |
| KPiR | B, C, F, exchange differences (art. 24c), fees as costs, VAT-9M VAT as cost when paid (exempt payer) |
| Active VAT payer | JPK_V7M from A–D, plus import of services from the Stripe/PayPal invoices |
| VAT-exempt | daily art. 109 ust. 1 record, VAT-9M, VAT-UE registration (and VAT-UE returns if any EU B2B sales) |

**On the script question.** Most per-transaction fields already exist in Stripe's own exports (Balance/transactions, Payouts). The value a script (e.g., in Go) adds is:
- Warsaw-date bucketing;
- NBP rate lookup;
- VAT/ryczałt classification;
- threshold tracking;
- generating the daily dowód wewnętrzny/zestawienie.

A summary-only email would not meet the per-day and per-transaction requirements above.

### Gaps
- I found no published checklist from wFirma, inFakt, iFirma or Fakturownia specifically for Stripe monthly handovers. The Stripe guidance found came from one accounting office (General Office) and invoicing/bookkeeping vendors (Stripto, KsięgaI, ifirma 2021).
- Preferred file formats (CSV vs XLSX vs JPK-like XML) vary by office and were not documented in the sources. Ask the accountant.
- **Open points to confirm with the accountant:**
  1. Revenue-date method (charge date vs billing-period end) and NBP date.
  2. Ryczałt rate / PKWiU for the SaaS.
  3. Whether Stripe/PayPal payments are treated as satisfying poz. 42, or the 20k limit is relied on instead.
  4. The VAT status path (exempt vs active; OSS registration when crossing EUR 10k).
  5. The VAT classification of each Stripe fee line (processing vs Billing/Tax/Radar) and of PayPal fees.
  6. Whether exempt imports are still reported (art. 17 ust. 1j status).
  7. Treatment of FX differences from Stripe's conversion.
  8. Whether OSS revenue for PIT is net of foreign VAT.
  9. B2C invoice policy (issue only on request).
  10. Who produces the daily dowód wewnętrzny/zestawienie and when.
  11. The JPK_PKPIR/JPK_EWP start year for this taxpayer (2026 vs 2027).
