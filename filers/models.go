package filers

import (
	"github.com/webdevelop-pro/go-common/db"
	"github.com/webdevelop-pro/i-models/pgtype"
)

// FilerFiler is an object representing the database table.
type FilerFiler struct {
	ID          int                `json:"id" yaml:"id"`
	UserID      int                `json:"user_id,omitempty" yaml:"user_id,omitempty"`
	GroupID     int                `json:"group_id,omitempty" yaml:"group_id,omitempty"`
	Filename    string             `json:"filename" yaml:"filename"`
	URL         string             `json:"url" yaml:"url"`
	Mime        string             `json:"mime" yaml:"mime"`
	Name        string             `json:"name" yaml:"name"`
	Description string             `json:"description" yaml:"description"`
	BucketName  string             `json:"bucket_name" yaml:"bucket_name"`
	BucketPath  string             `json:"bucket_path" yaml:"bucket_path"`
	MetaData    map[string]any     `json:"meta_data" yaml:"meta_data"`
	CreatedAt   pgtype.Timestamptz `json:"created_at" yaml:"created_at"`
	UpdatedAt   pgtype.Timestamptz `json:"updated_at" yaml:"updated_at"`

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
		"bucket_name": model.BucketName,
		"bucket_path": model.BucketPath,
		"meta_data":   model.MetaData,
		"created_at":  model.CreatedAt,
		"updated_at":  model.UpdatedAt,
	}
}

func (model FilerFiler) Fields() []string {
	return []string{
		"ID",
		"UserID",
		"GroupID",
		"Filename",
		"URL",
		"Mime",
		"Name",
		"Description",
		"BucketName",
		"BucketPath",
		"MetaData",
		"CreatedAt",
		"UpdatedAt",
	}
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
