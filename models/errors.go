package models

import (
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

var (
	ErrNotFound = errors.Wrapf(pgx.ErrNoRows, "") // so we have stack trace and error message from pgx for errors.Is
)

const (
	ErrSQLPrepare  = "error during sql prepare"
	ErrSQLRequest  = "sql error happen"
	ErrJSONMarshal = "json marshal error"

	ErrRetrieveOne = "cannot retrieve one record"
	ErrRetrieveAll = "cannot retrieve records"
	ErrCreate      = "cannot create record"
	ErrUpdate      = "cannot update record"
	ErrExist       = "cannot check existence"

	ErrNotUpdated = "not record updated"
	ErrIDEmpty    = "model id is not set"
)
