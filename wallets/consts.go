package wallets

import "github.com/friendsofgo/errors"

type WalletStatusT string

// Enum values for WalletStatusT
const (
	WalletStatusTCreated  WalletStatusT = "created"
	WalletStatusTVerified WalletStatusT = "verified"
	WalletStatusTError    WalletStatusT = "error"
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
