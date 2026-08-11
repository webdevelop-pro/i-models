# `notifications`

This package models `notification_notifications`. `Notification` stores the
target user, content-type/object relation, message data, read state, and
timestamps. `NotificationStatusT` and `NotificationTypeT` define the canonical
status and delivery category labels.

Verified consumers:

- `evm-api` creates internal notifications during wallet workflows;
- `esign-api` creates notification records around signature flows;
- `email-worker` retrieves notification context and coordinates email output.

```go
notification, err := models.Create[notifications.Notification](ctx, repo,
    map[string]any{
        "user_id": userID,
        "type": notifications.NotificationTypeTInternal,
        "status": notifications.NotificationStatusTUnread,
    },
)
```

Creating the row does not itself deliver email or push traffic.

See the [root guide](../README.md) for persistence rules.
