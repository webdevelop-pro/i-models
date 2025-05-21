package evmwallets

import "github.com/pkg/errors"

type WalletStatusT string

// Enum values for WalletStatusT
const (
	WalletStatusTCreated  WalletStatusT = "created"
	WalletStatusTPending  WalletStatusT = "pending"
	WalletStatusTVerified WalletStatusT = "verified"
	WalletStatusTError    WalletStatusT = "error"

	AppLabel  = "evm"
	ModelName = "wallet"
	TableName = "evm_wallets"

	pkgName = "models/emvwallets"
)

func AllWalletStatusT() []WalletStatusT {
	return []WalletStatusT{
		WalletStatusTCreated,
		WalletStatusTPending,
		WalletStatusTVerified,
		WalletStatusTError,
	}
}

func (e WalletStatusT) IsValid() error {
	switch e {
	case WalletStatusTCreated, WalletStatusTPending, WalletStatusTVerified, WalletStatusTError:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e WalletStatusT) String() string {
	return string(e)
}
