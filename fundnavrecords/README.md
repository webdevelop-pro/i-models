# `fundnavrecords`

This package contains `FundNAVRecord`, the projection for `fund_nav_records`.
It represents a fund's dated net-asset-value observation and implements the
standard model contract for generic reads and writes.

```go
records, err := models.RetrieveAll[fundnavrecords.FundNAVRecord](
    ctx, repo, map[string]any{"offer_id": offerID},
)
```

No direct package import was found in the checked sibling services. `evm-api`
does actively join and update `fund_nav_records` through its vault strategy,
fulfillment, economics, and API repository SQL. The shared type is therefore a
base/schema projection, not that workflow's repository abstraction. Ordering
by valuation time must be added through the lower-level ORM or SQL because the
compatibility facade accepts equality filters only.

See the [root guide](../README.md) for advanced-query guidance.
