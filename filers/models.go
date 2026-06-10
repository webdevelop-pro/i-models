package filers

import (
	"github.com/webdevelop-pro/go-common/db"
	"github.com/webdevelop-pro/go-common/orm"
	"github.com/webdevelop-pro/go-common/orm/pgtype"
)

// FilerFiler is an object representing the database table.
type FilerFiler struct {
	ID          int                `db:"id" json:"id" yaml:"id"`
	UserID      int                `db:"user_id" json:"user_id,omitempty" yaml:"user_id,omitempty"`
	GroupID     int                `db:"group_id" json:"group_id,omitempty" yaml:"group_id,omitempty"`
	Filename    string             `db:"filename" json:"filename" yaml:"filename"`
	URL         string             `db:"url" json:"url" yaml:"url"`
	Mime        string             `db:"mime" json:"mime" yaml:"mime"`
	Name        string             `db:"name" json:"name" yaml:"name"`
	Description string             `db:"description" json:"description" yaml:"description"`
	MetaData    map[string]any     `db:"meta_data" json:"meta_data" yaml:"meta_data"`
	CreatedAt   pgtype.Timestamptz `db:"created_at" json:"created_at" yaml:"created_at"`
	UpdatedAt   pgtype.Timestamptz `db:"updated_at" json:"updated_at" yaml:"updated_at"`

	updatedFields []string       `db:"-" json:"-"`
	fns           map[string]any `db:"-" json:"-"`
	db            db.Repository  `db:"-" json:"-"`
}

func New(db db.Repository) *FilerFiler {
	return &FilerFiler{
		db:  db,
		fns: map[string]any{},
	}
}

func (model FilerFiler) ToJSON() map[string]any {
	return map[string]any{
		"id":          model.ID,
		"user_id":     model.UserID,
		"group_id":    model.GroupID,
		"filename":    model.Filename,
		"url":         model.URL,
		"mime":        model.Mime,
		"name":        model.Name,
		"description": model.Description,
		"meta_data":   model.MetaData,
		"created_at":  model.CreatedAt,
		"updated_at":  model.UpdatedAt,
	}
}

func (model FilerFiler) Fields() []string {
	return orm.DefaultFields(&model)
}

func (model FilerFiler) Table() string {
	return TableName
}

func (model FilerFiler) GetID() any {
	return model.ID
}

func (model *FilerFiler) SetID(id any) {
	model.ID = id.(int)
}

func (model *FilerFiler) SetDB(db db.Repository) {
	model.db = db
}
