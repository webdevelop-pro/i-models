# `userinvitations`

This package models invitations in `user_user_invitations`. `UserInvitation`
contains inviter/invitee, profile/investment relations, token and lifecycle
data. `Kind`, `Role`, and `Status` are validated enum-like string types with
`AllKinds`, `AllRoles`, and `AllStatuses` registries.

`email-worker` is the verified consumer. It uses invitation kind/role/status
and investment context to qualify domain events and build invitation email
workflows.

```go
invitation, err := models.RetrieveOne[userinvitations.UserInvitation](
    ctx, repo, map[string]any{"token": token},
)
```

Validate externally supplied enum values with `IsValid`; casting a string to
one of these types does not validate it.

See the [root guide](../README.md) for enum rules.
