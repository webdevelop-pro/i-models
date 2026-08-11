# `pgtype`

This package preserves the historical `i-models/pgtype` import path while
delegating timestamp semantics to `go-common/orm/v2/pgtype`.

It aliases `Status` and `InfinityModifier`, re-exports their constants, and
provides `Timestamptz` with pgx scanning, `database/sql` value support, JSON
encoding, finite timestamp handling, and PostgreSQL positive/negative infinity.

```go
ts := pgtype.Timestamptz{
    Time:   time.Now().UTC(),
    Status: pgtype.Present,
}
```

`wallet-api` imports this path in DTO/validator projections. Many structs in
this module use the type as part of their shared database contract. New code
outside that compatibility boundary may import the canonical go-common type
directly.

See the [root guide](../README.md) for serialization and projection rules.
