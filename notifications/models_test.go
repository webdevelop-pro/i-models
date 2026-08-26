package notifications

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"
)

func TestNovuPushStateIsPersistenceOnly(t *testing.T) {
	model := Notification{NovuPushAttemptCount: 3}

	encoded, err := json.Marshal(model)
	if err != nil {
		t.Fatalf("marshal notification: %v", err)
	}
	if len(encoded) == 0 || bytes.Contains(encoded, []byte("novu_push")) {
		t.Fatalf("Novu push state leaked into JSON: %s", encoded)
	}

	for _, field := range model.Fields() {
		if field == "novu_push_processed_at" || field == "novu_push_attempt_count" {
			t.Fatalf("Novu push state %q must not be a default ORM field", field)
		}
	}

	userID, ok := reflect.TypeOf(model).FieldByName("UserID")
	if !ok || userID.Type.Kind() != reflect.Int {
		t.Fatalf("UserID must match the NOT NULL integer schema, got %v", userID.Type)
	}
}
