package wallets

import (
	"github.com/pkg/errors"
)

type WalletStatusT string

// Enum values for WalletStatusT
const (
	WalletStatusTCreated        WalletStatusT = "created"
	WalletStatusTPending        WalletStatusT = "pending"
	WalletStatusTVerified       WalletStatusT = "verified"
	WalletStatusTError          WalletStatusT = "error"
	WalletStatusTErrorDocument  WalletStatusT = "error_document"
	WalletStatusTErrorPending   WalletStatusT = "error_pending"
	WalletStatusTErrorRetry     WalletStatusT = "error_retry"
	WalletStatusTErrorSuspended WalletStatusT = "error_suspended"

	AppLabel  = "wallet"
	ModelName = "wallet"

	TableName = "wallet_wallets"
)

func AllWalletStatusT() []WalletStatusT {
	return []WalletStatusT{
		WalletStatusTCreated,
		WalletStatusTPending,
		WalletStatusTVerified,
		WalletStatusTError,
		WalletStatusTErrorDocument,
		WalletStatusTErrorPending,
		WalletStatusTErrorRetry,
		WalletStatusTErrorSuspended,
	}
}

func (e WalletStatusT) IsValid() error {
	switch e {
	case WalletStatusTCreated, WalletStatusTPending, WalletStatusTVerified, WalletStatusTError, WalletStatusTErrorDocument, WalletStatusTErrorPending, WalletStatusTErrorRetry, WalletStatusTErrorSuspended:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e WalletStatusT) String() string {
	return string(e)
}
