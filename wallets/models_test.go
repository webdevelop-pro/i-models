package wallets

import "testing"

func TestWalletToMapUsesDatabaseKeys(t *testing.T) {
	userID := 12
	model := Wallet{
		ID:         7,
		UserID:     &userID,
		Balance:    10.5,
		IncBalance: 2.5,
		OutBalance: 1.25,
		Status:     WalletStatusTVerified,
		ObjectID:   "wallet-object",
	}

	got := model.ToMap()
	if got["id"] != 7 {
		t.Fatalf("unexpected id: %#v", got["id"])
	}
	if got["user_id"] != &userID {
		t.Fatalf("unexpected user_id: %#v", got["user_id"])
	}
	if got["balance"] != 10.5 {
		t.Fatalf("unexpected balance: %#v", got["balance"])
	}
	if got["status"] != WalletStatusTVerified {
		t.Fatalf("unexpected status: %#v", got["status"])
	}
	if _, ok := got["ID"]; ok {
		t.Fatalf("ToMap returned Go field key ID: %#v", got)
	}
}
