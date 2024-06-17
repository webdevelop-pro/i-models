package fundingsource

import "github.com/jackc/pgtype"

type FoundingSource struct {
	ID int `json:"id"`

	Status   Status `json:"status"`
	Type     Type   `json:"type"`
	BankName string `json:"bank_name"`
	Name     string `json:"name"`

	WalletID int    `json:"-"`
	DwollaID string `json:"-"`

	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
}
