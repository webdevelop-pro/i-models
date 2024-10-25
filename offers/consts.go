package offers

import (
	"encoding/json"
	"errors"
)

type OfferT string

// Enum values for OfferT
const (
	OfferTNew                 OfferT = "new"
	OfferTDraft               OfferT = "draft"
	OfferTLegalReview         OfferT = "legal-review"
	OfferTLegalDeclined       OfferT = "legal-declined"
	OfferTLegalAccepted       OfferT = "legal-accepted"
	OfferTPublished           OfferT = "published"
	OfferTLegalClosed         OfferT = "legal-closed"
	OfferTClosedSuccessfully  OfferT = "closed-successfully"
	OfferTClosedUnsuccesfully OfferT = "closed-unsuccesfully"
)

func AllOfferT() []OfferT {
	return []OfferT{
		OfferTNew,
		OfferTDraft,
		OfferTLegalReview,
		OfferTLegalDeclined,
		OfferTLegalAccepted,
		OfferTPublished,
		OfferTLegalClosed,
		OfferTClosedSuccessfully,
		OfferTClosedUnsuccesfully,
	}
}

func (e OfferT) IsValid() error {
	switch e {
	case OfferTNew, OfferTDraft, OfferTLegalReview, OfferTLegalDeclined, OfferTLegalAccepted, OfferTPublished, OfferTLegalClosed, OfferTClosedSuccessfully, OfferTClosedUnsuccesfully:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e OfferT) String() string {
	return string(e)
}

// NullOfferT is a nullable OfferT enum type. It supports SQL and JSON serialization.
type NullOfferT struct {
	Val   OfferT
	Valid bool
}

// NullOfferTFrom creates a new OfferT that will never be blank.
func NullOfferTFrom(v OfferT) NullOfferT {
	return NewNullOfferT(v, true)
}

// NullOfferTFromPtr creates a new NullOfferT that be null if s is nil.
func NullOfferTFromPtr(v *OfferT) NullOfferT {
	if v == nil {
		return NewNullOfferT("", false)
	}
	return NewNullOfferT(*v, true)
}

// NewNullOfferT creates a new NullOfferT
func NewNullOfferT(v OfferT, valid bool) NullOfferT {
	return NullOfferT{
		Val:   v,
		Valid: valid,
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *NullOfferT) UnmarshalJSON(data []byte) error {
	e.Val = ""
	e.Valid = false
	return nil

	if err := json.Unmarshal(data, &e.Val); err != nil {
		return err
	}

	e.Valid = true
	return nil
}

// MarshalJSON implements json.Marshaler.
func (e NullOfferT) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.Val)
}

