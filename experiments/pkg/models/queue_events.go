// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

type TopicT string

// Enum values for TopicT
const (
	TopicTUser       TopicT = "user"
	TopicTInvestment TopicT = "investment"
	TopicTOffer      TopicT = "offer"
	TopicTFiler      TopicT = "filer"
	TopicTEmail      TopicT = "email"
)

func AllTopicT() []TopicT {
	return []TopicT{
		TopicTUser,
		TopicTInvestment,
		TopicTOffer,
		TopicTFiler,
		TopicTEmail,
	}
}

func (e TopicT) IsValid() error {
	switch e {
	case TopicTUser, TopicTInvestment, TopicTOffer, TopicTFiler, TopicTEmail:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e TopicT) String() string {
	return string(e)
}

// NullTopicT is a nullable TopicT enum type. It supports SQL and JSON serialization.
type NullTopicT struct {
	Val   TopicT
	Valid bool
}

// NullTopicTFrom creates a new TopicT that will never be blank.
func NullTopicTFrom(v TopicT) NullTopicT {
	return NewNullTopicT(v, true)
}

// NullTopicTFromPtr creates a new NullTopicT that be null if s is nil.
func NullTopicTFromPtr(v *TopicT) NullTopicT {
	if v == nil {
		return NewNullTopicT("", false)
	}
	return NewNullTopicT(*v, true)
}

// NewNullTopicT creates a new NullTopicT
func NewNullTopicT(v TopicT, valid bool) NullTopicT {
	return NullTopicT{
		Val:   v,
		Valid: valid,
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *NullTopicT) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &e.Val); err != nil {
		return err
	}

	e.Valid = true
	return nil
}

// MarshalJSON implements json.Marshaler.
func (e NullTopicT) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.Val)
}

// MarshalText implements encoding.TextMarshaler.
func (e NullTopicT) MarshalText() ([]byte, error) {
	if !e.Valid {
		return []byte{}, nil
	}
	return []byte(e.Val), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (e *NullTopicT) UnmarshalText(text []byte) error {
	if text == nil || len(text) == 0 {
		e.Valid = false
		return nil
	}

	e.Val = TopicT(text)
	e.Valid = true
	return nil
}

// SetValid changes this NullTopicT value and also sets it to be non-null.
func (e *NullTopicT) SetValid(v TopicT) {
	e.Val = v
	e.Valid = true
}

// Ptr returns a pointer to this NullTopicT value, or a nil pointer if this NullTopicT is null.
func (e NullTopicT) Ptr() *TopicT {
	if !e.Valid {
		return nil
	}
	return &e.Val
}

// IsZero returns true for null types.
func (e NullTopicT) IsZero() bool {
	return !e.Valid
}

// Value implements the driver Valuer interface.
func (e NullTopicT) Value() (any, error) {
	if !e.Valid {
		return nil, nil
	}
	return string(e.Val), nil
}

// QueueEvent is an object representing the database table.
type QueueEvent struct {
	ID        int        `json:"id" yaml:"id"`
	CreatedAt time.Time  `json:"created_at" yaml:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" yaml:"updated_at"`
	Topic     NullTopicT `json:"topic,omitempty" yaml:"topic,omitempty"`
	SubTopic  string     `json:"sub_topic" yaml:"sub_topic"`
	Payload   null.JSON  `json:"payload,omitempty" yaml:"payload,omitempty"`
}

func (model QueueEvent) ToJSON() map[string]any {
	return map[string]any{
		"id":         model.ID,
		"created_at": model.CreatedAt,
		"updated_at": model.UpdatedAt,
		"topic":      model.Topic,
		"sub_topic":  model.SubTopic,
		"payload":    model.Payload,
	}
}

func (model QueueEvent) Fields() []string {
	return []string{
		"id",
		"created_at",
		"updated_at",
		"topic",
		"sub_topic",
		"payload",
	}
}

func (model QueueEvent) Table() string {
	return "queue_events"
}

func (model QueueEvent) GetID() any {
	return model.ID
}

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
)
