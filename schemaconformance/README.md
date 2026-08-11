# `schemaconformance`

This is a test-only package that guards the contract between model projections
and PostgreSQL. It contains the canonical registry of model tables and enum
families used by the repository's conformance tests.

The normal tests verify that `Fields()` contains valid, unique projections and
that enum registries remain internally consistent. The optional live test
compares those expectations with a configured PostgreSQL schema:

```sh
I_MODELS_SCHEMA_TEST=1 go test ./schemaconformance
```

No sibling service imports this package at runtime. Add new persistent models
and enum families to the registry when they become part of the supported
schema. The retired `evmtransfers` model is intentionally excluded.

See the [root guide](../README.md) for the complete validation workflow.
