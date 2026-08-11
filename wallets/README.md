# `wallets`

This package models the legacy custodial/Dwolla wallet in `wallet_wallets`.
`Wallet` contains provider identifiers, available/incoming/outgoing balances,
content-type linkage, and `WalletStatusT`. It is distinct from the embedded
on-chain wallet in [`evmwallets`](../evmwallets/README.md).

`wallet-api` is the primary consumer for account creation, balance updates,
Dwolla webhooks, transaction limits, and API projections. `email-worker` loads
wallet context for notifications. [`fundingsources`](../fundingsources/README.md)
reuses the status enum.

`Get` and `GetByID` inject the repository so tracked setters followed by `Save`
can update the row. `Save` also attempts notification and Django history-log
side effects, but ignores errors returned by those hooks. Use explicit,
transaction-scoped writes when those side effects must be atomic.

Important legacy defect: `SetOutBalanceFN` and `SetBalanceFN` currently add the
delta to the in-memory `IncBalance` field even though their SQL expressions
update `out_balance` and `balance`. Do not rely on their in-memory result.

Balances are `float64` for compatibility. Convert from a validated money type
at the boundary and avoid repeated floating-point arithmetic.

See the [root guide](../README.md) for preferred mutation patterns.
