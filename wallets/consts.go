package wallets

import "github.com/pkg/errors"

type WalletStatusT string

// Enum values for WalletStatusT
const (
	WalletStatusTCreated  WalletStatusT = "created"
	WalletStatusTPending  WalletStatusT = "pending"
	WalletStatusTVerified WalletStatusT = "verified"
	WalletStatusTError    WalletStatusT = "error"

	WalletAppLabel  = "wallet"
	WalletModelName = "wallet"
)

func AllWalletStatusT() []WalletStatusT {
	return []WalletStatusT{
		WalletStatusTCreated,
		WalletStatusTVerified,
		WalletStatusTError,
	}
}

func (e WalletStatusT) IsValid() error {
	switch e {
	case WalletStatusTCreated, WalletStatusTVerified, WalletStatusTError:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e WalletStatusT) String() string {
	return string(e)
}
