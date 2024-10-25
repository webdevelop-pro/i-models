package models

import (
	"github.com/friendsofgo/errors"
	"github.com/jackc/pgx/v5"
)

var (
	ErrNotFound = errors.Wrapf(pgx.ErrNoRows, "") // so we have stack trace and error message from pgx for errors.Is
)

const (
	ErrSQLPrepare  = "error during sql prepare"
	ErrSQLRequest  = "sql error happen"
	ErrJSONMarshal = "json marshal error"

	ErrRetrieveOne = "cannot retrieve one element"
	ErrRetrieveAll = "cannot retrieve elements"
	ErrCreate      = "cannot create record"
)