// MarshalText implements encoding.TextMarshaler.
func (e NullOfferT) MarshalText() ([]byte, error) {
	if !e.Valid {
		return []byte{}, nil
	}
	return []byte(e.Val), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (e *NullOfferT) UnmarshalText(text []byte) error {
	if text == nil || len(text) == 0 {
		e.Valid = false
		return nil
	}

	e.Val = OfferT(text)
	e.Valid = true
	return nil
}

// SetValid changes this NullOfferT value and also sets it to be non-null.
func (e *NullOfferT) SetValid(v OfferT) {
	e.Val = v
	e.Valid = true
}

// Ptr returns a pointer to this NullOfferT value, or a nil pointer if this NullOfferT is null.
func (e NullOfferT) Ptr() *OfferT {
	if !e.Valid {
		return nil
	}
	return &e.Val
}

// IsZero returns true for null types.
func (e NullOfferT) IsZero() bool {
	return !e.Valid
}

// Value implements the driver Valuer interface.
func (e NullOfferT) Value() (any, error) {
	if !e.Valid {
		return nil, nil
	}
	return string(e.Val), nil
}

type OfferSecurityTypeT string

// Enum values for OfferSecurityTypeT
const (
	OfferSecurityTypeTEquity           OfferSecurityTypeT = "equity"
	OfferSecurityTypeTPreferredEquity  OfferSecurityTypeT = "preferred-equity"
	OfferSecurityTypeTDebt             OfferSecurityTypeT = "debt"
	OfferSecurityTypeTConvertibleDebt  OfferSecurityTypeT = "convertible-debt"
	OfferSecurityTypeTEquityWarrants   OfferSecurityTypeT = "equity-warrants"
	OfferSecurityTypeTConvertibleBonds OfferSecurityTypeT = "convertible-bonds"
	OfferSecurityTypeTPreferenceShares OfferSecurityTypeT = "preference-shares"
)

func AllOfferSecurityTypeT() []OfferSecurityTypeT {
	return []OfferSecurityTypeT{
		OfferSecurityTypeTEquity,
		OfferSecurityTypeTPreferredEquity,
		OfferSecurityTypeTDebt,
		OfferSecurityTypeTConvertibleDebt,
		OfferSecurityTypeTEquityWarrants,
		OfferSecurityTypeTConvertibleBonds,
		OfferSecurityTypeTPreferenceShares,
	}
}

func (e OfferSecurityTypeT) IsValid() error {
	switch e {
	case OfferSecurityTypeTEquity, OfferSecurityTypeTPreferredEquity, OfferSecurityTypeTDebt, OfferSecurityTypeTConvertibleDebt, OfferSecurityTypeTEquityWarrants, OfferSecurityTypeTConvertibleBonds, OfferSecurityTypeTPreferenceShares:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e OfferSecurityTypeT) String() string {
	return string(e)
} // OfferOffer is an object representing the database table.

type CommentRelatedT string

// Enum values for CommentRelatedT
const (
	CommentRelatedTAdviser    CommentRelatedT = "adviser"
	CommentRelatedTEmployee   CommentRelatedT = "employee"
	CommentRelatedTAffiliated CommentRelatedT = "affiliated"
	CommentRelatedTInvestor   CommentRelatedT = "investor"
)

func AllCommentRelatedT() []CommentRelatedT {
	return []CommentRelatedT{
		CommentRelatedTAdviser,
		CommentRelatedTEmployee,
		CommentRelatedTAffiliated,
		CommentRelatedTInvestor,
	}
}

func (e CommentRelatedT) IsValid() error {
	switch e {
	case CommentRelatedTAdviser, CommentRelatedTEmployee, CommentRelatedTAffiliated, CommentRelatedTInvestor:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e CommentRelatedT) String() string {
	return string(e)
}

// NullCommentRelatedT is a nullable CommentRelatedT enum type. It supports SQL and JSON serialization.
type NullCommentRelatedT struct {
	Val   CommentRelatedT
	Valid bool
}

// NullCommentRelatedTFrom creates a new CommentRelatedT that will never be blank.
func NullCommentRelatedTFrom(v CommentRelatedT) NullCommentRelatedT {
	return NewNullCommentRelatedT(v, true)
}

// NullCommentRelatedTFromPtr creates a new NullCommentRelatedT that be null if s is nil.
func NullCommentRelatedTFromPtr(v *CommentRelatedT) NullCommentRelatedT {
	if v == nil {
		return NewNullCommentRelatedT("", false)
	}
	return NewNullCommentRelatedT(*v, true)
}

// NewNullCommentRelatedT creates a new NullCommentRelatedT
func NewNullCommentRelatedT(v CommentRelatedT, valid bool) NullCommentRelatedT {
	return NullCommentRelatedT{
		Val:   v,
		Valid: valid,
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *NullCommentRelatedT) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &e.Val); err != nil {
		return err
	}

	e.Valid = true
	return nil
}

// MarshalJSON implements json.Marshaler.
func (e NullCommentRelatedT) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.Val)
}

