package offers

import (
	"reflect"
	"testing"
)

func TestOfferFieldsIncludeVaultDeploymentSelectors(t *testing.T) {
	t.Parallel()

	fields := make(map[string]struct{})
	for _, field := range (OfferOffer{}).Fields() {
		fields[field] = struct{}{}
	}

	for _, field := range []string{"tokenization_engine", "fund_structure"} {
		if _, ok := fields[field]; !ok {
			t.Fatalf("OfferOffer.Fields must select %q for Vault deployment decisions", field)
		}
	}
}

func TestOfferModelExcludesDroppedTokenizationFramework(t *testing.T) {
	t.Parallel()

	if _, ok := reflect.TypeOf(OfferOffer{}).FieldByName("TokenizationFramework"); ok {
		t.Fatal("OfferOffer must not expose the dropped tokenization_framework column")
	}

	for _, field := range (OfferOffer{}).Fields() {
		if field == "tokenization_framework" {
			t.Fatal("OfferOffer.Fields must not select the dropped tokenization_framework column")
		}
	}
}
