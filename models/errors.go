package models

import (
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

var (
	ErrRecordNotFound = errors.Wrap(pgx.ErrNoRows, "record not found")
	ErrNotFound       = ErrRecordNotFound
)

const (
	ErrSQLPrepare  = "error during sql prepare"
	ErrSQLRequest  = "sql error happen"
	ErrJSONMarshal = "json marshal error"

	ErrRetrieveOne = "cannot retrieve one record"
	ErrRetrieveAll = "cannot retrieve records"
	ErrCreate      = "cannot create record"
	ErrUpdate      = "cannot update record"
	ErrDelete      = "cannot delete record"
	ErrExist       = "cannot check existence"

	ErrEmptyPredicate = "empty predicate is not allowed"
	ErrNotUpdated     = "not record updated"
	ErrIDEmpty        = "model id is not set"
)
