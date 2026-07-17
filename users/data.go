package users

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/global-torque/go-common/db/v2"
	"github.com/global-torque/go-common/orm/v2"
)

// MergeDataByID atomically merges a JSON object into user_users.data.
func MergeDataByID(ctx context.Context, repo db.Repository, userID int, data any) error {
	dataJSON, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("marshal user data for user %d: %w", userID, err)
	}

	const sql = `UPDATE user_users
		SET data=COALESCE(data, '{}'::jsonb) || $1::jsonb
		WHERE id=$2`
	result, err := repo.Exec(ctx, sql, string(dataJSON), userID)
	if err != nil {
		return fmt.Errorf("merge data for user %d: %w", userID, err)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("merge data for user %d: %w", userID, orm.ErrNoRowsAffected)
	}

	return nil
}
