package offers

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type OfferFilerT string

// Enum values for OfferFilerT
const (
	OfferFilerTCompany              OfferFilerT = "company"
	OfferFilerTTax                  OfferFilerT = "tax"
	OfferFilerTInvestmentAgreements OfferFilerT = "investment-agreements"
	OfferFilerTInvestorUpdates      OfferFilerT = "investor-updates"
	OfferFilerTMedia                OfferFilerT = "media"
	OfferFilerTOther                OfferFilerT = "other"
)

func AllOfferFilerT() []OfferFilerT {
	return []OfferFilerT{
		OfferFilerTCompany,
		OfferFilerTTax,
		OfferFilerTInvestmentAgreements,
		OfferFilerTInvestorUpdates,
		OfferFilerTMedia,
		OfferFilerTOther,
	}
}

func (e OfferFilerT) IsValid() error {
	switch e {
	case OfferFilerTCompany, OfferFilerTTax, OfferFilerTInvestmentAgreements, OfferFilerTInvestorUpdates, OfferFilerTMedia, OfferFilerTOther:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e OfferFilerT) String() string {
	return string(e)
}

// NullOfferFilerT is a nullable OfferFilerT enum type. It supports SQL and JSON serialization.
type NullOfferFilerT struct {
	Val   OfferFilerT
	Valid bool
}

// NullOfferFilerTFrom creates a new OfferFilerT that will never be blank.
func NullOfferFilerTFrom(v OfferFilerT) NullOfferFilerT {
	return NewNullOfferFilerT(v, true)
}

// NullOfferFilerTFromPtr creates a new NullOfferFilerT that be null if s is nil.
func NullOfferFilerTFromPtr(v *OfferFilerT) NullOfferFilerT {
	if v == nil {
		return NewNullOfferFilerT("", false)
	}
	return NewNullOfferFilerT(*v, true)
}

// NewNullOfferFilerT creates a new NullOfferFilerT
func NewNullOfferFilerT(v OfferFilerT, valid bool) NullOfferFilerT {
	return NullOfferFilerT{
		Val:   v,
		Valid: valid,
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *NullOfferFilerT) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &e.Val); err != nil {
		return err
	}

	e.Valid = true
	return nil
}

// MarshalJSON implements json.Marshaler.
func (e NullOfferFilerT) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.Val)
}

// MarshalText implements encoding.TextMarshaler.
func (e NullOfferFilerT) MarshalText() ([]byte, error) {
	if !e.Valid {
		return []byte{}, nil
	}
	return []byte(e.Val), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (e *NullOfferFilerT) UnmarshalText(text []byte) error {
	if text == nil || len(text) == 0 {
		e.Valid = false
		return nil
	}

	e.Val = OfferFilerT(text)
	e.Valid = true
	return nil
}

// SetValid changes this NullOfferFilerT value and also sets it to be non-null.
func (e *NullOfferFilerT) SetValid(v OfferFilerT) {
	e.Val = v
	e.Valid = true
}

// Ptr returns a pointer to this NullOfferFilerT value, or a nil pointer if this NullOfferFilerT is null.
func (e NullOfferFilerT) Ptr() *OfferFilerT {
	if !e.Valid {
		return nil
	}
	return &e.Val
}

// IsZero returns true for null types.
func (e NullOfferFilerT) IsZero() bool {
	return !e.Valid
}

// Value implements the driver Valuer interface.
func (e NullOfferFilerT) Value() (any, error) {
	if !e.Valid {
		return nil, nil
	}
	return string(e.Val), nil
}

// OfferOfferFiler is an object representing the database table.
type OfferOfferFiler struct {
	ID      int             `json:"id" yaml:"id"`
	OfferID int             `json:"offer_id,omitempty" yaml:"offer_id,omitempty"`
	FilerID int             `json:"filer_id,omitempty" yaml:"filer_id,omitempty"`
	Type    NullOfferFilerT `json:"type,omitempty" yaml:"type,omitempty"`
}

func (model OfferOfferFiler) ToJSON() map[string]any {
	return map[string]any{
		"id":       model.ID,
		"offer_id": model.OfferID,
		"filer_id": model.FilerID,
		"type":     model.Type,
	}
}

func (model OfferOfferFiler) Fields() []string {
	return []string{
		"ID",
		"OfferID",
		"FilerID",
		"Type",
	}
}

func (model OfferOfferFiler) Table() string {
	return "offer_offer_filers"
}

func (model OfferOfferFiler) GetID() any {
	return model.ID
}
