# KSeF obligations (state as of 19 Oct 2026) and Polish invoicing/accounting tools that integrate with Stripe — for a JDG selling SaaS

Scope note: KSeF law + Stripe consequences + tool comparison + "document per payment vs monthly summary" trade-off. General PIT/VAT record-keeping, Stripe's own reports and Patreon are covered by other researchers. "MF" = Ministerstwo Finansów; "FAQ" = official "Pytania i odpowiedzi KSeF 2.0" on ksef.podatki.gov.pl.

## 1. KSeF timeline and scope actually in force in October 2026 (dates, deferrals, penalties, VAT-exempt, foreign buyers, B2C, "do paragonu", 2025–2026 changes)

### Takeaway
Under current law, issuing invoices in KSeF has been mandatory since 1 Feb 2026 for taxpayers with 2024 sales above PLN 200 m and since 1 Apr 2026 for everyone else, VAT-exempt taxpayers included whenever they issue an invoice; until 31 Dec 2026 only taxpayers with ≤ PLN 10,000 gross of B2B-invoiced sales per month and cash-register invoices may stay outside, and from 1 Jan 2027 there are no exceptions (receiving via KSeF has been mandatory for all since 1 Feb 2026). Consumer invoices are voluntary in KSeF, invoices to foreign businesses (EU and non-EU) must be issued in KSeF and delivered outside it with a QR code, and penalties (art. 106ni) start 1 Jan 2027 under current law, although an MF draft of ~23 Sep 2026 (not yet law as of 8 Oct 2026) would defer them to 1 Jan 2028.

### Cited Findings

