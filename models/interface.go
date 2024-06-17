package models

type Model interface {
	Table() string
	Fields() []string
	// UpdatedFields() []string
	PrimaryFieldKey() string
	PrimaryFieldValue() any
	// New(*db.DB, string)
	// Select(context.Context, []string, any) error
}
