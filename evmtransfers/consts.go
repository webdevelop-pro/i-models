package evmtransfers

import "github.com/pkg/errors"

type StatusT string

// Enum values for StatusT
const (
	StatusTCreated  StatusT = "created"
	StatusTPending  StatusT = "pending"
	StatusTDeployed StatusT = "deployed"
	StatusTError    StatusT = "error"

	AppLabel  = "evm"
	ModelName = "transfer"
	TableName = "evm_transfers"

	pkgName = "models/evmtransfers"
)

func AllStatusT() []StatusT {
	return []StatusT{
		StatusTCreated,
		StatusTPending,
		StatusTDeployed,
		StatusTError,
	}
}

func (e StatusT) IsValid() error {
	switch e {
	case StatusTCreated, StatusTPending, StatusTDeployed, StatusTError:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e StatusT) String() string {
	return string(e)
}
