# `evmallowancesequences`

This package models `evm_erc20_allowance_sequences`, the durable coordination
row for ERC-20 allowance workflows. `AllowanceSequence` supplies the ORM model
contract and `PurposeT`/`AllPurposeT` define the supported sequence purposes.

Use the shared type for simple equality-based CRUD:

```go
sequence, err := models.RetrieveOne[evmallowancesequences.AllowanceSequence](
    ctx, repo, map[string]any{"id": sequenceID},
)
```

No sibling service currently imports this package directly. `evm-api` does use
the table extensively in vault deposit and strategy-return workflows, but it
owns a richer local `AllowanceSequence` domain type and transaction-specific
SQL. Those operations require ordering, locking, and atomic state transitions,
so the generic model is a schema reference rather than the workflow repository.

See the [root guide](../README.md) for generic CRUD and schema-change rules.
