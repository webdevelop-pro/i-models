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
} // WalletFundingSource is an object representing the database table.
type WalletFundingSource struct {
	ID        int             `json:"id" yaml:"id"`
	WalletID  int        `json:"wallet_id,omitempty" yaml:"wallet_id,omitempty"`
	EntityID  string          `json:"entity_id" yaml:"entity_id"`
	Type      FoundingSourceT `json:"type" yaml:"type"`
	BankName  string          `json:"bank_name" yaml:"bank_name"`
	Status    WalletStatusT   `json:"status" yaml:"status"`
	Name      string          `json:"name" yaml:"name"`
	CreatedAt time.Time       `json:"created_at" yaml:"created_at"`
	UpdatedAt time.Time       `json:"updated_at" yaml:"updated_at"`
}

func (model WalletFundingSource) ToJSON() map[string]any {
	return map[string]any{
		"id":         model.ID,
		"wallet_id":  model.WalletID,
		"entity_id":  model.EntityID,
		"type":       model.Type,
		"bank_name":  model.BankName,
		"status":     model.Status,
		"name":       model.Name,
		"created_at": model.CreatedAt,
		"updated_at": model.UpdatedAt,
	}
}

func (model WalletFundingSource) Fields() []string {
	return []string{
		"id",
		"wallet_id",
		"entity_id",
		"type",
		"bank_name",
		"status",
		"name",
		"created_at",
		"updated_at",
	}
}

func (model WalletFundingSource) Table() string {
	return "wallet_funding_source"
}

func (model WalletFundingSource) GetID() any {
	return model.ID
}

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
)
