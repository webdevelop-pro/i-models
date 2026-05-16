package pubsubactivities

import (
	"github.com/webdevelop-pro/go-common/db/dbtests"
	"github.com/webdevelop-pro/go-common/tests"
)

// ClearAction returns a test action that wipes the pubsub_activities table.
//
// Use it as the FIRST action in every pull-mode worker test scenario. Such
// tests publish through the Pub/Sub emulator and recreate the topic per
// RunTableTest (qtests); the emulator restarts message IDs at 1 on topic
// recreation, so a row left by an earlier test would make the message-level
// deduper wrongly skip a later test's event. Real Pub/Sub never reuses
// message IDs, so this is purely an emulator-only concern.
//
// It relies on the dbtests-injected pool (the scenario's FixturesManager list
// must include dbtests.NewFixturesManager), the convention these suites
// already follow.
func ClearAction() tests.SomeAction {
	return dbtests.RawSQL("DELETE FROM pubsub_activities")
}
