# `evmcontracts`

This package models deployed EVM contracts in `evm_contracts`. `Contract`
contains the shared projection, `StatusT` describes deployment lifecycle, and
`DeploymentLegT` identifies deployment legs. Constants also expose the table
names for deployment operations, transaction history, and signer nonce
reservations.

`evm-api` is the verified consumer. Its wallet and exchange paths use these
table-name constants in transaction-specific SQL joined to `offer_offers`; the
shared status values also keep deployment state aligned across projections.

```go
contract, err := models.RetrieveOne[evmcontracts.Contract](
    ctx, repo, map[string]any{"id": contractID},
)
```

`Contract.Save` is compatibility-only: the package currently exposes no public
setter that fills its private update map. Prefer explicit `models.Update` or
transaction-scoped SQL for writes.

See the [root guide](../README.md) for mutation guidance.
