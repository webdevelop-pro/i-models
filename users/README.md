# `users`

This package defines `UserUser` for `user_users` and `UserData` for structured
JSONB account metadata. It intentionally represents persistence identity, not
authentication/session policy.

`MergeDataByID` atomically merges an object into the user's JSONB `data`
column. `HasGroup` checks membership by group name using the supplied
repository.

Verified consumers are `evm-api`, `wallet-api`, `escrow-api`, `esign-api`,
`kyc-api`, and `email-worker`. Typical uses are local projection embedding,
user/profile lookup, event context, and JSONB provider-state updates.

```go
if err := users.MergeDataByID(ctx, repo, userID, patch); err != nil {
    return err
}
```

Do not put credentials or service secrets in `UserData`; API serialization and
database persistence are separate trust boundaries.

See the [root guide](../README.md) for serialization conventions.
