package transactions

import (
	"context"

	"github.com/pkg/errors"
	"github.com/webdevelop-pro/go-common/db"
	"github.com/webdevelop-pro/go-common/logger"
	"github.com/webdevelop-pro/i-models/models"
	"github.com/webdevelop-pro/i-models/pgtype"
)

const Table = "wallet_transactions"
const pkgName = "models/transactions"

// Transaction is an object representing the database table.
type Transaction struct {
	ID              int                 `json:"id" yaml:"id"`
	SourceWalletID  int                 `json:"source_wallet_id" yaml:"source_wallet_id,omitempty"`
	DestWalletID    int                 `json:"dest_wallet_id" yaml:"dest_wallet_id,omitempty"`
	SourceFundingID *int                `json:"source_funding_id,omitempty" yaml:"source_wallet_id,omitempty"`
	DestFundingID   *int                `json:"dest_funding_id,omitempty" yaml:"dest_wallet_id,omitempty"`
	EntityID        *string             `json:"entity_id,omitempty" yaml:"entity_id,omitempty"`
	Type            TransactionsTypeT   `json:"type" yaml:"type"`
	Amount          float64             `json:"amount" yaml:"amount"`
	Status          TransactionsStatusT `json:"status" yaml:"status"`
	CreatedAt       pgtype.Timestamptz  `json:"created_at" yaml:"created_at"`
	UpdatedAt       pgtype.Timestamptz  `json:"updated_at" yaml:"updated_at"`

	updatedFields []string      `db:"-" json:"-"`
	db            db.Repository `db:"-" json:"-"`
}

func New(db db.Repository) *Transaction {
	return &Transaction{
		db: db,
	}
}

func (model Transaction) Fields() []string {
	return []string{
		"id",
		"source_wallet_id",
		"dest_wallet_id",
		"source_funding_id",
		"dest_funding_id",
		"entity_id",
		"type",
		"amount",
		"status",
		"created_at",
		"updated_at",
	}
}

func (model Transaction) GetField(field string) any {
	switch field {
	case "ID":
		return model.ID
	case "SourceWalletID":
		return model.SourceWalletID
	case "DestWalletID":
		return model.DestWalletID
	case "SourceFundingID":
		return model.SourceFundingID
	case "DestFundingID":
		return model.DestFundingID
	case "EntityID":
		return model.EntityID
	case "Type":
		return model.Type
	case "Amount":
		return model.Amount
	case "Status":
		return model.Status
	case "CreatedAt":
		return model.CreatedAt
	case "UpdatedAt":
		return model.UpdatedAt
	}
	return nil
}

func (model Transaction) GetValueByTag(field string) any {
	switch field {
	case "id":
		return model.ID
	case "source_wallet_id":
		return model.SourceWalletID
	case "dest_wallet_id":
		return model.DestWalletID
	case "source_funding_id":
		return model.SourceFundingID
	case "dest_funding_id":
		return model.DestFundingID
	case "entity_id":
		return model.EntityID
	case "type":
		return model.Type
	case "amount":
		return model.Amount
	case "status":
		return model.Status
	case "created_at":
		return model.CreatedAt
	case "updated_at":
		return model.UpdatedAt
	}
	return nil
}

func (model *Transaction) SetStatus(val TransactionsStatusT) {
	model.Status = val
	model.updatedFields = append(model.updatedFields, "status")
}

func (model Transaction) Save(ctx context.Context) error {
	if model.ID == 0 {
		err := errors.Errorf("%s: Transaction %d", models.ErrIDEmpty, model.ID)
		logger.FromCtx(ctx, pkgName).Error().Stack().Err(err).Msg(models.ErrIDEmpty)
		return err
	}

	updates := map[string]any{}
	for _, field := range model.updatedFields {
		updates[field] = model.GetValueByTag(field)
	}
	updated, err := models.Update[Transaction](
		ctx,
		model.db,
		map[string]any{
			"id": model.ID,
		},
		updates,
	)
	if err != nil {
		err = errors.Wrapf(err, "cannot update %d", model.ID)
		return err
	}
	if updated == false {
		err := errors.Errorf("%s: Transaction %d", models.ErrNotUpdated, model.ID)
		logger.FromCtx(ctx, pkgName).Error().Stack().Err(err).Msg(models.ErrNotUpdated)
		return err
	}
	return nil
}

func Get(ctx context.Context, db db.Repository, where map[string]any) (*Transaction, error) {
	model, err := models.RetriveOne[Transaction](
		ctx,
		db,
		where,
	)
	if err != nil {
		err = errors.Wrapf(err, "cannot get model")
		return nil, err
	}
	model.db = db
	return model, nil
}

func GetByID(ctx context.Context, db db.Repository, id int) (*Transaction, error) {
	model, err := models.RetriveOne[Transaction](
		ctx,
		db,
		map[string]any{
			"id": id,
		},
	)
	if err != nil {
		err = errors.Wrapf(err, "cannot get model")
		return nil, err
	}
	model.db = db
	return model, nil
}

func (model Transaction) ToMap() map[string]any {
	res := map[string]any{}
	fields := model.Fields()
	for _, key := range fields {
		res[key] = model.GetField(key)
	}
	return res
}

func (model Transaction) Table() string {
	return Table
}

func (model Transaction) GetID() any {
	return model.ID
}

func (model *Transaction) SetID(id any) {
	model.ID = id.(int)
}
