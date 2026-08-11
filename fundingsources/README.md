# `fundingsources`

This package models linked bank/funding sources in `wallet_funding_source`.
`WalletFundingSource` includes the provider entity ID, bank display data, type,
status, and optional legacy wallet relation. `FoundingSourceT` is the
historically named funding-source enum; renaming it would be a public API break.

Verified consumers:

- `wallet-api` embeds the model in DTOs, creates/deletes rows during bank-link
  flows, and uses it in Dwolla/Plaid adapters;
- `email-worker` retrieves funding-source context for notification templates.

The `Status` field deliberately uses `wallets.WalletStatusT`, so changes to
legacy wallet states can affect this package too.

See the [root guide](../README.md) for CRUD examples.
