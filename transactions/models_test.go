package transactions

import "testing"

func TestTransactionToMapUsesDatabaseKeys(t *testing.T) {
	sourceFundingID := 3
	destFundingID := 4
	entityID := "entity-1"
	model := Transaction{
		ID:              7,
		SourceWalletID:  ptr(8),
		DestWalletID:    ptr(9),
		SourceFundingID: &sourceFundingID,
		DestFundingID:   &destFundingID,
		EntityID:        &entityID,
		Type:            TransactionsTypeTInvestment,
		Amount:          10.5,
		Status:          TransactionsStatusTProcessed,
	}

	got := model.ToMap()
	if got["id"] != 7 {
		t.Fatalf("unexpected id: %#v", got["id"])
	}
	if got["source_wallet_id"] != model.SourceWalletID {
		t.Fatalf("unexpected source_wallet_id: %#v", got["source_wallet_id"])
	}
	if got["dest_wallet_id"] != model.DestWalletID {
		t.Fatalf("unexpected dest_wallet_id: %#v", got["dest_wallet_id"])
	}
	if got["type"] != TransactionsTypeTInvestment {
		t.Fatalf("unexpected type: %#v", got["type"])
	}
	if _, ok := got["COALESCE(source_wallet_id, 0) AS source_wallet_id"]; ok {
		t.Fatalf("ToMap returned select expression key: %#v", got)
	}
	if _, ok := got["SourceWalletID"]; ok {
		t.Fatalf("ToMap returned Go field key SourceWalletID: %#v", got)
	}
}

func ptr[T any](value T) *T { return &value }
