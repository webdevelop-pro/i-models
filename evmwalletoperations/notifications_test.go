package evmwalletoperations

import "testing"

func TestIsZeroValueNativeETHOperation(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		asset        string
		assetAddress string
		amount       string
		want         bool
	}{
		{name: "native eth zero", asset: "ETH", amount: "0", want: true},
		{name: "native ethereum zero decimal", asset: "Ethereum", amount: "0.000000000000000000", want: true},
		{name: "native eth zero hex", asset: "ETH", amount: "0x0", want: true},
		{name: "native eth positive", asset: "ETH", amount: "0.0001"},
		{name: "erc20 zero is not native", asset: "ETH", assetAddress: "0x0000000000000000000000000000000000000001", amount: "0"},
		{name: "token zero is not eth", asset: "VSB", amount: "0"},
		{name: "empty amount is not zero", asset: "ETH"},
		{name: "invalid amount is not zero", asset: "ETH", amount: "not-a-number"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := IsZeroValueNativeETHOperation(tt.asset, tt.assetAddress, tt.amount); got != tt.want {
				t.Fatalf("IsZeroValueNativeETHOperation() = %t, want %t", got, tt.want)
			}
		})
	}
}
