# `models`

This package is the small compatibility CRUD facade over
`github.com/global-torque/go-common/orm/v2`. It aliases `db.Repository`, adapts
map filters to equality predicates, and injects the repository into returned
models that implement `SetDB`.

## API

`RetrieveOne`, `RetrieveAll`, `Create`, `Update`, `Exists`, and `Delete` operate
on any model whose pointer implements `SetID`, `Fields`, and `Table`.
`DefaultFields` derives a projection from struct tags. `RetriveOne` and
`RetriveAll` are deprecated misspellings retained for existing services.

```go
wallet, err := models.RetrieveOne[wallets.Wallet](
    ctx, repo, map[string]any{"id": walletID},
)

updated, err := models.Update[wallets.Wallet](
    ctx,
    repo,
    map[string]any{"id": walletID},
    map[string]any{"status": wallets.WalletStatusTVerified},
)
```

The filter is always equality-based. Use `go-common/orm/v2`, Squirrel, or SQL
for joins, ordering, ranges, locking clauses, and other predicates.

## Consumers

`wallet-api` uses these helpers extensively with both shared models and local
validator projections. `email-worker` imports the package in its PostgreSQL
adapter. Newer services commonly use `go-common/orm` directly while retaining
the model types from this module.

See the [root guide](../README.md) for the complete model contract.
