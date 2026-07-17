package filers

import (
	"github.com/global-torque/go-common/db/v2"
	"github.com/global-torque/go-common/orm/v2/pgtype"
)

// FilerFiler is an object representing the database table.
type FilerFiler struct {
	ID               int                `db:"id" json:"id" yaml:"id"`
	UserID           *int               `db:"user_id" json:"user_id,omitempty" yaml:"user_id,omitempty"`
	GroupID          *int               `db:"group_id" json:"group_id,omitempty" yaml:"group_id,omitempty"`
	Filename         string             `db:"filename" json:"filename" yaml:"filename"`
	OriginalFilename string             `db:"original_filename" json:"original_filename" yaml:"original_filename"`
	OriginalExt      string             `db:"original_ext" json:"original_ext" yaml:"original_ext"`
	URL              string             `db:"url" json:"url" yaml:"url"`
	Mime             string             `db:"mime" json:"mime" yaml:"mime"`
	Name             string             `db:"name" json:"name" yaml:"name"`
	Description      string             `db:"description" json:"description" yaml:"description"`
	Path             string             `db:"path" json:"path" yaml:"path"`
	IsPublic         bool               `db:"is_public" json:"is_public" yaml:"is_public"`
	Type             Type               `db:"type" json:"type" yaml:"type"`
	CreatedBy        *int               `db:"created_by" json:"created_by,omitempty" yaml:"created_by,omitempty"`
	BucketName       string             `db:"-" json:"bucket_name,omitempty" yaml:"bucket_name,omitempty"`
	BucketPath       string             `db:"-" json:"bucket_path,omitempty" yaml:"bucket_path,omitempty"`
	MetaData         map[string]any     `db:"meta_data" json:"meta_data" yaml:"meta_data"`
	CreatedAt        pgtype.Timestamptz `db:"created_at" json:"created_at" yaml:"created_at"`
	UpdatedAt        pgtype.Timestamptz `db:"updated_at" json:"updated_at" yaml:"updated_at"`

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
		"id":                model.ID,
		"user_id":           model.UserID,
		"group_id":          model.GroupID,
		"filename":          model.Filename,
		"original_filename": model.OriginalFilename,
		"original_ext":      model.OriginalExt,
		"url":               model.URL,
		"mime":              model.Mime,
		"name":              model.Name,
		"description":       model.Description,
		"path":              model.Path,
		"is_public":         model.IsPublic,
		"type":              model.Type,
		"created_by":        model.CreatedBy,
		"meta_data":         model.MetaData,
		"created_at":        model.CreatedAt,
		"updated_at":        model.UpdatedAt,
	}
}

func (model FilerFiler) Fields() []string {
	return []string{
		"id", "user_id", "group_id", "filename", "url", "mime", "name",
		"description", "meta_data", "created_at", "updated_at",
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
