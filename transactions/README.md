# `transactions`

This package models legacy custodial movements in `wallet_transactions`.
`Transaction` relates wallets, funding sources, and investments; the
`TransactionsTypeT` and `TransactionsStatusT` enums describe movement kind and
lifecycle.

`wallet-api` is the primary consumer. It creates deposits/withdrawals,
reconciles Dwolla webhooks, applies direct investment funding, exposes API DTOs,
and uses `Get`/`GetByID`. `email-worker` loads transaction context for mail.

## Mutation styles

Prefer an explicit update for new code:

```go
updated, err := models.Update[transactions.Transaction](
    ctx, tx,
    map[string]any{"id": id, "status": transactions.TransactionsStatusTPending},
    map[string]any{"status": transactions.TransactionsStatusTProcessed},
)
```

The legacy stateful path is `Get`/`GetByID`, then `SetStatus` or `SetEntityID`,
then `Save`. `Save` creates notification/history side effects after updating,
but currently passes hard-coded actor ID `1` and ignores the post-update error.
It is therefore unsuitable when accurate audit attribution or atomic side
effects are required.

Amounts are `float64` for compatibility. Financial code should validate and
convert at its boundary rather than accumulate binary floating-point values.

See the [root guide](../README.md) for transaction and compatibility guidance.
