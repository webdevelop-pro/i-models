package fundingsources

import "github.com/pkg/errors"

type FoundingSourceT string

// Enum values for FoundingSourceT
const (
	FoundingSourceTChecking FoundingSourceT = "checking"
	FoundingSourceTBank     FoundingSourceT = "bank"
	FoundingSourceTSavings  FoundingSourceT = "savings"
)

func AllFoundingSourceT() []FoundingSourceT {
	return []FoundingSourceT{
		FoundingSourceTChecking,
		FoundingSourceTBank,
		FoundingSourceTSavings,
	}
}

func (e FoundingSourceT) IsValid() error {
	switch e {
	case FoundingSourceTChecking, FoundingSourceTBank, FoundingSourceTSavings:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e FoundingSourceT) String() string {
	return string(e)
}
