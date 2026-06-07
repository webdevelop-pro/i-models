package profiles

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/webdevelop-pro/go-common/db"
)

const kycTimestampLayout = "2006-01-02 15:04:05-07:00"

// AppendKYCByID updates the profile KYC fields and prepends a KYC audit entry.
func AppendKYCByID(ctx context.Context, repo db.Repository, profile *Profile, data *ProfileData) error {
	if profile == nil {
		return fmt.Errorf("profile is nil")
	}
	if profile.KycID == nil {
		return fmt.Errorf("profile %d has nil kyc id", profile.ID)
	}

	now := time.Now()
	completedAt := "-infinity"
	if profile.KycStatus == KycApproved {
		completedAt = now.Format(kycTimestampLayout)
	}

	kycEvent, err := json.Marshal(map[string]string{
		"id":           *profile.KycID,
		"status":       string(profile.KycStatus),
		"created_at":   now.Format(kycTimestampLayout),
		"completed_at": completedAt,
	})
	if err != nil {
		return fmt.Errorf("marshal kyc event for profile %d: %w", profile.ID, err)
	}

	args := []any{
		profile.KycStatus,
		profile.KycID,
		string(kycEvent),
		completedAt,
	}
	sql := `UPDATE investment_profiles SET
		kyc_status=$1,
		kyc_id=$2,
		kyc_data=jsonb_insert(COALESCE(kyc_data, '[]'::jsonb), '{0}', $3::jsonb),
		kyc_at=$4`

	if data != nil {
		dataJSON, err := json.Marshal(data)
		if err != nil {
			return fmt.Errorf("marshal profile data for profile %d: %w", profile.ID, err)
		}
		args = append(args, string(dataJSON))
		sql += fmt.Sprintf(", data=COALESCE(data, '{}'::jsonb) || $%d::jsonb", len(args))
	}

	args = append(args, profile.ID)
	sql += fmt.Sprintf(" WHERE id=$%d", len(args))

	_, err = repo.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("update kyc for profile %d: %w", profile.ID, err)
	}
	return nil
}

// MergeDataByID merges JSON profile data into investment_profiles.data.
func MergeDataByID(ctx context.Context, repo db.Repository, profileID int, data any) error {
	dataJSON, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("marshal profile data for profile %d: %w", profileID, err)
	}

	const sql = `UPDATE investment_profiles
		SET data=COALESCE(data, '{}'::jsonb) || $1::jsonb
		WHERE id=$2`
	if _, err := repo.Exec(ctx, sql, string(dataJSON), profileID); err != nil {
		return fmt.Errorf("merge data for profile %d: %w", profileID, err)
	}
	return nil
}

// UpdateRelatedKYCByProfile propagates an individual profile's KYC status to
// related IRA profiles, or propagates entity/trust status to child profiles.
func UpdateRelatedKYCByProfile(ctx context.Context, repo db.Repository, profile *Profile) ([]Profile, error) {
	if profile == nil {
		return nil, fmt.Errorf("profile is nil")
	}

	completedAt := "-infinity"
	if profile.KycStatus == KycApproved {
		completedAt = time.Now().Format(kycTimestampLayout)
	}

	args := []any{profile.KycStatus, profile.KycID, completedAt}
	sql := `UPDATE investment_profiles SET
		kyc_status=$1,
		kyc_id=CASE WHEN "type"='individual' THEN $2 ELSE null END,
		kyc_at=$3`

	if profile.Type == Individual {
		args = append(args, profile.UserID)
		sql += ` WHERE user_id=$4 and ("type"='individual' or "type"='solo401k' or "type"='sdira')`
	} else {
		args = append(args, profile.ID)
		sql += " WHERE profile_id=$4"
	}
	sql += " RETURNING id"

	rows, err := repo.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("update related kyc for profile %d: %w", profile.ID, err)
	}
	defer rows.Close()

	related := []Profile{}
	for rows.Next() {
		var prof Profile
		if err := rows.Scan(&prof.ID); err != nil {
			return related, fmt.Errorf("scan related profile for %d: %w", profile.ID, err)
		}
		related = append(related, prof)
	}
	if err := rows.Err(); err != nil {
		return related, fmt.Errorf("iterate related profiles for %d: %w", profile.ID, err)
	}

	return related, nil
}
