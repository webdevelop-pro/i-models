# `pubsubactivities`

This package implements durable Pub/Sub delivery deduplication and fenced
domain-reaction processing in PostgreSQL. Unlike most `i-models` packages, it
is an active coordination repository rather than only a row projection.

`Status` supports processing, processed, failed, reacted, rejected, and
resolved states. A claim returns a UUID fencing token; terminal writes require
that token so an expired worker cannot overwrite a newer claim.

## Correct layer

The application event handler owns the reaction lifecycle. It decides whether
an event is actionable, claims before executing business work, and records the
outcome. It should depend on `DomainReactionStore`, not on a PostgreSQL pool or
`TransactionalRepository`.

```text
Pub/Sub port
    │ decode envelope and transport metadata
    ▼
Application event handler
    │ plan → claim → execute → finalize
    ▼
DomainReactionStore
    ▼
PostgreSQL adapter
    │ binds service name and repository
    ▼
pubsubactivities SQL functions
```

The domain layer must not receive message IDs, topics, delivery attempts,
claim tokens, repositories, or `pubsubactivities` types.

## Application-facing interface

`DomainReactionStore` is the complete persistence port for one service's
domain-event reaction lifecycle:

```go
type Handler struct {
    reactions pubsubactivities.DomainReactionStore
}

func (handler *Handler) Handle(
    ctx context.Context,
    event DomainEvent,
    delivery DeliveryMetadata,
) error {
    plan, err := handler.plan(ctx, event) // no business side effects
    if err != nil {
        if !isPermanent(err) {
            return err // transient planning failure: NACK and retry
        }
        if recordErr := handler.reactions.RecordApplicationRejection(
            ctx, event.ID, delivery.Topic, delivery.Attempt, err,
        ); recordErr != nil {
            return fmt.Errorf("record application rejection: %w", recordErr)
        }
        return err // the transport policy recognizes this as permanent
    }

    if !plan.Actionable() {
        _, err = handler.reactions.ResolveDomainReaction(ctx, event.ID)
        return err
    }

    token, claimed, err := handler.reactions.ClaimDomainReaction(
        ctx, event.ID, delivery.Topic, delivery.MessageID, delivery.Attempt,
    )
    if err != nil || !claimed {
        return err
    }

    if err = handler.execute(ctx, plan); err != nil {
        if isPermanent(err) {
            if markErr := handler.reactions.MarkDomainReactionRejected(
                ctx, event.ID, token, err,
            ); markErr != nil {
                return fmt.Errorf("mark rejected: %w", markErr)
            }
            return err
        }

        if markErr := handler.reactions.MarkDomainReactionFailed(
            ctx, event.ID, token, err,
        ); markErr != nil {
            return errors.Join(err, markErr)
        }
        return err // preserve the retryable error so the delivery is NACKed
    }

    return handler.reactions.MarkDomainReactionReacted(ctx, event.ID, token)
}
```

The transport maps successful, duplicate, ignored, and permanent-rejection
outcomes to ACK, and retryable errors or `ErrInProgress` to NACK. Transport
decoding failures are recorded by the Pub/Sub port with
`RecordTransportRejection`; they never enter the business handler.

Because Go interfaces are structural, a use case that needs only part of the
lifecycle may declare a narrower local interface. Adapters implementing the
full shared contract satisfy both.

## PostgreSQL adapter

The adapter owns the concrete repository and immutable service identifier. It
implements `DomainReactionStore` by delegating to this package's functions:

| Interface method | Package function |
|---|---|
| `ClaimDomainReaction` | `ClaimDomainReaction(ctx, repo, eventID, service, topic, messageID, attempt)` |
| `ResolveDomainReaction` | `ResolveDomainReaction(ctx, repo, eventID, service)` |
| `MarkDomainReactionReacted` | `MarkReactedLease(ctx, repo, eventID, service, token)` |
| `MarkDomainReactionRejected` | `MarkRejectedLease(ctx, repo, eventID, service, token, cause.Error())` |
| `MarkDomainReactionFailed` | `MarkFailedLease(ctx, repo, eventID, service, token, cause.Error())` |
| `RecordTransportRejection` | `RecordTransportRejection(ctx, repo, messageID, service, topic, attempt, cause.Error())` |
| `RecordApplicationRejection` | `RecordApplicationRejection(ctx, repo, eventID, service, topic, attempt, cause.Error())` |

Only the adapter should need `TransactionalRepository`. `ClaimDomainReaction`
uses it to atomically claim the event and record the concrete Pub/Sub delivery
before business execution begins.

Adapters should convert a nil `cause` safely rather than calling `Error` on it.
Terminal persistence should use a short context detached from request
cancellation. `MarkReactedLease` and `MarkRejectedLease` already do this;
adapters must provide equivalent protection around `MarkFailedLease`.

## Required lifecycle

1. The Pub/Sub port decodes and validates the transport envelope. A malformed
   envelope is recorded with `RecordTransportRejection`.
2. The application handler performs side-effect-free planning. A permanent
   application-contract failure is recorded with
   `RecordApplicationRejection`.
3. A non-actionable event calls `ResolveDomainReaction`. This only resolves
   prior failed/expired delivery evidence; the first ignored delivery creates
   no activity row.
4. An actionable event calls `ClaimDomainReaction` before executing its
   business command.
5. `claimed == false` means the event already reached a terminal state and is
   a duplicate. Do not execute it.
6. Success calls `MarkDomainReactionReacted`; permanent failure calls
   `MarkDomainReactionRejected`; retryable failure calls
   `MarkDomainReactionFailed`.

`ErrInProgress` should result in a NACK/retry. `ErrClaimLost` means the worker's
lease was reclaimed; it must not finalize or repeat unprotected side effects.
`ErrDeliveryIdentityConflict` indicates the same service/message identity was
attached to a different immutable event and must fail closed.

## Generic message deduplication

Non-domain messages can use the lower-level fenced lease API directly at the
transport/application boundary:

```go
token, claimed, err := pubsubactivities.ClaimLease(
    ctx, repo, messageID, service, topic, attempt,
)
if err != nil || !claimed {
    return err
}

// Execute the handler.
return pubsubactivities.MarkProcessedLease(ctx, repo, messageID, service, token)
```

Use `MarkFailedLease` on retryable failure. `RecordDomainDelivery`,
`RecordFailedDelivery`, and unfenced `Claim`/`Mark*` functions remain only for
compatibility; new domain-event consumers must use the fenced reaction API.

The package is used by event consumers in `evm-api`, `wallet-api`,
`payment-api`, `escrow-api`, `kyc-api`, `email-worker`, and `filer-api`.

See the [root guide](../README.md) for transaction rules.
