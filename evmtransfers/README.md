# `evmtransfers`

This package is the retired legacy model for `evm_transfers`. It contains the
`Transfer` projection and its historical `StatusT` values.

No current sibling service imports the package, and the table is intentionally
excluded from live schema conformance. New EVM movement is represented by
`evmwalletoperations` and `evmwalletoperationeffects`; do not add new workflow
dependencies to `evmtransfers`.

`Transfer.Save` remains for source compatibility but has no public setter that
populates its internal update map. Treat the entire package as read/compatibility
surface pending deletion with the corresponding database retirement.

See the [root guide](../README.md) for the active EVM packages.
