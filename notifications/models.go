package notifications

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"strings"

	"github.com/Joker/jade"
	"github.com/jackc/pgtype"
	"github.com/pkg/errors"
)

type Notification struct {
	ID        int64              `json:"id"`
	UserID    int                `json:"user_id"`
	Content   string             `json:"content"`
	Data      interface{}        `json:"data"`
	Status    StatusType         `json:"status"`
	Type      Type               `json:"type"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
}

type Notifications []Notification

func (notif *Notification) RenderPugTemplate(name string, data map[string]any) (string, error) {
	workDir, err := os.Getwd()
	if err != nil {
		return "", errors.Errorf("renderTemplate: get current workdir error: %v", err)
	}
	name = fmt.Sprintf("%s/templates/%s", workDir, name)
	b := new(bytes.Buffer)

	jadeTpl, err := jade.ParseFile(name)
	if err != nil {
		return "", errors.Errorf("renderTemplate: jade parse error: %v", err)
	}

	goTpl, err := template.New("content").Parse(jadeTpl)
	if err != nil {
		return "", errors.Errorf("renderTemplate: template parse error: %v", err)
	}

	if err := goTpl.Execute(b, data); err != nil {
		return "", errors.Errorf("renderTemplate: template %v write error: %v", goTpl, err)
	}

	res := strings.ReplaceAll(
		strings.ReplaceAll(
			strings.ReplaceAll(
				b.String(), "\"", "'",
			),
			"\t", ""),
		"\n", "")

	return res, nil
}
