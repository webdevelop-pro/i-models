package filers

import "github.com/webdevelop-pro/i-models/pgtype"

type FileMeta struct {
	BucketName string `json:"-"`
	BucketPath string `json:"-"`

	ObjectID   int               `json:"object-id"`
	ObjectName string            `json:"object-name"`
	ObjectType string            `json:"object-type"`
	ObjectData map[string]string `json:"object-data"`

	ID          int64             `json:"id,omitempty"`
	Filename    string            `json:"filename"`
	UserID      int64             `json:"user_id"`
	GroupID     int64             `json:"group_id,omitempty"`
	Mime        string            `json:"mime,omitempty"`
	Name        string            `json:"name"`
	Description string            `json:"description,omitempty"`
	MetaData    map[string]string `json:"meta_data"`

	// use to insert link on separate services, like pandadoc and etc
	URL string `json:"url"`

	CreatedAt pgtype.Timestamptz `json:"created_at,omitempty"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at,omitempty"`
}
