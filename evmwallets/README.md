# `evmwallets`

This package contains the shared `Wallet` projection for `evm_wallets` and the
`WalletStatusT` lifecycle enum. It represents the embedded/on-chain wallet,
which is distinct from the legacy custodial wallet in [`wallets`](../wallets/README.md).

`evm-api` is the verified consumer. Its `modelprojections.Wallet` owns a wider
provider/Turnkey projection, but reuses this package's table name and status
type so API validation and persistence agree on core states.

Simple consumers can use the base type with the generic helpers. For example:

```go
wallet, err := models.RetrieveOne[evmwallets.Wallet](
    ctx, repo, map[string]any{"user_id": userID},
)
```

`Wallet.Save` is compatibility-only because no public setter currently records
updated fields. Prefer explicit update maps or the owning service's projection.

See the [root guide](../README.md) for mutation guidance.
