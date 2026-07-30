package investments

import (
	"encoding/json"
	"testing"
)

func TestInvestmentAmountJSONUsesExactStringAndPreservesNullDraft(t *testing.T) {
	t.Parallel()

	amount := "12345678901234567890.123456"
	model := InvestmentInvestment{
		Amount:         &amount,
		NumberOfShares: "98765432109876543210.123456789012345678",
	}

	encoded, err := json.Marshal(model)
	if err != nil {
		t.Fatalf("marshal investment: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(encoded, &payload); err != nil {
		t.Fatalf("unmarshal investment JSON: %v", err)
	}
	if got := payload["amount"]; got != amount {
		t.Fatalf("amount must be an exact JSON string: got %#v, want %q", got, amount)
	}
	if got := payload["number_of_shares"]; got != model.NumberOfShares {
		t.Fatalf("number_of_shares must be an exact JSON string: got %#v, want %q", got, model.NumberOfShares)
	}

	encoded, err = json.Marshal(InvestmentInvestment{})
	if err != nil {
		t.Fatalf("marshal draft investment: %v", err)
	}
	if err := json.Unmarshal(encoded, &payload); err != nil {
		t.Fatalf("unmarshal draft investment JSON: %v", err)
	}
	if got, ok := payload["amount"]; !ok || got != nil {
		t.Fatalf("draft amount must remain explicit JSON null: present=%v value=%#v", ok, got)
	}
}

func TestInvestmentToJSONPreservesExactProtocolAmounts(t *testing.T) {
	t.Parallel()

	amount := "10000000000000000000.000001"
	assetRaw := "10000000000000000000000001"
	shareRaw := "99999999999999999999999999999999999999"
	model := InvestmentInvestment{
		Amount:         &amount,
		AssetAmountRaw: &assetRaw,
		ShareAmountRaw: &shareRaw,
	}

	payload := model.ToJSON()
	for field, want := range map[string]string{
		"amount":           amount,
		"asset_amount_raw": assetRaw,
		"share_amount_raw": shareRaw,
	} {
		got, ok := payload[field].(*string)
		if !ok || got == nil || *got != want {
			t.Fatalf("%s must retain its exact string pointer: got %#v, want %q", field, payload[field], want)
		}
	}
}

func TestAllFundingTypesAreValid(t *testing.T) {
	t.Parallel()

	for _, fundingType := range AllFundingT() {
		if err := fundingType.IsValid(); err != nil {
			t.Fatalf("AllFundingT contains invalid value %q: %v", fundingType, err)
		}
	}
}
