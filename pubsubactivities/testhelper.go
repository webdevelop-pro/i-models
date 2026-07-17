package pubsubactivities

import (
	"context"
	"os"

	"github.com/global-torque/go-common/db/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

// FixturesManager is the structural fixture contract used by both legacy and
// v2 test runners. Keeping it local avoids importing a test-framework module
// into the released models module.
type FixturesManager interface {
	CleanAndApply() error
	SetCTX(context.Context) context.Context
}

// cleanFixtures is a tests.FixturesManager that wipes pubsub_activities.
type cleanFixtures struct {
	db *pgxpool.Pool
}

// NewCleanFixtures returns a tests.FixturesManager that deletes every
// pubsub_activities row once per RunTableTest (its CleanAndApply runs
// alongside the dbtests/qtests managers, before any TestAction).
//
// Why this is needed for pull-mode worker tests: qtests recreates the
// Pub/Sub topic per test function and the local emulator restarts message
// IDs at 1, so a (msg_id, service) row left by an earlier test would make
// the message-level deduper wrongly skip a later test's events. Real Pub/Sub
// never reuses message IDs, so this is purely an emulator-test concern.
//
// Add it to the RunTableTest []tests.FixturesManager list, e.g.:
//
//	[]tests.FixturesManager{
//	    dbtests.NewFixturesManager(ctx, ...),
//	    qtests.NewFixturesManager(ctx, ...),
//	    pubsubactivities.NewCleanFixtures(ctx),
//	}
func NewCleanFixtures(ctx context.Context) FixturesManager {
	_ = os.Setenv("TZ", "UTC")

	pool, err := db.NewPool(ctx)
	if err != nil {
		panic(err)
	}

	return cleanFixtures{db: pool}
}

func (f cleanFixtures) CleanAndApply() error {
	_, err := f.db.Exec(context.Background(), "DELETE FROM pubsub_activities")
	return err
}

func (f cleanFixtures) SetCTX(ctx context.Context) context.Context {
	return ctx
}
