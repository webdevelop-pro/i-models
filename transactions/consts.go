package transactions

import "github.com/pkg/errors"

type TransactionsTypeT string

// Enum values for TransactionsTypeT
const (
	TransactionsTypeTDeposit    TransactionsTypeT = "deposit"
	TransactionsTypeTInvestment TransactionsTypeT = "investment"
	TransactionsTypeTWithdrawal TransactionsTypeT = "withdrawal"

	AppLabel  = "wallet"
	ModelName = "transaction"
)

func AllTransactionsTypeT() []TransactionsTypeT {
	return []TransactionsTypeT{
		TransactionsTypeTDeposit,
		TransactionsTypeTInvestment,
		TransactionsTypeTWithdrawal,
	}
}

func (e TransactionsTypeT) IsValid() error {
	switch e {
	case TransactionsTypeTDeposit, TransactionsTypeTInvestment, TransactionsTypeTWithdrawal:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e TransactionsTypeT) String() string {
	return string(e)
}

type TransactionsStatusT string

// Enum values for TransactionsStatusT
const (
	TransactionsStatusTCreated   TransactionsStatusT = "created"
	TransactionsStatusTPending   TransactionsStatusT = "pending"
	TransactionsStatusTProcessed TransactionsStatusT = "processed"
	TransactionsStatusTCancelled TransactionsStatusT = "cancelled"
	TransactionsStatusTWait      TransactionsStatusT = "wait"
	TransactionsStatusTFailed    TransactionsStatusT = "failed"
)

func AllTransactionsStatusT() []TransactionsStatusT {
	return []TransactionsStatusT{
		TransactionsStatusTCreated,
		TransactionsStatusTPending,
		TransactionsStatusTProcessed,
		TransactionsStatusTCancelled,
		TransactionsStatusTWait,
		TransactionsStatusTFailed,
	}
}

func (e TransactionsStatusT) IsValid() error {
	switch e {
	case TransactionsStatusTCreated, TransactionsStatusTPending, TransactionsStatusTProcessed, TransactionsStatusTCancelled, TransactionsStatusTWait, TransactionsStatusTFailed:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e TransactionsStatusT) String() string {
	return string(e)
}
