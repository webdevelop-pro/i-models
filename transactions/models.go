package transactions

import "github.com/jackc/pgtype"

type Transaction struct {
	ID     int     `json:"id"`
	Amount float64 `json:"amount"`
	Status Status  `json:"status"`

	DestWalletID   int  `json:"-"`
	SourceWalletID int  `json:"-"`
	Type           Type `json:"type"`

	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
}
