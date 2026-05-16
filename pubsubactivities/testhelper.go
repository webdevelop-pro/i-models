package pubsubactivities

import (
	"context"
	"os"

	"github.com/webdevelop-pro/go-common/configurator"
	"github.com/webdevelop-pro/go-common/db"
	"github.com/webdevelop-pro/go-common/tests"
)

// cleanFixtures is a tests.FixturesManager that wipes pubsub_activities.
type cleanFixtures struct {
	db *db.DB
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
func NewCleanFixtures(ctx context.Context) tests.FixturesManager {
	_ = os.Setenv("TZ", "UTC")

	cfg := db.Config{}
	if err := configurator.NewConfiguration(&cfg, "DB"); err != nil {
		panic(err)
	}

	return cleanFixtures{db: db.New(ctx)}
}

func (f cleanFixtures) CleanAndApply() error {
	_, err := f.db.Exec(context.Background(), "DELETE FROM pubsub_activities")
	return err
}

func (f cleanFixtures) SetCTX(ctx context.Context) context.Context {
	return ctx
}
