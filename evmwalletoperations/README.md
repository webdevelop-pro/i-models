# `evmwalletoperations`

This package defines the canonical base row for `evm_wallet_operations` and the
`OperationTypeT`, `OperationStatusT`, and `OperationSourceT` enum families. An
operation represents an on-chain intent and its submission/confirmation
lifecycle; service-local projections may add newer operational columns.

Verified consumers:

- `evm-api` reuses `TableName` and defines a wider local projection for wallet,
  withdrawal, exchange, reconciliation, and webhook workflows.
- `payment-api` scans the enum types from settlement SQL and gates settlement
  on confirmed platform operations.
- `email-worker` uses `IsZeroValueNativeETHOperation` to suppress misleading
  zero-value native-ETH notifications.

```go
if evmwalletoperations.IsZeroValueNativeETHOperation(asset, address, amount) {
    return nil // no user-facing transfer notification
}
```

The helper only classifies the asset/amount tuple; it does not validate a whole
operation or its chain state.

See the [root guide](../README.md) for service-local projection guidance.
