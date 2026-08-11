# `offers`

This package is the shared offering schema:

- `OfferOffer` maps `offer_offers`;
- `OfferOfferFiler` maps `offer_offer_filers`;
- `OfferComment` maps `offer_comments`.

It also defines status, security, regulation, jurisdiction, tokenization, fund
structure, filer, and comment enums. Nullable generated-style wrappers provide
SQL `Valuer`/`Scanner` and JSON/text encoding for optional enum columns.

`Get` retrieves one `OfferOffer` by equality filters. Generic helpers can be
used for the other model types.

Verified consumers include `evm-api` (token/contract and exchange rules),
`wallet-api` (investment validation), `payment-api` (settlement eligibility),
`escrow-api` (offer administration), `esign-api` (document validation), and
`email-worker` (event context). Several consumers embed `OfferOffer` into a
narrower/wider DTO and query it through their own ORM layer.

Treat enum additions as database migrations and cross-service contract changes.

See the [root guide](../README.md) for enum validation guidance.