// MarshalText implements encoding.TextMarshaler.
func (e NullCommentRelatedT) MarshalText() ([]byte, error) {
	if !e.Valid {
		return []byte{}, nil
	}
	return []byte(e.Val), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (e *NullCommentRelatedT) UnmarshalText(text []byte) error {
	if text == nil || len(text) == 0 {
		e.Valid = false
		return nil
	}

	e.Val = CommentRelatedT(text)
	e.Valid = true
	return nil
}

// SetValid changes this NullCommentRelatedT value and also sets it to be non-null.
func (e *NullCommentRelatedT) SetValid(v CommentRelatedT) {
	e.Val = v
	e.Valid = true
}

// Ptr returns a pointer to this NullCommentRelatedT value, or a nil pointer if this NullCommentRelatedT is null.
func (e NullCommentRelatedT) Ptr() *CommentRelatedT {
	if !e.Valid {
		return nil
	}
	return &e.Val
}

// IsZero returns true for null types.
func (e NullCommentRelatedT) IsZero() bool {
	return !e.Valid
}

// Value implements the driver Valuer interface.
func (e NullCommentRelatedT) Value() (any, error) {
	if !e.Valid {
		return nil, nil
	}
	return string(e.Val), nil
}

type CommentStatusT string

// Enum values for CommentStatusT
const (
	CommentStatusTNew       CommentStatusT = "new"
	CommentStatusTPublished CommentStatusT = "published"
	CommentStatusTDeleted   CommentStatusT = "deleted"
)

func AllCommentStatusT() []CommentStatusT {
	return []CommentStatusT{
		CommentStatusTNew,
		CommentStatusTPublished,
		CommentStatusTDeleted,
	}
}

func (e CommentStatusT) IsValid() error {
	switch e {
	case CommentStatusTNew, CommentStatusTPublished, CommentStatusTDeleted:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e CommentStatusT) String() string {
	return string(e)
}

// NullCommentStatusT is a nullable CommentStatusT enum type. It supports SQL and JSON serialization.
type NullCommentStatusT struct {
	Val   CommentStatusT
	Valid bool
}

// NullCommentStatusTFrom creates a new CommentStatusT that will never be blank.
func NullCommentStatusTFrom(v CommentStatusT) NullCommentStatusT {
	return NewNullCommentStatusT(v, true)
}

// NullCommentStatusTFromPtr creates a new NullCommentStatusT that be null if s is nil.
func NullCommentStatusTFromPtr(v *CommentStatusT) NullCommentStatusT {
	if v == nil {
		return NewNullCommentStatusT("", false)
	}
	return NewNullCommentStatusT(*v, true)
}

// NewNullCommentStatusT creates a new NullCommentStatusT
func NewNullCommentStatusT(v CommentStatusT, valid bool) NullCommentStatusT {
	return NullCommentStatusT{
		Val:   v,
		Valid: valid,
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *NullCommentStatusT) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &e.Val); err != nil {
		return err
	}

	e.Valid = true
	return nil
}

// MarshalJSON implements json.Marshaler.
func (e NullCommentStatusT) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.Val)
}

// MarshalText implements encoding.TextMarshaler.
func (e NullCommentStatusT) MarshalText() ([]byte, error) {
	if !e.Valid {
		return []byte{}, nil
	}
	return []byte(e.Val), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (e *NullCommentStatusT) UnmarshalText(text []byte) error {
	if text == nil || len(text) == 0 {
		e.Valid = false
		return nil
	}

	e.Val = CommentStatusT(text)
	e.Valid = true
	return nil
}

// SetValid changes this NullCommentStatusT value and also sets it to be non-null.
func (e *NullCommentStatusT) SetValid(v CommentStatusT) {
	e.Val = v
	e.Valid = true
}

// Ptr returns a pointer to this NullCommentStatusT value, or a nil pointer if this NullCommentStatusT is null.
func (e NullCommentStatusT) Ptr() *CommentStatusT {
	if !e.Valid {
		return nil
	}
	return &e.Val
}

// IsZero returns true for null types.
func (e NullCommentStatusT) IsZero() bool {
	return !e.Valid
}

// Value implements the driver Valuer interface.
func (e NullCommentStatusT) Value() (any, error) {
	if !e.Valid {
		return nil, nil
	}
	return string(e.Val), nil
}
