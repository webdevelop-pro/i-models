package evmwalletoperations

import (
	"math/big"
	"strings"
)

// IsZeroValueNativeETHOperation reports whether an operation payload describes
// a native ETH movement with no value. These are provider trace artifacts and
// should not create end-user notifications.
func IsZeroValueNativeETHOperation(asset, assetAddress, amount string) bool {
	if strings.TrimSpace(assetAddress) != "" {
		return false
	}
	if !isNativeETHAsset(asset) {
		return false
	}
	return isZeroAmountLiteral(amount)
}

func isNativeETHAsset(asset string) bool {
	switch strings.ToUpper(strings.TrimSpace(asset)) {
	case "ETH", "ETHEREUM":
		return true
	default:
		return false
	}
}

func isZeroAmountLiteral(amount string) bool {
	amount = strings.TrimSpace(amount)
	if amount == "" {
		return false
	}

	if strings.HasPrefix(amount, "0x") || strings.HasPrefix(amount, "0X") {
		amount = amount[2:]
		if amount == "" {
			return true
		}
		parsed, ok := new(big.Int).SetString(amount, 16)
		return ok && parsed.Sign() == 0
	}

	parsed, ok := new(big.Rat).SetString(amount)
	return ok && parsed.Sign() == 0
}
