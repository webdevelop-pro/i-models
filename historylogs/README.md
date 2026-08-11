# `historylogs`

This package provides `HistoryLog`, a compatibility projection for Django's
`django_admin_log`, and the shared action constants used to record additions,
changes, and deletions.

`kyc-api` is the verified direct consumer for manual KYC audit history. The
legacy [`wallets`](../wallets/README.md) and
[`transactions`](../transactions/README.md) packages also create these rows in
their post-update hooks.

```go
_, err := models.Create[historylogs.HistoryLog](ctx, repo, map[string]any{
    "user_id":        actorID,
    "action_flag":    historylogs.ActionChange,
    "change_message": message,
})
```

The row is an audit record, not application logging. Writes should identify the
real actor and must not silently replace it with a system/default user ID.

See the [root guide](../README.md) for audit and transaction boundaries.
