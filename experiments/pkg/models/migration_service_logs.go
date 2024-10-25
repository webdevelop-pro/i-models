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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// MigrationServiceLog is an object representing the database table.
type MigrationServiceLog struct {
	ID                    int       `json:"id" yaml:"id"`
	MigrationServicesName string    `json:"migration_services_name" yaml:"migration_services_name"`
	Priority              int       `json:"priority" yaml:"priority"`
	Version               int       `json:"version" yaml:"version"`
	FileName              string    `json:"file_name" yaml:"file_name"`
	SQL                   string    `json:"sql" yaml:"sql"`
	Hash                  string    `json:"hash" yaml:"hash"`
	CreatedAt             time.Time `json:"created_at" yaml:"created_at"`
	UpdatedAt             time.Time `json:"updated_at" yaml:"updated_at"`
}

func (model MigrationServiceLog) ToJSON() map[string]any {
	return map[string]any{
		"id":                      model.ID,
		"migration_services_name": model.MigrationServicesName,
		"priority":                model.Priority,
		"version":                 model.Version,
		"file_name":               model.FileName,
		"sql":                     model.SQL,
		"hash":                    model.Hash,
		"created_at":              model.CreatedAt,
		"updated_at":              model.UpdatedAt,
	}
}

func (model MigrationServiceLog) Fields() []string {
	return []string{
		"id",
		"migration_services_name",
		"priority",
		"version",
		"file_name",
		"sql",
		"hash",
		"created_at",
		"updated_at",
	}
}

func (model MigrationServiceLog) Table() string {
	return "migration_service_logs"
}

func (model MigrationServiceLog) GetID() any {
	return model.ID
}

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
)