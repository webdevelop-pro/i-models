package pubsubactivities

import "context"

// DomainReactionStore is the application-facing persistence contract for one
// service's reaction to immutable domain events. Implementations bind the
// service identity and persistence repository so application handlers only
// coordinate planning, command execution, and the resulting lifecycle state.
//
// A handler must claim an actionable event before executing its command and
// must finalize the returned fencing token as reacted, rejected, or failed.
// Transport and application rejections are recorded before a command begins
// and therefore do not require a claim token.
type DomainReactionStore interface {
	ClaimDomainReaction(
		ctx context.Context,
		eventID string,
		topic string,
		messageID string,
		attempt int,
	) (claimToken string, claimed bool, err error)
	ResolveDomainReaction(ctx context.Context, eventID string) (resolved bool, err error)
	MarkDomainReactionReacted(ctx context.Context, eventID string, claimToken string) error
	MarkDomainReactionRejected(ctx context.Context, eventID string, claimToken string, cause error) error
	MarkDomainReactionFailed(ctx context.Context, eventID string, claimToken string, cause error) error
	RecordTransportRejection(
		ctx context.Context,
		messageID string,
		topic string,
		attempt int,
		cause error,
	) error
	RecordApplicationRejection(
		ctx context.Context,
		eventID string,
		topic string,
		attempt int,
		cause error,
	) error
}
