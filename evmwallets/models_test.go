package evmwallets

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/webdevelop-pro/go-common/orm/pgtype"
	"gopkg.in/yaml.v3"
)

func TestWalletToJSONUsesDatabaseKeysAndExcludesPrivateKey(t *testing.T) {
	model := Wallet{
		ID:         7,
		UserID:     8,
		PublicKey:  "public",
		PrivateKey: "secret",
		Balance:    10.5,
		Status:     WalletStatusTVerified,
	}

	got := model.ToJSON()
	if got["id"] != 7 {
		t.Fatalf("unexpected id: %#v", got["id"])
	}
	if got["public_key"] != "public" {
		t.Fatalf("unexpected public_key: %#v", got["public_key"])
	}
	if _, ok := got["private_key"]; ok {
		t.Fatalf("ToJSON returned private_key: %#v", got)
	}
	if _, ok := got["PublicKey"]; ok {
		t.Fatalf("ToJSON returned Go field key PublicKey: %#v", got)
	}
}

func TestWalletSerializationExcludesPrivateKey(t *testing.T) {
	nullTime := pgtype.Timestamptz{Status: pgtype.Null}
	model := Wallet{
		PublicKey:  "public",
		PrivateKey: "secret",
		CreatedAt:  nullTime,
		UpdatedAt:  nullTime,
	}

	jsonPayload, err := json.Marshal(model)
	if err != nil {
		t.Fatalf("json marshal: %v", err)
	}
	if strings.Contains(string(jsonPayload), "private_key") || strings.Contains(string(jsonPayload), "secret") {
		t.Fatalf("json payload includes private key: %s", jsonPayload)
	}

	yamlPayload, err := yaml.Marshal(model)
	if err != nil {
		t.Fatalf("yaml marshal: %v", err)
	}
	if strings.Contains(string(yamlPayload), "private_key") || strings.Contains(string(yamlPayload), "secret") {
		t.Fatalf("yaml payload includes private key: %s", yamlPayload)
	}
}
