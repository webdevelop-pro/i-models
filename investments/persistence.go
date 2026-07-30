package investments

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/global-torque/go-common/db/v2"
	"github.com/global-torque/go-common/orm/v2"
)

// RetrieveInvestmentForUpdate loads and row-locks one investment. Callers must
// use a transaction-scoped repository so the lock spans their recheck and
// conditional write.
func RetrieveInvestmentForUpdate(
	ctx context.Context,
	repo db.Repository,
	investmentID int,
) (*InvestmentInvestment, error) {
	return orm.RetrieveOne[InvestmentInvestment, *InvestmentInvestment](
		ctx,
		repo,
		sq.Eq{"id": investmentID},
		sq.Expr("FOR UPDATE"),
	)
}

// UpdateUnlockedInvestment conditionally changes one business row only while
// its expected status and deposit request-path lock still match. Additional
// predicates carry service-owned custody/readiness checks. The boolean is true
// only when exactly one row was affected and must be checked by every caller.
func UpdateUnlockedInvestment(
	ctx context.Context,
	repo db.Repository,
	investmentID int,
	expectedStatus InvestmentT,
	updates map[string]any,
	additionalPredicates ...sq.Sqlizer,
) (bool, error) {
	predicates := append(
		[]sq.Sqlizer{sq.Expr("vault_request_locked_at IS NULL")},
		additionalPredicates...,
	)
	return orm.Update[InvestmentInvestment, *InvestmentInvestment](
		ctx,
		repo,
		map[string]any{
			"id":     investmentID,
			"status": expectedStatus,
		},
		updates,
		predicates...,
	)
}

// RetrieveRedemptionForUpdate loads and row-locks one redemption. Callers must
// use a transaction-scoped repository so preparation and cancellation share
// the same business-row lock.
func RetrieveRedemptionForUpdate(
	ctx context.Context,
	repo db.Repository,
	redemptionID int64,
) (*InvestmentRedemption, error) {
	return orm.RetrieveOne[InvestmentRedemption, *InvestmentRedemption](
		ctx,
		repo,
		sq.Eq{"id": redemptionID},
		sq.Expr("FOR UPDATE"),
	)
}

// UpdateUnlockedRedemption conditionally changes one business row only while
// its expected status and request-path lock still match. Additional predicates
// carry service-owned operation-fence checks. The boolean is true only when
// exactly one row was affected and must be checked by every caller.
func UpdateUnlockedRedemption(
	ctx context.Context,
	repo db.Repository,
	redemptionID int64,
	expectedStatus RedemptionStatusT,
	updates map[string]any,
	additionalPredicates ...sq.Sqlizer,
) (bool, error) {
	predicates := append(
		[]sq.Sqlizer{sq.Expr("request_locked_at IS NULL")},
		additionalPredicates...,
	)
	return orm.Update[InvestmentRedemption, *InvestmentRedemption](
		ctx,
		repo,
		map[string]any{
			"id":     redemptionID,
			"status": expectedStatus,
		},
		updates,
		predicates...,
	)
}
