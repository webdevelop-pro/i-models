package users

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/global-torque/go-common/orm/v2/pgtype"
	"gopkg.in/yaml.v3"
)

func TestUserSerializationExcludesPassword(t *testing.T) {
	nullTime := pgtype.Timestamptz{Status: pgtype.Null}
	model := UserUser{
		ID:        7,
		Password:  "secret",
		Email:     "user@example.com",
		LastLogin: nullTime,
		CreatedAt: nullTime,
		UpdatedAt: nullTime,
	}

	got := model.ToJSON()
	if _, ok := got["password"]; ok {
		t.Fatalf("ToJSON returned password: %#v", got)
	}

	jsonPayload, err := json.Marshal(model)
	if err != nil {
		t.Fatalf("json marshal: %v", err)
	}
	if strings.Contains(string(jsonPayload), "password") || strings.Contains(string(jsonPayload), "secret") {
		t.Fatalf("json payload includes password: %s", jsonPayload)
	}

	yamlPayload, err := yaml.Marshal(model)
	if err != nil {
		t.Fatalf("yaml marshal: %v", err)
	}
	if strings.Contains(string(yamlPayload), "password") || strings.Contains(string(yamlPayload), "secret") {
		t.Fatalf("yaml payload includes password: %s", yamlPayload)
	}
}
