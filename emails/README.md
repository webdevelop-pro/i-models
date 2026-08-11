# `emails`

This package models outbound messages in `email_emails`. `EmailEmail` contains
recipient, template/content, delivery-attempt, provider, status, and timestamp data;
`EmailStatusT` is the shared delivery lifecycle enum.

`Provider` is the minimal sending boundary:

```go
type Provider interface {
    Send(email *EmailEmail) error
    Cancel(email *EmailEmail) error
}
```

`email-worker` is the verified consumer. Its PostgreSQL adapter creates and
updates email rows around the worker's delivery flow. The model's injected
repository is compatibility state; the package itself does not send network
requests.

`DomainEventID` is present on the struct and `ToJSON`, but is not in the current
default `Fields()` projection. Code that must load that value should use a
service-specific projection or explicit query until the shared projection is
changed deliberately.

See the [root guide](../README.md) for generic create/retrieve examples.
