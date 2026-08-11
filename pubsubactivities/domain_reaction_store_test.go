package pubsubactivities

import "context"

// This compile-time assertion pins the public application-facing contract.
var _ DomainReactionStore = (*domainReactionStoreContract)(nil)

type domainReactionStoreContract struct{}

func (*domainReactionStoreContract) ClaimDomainReaction(
	context.Context,
	string,
	string,
	string,
	int,
) (string, bool, error) {
	return "", false, nil
}

func (*domainReactionStoreContract) ResolveDomainReaction(context.Context, string) (bool, error) {
	return false, nil
}

func (*domainReactionStoreContract) MarkDomainReactionReacted(context.Context, string, string) error {
	return nil
}

func (*domainReactionStoreContract) MarkDomainReactionRejected(
	context.Context,
	string,
	string,
	error,
) error {
	return nil
}

func (*domainReactionStoreContract) MarkDomainReactionFailed(
	context.Context,
	string,
	string,
	error,
) error {
	return nil
}

func (*domainReactionStoreContract) RecordTransportRejection(
	context.Context,
	string,
	string,
	int,
	error,
) error {
	return nil
}

func (*domainReactionStoreContract) RecordApplicationRejection(
	context.Context,
	string,
	string,
	int,
	error,
) error {
	return nil
}
