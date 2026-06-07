package investments

import (
	"context"
	"fmt"

	"github.com/webdevelop-pro/go-common/db"
	"github.com/webdevelop-pro/i-models/profiles"
)

// MarkLegallyConfirmedForProfile moves confirmed investments to
// legally_confirmed when KYC/accreditation requirements are satisfied.
func MarkLegallyConfirmedForProfile(ctx context.Context, repo db.Repository, profile *profiles.Profile) ([]InvestmentInvestment, error) {
	if profile == nil {
		return nil, fmt.Errorf("profile is nil")
	}

	var (
		sql  string
		args []any
	)
	if profile.Type == profiles.Individual {
		sql = `
			UPDATE investment_investments as ii SET status=$1
			FROM investment_profiles AS ip
			WHERE ii.user_id=$2 and ii.status=$3
			AND ip.type in ('individual', 'sdira', 'solo401k')
			RETURNING ii.id,ii.status`
		args = []any{
			InvestmentTLegallyConfirmed, profile.UserID, InvestmentTConfirmed,
		}
	} else {
		sql = `
			UPDATE investment_investments SET status=$1 WHERE profile_id=$2 and status=$3
			RETURNING id,status`
		args = []any{
			InvestmentTLegallyConfirmed, profile.ID, InvestmentTConfirmed,
		}
	}

	rows, err := repo.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("update investments for profile %d: %w", profile.ID, err)
	}
	defer rows.Close()

	investments := []InvestmentInvestment{}
	for rows.Next() {
		var investment InvestmentInvestment
		if err := rows.Scan(&investment.ID, &investment.Status); err != nil {
			return investments, fmt.Errorf("scan investment for profile %d: %w", profile.ID, err)
		}
		investments = append(investments, investment)
	}
	if err := rows.Err(); err != nil {
		return investments, fmt.Errorf("iterate investments for profile %d: %w", profile.ID, err)
	}

	return investments, nil
}
