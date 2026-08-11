# `evmchainaccounts`

This package defines `EvmChainAccount` for `evm_chain_accounts`, plus the
canonical `ChainT`, `AccountModeT`, and `StatusT` enum families. The row maps an
owner/controller to a chain address and its lifecycle state.

`NewEvmChainAccount` injects a repository for compatibility, while the standard
generic helpers can retrieve or create the type directly.

No sibling service currently imports this package. `evm-api` maintains a local
`WalletChainAccount` projection and local domain enums for newer columns, and
both `evm-api` and `payment-api` join `evm_chain_accounts` directly in vault and
custody-settlement SQL. Changes here therefore still need to be checked against
those local projections and queries even though import search reports no direct
consumer.

See the [root guide](../README.md) for the difference between table use and a
direct package dependency.
