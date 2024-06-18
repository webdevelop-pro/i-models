package emails

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strings"
	"text/template"

	"github.com/Joker/jade"
	"github.com/jackc/pgtype"
	"github.com/pkg/errors"
)

type Email struct {
	ID               int     `json:"id"`
	Status           string  `json:"status"`
	UserID           *int    `json:"user_id"`
	TransactionID    *string `json:"transaction_id"`
	SentAttemptCount int     `json:"sent_attempt_count"`
	RecipientEmail   string  `json:"recipient_email"`
	RecipientName    string  `json:"recipient_name"`
	SenderEmail      string  `json:"sender_email"`
	SenderName       string  `json:"sender_name"`
	Subject          string  `json:"subject"`
	Template         string  `json:"template"`
	ContentText      string  `json:"-"`
	ContentHTML      string  `json:"-"`

	Data              interface{}            `json:"data"`
	MetaData          map[string]interface{} `json:"meta_data"`
	RecipientLocation map[string]interface{} `json:"recipient_location"`

	SentAttempAt pgtype.Timestamptz `json:"sent_attemp_at"`
	CreatedAt    pgtype.Timestamptz `json:"created_at"`
	UpdatedAt    pgtype.Timestamptz `json:"updated_at"`
}

func (e *Email) RemoveHTML() string {
	re := regexp.MustCompile("<[^>]*>")
	return re.ReplaceAllLiteralString(e.ContentHTML, "")
}

func (e *Email) RenderSubject() (string, error) {
	b := new(bytes.Buffer)

	tmpl, err := template.New("subject").Parse(e.Subject)
	if err != nil {
		return "", errors.Errorf("renderTemplate: template parse error: %v", err)
	}

	mergedData := e.Data // + merge incoming data
	if err := tmpl.Execute(b, mergedData); err != nil {
		return "", errors.Errorf("renderTemplate: template %v write error: %v", tmpl, err)
	}

	return b.String(), nil
}

func (e *Email) RenderPugTemplate() (string, error) {
	workDir, err := os.Getwd()
	if err != nil {
		return "", errors.Errorf("renderTemplate: get current workdir error: %v", err)
	}

	b := new(bytes.Buffer)

	jadeTpl, err := jade.ParseFile(fmt.Sprintf("%s/templates/%s", workDir, e.Template))
	if err != nil {
		return "", errors.Errorf("renderTemplate: jade parse error: %v", err)
	}

	// ToDo
	// Accept additional parameters
	mergedData := e.Data // + merge incoming data
	/*
			// merge two interfaces
			mp3 := make(map[string]interface{})
		    for k, v := range mp1 {
		        if _, ok := mp1[k]; ok {
		            mp3[k] = v
		        }
		    }

		    for k, v := range mp2 {
		        if _, ok := mp2[k]; ok {
		            mp3[k] = v
		        }
		    }

			// workaround to add site_url to the data
			data.(map[string]interface{})["site_url"] = a.settings.MainURL
			data.(map[string]interface{})["admin_panel_url"] = a.settings.AdminPanelURL
			data.(map[string]interface{})["current_year"] = time.Now().Year()
	*/

	goTpl, err := template.New("content").Parse(jadeTpl)
	if err != nil {
		return "", errors.Errorf("renderTemplate: template parse error: %v", err)
	}

	if err := goTpl.Execute(b, mergedData); err != nil {
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
