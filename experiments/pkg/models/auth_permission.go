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

// AuthPermission is an object representing the database table.
type AuthPermission struct {
	ID            int    `json:"id" yaml:"id"`
	Name          string `json:"name" yaml:"name"`
	ContentTypeID int    `json:"content_type_id" yaml:"content_type_id"`
	Codename      string `json:"codename" yaml:"codename"`
}

func (model AuthPermission) ToJSON() map[string]any {
	return map[string]any{
		"id":              model.ID,
		"name":            model.Name,
		"content_type_id": model.ContentTypeID,
		"codename":        model.Codename,
	}
}

func (model AuthPermission) Fields() []string {
	return []string{
		"id",
		"name",
		"content_type_id",
		"codename",
	}
}

func (model AuthPermission) Table() string {
	return "auth_permission"
}

func (model AuthPermission) GetID() any {
	return model.ID
}

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
)