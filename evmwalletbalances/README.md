# `evmwalletbalances`

This package defines `WalletBalance` for `evm_wallet_balances`, a per-wallet,
per-chain asset balance projection. It carries asset metadata, raw/display
amounts, pricing data, and synchronization timestamps.

`evm-api` aliases this model in its `modelprojections` layer and reuses
`TableName`. Wallet endpoints retrieve cached balances with local ORM helpers
and refresh them after provider reconciliation.

```go
balances, err := models.RetrieveAll[evmwalletbalances.WalletBalance](
    ctx, repo, map[string]any{"wallet_id": walletID},
)
```

Balance refresh is a service workflow and often needs provider and transaction
logic; this package only supplies the row contract.

See the [root guide](../README.md) for projection conventions.
