# `distributions`

This package defines three related table projections:

- `DistributionDistribution` for `distribution_distributions`;
- `DistributionDistributionFiler` for `distribution_distribution_filers`;
- `DistributionDistributionReport` for `distribution_distribution_reports`.

It also owns distribution/report classifications and the nullable
`DistributionFilerT` SQL/JSON wrapper used by filer links.

```go
rows, err := models.RetrieveAll[distributions.DistributionDistribution](
    ctx, repo, map[string]any{"offer_id": offerID},
)
```

No direct import was found in the checked sibling services. Keep this package
aligned with migrations and schema conformance; verify prospective consumers
before extending its API.

See the [root guide](../README.md) for model conventions.
