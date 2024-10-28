package distributions

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type DistributionT string

// Enum values for DistributionT
const (
	DistributionTNew         DistributionT = "new"
	DistributionTInProgress  DistributionT = "in_progress"
	DistributionTSuccess     DistributionT = "success"
	DistributionTSystemError DistributionT = "system_error"
)

func AllDistributionT() []DistributionT {
	return []DistributionT{
		DistributionTNew,
		DistributionTInProgress,
		DistributionTSuccess,
		DistributionTSystemError,
	}
}

func (e DistributionT) IsValid() error {
	switch e {
	case DistributionTNew, DistributionTInProgress, DistributionTSuccess, DistributionTSystemError:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e DistributionT) String() string {
	return string(e)
}

type DistributionFilerT string

// Enum values for DistributionFilerT
const (
	DistributionFilerTTax   DistributionFilerT = "tax"
	DistributionFilerTOther DistributionFilerT = "other"
)

func AllDistributionFilerT() []DistributionFilerT {
	return []DistributionFilerT{
		DistributionFilerTTax,
		DistributionFilerTOther,
	}
}

func (e DistributionFilerT) IsValid() error {
	switch e {
	case DistributionFilerTTax, DistributionFilerTOther:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e DistributionFilerT) String() string {
	return string(e)
}

// NullDistributionFilerT is a nullable DistributionFilerT enum type. It supports SQL and JSON serialization.
type NullDistributionFilerT struct {
	Val   DistributionFilerT
	Valid bool
}

// NullDistributionFilerTFrom creates a new DistributionFilerT that will never be blank.
func NullDistributionFilerTFrom(v DistributionFilerT) NullDistributionFilerT {
	return NewNullDistributionFilerT(v, true)
}

// NullDistributionFilerTFromPtr creates a new NullDistributionFilerT that be null if s is nil.
func NullDistributionFilerTFromPtr(v *DistributionFilerT) NullDistributionFilerT {
	if v == nil {
		return NewNullDistributionFilerT("", false)
	}
	return NewNullDistributionFilerT(*v, true)
}

// NewNullDistributionFilerT creates a new NullDistributionFilerT
func NewNullDistributionFilerT(v DistributionFilerT, valid bool) NullDistributionFilerT {
	return NullDistributionFilerT{
		Val:   v,
		Valid: valid,
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *NullDistributionFilerT) UnmarshalJSON(data []byte) error {

	if err := json.Unmarshal(data, &e.Val); err != nil {
		return err
	}

	e.Valid = true
	return nil
}

// MarshalJSON implements json.Marshaler.
func (e NullDistributionFilerT) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.Val)
}

// MarshalText implements encoding.TextMarshaler.
func (e NullDistributionFilerT) MarshalText() ([]byte, error) {
	return []byte(e.Val), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (e *NullDistributionFilerT) UnmarshalText(text []byte) error {
	if text == nil || len(text) == 0 {
		e.Valid = false
		return nil
	}

	e.Val = DistributionFilerT(text)
	e.Valid = true
	return nil
}

// SetValid changes this NullDistributionFilerT value and also sets it to be non-null.
func (e *NullDistributionFilerT) SetValid(v DistributionFilerT) {
	e.Val = v
	e.Valid = true
}

// Ptr returns a pointer to this NullDistributionFilerT value, or a nil pointer if this NullDistributionFilerT is null.
func (e NullDistributionFilerT) Ptr() *DistributionFilerT {
	if !e.Valid {
		return nil
	}
	return &e.Val
}

// IsZero returns true for null types.
func (e NullDistributionFilerT) IsZero() bool {
	return !e.Valid
}

// Value implements the driver Valuer interface.
func (e NullDistributionFilerT) Value() (any, error) {
	if !e.Valid {
		return nil, nil
	}
	return string(e.Val), nil
}