**Legal basis**
- The original KSeF act is the ustawa z 16 czerwca 2023 r. (Dz.U. 2023 poz. 1598). The current timeline and the transitional no-penalty period come from the ustawa z 5 sierpnia 2025 r. (Dz.U. 2025 poz. 1203). — [VATAX, updated 7.10.2026](https://www.vatax.pl/blog/ksef-fiskus-nie-bedzie-karac)
- The exclusions are set by the Rozporządzenie Ministra Finansów i Gospodarki z 7 grudnia 2025 r. w sprawie przypadków, w których nie ma obowiązku wystawiania faktur ustrukturyzowanych (Dz.U. poz. 1740). Examples: self-billing where neither party has a Polish NIP, and single-use passenger tickets. — [MF: Zasady obowiązywania KSeF i przepisy prawne, 27.04.2026](https://ksef.podatki.gov.pl/ksef-news/zasady-obowiazywania-ksef-i-przepisy-prawne/)
- Earlier history: the mandatory start was postponed to 1 Feb 2026, per the MF notice "Obowiązkowy KSeF odroczony do 1 lutego 2026 r." — [ksef.podatki.gov.pl](https://ksef.podatki.gov.pl/wyjasnienia/obowiazkowy-ksef-odroczony-do-1-lutego-2026-r/)
- The official manual "Podręcznik KSeF 2.0" is published in parts. Parts I and II carry 9 Feb 2026 versions. — [Podręcznik cz. I](https://ksef.podatki.gov.pl/media/jzrevse3/podrecznik-ksef-20-cz-i-rozpoczecie-korzystania-z-ksef-20260209.pdf), [Podręcznik cz. II](https://ksef.podatki.gov.pl/media/qzlha3vx/podrecznik-ksef-20-cz-ii-wystawianie-i-otrzymywanie-faktur-w-ksef-20260209.pdf)

**Dates**
- **1 Feb 2026:** issuing becomes mandatory for taxpayers whose sales including VAT exceeded PLN 200 m in 2024. Receiving via KSeF becomes mandatory for all taxpayers. — [MF: Zakres obowiązkowego KSeF (modified 06.02.2026)](https://ksef.podatki.gov.pl/informacje-ogolne-ksef-20/zakres-obowiazkowego-ksef/)
- **1 Apr 2026:** issuing becomes mandatory for all other taxpayers, except those under the PLN 10,000 monthly rule. — [MF: Zakres obowiązkowego KSeF](https://ksef.podatki.gov.pl/informacje-ogolne-ksef-20/zakres-obowiazkowego-ksef/); [MF: Zasady obowiązywania](https://ksef.podatki.gov.pl/ksef-news/zasady-obowiazywania-ksef-i-przepisy-prawne/)
- **PLN 10,000 rule** (statutory basis art. 145m; applies 1 Apr 2026 – 31 Dec 2026) — [FAQ Q61](https://ksef.podatki.gov.pl/pytania-i-odpowiedzi-ksef-20/), [MF: Poniżej 10 000 zł (modified 20.11.2025)](https://ksef.podatki.gov.pl/ponizej-10-000-zl/):
  - Invoices may be issued outside KSeF (paper or electronic) while the gross value of sales documented by invoices in a calendar month is ≤ PLN 10,000.
  - Only invoices that would otherwise have to be in KSeF (B2B) count toward the limit.
  - Consumer invoices, cash-register-only consumer sales, cash-register invoices and receipts ≤ PLN 450 do not count.
  - Once the limit is exceeded, the crossing invoice and every later invoice must go through KSeF. Falling back under the limit in a later month does not switch the obligation off.
  - The month is set by the invoice issue date, not the sale date.
- **1 Jan 2027:** "Od 1 stycznia 2027 r. system KSeF będzie obowiązkowy dla wszystkich podatników". Invoices outside KSeF under the 10k rule are possible only until 31 Dec 2026, and a receipt with NIP stops counting as a simplified invoice. — [IAS Wrocław (gov.pl), 9.10.2026](https://www.gov.pl/web/ias-wroclaw/krajowy-system-e-faktur-ksef--iii-etap-wdrozenia-od-1-stycznia-2027-r-co-warto-wiedziec)
- **Cash registers:** until the end of 2026, obligated taxpayers may still issue invoices from cash registers, and receipts with the buyer's NIP up to PLN 450 / EUR 100 count as simplified invoices. — [MF: Zakres obowiązkowego KSeF](https://ksef.podatki.gov.pl/informacje-ogolne-ksef-20/zakres-obowiazkowego-ksef/)
  - From 2027, an invoice for a transaction recorded on an online cash register must go through KSeF if the buyer is a taxpayer. — [FAQ Q3](https://ksef.podatki.gov.pl/pytania-i-odpowiedzi-ksef-20/)
  - Under current law, invoices from cash registers will no longer be possible from 2027. — [money.pl, 23.09.2026](https://www.money.pl/biznes/zwrot-ws-kar-za-bledy-w-ksef-jest-projekt-7332618221013248a.html)
- **KSeF number in transfers:** from 2027 the KSeF number must appear in bank-transfer titles, including split payment (MPP). This matters to B2B payers paying by bank transfer, not to Stripe card/PayPal collections. — [VATAX](https://www.vatax.pl/blog/ksef-fiskus-nie-bedzie-karac)
- **Storage:** structured invoices are kept in KSeF for 10 years from the end of the year of issue. — [MF: Zasady obowiązywania](https://ksef.podatki.gov.pl/ksef-news/zasady-obowiazywania-ksef-i-przepisy-prawne/)

**Penalties**
- **Current law:** "Kary te będą stosowane dopiero od 1 stycznia 2027 r." and sanctions do not apply retroactively to 2026. — [MF: Zasady obowiązywania](https://ksef.podatki.gov.pl/ksef-news/zasady-obowiazywania-ksef-i-przepisy-prawne/); FAQ Q23 also says penalties were postponed to 1 Jan 2027 — [FAQ](https://ksef.podatki.gov.pl/pytania-i-odpowiedzi-ksef-20/)
- **Penalty amounts (art. 106ni):** up to 100% of the VAT shown on the invoice, or up to 18.7% of the total amount for invoices without VAT. — [VATAX](https://www.vatax.pl/blog/ksef-fiskus-nie-bedzie-karac)
  - Violations covered by the deferral discussion: issuing an invoice outside KSeF, late submission of an offline invoice, and an invoice not matching the required schema. — [pit.pl, 8.10.2026](https://www.pit.pl/aktualnosci/ksef-bez-kar-za-bledy-przez-kolejny-rok-ministerstwo-finansow-zapowiada-zmiane)
- **16 Sep 2026 announcement:** the MF said it plans to extend the no-penalty period to 31 Dec 2027, so penalties would start 1 Jan 2028. — [Akademia LTCA, 16.09.2026](https://akademialtca.pl/blog/ministerstwo-finansow-opublikowalo-informacje-dotyczaca-przedluzenia-odroczenia-kar-za-bledy-w-stosowaniu-ksef-do-konca-2027-r); [pit.pl](https://www.pit.pl/aktualnosci/ksef-bez-kar-za-bledy-przez-kolejny-rok-ministerstwo-finansow-zapowiada-zmiane)
  - The MF says "odroczenie kar nie oznacza odroczenia samego obowiązku korzystania z KSeF" (deferring penalties does not defer the obligation).
  - The tax administration will still react to invoices issued outside KSeF and check whether tax was reported correctly.
- **Draft on RCL (~23 Sep 2026):** a draft amendment of the VAT Act was published on the RCL site. — [money.pl, 23.09.2026](https://www.money.pl/biznes/zwrot-ws-kar-za-bledy-w-ksef-jest-projekt-7332618221013248a.html)
  - It postpones KSeF penalties by one year, to the end of 2027.
  - Planned entry into force: 31 Dec 2026.
  - It also extends reverse charge for gas, electricity and emission allowances.
  - MF's rationale: it "zapewni czas na dostosowanie się przez podatników" (gives taxpayers time to adapt).
- **Status as of 7–8 Oct 2026:** not yet law; the full legislative process is still needed. — [VATAX, updated 7.10.2026](https://www.vatax.pl/blog/ksef-fiskus-nie-bedzie-karac); [pit.pl, 8.10.2026](https://www.pit.pl/aktualnosci/ksef-bez-kar-za-bledy-przez-kolejny-rok-ministerstwo-finansow-zapowiada-zmiane)

**VAT-exempt taxpayers (zwolnieni podmiotowo, art. 113)**
- FAQ: "każdy podatnik (czynny i zwolniony), który ma obowiązek wystawienia faktury, będzie wystawiał ją w KSeF". — [FAQ Q10](https://ksef.podatki.gov.pl/pytania-i-odpowiedzi-ksef-20/)
- MF also states: "podatnicy zwolnieni z VAT przedmiotowo lub podmiotowo mają obowiązek wystawienia faktur przy użyciu KSeF". The 10k rule applies to them too: those at ≤ PLN 10,000 gross a month start on 1 Jan 2027. — [MF: Zasady obowiązywania](https://ksef.podatki.gov.pl/ksef-news/zasady-obowiazywania-ksef-i-przepisy-prawne/)
- Under the art. 113 exemption, an invoice is issued only on the buyer's demand, or voluntarily. The exemption limit is PLN 240,000. (Seen in a search extract of the PDF; I did not read the whole PDF.) — [Podręcznik KSeF 2.0 cz. II (09.02.2026)](https://ksef.podatki.gov.pl/media/qzlha3vx/podrecznik-ksef-20-cz-ii-wystawianie-i-otrzymywanie-faktur-w-ksef-20260209.pdf)
- A KIS individual interpretation: if no buyer demands an invoice for an exempt sale, no structured invoice has to be issued. (Search extract.) — [Eureka 696263](https://eureka.mf.gov.pl/informacje/podglad/696263)
- Invoices under the EU SME cross-border scheme (art. 113a) are not issued in KSeF, and voluntary issue there is not possible. — [MF: Zakres obowiązkowego KSeF](https://ksef.podatki.gov.pl/informacje-ogolne-ksef-20/zakres-obowiazkowego-ksef/)

**Consumers (B2C)**
- "wystawianie faktur B2C (na rzecz osób prywatnych – konsumentów) w KSeF jest dobrowolne". — [FAQ Q5](https://ksef.podatki.gov.pl/pytania-i-odpowiedzi-ksef-20/); also FAQ Q54 and Q80
- Invoices to natural persons not running a business are excluded from the obligation; issuing them in KSeF voluntarily is possible. — [MF: Zakres obowiązkowego KSeF](https://ksef.podatki.gov.pl/informacje-ogolne-ksef-20/zakres-obowiazkowego-ksef/)
- If issued voluntarily in KSeF, the invoice is delivered to the consumer in an agreed form, e.g. a printout with the QR code and KSeF number. — [MF: Zasady obowiązywania](https://ksef.podatki.gov.pl/ksef-news/zasady-obowiazywania-ksef-i-przepisy-prawne/)
- Commentary by Mikołaj Ratajczak (TPA), 23.03.2026: [infor.pl](https://ksiegowosc.infor.pl/podatki/vat/faktura/7529636,faktury-dla-konsumentow-w-2026-r-czy-mozna-wystawiac-w-ksef-jak-przekazac-nabywcy.html)
  - The legal basis for the exclusion is art. 106ga ust. 2 pkt 4.
  - A consumer invoice is owed only on request made within 3 months; KSeF does not change that.
  - Outside KSeF it can be paper or a PDF sent by e-mail or through a portal.
  - For online sales, set the delivery method in the terms of service.

**Foreign buyers (EU and non-EU, no Polish NIP)**
- FAQ Q9: "Jak wystawić fakturę dla klienta spoza UE? — Wystawiasz ją w KSeF, a następnie przekazujesz nabywcy w ustalony sposób (np. PDF)." — [FAQ](https://ksef.podatki.gov.pl/pytania-i-odpowiedzi-ksef-20/)
- FAQ Q36: a foreign recipient "nie będzie mógł zalogować się do systemu i jej pobrać". The invoice "powinna być opatrzona kodem QR" so the buyer can verify it exists in KSeF. — [FAQ](https://ksef.podatki.gov.pl/pytania-i-odpowiedzi-ksef-20/)
- MF example: a Polish service provider invoicing a Swedish VAT taxpayer "wystawia tę fakturę w KSeF". The invoice is passed to the buyer "w sposób z nim uzgodniony (np. w formie wydruku lub pliku pdf, opatrzoną kodem QR)". Receipt through KSeF does not apply to entities listed in art. 106gb ust. 4, including foreign entities and consumers. — [MF: Zasady obowiązywania](https://ksef.podatki.gov.pl/ksef-news/zasady-obowiazywania-ksef-i-przepisy-prawne/)
- QR rules for invoices delivered outside KSeF to art. 106gb ust. 4 buyers — [MF: Potwierdzenie transakcji (modified 06.02.2026)](https://ksef.podatki.gov.pl/informacje-ogolne-ksef-20/potwierdzenie-transakcji/):
  - Online invoice: one QR code with the KSeF number.
  - Offline invoice without a KSeF number yet: two QR codes, labelled "OFFLINE" and "CERTYFIKAT".
  - A voluntary "potwierdzenie transakcji" (transaction confirmation) is explicitly "nie jest fakturą" (not an invoice).
- Generating the second (CERTYFIKAT) QR code requires a KSeF certificate of type 2. (Search extract of official FAQ.) — [FAQ](https://ksef.podatki.gov.pl/pytania-i-odpowiedzi-ksef-20/)
- Exclusions concern foreign *sellers* without a Polish seat or fixed establishment, the non-EU OSS, IOSS and SME procedures, and self-billing without a Polish NIP. A Polish seller invoicing a foreign *buyer* is not excluded. — [MF: Zakres obowiązkowego KSeF](https://ksef.podatki.gov.pl/informacje-ogolne-ksef-20/zakres-obowiazkowego-ksef/)
- **Foreign currency:** on KSeF invoices amounts stay in the foreign currency, and only the VAT amount (field P_14_xW) is converted to PLN, with the currency code and rate fields filled. — [FAQ Q19](https://ksef.podatki.gov.pl/pytania-i-odpowiedzi-ksef-20/)
  - This reflects the general rule in art. 106e ust. 11 that VAT on an invoice is shown in PLN. (Search-extract level for the commentary.) — [podatki.biz](https://www.podatki.biz/artykuly/faktura-w-walucie-obcej-pomiedzy-polskimi-podatnikami-vat_55_41139.htm)

**"Faktura do paragonu" / cash registers**
- An invoice to a receipt for a consumer may be issued in KSeF voluntarily. — [FAQ Q25](https://ksef.podatki.gov.pl/pytania-i-odpowiedzi-ksef-20/)
- Receipts count as invoices under art. 106e ust. 5 pkt 3 only until the end of the transition; after it "wszystkie faktury będą wystawiane przy użyciu KSeF". — [FAQ Q66, Q48, Q56](https://ksef.podatki.gov.pl/pytania-i-odpowiedzi-ksef-20/)
- Context for whether this applies to SaaS at all:
  - Item 37 of the cash-register exemption regulation exempts online services when payment comes entirely via post, bank or SKOK and the records show which service each payment was for.
  - iFirma's 2023 article says PayU, Przelewy24 and PayPal flows qualify. It cites interpretation 0111-KDIB3-2.4012.548.2021.2.MN (card payments). Stripe is not mentioned. — [iFirma blog, 21.04.2023](https://www.ifirma.pl/blog/sprzedaz-w-internecie-z-kasa-fiskalna-czy-bez.html)

**Other 2025–2026 developments**
- On 9 Jun 2026 the MF held its first consultations after partial roll-out, on proposed changes to business-event models and invoice synchronisation. — [podatki.gov.pl, 03.06.2026](https://podatki.gov.pl/aktualnosci/zapraszamy-na-konsultacje-dotyczace-zmian-w-ksef-03062026/)
- The MF proposed keeping KSeF authentication tokens after 31 Dec 2026; they were previously due to be replaced by certificates (interpelacja nr 18129). (Search extract.) — [podatki.biz](https://www.podatki.biz/artykuly/beda-zmiany-w-ksef-tokeny-maja-pozostac_55_62973.htm)
- A PiS MPs' draft (no. 5347, Feb/Mar 2026) would let microentrepreneurs issue paper or e-invoices instead of KSeF invoices from 1 Apr 2026 to 31 Dec 2027. The source gives no legislative progress. — [Lege Artis](https://czasopismo.legeartis.org/2026/02/odroczenie-ksef-mikroprzedsiebiorcow-31-grudnia-2027-obowiazek-wystawiania-faktur/)
- **Accountant access:** you can grant KSeF permissions to the whole accounting firm rather than each employee. — [FAQ "Biuro rachunkowe wystawia moje faktury…"](https://ksef.podatki.gov.pl/pytania-i-odpowiedzi-ksef-20/)
  - More than one person can be authorised to issue and receive invoices. — [FAQ Q52](https://ksef.podatki.gov.pl/pytania-i-odpowiedzi-ksef-20/)

### Inferences
- **PIT form:** the choice of ryczałt, liniowy or skala has no bearing on KSeF. KSeF is a VAT Act obligation tied to issuing invoices, and the MF wording covers "każdy podatnik (czynny i zwolniony)". The PIT form matters only for record-keeping, which is covered by other researchers.
- **Branch matrix** for this JDG, derived from the cited rules:

| Seller status | Buyer | Invoice owed? | KSeF in Oct 2026 | KSeF from 1 Jan 2027 |
|---|---|---|---|---|
| czynny VAT | consumer (PL, EU, non-EU) | only on request (3 months) | voluntary | voluntary |
| czynny VAT | PL business | yes | mandatory, unless monthly B2B-invoiced gross ≤ PLN 10k (until the crossing invoice) | mandatory |
| czynny VAT | EU business (reverse charge) or non-EU business | yes | issue in KSeF (same 10k caveat); deliver PDF/print with QR outside KSeF | same, no 10k caveat |
| zwolniony podmiotowo (≤ PLN 240k) | business | on demand (or voluntary) | if issued, in KSeF (10k rule until end-2026) | if issued, in KSeF |
| zwolniony podmiotowo | consumer | on demand | voluntary | voluntary |

- **B2C-heavy seller in 2026:** a SaaS with few B2B invoices (≤ PLN 10k gross a month) may legally still be outside KSeF for issuing until 31 Dec 2026, because B2C invoices do not count toward the limit. Any setup chosen now must handle KSeF for B2B invoices by 1 Jan 2027 regardless of the penalty-deferral draft.
- **Receiving is already live:** receiving via KSeF is mandatory now, so purchase invoices from Polish suppliers arrive in KSeF and the accountant can pull them directly once given permissions. Foreign suppliers (e.g. Stripe's own fee invoices) remain outside KSeF.
- **Penalty deferral does not legalise skipping KSeF:** even if the draft passes, an obligated B2B invoice issued outside KSeF remains a breach. Only the fine is deferred, and the MF says it will still follow up.

### Gaps
- No UD number or text found for the ~23 Sep 2026 RCL draft; no information after 8 Oct 2026 on whether it passed the Sejm or was signed. The safe reading for 19 Oct 2026: penalties from 1 Jan 2027 under current law.
- Status of the PiS microentrepreneur draft (no. 5347) unknown. It was clearly not in force on 1 Apr 2026, but I found no source on rejection or withdrawal.
- A June 2026 draft amendment of the cash-register regulation (removing invoices from cash registers from 2027) appeared only in an unattributed search summary; unverified.
- The Podręcznik KSeF 2.0 PDFs were not read in full; the 240k limit and the "invoice on demand" points rely on search extracts of the official PDF.
- Whether a zwolniony podmiotowo JDG must invoice B2B e-services to EU businesses (place of supply abroad, reverse charge) is a VAT question; not verified here.

## 2. Consequences for Stripe: can Stripe-generated invoices still be the legal VAT invoice, and does Stripe offer KSeF integration?

### Takeaway
Once the seller is in KSeF (1 Apr 2026, or from the invoice crossing the 10k limit, at the latest 1 Jan 2027), a Stripe Billing/Invoicing PDF cannot be the legal VAT invoice for any business buyer, Polish or foreign; for consumers it can remain an invoice outside KSeF only if it meets art. 106e (e.g. VAT shown in PLN on foreign-currency invoices). Stripe has no native KSeF or e-invoicing support and points to Marketplace apps, of which S2K – Seamless KSeF Integration is the one KSeF app found; other routes are Fakturownia's Stripe app, the Stripto bridge and Invopop's API.

### Cited Findings
- Stripe (page updated 20 Aug 2026): "While Stripe Billing and Stripe Invoicing do not directly generate or send e-invoices", the Stripe App Marketplace "offers multiple options that allow you to use third-party providers". The page does not mention KSeF. — [Stripe: e-invoicing France guide](https://stripe.com/en-pl/resources/more/e-invoicing-france)
- Searches of docs.stripe.com, stripe.com and support.stripe.com for Poland/KSeF e-invoicing returned no Stripe document on KSeF. Stripe Invoicing lists BLIK/PLN for Poland, but no e-invoicing. — [Stripe Invoicing payment methods](https://docs.stripe.com/invoicing/payment-methods)
- **S2K – Seamless KSeF Integration** — [Stripe Marketplace listing](https://marketplace.stripe.com/apps/s2k---seamless-ksef-integration); details from the vendor site, because the listing did not render — [s2ksef.com](https://s2ksef.com/):
  - Claims it "Automatycznie wysyła faktury ze Stripe do KSeF bezpośrednio z dashboardu Stripe".
  - Authenticates via Profil Zaufany and generates certificates.
  - "Automatycznie wykrywa faktury z polskim NIP-em"; missing GTU, PKWiU or buyer data are supplied via Stripe metadata.
  - Price: €8/month (25 invoices), €20 (200), €50 (1,000), Enterprise on request; 14 days free in production.
  - The operator is not named, and the site says nothing about foreign or consumer buyers, QR codes, NBP conversion, OSS or accountant export.
- **Stripto** (Polish service), vendor page dated "2026" — [stripto.pl](https://stripto.pl/stripe-integracja-ksef):
  - Mode 1: listens to Stripe payments and creates the invoice in Fakturownia, inFakt or wFirma, which then send it to KSeF.
  - Mode 2: builds FA(3) XML and sends it directly to KSeF using a KSeF token and NIP; then no copy lands in an accounting program.
  - Price: Mini 19 zł (25 invoices), Starter 39 zł (100), Pro 79 zł (1,000) per month; annual from 190 zł; one integration active at a time.
  - Issues invoices in the payment currency, with PLN conversion at the NBP rate from the day before the sale done by the accounting tool.
  - Detects EU reverse charge and supports 23/8/5%, zw and np rates; OSS not mentioned.
  - Can skip invoicing buyers without a NIP.
  - Claims Stripe PDFs "nie są uznawane jako ważne prawnie w Polsce od momentu, gdy KSeF stał się obowiązkowy", without citing a legal source.
  - Also claims (competitor statement) that S2K and "KSeF Kit" depend on Stripe Invoicing at 0.4% per invoice and miss payments with no Stripe invoice, such as Payment Links.
- **Invopop** — [Invopop KSeF Poland app docs](https://docs.invopop.com/apps/ksef-poland); [Invopop Poland coverage](https://invopop.com/coverage/poland); [Invopop KSeF guide](https://docs.invopop.com/guides/pl-ksef-v2); [Invopop Stripe app docs](https://docs.invopop.com/apps/stripe):
  - Its KSeF 2.0 app issues FA(3) invoices, credit notes (KOR), prepayment invoices, reverse-charge invoices, voluntary B2C invoices and self-billing.
  - It generates QR codes and says foreign recipients must get the invoice outside KSeF.
  - It requires a .p12 KSeF certificate with InvoiceRead/InvoiceWrite permissions; no pricing is published.
  - Its Stripe app converts finalized Stripe invoices and credit notes into e-invoices, but the Stripe-app page names Spain, XRechnung, Peppol and France, not Poland. The Stripe→KSeF combination is therefore unverified.
- **Fakturownia (as a Stripe→KSeF route):** its help tells users to switch off Stripe's "After Payment > Create an invoice PDF" to avoid double invoicing. Documents created from Stripe payments go to KSeF like any other Fakturownia document. — [Fakturownia help](https://pomoc.fakturownia.pl/integracja-fakturowni-ze-stripe-automatyczne-fakturowanie-platnosci); [Fakturownia + Stripe](https://fakturownia.pl/apps/stripe-app)
- **Freemius** (a Merchant-of-Record competitor) argues in marketing copy that Stripe's invoice handling clashes with KSeF: refunds alter existing invoices rather than issuing correction notes, and invoice finalization follows payment timing. Treat as vendor framing. — [Freemius](https://freemius.com/poland-ksef-stripe-saas-billing-compliance/)
- **Rule references from Section 1:**
  - B2B invoices, including to foreign businesses, must be issued in KSeF and delivered outside it with a QR code. — [FAQ Q9/Q36](https://ksef.podatki.gov.pl/pytania-i-odpowiedzi-ksef-20/); [MF: Zasady obowiązywania](https://ksef.podatki.gov.pl/ksef-news/zasady-obowiazywania-ksef-i-przepisy-prawne/)
  - Consumer invoices are voluntary in KSeF. — [FAQ Q5](https://ksef.podatki.gov.pl/pytania-i-odpowiedzi-ksef-20/)
  - Foreign-currency invoices must show VAT in PLN (P_14_xW). — [FAQ Q19](https://ksef.podatki.gov.pl/pytania-i-odpowiedzi-ksef-20/)

### Inferences
- **B2B customers (PL, EU, non-EU):** the Stripe PDF must stop being "the invoice".
  - The legal document is the KSeF invoice, issued by a tool (Fakturownia, Stripto, S2K, own code) and e-mailed to foreign buyers as a PDF with the KSeF QR code.
  - Stripe's own invoice e-mail/PDF should be disabled or reduced to a payment receipt, so there are not two documents titled "invoice" for one sale (the double-invoicing concern Fakturownia flags).
  - Refunds to business buyers need a KSeF corrective invoice (faktura korygująca); a Stripe credit note alone will not do.
- **Consumers:** a Stripe PDF can legally remain the invoice outside KSeF, but only if it satisfies art. 106e. I did not verify whether Stripe's template can show VAT in PLN for EUR/USD invoices; this is the likeliest content gap and should be checked before relying on Stripe PDFs as consumer invoices.
- **PayPal-via-Stripe:** this changes nothing legally. The invoicing duty depends on the buyer, not the payment method. Tools triggered by Stripe payment or invoice events probably treat PayPal charges like card charges, but no vendor documents this; test it in Stripe test mode.
- **Building it yourself:** writing a direct KSeF integration (e.g. in Go) means handling FA(3) XML, KSeF 2.0 authentication (tokens/certificates), QR generation, offline/emergency modes and corrections. The bridges cost roughly €8–50 or 19–79 zł a month, so building only makes sense if there is a strong reason.

### Gaps
- No official Stripe statement on KSeF found.
- S2K's operator, launch date, reviews and handling of foreign buyers, QR codes, consumers and corrections are unknown; the Marketplace page did not render.
- Whether Invopop's Stripe app can target KSeF is undocumented.
- "KSeF Kit" as a Stripe app was not found; only a .NET library of that name ([nuget](https://www.nuget.org/packages/MarcinBorecki.KsefKit)).
- Striptu ([striptu.com](http://striptu.com/pl/integracje/integracja-stripe-ksef)) advertises a Stripe–KSeF connection; not fetched or verified.
- Stripe Invoicing's per-invoice price (0.4%) comes only from a competitor's page; not checked against Stripe.

## 3. Polish invoicing/accounting products with Stripe integration — what they automate, KSeF, FX/NBP, OSS, PayPal-via-Stripe, accountant hand-off, 2026 price, limitations

### Takeaway
Of the big four, only **Fakturownia** has a native, documented Stripe → invoice/receipt flow (every successful payment including renewals, invoices for companies and receipts for consumers, KSeF sending, OSS rates, NBP rates, free accountant access); inFakt works only through a beta Zapier recipe that creates drafts, wFirma has no Stripe integration and iFirma lists Stripe without documenting what it does. Bridges (Stripto 19–79 zł/month, S2K €8–50/month) add KSeF on top of Stripe, no Stripe integration was found for Faktura.pl, Saldeo SMART, Taxxo, Comarch ERP XT or InsERT, and no vendor mentions PayPal-via-Stripe.

### Cited Findings

**Fakturownia (fakturownia.pl) — native integration "Fakturownia + Stripe"**
- **Setup:** paste Stripe publishable and secret keys (test or live); "Fakturownia automatycznie utworzy webhook w Stripe". — [Fakturownia + Stripe](https://fakturownia.pl/apps/stripe-app)
- **What it automates:** "Po każdej udanej płatności w Stripe, system automatycznie wystawia dokument sprzedaży." — [Fakturownia help](https://pomoc.fakturownia.pl/integracja-fakturowni-ze-stripe-automatyczne-fakturowanie-platnosci)
  - Covers Checkout, embedded page, Payment Links and site forms.
  - For subscriptions it issues a document after the first payment and after each automatic renewal.
  - Flow: webhook → payment fetched via Stripe API → a payment marked "Stripe" appears in the Płatności tab → document issued.
  - Documents are issued automatically; no draft mode is described.
- **Document types:** VAT invoice per payment, proforma, receipt (paragon) per payment, or the mixed "Faktura VAT dla firm, paragon dla osób prywatnych" (invoice if a NIP is given, receipt if not). You set the department (dział) and a default VAT rate. — [Fakturownia help](https://pomoc.fakturownia.pl/integracja-fakturowni-ze-stripe-automatyczne-fakturowanie-platnosci)
- **Buyer data:** Stripe must be configured to collect it (`tax_id_collection`, `billing_address_collection: 'required'`, `name_collection`, or the Payment Link equivalents). With a foreign NIP the document "zawsze zostanie wygenerowany" with ID type VAT UE; optional NIP validation is for Polish NIPs only. — [Fakturownia help](https://pomoc.fakturownia.pl/integracja-fakturowni-ze-stripe-automatyczne-fakturowanie-platnosci)
- **Limitations Fakturownia states:** [Fakturownia help](https://pomoc.fakturownia.pl/integracja-fakturowni-ze-stripe-automatyczne-fakturowanie-platnosci)
  - Stripe coupons "nie są obecnie uwzględniane" (not currently taken into account).
  - Stripe's own "Create an invoice PDF" must be off.
  - With Stripe Tax, the integration's VAT rate must match Stripe's.
  - Refunds/corrections, PayPal and multiple Stripe accounts are not covered on the page.
- **Other Stripe apps:**
  - "Import płatności Stripe" imports past payments from a Stripe CSV for matching to invoices. It issues no documents, and it is not stated whether payouts or fees are included. — [Fakturownia app](https://fakturownia.pl/apps/stripe-import-app)
  - "Płatności Stripe" only adds a pay-online button to documents. — [Fakturownia app](https://fakturownia.pl/apps/stripe_payments)
- **KSeF:**
  - Included in all plans. — [Fakturownia cennik](https://fakturownia.pl/cennik)
  - Stripe-generated documents go to KSeF automatically or manually according to KSeF settings. — [Fakturownia + Stripe](https://fakturownia.pl/apps/stripe-app)
  - Sales invoices are sent automatically and purchase invoices fetched; UPO, KSeF number and XML show on the document. — [VATAX comparison, updated 15.10.2026](https://www.vatax.pl/blog/obsluga-ksef-porownanie-cen)
  - **Trap:** foreign-currency invoices are rejected for KSeF unless converted to PLN. The error reads "Do wysyłki do KSeF waluta faktury musi zostać przeliczona na PLN lub jej główna waluta powinna być PLN". — [Fakturownia help](https://pomoc.fakturownia.pl/blad-wysylki-do-ksef-do-wysylki-do-ksef-waluta-faktury-musi-zostac-przeliczona-na-pln-lub-jej-glowna-waluta-powinna-byc-pln)
- **FX/NBP:**
  - Any currency is allowed in all plans. — [Fakturownia cennik](https://fakturownia.pl/cennik)
  - The default rate is from the business day before the sale date, or before the issue date if that is earlier; NBP, ECB or other banks or a manual rate can be chosen. — [Fakturownia help](https://pomoc.fakturownia.pl/896746-Kurs-wedlug-Narodowego-Banku-Polskiego-zgodne-z-polskimi-przepisami-)
  - Users have reported rate mismatches. (Search extract.) — [Fakturownia help thread](https://pomoc.fakturownia.pl/153339189-Wybor-kursu-przeliczenia-na-fakturze-walutowej)
- **OSS:** — [Fakturownia: VAT OSS](https://pomoc.fakturownia.pl/1236977-vat-oss)
  - Enabled with "Faktury OSS" (optionally as the default). The VAT field then offers the buyer country's rates, kept current and archived by date.
  - The OSS box cannot be ticked when the buyer's NIP is in VIES, receipts cannot be marked OSS, and OSS invoices are left out of JPK_V7.
  - Whether the Stripe integration sets the OSS flag automatically is not documented.
- **Accountant hand-off:**
  - Free accountant access. — [Fakturownia cennik](https://fakturownia.pl/cennik)
  - XML export to Comarch ERP Optima, max 100 documents per export, KSeF numbers included. — [Fakturownia help](https://pomoc.fakturownia.pl/182024586-Eksport-pliku-Fakturowni-do-formatu-ERP-Optima)
  - Export to Symfonia FK (GTU codes added Aug 2026). — [Fakturownia exports](https://fakturownia.pl/eksporty-ksiegowe)
  - Rewizor only via iVend. — [Fakturownia apps: exports](https://fakturownia.pl/apps/c/exports)
- **Price:**
  - Micro plan: 0 zł for 3 sales documents a month. API, KSeF and shop integrations are in every plan; online payments and the bank connection start at Standard. — [Fakturownia cennik](https://fakturownia.pl/cennik)
  - Paid-plan amounts (Start/Standard/Pro/Pro Plus) were not readable in the page text.
  - A notice says Start and Standard went up and PRO/MAX went down, with no amounts. — [Fakturownia news](https://fakturownia.pl/aktualnosci/zmiany-w-abonamentach)
  - The plan required for the Stripe app is not stated.
- **User reports:** the only forum thread found on the Stripe integration has one question and no reply. — [Fakturownia forum](https://sugester.fakturownia.pl/115802371-Integracja-z-Stripe)
- **Zapier:** Fakturownia has no official Zapier app. — [Zapier community](https://community.zapier.com/how-do-i-3/add-fakturownia-invoicing-system-53531)

**inFakt (infakt.pl) — no native Stripe integration; Zapier or API** (inFakt help pages returned 403, so items marked "search extract" are from search summaries)
- **Integration:** inFakt's help shows a Zapier recipe: Stripe "New Payment" → inFakt "Create Invoice" → "Mark Invoice as Paid". (Search extract.) — [inFakt help](https://pomoc.infakt.pl/hc/pl/articles/34618561424914-Popularne-przyk%C5%82ady-automatyzacji-z-inFakt)
  - The inFakt Zapier app is labelled Beta. — [Zapier](https://zapier.com/apps/infakt/integrations/stripe)
- **How the recipe behaves** (search extract) — [inFakt help](https://pomoc.infakt.pl/hc/pl/articles/34618355010066-Jak-automatycznie-tworzy%C4%87-faktury-przez-Zapier):
  - Stripe amounts are in cents, so they must be divided by 100 with Formatter.
  - The API key needs `api:invoices:read`, `api:invoices:write` and `api:account:read`.
  - "Faktura pojawia się natychmiast ze statusem 'draft' (szkic)", and e-mailing is off by default.
- **July 2026 update** (search extract) — [inFakt help](https://pomoc.infakt.pl/hc/pl/articles/37414278043410-Co-nowego-w-integracji-inFakt-Zapier):
  - A "Create Invoice (Async)" action was added to avoid numbering conflicts.
  - "Mark as paid" previously "mogła w tle nie powieść się, a mimo to Zap kończył się 'sukcesem'" (could fail silently while the Zap reported success); it now reports the error.
- **Not found:** subscription handling, refunds, payout/fee import, PayPal.
- **OSS:** inFakt says it applies the buyer country's VAT rate on OSS invoices (article 1–2 years old). — [inFakt blog](https://www.infakt.pl/blog/jak-wystawic-fakture-oss/)
- **KSeF:** full integration, automatic sending and receiving, no extra fee. — [VATAX comparison, 15.10.2026](https://www.vatax.pl/blog/obsluga-ksef-porownanie-cen)
- **Accountant hand-off:**
  - "Dostęp Twojego księgowego": limited, revocable access. — [inFakt help](https://pomoc.infakt.pl/hc/pl/articles/115000664384-Jak-doda%C4%87-swojego-ksi%C4%99gowego-do-aplikacji-inFakt)
  - Sales invoices export as txt/csv for Optima, Insert and enova (older blog). — [inFakt blog](https://www.infakt.pl/blog/udostepnij-faktury-swojemu-ksiegowemu/)
  - inFakt's own bookkeeping service for a JDG "from 200 PLN" (page dateModified 2026-09-16; net/gross and per-document rules not shown). — [inFakt cennik](https://www.infakt.pl/cennik)
- **Price conflict:**
  - 4,99 zł net/month for basic invoicing with no document limit. — [VATAX comparison](https://www.vatax.pl/blog/obsluga-ksef-porownanie-cen)
  - Contradicted by 29/49/79 PLN, seen only in a search summary. — [kluczesoft](https://kluczesoft.pl/pomoc/porownanie/najlepszy-program-do-wystawiania-faktur)

**wFirma (wfirma.pl) — no Stripe integration found**
- **Integrations:** the page lists no Stripe, PayPal or Zapier/Make. It names a generic API, unnamed online payment operators, IdoSell, Apilo, eParagony, 12 banks, KSeF and VIES/white list. — [wFirma integracje](https://wfirma.pl/integracje-online)
  - Pay-links: Blue Media and DotPay. — [wFirma](https://wfirma.pl/szybkie-platnosci)
  - A help article covers booking sales paid via payment intermediaries in full accounting. — [wFirma help](https://pomoc.wfirma.pl/-ksiegowanie-sprzedazy-wykonanej-z-wykorzystaniem-platnosci-online-posrednicy-platnosci-w-pelnej-ksiegowosci)
- **OSS:** — [wFirma help](https://pomoc.wfirma.pl/-faktura-vat-oss-jak-wystawic)
  - Set "PODATNIK VAT ZAREJESTROWANY W OSS"; the buyer identifier is "BRAK" and the buyer country's rate is used.
  - At least two non-contradictory pieces of evidence of the consumer's country must be recorded.
  - Sales go to KPiR or the revenue register and to the VAT OSS register; the VIU-DO declaration is generated with automatic EUR conversion.
  - Available in "Fakturowanie" and "Księgowość online".
- **FX:** automatic NBP average rate from the last business day before the earlier of the sale and issue dates. — [wFirma help](https://pomoc.wfirma.pl/-po-jakim-kursie-nalezy-wystawic-fakture-w-obcej-walucie)
- **KSeF:** dedicated integration. — [wFirma KSeF](https://wfirma.pl/integracja-z-ksef)
- **Accountant hand-off:**
  - Sales invoices export to Insert, Streamsoft Ala and Comarch Optima. — [wFirma help](https://pomoc.wfirma.pl/-eksport-do-innych-programow)
  - No plans for a Symfonia export (forum). — [wFirma forum](https://forum.wfirma.pl/temat/7849-eksport-faktur-do-symfonia-finanse-i-ksiegowosc)
  - Accounting offices can connect to a client account via "Panel klienta". — [wFirma help](https://pomoc.wfirma.pl/-import-i-eksport-dokumentow)
- **Price:** 19 zł/month for invoicing incl. KSeF and 54 zł/month with accounting (third-party, net/gross not stated; 30-day trial). — [VATAX comparison](https://www.vatax.pl/blog/obsluga-ksef-porownanie-cen)

**iFirma (ifirma.pl) — Stripe listed, function undocumented**
- **Integration:** Stripe and PayPal appear under "Platformy płatnicze" with generic text ("automatyczne pobieranie danych o transakcjach, masowe przesyłanie zleceń płatniczych…"). There is no Stripe help page and no invoice-per-payment description. — [iFirma integracje](https://www.ifirma.pl/integracje/)
- **API and e-commerce limits:** — [iFirma integracje](https://www.ifirma.pl/integracje/)
  - Issuing documents through the API is included up to 1,000 a month, charged above that.
  - E-commerce integration is free up to 100 orders a month.
  - For clients of iFirma's own accounting office it costs 29 zł a month. — [iFirma cennik](https://www.ifirma.pl/cennik)
- **Stripe fee invoices:** iFirma's guide treats them as import of services, VAT-exempt as payment intermediation, converted at the NBP rate from the last business day before the invoice date and reported in JPK_V7 P_27/P_28 or P_29/P_30. — [iFirma blog](https://www.ifirma.pl/blog/aktualnosci/rozliczenie-faktury-od-stripe/)
- **OSS:**
  - Help pages exist for OSS sales documents and services to EU consumers. — [iFirma help OSS](https://pomoc.ifirma.pl/pomoc-artykul/dokumenty-sprzedazy-w-ramach-procedury-oss/), [iFirma help EU consumers](https://pomoc.ifirma.pl/pomoc-artykul/swiadczenie-uslug-na-rzecz-osob-fizycznych-z-terytorium-unii-europejskiej/)
  - iFirma (legal state 3 Jul 2025) says OSS sales need not be invoiced even if the buyer asks. — [iFirma blog](https://www.ifirma.pl/blog/czym-jest-vat-oss-ile-wynosi-limit-oraz-jak-rozliczac-sprzedaz-na-terenie-ue/)
- **KSeF:** "Pełna integracja z KSeF w każdym pakiecie". — [iFirma cennik](https://www.ifirma.pl/cennik)
- **Price** (net; the page cites effective dates "09/2026" and "03/2026"; the plan that includes the API or multi-currency is not stated) — [iFirma cennik](https://www.ifirma.pl/cennik):

| Item | Price |
|---|---|
| App | from 0 zł; "Faktura+" 12,50 zł/month paid annually |
| Accounting office, ryczałt VAT-exempt, 0–10 cost documents | 149 zł |
| Accounting office, Standard, 0–10 cost documents | 211 zł |
| Each further 5 cost documents | +33 zł |
| "Aplikacja + Uproszczona księgowość" | 149 zł/month |
| More than 70 cost documents | individual quote |
| FX differences (różnice kursowe) | 7 zł per item |

**Bridges and connectors**
- Stripto, S2K, Invopop and Fakturownia + Stripe: see Section 2.
- Explainer articles found but not fetched: [oneclick-workflow](https://oneclick-workflow.pl/stripe-a-ksef-faktury-z-platnosci-online/), [besidesfinance](https://www.besidesfinance.com/en/blog/ksef-i-stripe), [sprytnafirma](https://sprytnafirma.pl/podatki-i-ksiegowosc/stripe-faktura-ksef-2026-jak-wystawic).

**Other products: no Stripe integration found**
- **Faktura.pl:** no Stripe integration found (searches returned Fakturownia). KSeF module with automatic sending and receiving, from 18.99 zł/month. — [VATAX comparison](https://www.vatax.pl/blog/obsluga-ksef-porownanie-cen)
- **Saldeo SMART:** no Stripe integration or price found. Authorises in KSeF with one token holding maximum permissions, narrowed inside Saldeo. (Search extract.) — [Saldeo help attachment](https://pomoc.saldeosmart.pl/hc/pl/article_attachments/28755097763613)
- **Comarch ERP XT:** no Stripe integration found; described as invoicing, warehouse, simplified accounting and an online store. — [Comarch](https://comarch.com/erp/comarch-erp-xt)
- **Taxxo, InsERT (Subiekt/Rewizor nexo):** nothing found on Stripe.
- **KSeF costs on accountants' systems** (relevant to what the accountant pays) — [VATAX comparison, updated 15.10.2026; ERP prices as of 3.09.2025](https://www.vatax.pl/blog/obsluga-ksef-porownanie-cen):
  - Comarch ERP Optima KSeF: paid module.
  - Symfonia: "KSeF Plus" sold in operation packs.
  - InsERT GT: 1 InsPunkt per document sent or received.
  - enova365: KSeF included.
  - FakturaXL: free up to 10 invoices/month with KSeF.
  - Note: VATAX promotes its own Claro app.

### Inferences
- **Buy, turnkey:** Fakturownia's Stripe app is the only turnkey "Stripe payment → Polish document → KSeF → accountant" chain, renewals included. To keep document count low, configure it for invoices only to companies (receipts or nothing for consumers). Test in Stripe test mode how it handles refunds, PayPal-via-Stripe charges, coupons and EUR/USD conversion to PLN for KSeF.
- **Buy, add-on:** Stripto (19–79 zł) sitting on top of Fakturownia, inFakt or wFirma, or S2K (€8–50) working off Stripe invoices, covers the B2B KSeF requirement. Stripto's accounting-tool mode leaves documents where the accountant can reach them; S2K's does not.
- **Gaps no tool closes:** no product found reconciles Stripe payouts, Stripe fees or FX differences into an accountant-ready monthly summary. Fakturownia only imports a payments CSV for matching; iFirma's Stripe function is unclear. A small custom monthly report (the user's Go script) is still useful for B2C totals, fees, payouts and Patreon even if a tool handles B2B KSeF invoices.
- **inFakt and wFirma:** inFakt via Zapier works but adds a paid Zapier dependency, creates drafts, and has had silent failures. wFirma would need custom API work, which is close to "build".

### Gaps
- Fakturownia 2026 paid-plan prices not verified. inFakt's price is contradictory (4,99 vs 29/49/79 zł). wFirma's price is from a third party and net/gross is unknown.
- No vendor documents PayPal-via-Stripe.
- Fakturownia's Stripe app documents neither refunds/corrections, the document currency and conversion, multiple Stripe accounts, nor the plan required.
- iFirma's Stripe integration function is unclear.
- No independent user reviews of any Stripe integration found.
- Taxxo, InsERT, Saldeo and Faktura.pl were judged only by absence of evidence.
- Avalara, Fonoa, Storecove, Pagero and Edicom were not checked individually for a Stripe→KSeF path.
- Make templates were not checked.

## 4. Trade-offs for a B2C-heavy SaaS: a document for every Stripe payment vs. a monthly summary (zestawienie) without invoices

### Takeaway
KSeF does not push a B2C SaaS toward per-payment documents: consumer invoices are voluntary in KSeF, owed only on request, excluded from the 10k limit, and per iFirma not required for OSS sales at all. Because Polish accountants price in document bands (e.g. Warsaw KPiR 501 zł gross for 1–10 documents vs 1,760 zł for 90–100), the lean pattern is KSeF invoices only for business buyers and consumers who ask, plus a monthly summary of B2C Stripe sales, fees and payouts.

### Cited Findings
- B2C invoices are voluntary in KSeF. — [FAQ Q5](https://ksef.podatki.gov.pl/pytania-i-odpowiedzi-ksef-20/)
- A consumer invoice is owed only on request within 3 months. — [infor.pl (TPA), 23.03.2026](https://ksiegowosc.infor.pl/podatki/vat/faktura/7529636,faktury-dla-konsumentow-w-2026-r-czy-mozna-wystawiac-w-ksef-jak-przekazac-nabywcy.html)
- Consumer invoices do not count toward the PLN 10,000 monthly limit. — [MF: Poniżej 10 000 zł (Q9, Q11)](https://ksef.podatki.gov.pl/ponizej-10-000-zl/)
- iFirma (legal state 3 Jul 2025): OSS sales need not be invoiced even if the buyer asks. — [iFirma blog](https://www.ifirma.pl/blog/czym-jest-vat-oss-ile-wynosi-limit-oraz-jak-rozliczac-sprzedaz-na-terenie-ue/)
- **Accountant price bands by documents per month** — kb.pl, 7.01.2026, KPiR, gross per month, estimates from online offers; larger volumes priced individually — [kb.pl](https://kb.pl/porady/cennik-biur-rachunkowych-i-uslug-ksiegowych-w-calej-polsce/):

| City | 1–10 docs | 11–20 | 30–40 | 90–100 | 150–200 |
|---|---|---|---|---|---|
| Warszawa | 501 zł | 685 zł | 843 zł | 1,760 zł | 2,510 zł |
| Kraków | 483 zł | 639 zł | 799 zł | 1,710 zł | 2,460 zł |
| Lublin | 373 zł | 548 zł | 666 zł | 1,350 zł | 2,040 zł |

- cenauslug.pl (updated 27.05.2026, net per month) — [cenauslug.pl](https://cenauslug.pl/poradniki/ile-kosztuje-prowadzenie-ksiegowosci-porownanie-biur-rachunkowych-i-programow-online):
  - Ryczałt, up to 15 documents: 150–350 zł.
  - KPiR, not a VAT payer, up to 20 documents: 200–400 zł.
  - KPiR with VAT, up to 20 documents: 250–500 zł.
  - Price depends on document count, and "Ceny wzrosły... o 8-12% z powodu inflacji i wdrożenia KSeF" (prices rose 8–12% because of inflation and the KSeF roll-out).
- Snippet only, not fetched: self-service apps from 29 zł, online firms 80–250 zł, traditional firms 250–900 zł a month. — [freenance.io](https://freenance.io/ile-kosztuje/ile-kosztuje-ksiegowa-jdg-2026-biuro-rachunkowe-vs-samoobsluga/)
- iFirma's accounting office prices JDG packages by **cost** documents (149/211 zł for 0–10, +33 zł per 5 more, FX differences 7 zł per item). — [iFirma cennik](https://www.ifirma.pl/cennik)
- inFakt bookkeeping for a JDG from 200 PLN. — [inFakt cennik](https://www.infakt.pl/cennik)
- **Tool options that cut document count:**
  - Fakturownia: "invoice for companies, receipt for private persons" mode. — [Fakturownia help](https://pomoc.fakturownia.pl/integracja-fakturowni-ze-stripe-automatyczne-fakturowanie-platnosci)
  - Stripto can skip invoicing buyers without a NIP. — [Stripto](https://stripto.pl/stripe-integracja-ksef)
- OSS evidence: wFirma requires at least two non-contradictory pieces of evidence of the consumer's country for OSS sales. — [wFirma help](https://pomoc.wfirma.pl/-faktura-vat-oss-jak-wystawic)
- The accountant can be granted KSeF permissions as a firm and pull invoices directly. — [FAQ](https://ksef.podatki.gov.pl/pytania-i-odpowiedzi-ksef-20/)

### Inferences
- **Volume arithmetic:** at, say, 300 B2C renewals a month, per-payment invoices or receipts give 300+ documents. At kb.pl's Warsaw rates that is the "150–200+" band (≥ 2,510 zł gross a month), versus "1–10" or "11–20" (501–685 zł) for one monthly summary plus a few B2B KSeF invoices and cost documents.
  - Ask the accountant explicitly whether a tool-generated batch (JPK or Optima XML import, KSeF download) counts per document or as one entry.
  - iFirma-style pricing by cost documents would neutralise this.
- **Compliance trade-off:**
  - Per-payment documents give a per-transaction audit trail and automatic per-country VAT on OSS invoices, but create corrections work for refunds.
  - A monthly summary relies on Stripe's own data (exported by the user's script) and must still keep per-transaction evidence (customer country, OSS evidence, FX rate) available on request.
  - The legal basis for recording B2C e-service sales collectively, and the cash-register exemption for Stripe-collected payments, are covered by other researchers. The only source here (iFirma 2023) does not mention Stripe.
- **Recommended split for "buy vs build":**
  - Buy (or configure) a tool for the KSeF part, i.e. B2B invoices and consumer invoices on request, because that part is legally rigid and schema-heavy.
  - Build the monthly accountant summary yourself, because no tool found produces a Stripe payout/fee/FX/B2C-totals summary plus Patreon.
  - After KSeF, the accountant can fetch B2B sales invoices and Polish purchase invoices directly from KSeF, so the script needs to cover only what KSeF does not hold.

### Gaps
- No reputable source found specifically on documenting B2C Stripe sales with a periodic zestawienie or on the cash-register exemption for Stripe as a payment institution (deferred to the record-keeping researchers).
- No source gives a per-extra-document price, or says whether accountants count KSeF-imported or tool-exported documents individually.
- The volume figures above are illustrative; the user's actual monthly transaction count is unknown.
