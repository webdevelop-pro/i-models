# `evmwalletoperationeffects`

This package models `evm_wallet_operation_effects`, the ledger-like effects
derived from EVM operations. `WalletOperationEffect` records the affected asset
and amount; `EffectKindT` and `EffectDirectionT` classify what happened and
whether value moved in or out.

No sibling service imports the Go package directly. The table is nevertheless
active: `evm-api` uses it throughout vault lifecycle, fulfillment, recovery,
and reconciliation SQL, while `payment-api` uses it in custody-settlement SQL
and integration tests. Those callers need multi-row joins and locking, so they
use repository SQL rather than generic model CRUD.

Use this type for simple projections, fixtures, and schema-aligned tooling; do
not assume that creating one effect alone performs the surrounding accounting
transition.

See the [root guide](../README.md) for transaction boundaries.
