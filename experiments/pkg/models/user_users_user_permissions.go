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

// UserUsersUserPermission is an object representing the database table.
type UserUsersUserPermission struct {
	ID           int64 `json:"id" yaml:"id"`
	AccountID    int64 `json:"account_id" yaml:"account_id"`
	PermissionID int   `json:"permission_id" yaml:"permission_id"`
}

func (model UserUsersUserPermission) ToJSON() map[string]any {
	return map[string]any{
		"id":            model.ID,
		"account_id":    model.AccountID,
		"permission_id": model.PermissionID,
	}
}

func (model UserUsersUserPermission) Fields() []string {
	return []string{
		"id",
		"account_id",
		"permission_id",
	}
}

func (model UserUsersUserPermission) Table() string {
	return "user_users_user_permissions"
}

func (model UserUsersUserPermission) GetID() any {
	return model.ID
}

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
)
