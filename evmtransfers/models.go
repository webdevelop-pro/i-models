package evmtransfers

import (
	"context"

	"github.com/pkg/errors"
	"github.com/webdevelop-pro/go-common/db"
	"github.com/webdevelop-pro/go-common/logger"
	"github.com/webdevelop-pro/go-common/queue/pclient"
	"github.com/webdevelop-pro/i-models/models"
	"github.com/webdevelop-pro/i-models/pgtype"
)

// Transfer is an object representing the database table.
type Transfer struct {
	ID             int  `db:"id" json:"id" yaml:"id"`
	UserID         int  `db:"user_id" json:"user_id" yaml:"user_id"`
	DestWalletID   *int `db:"dest_wallet_id" json:"dest_wallet_id" yaml:"dest_wallet_id"`
	SourceWalletID *int `db:"source_wallet_id" json:"source_wallet_id" yaml:"source_wallet_id"`
	InvestmentID   *int `db:"investment_id" json:"investment_id" yaml:"investment_id"`

	Status        string `db:"status" json:"status" yaml:"status"`
	TransactionTX string `db:"transaction_tx" json:"transaction_tx" yaml:"transaction_tx"`

	CreatedAt pgtype.Timestamptz `db:"created_at" json:"created_at" yaml:"created_at"`
	UpdatedAt pgtype.Timestamptz `db:"updated_at" json:"updated_at" yaml:"updated_at"`

	updatedFields []string       `db:"-" json:"-"`
	fns           map[string]any `db:"-" json:"-"`
	db            db.Repository  `db:"-" json:"-"`
}

func New(db db.Repository) *Transfer {
	return &Transfer{
		db:  db,
		fns: map[string]any{},
	}
}

func (model Transfer) GetField(name string) any {
	switch name {
	case "ID":
		return model.ID
	case "UserID":
		return model.UserID
	case "DestWalletID":
		return model.DestWalletID
	case "SourceWalletID":
		return model.SourceWalletID
	case "InvestmentID":
		return model.InvestmentID
	case "TransactionTX":
		return model.TransactionTX
	case "Status":
		return model.Status
	case "CreatedAt":
		return model.CreatedAt
	case "UpdatedAt":
		return model.UpdatedAt
	}
	return nil
}

func (model Transfer) GetValueByTag(name string) any {
	switch name {
	case "id":
		return model.ID
	case "investment_id":
		return model.InvestmentID
	case "transaction_tx":
		return model.TransactionTX
	case "status":
		return model.Status
	case "created_at":
		return model.CreatedAt
	case "updated_at":
		return model.UpdatedAt
	}
	return nil
}

func (model Transfer) ToJSON() map[string]any {
	res := map[string]any{}
	fields := model.Fields()
	for _, key := range fields {
		res[key] = model.GetField(key)
	}
	return res
}

func (model Transfer) Fields() []string {
	return models.DefaultFields(&model)
}

func (model Transfer) Table() string {
	return TableName
}

func (model Transfer) GetID() any {
	return model.ID
}

func (model *Transfer) SetID(id any) {
	model.ID = id.(int)
}

func (model *Transfer) SetDB(db db.Repository) {
	model.db = db
}

func (model Transfer) Save(ctx context.Context, postUpdate func(ctx context.Context, msg pclient.Event) error) error {
	if model.ID == 0 {
		err := errors.Errorf("%s: Transfer %d", models.ErrIDEmpty, model.ID)
		logger.FromCtx(ctx, pkgName).Error().Stack().Err(err).Msg(models.ErrIDEmpty)
		return err
	}

	updates := map[string]any{}
	for _, field := range model.updatedFields {
		updates[field] = model.GetValueByTag(field)
	}
	updated, err := models.Update[Transfer](
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
		err := errors.Errorf("%s: Transfer %d", models.ErrNotUpdated, model.ID)
		logger.FromCtx(ctx, pkgName).Error().Stack().Err(err).Msg(models.ErrNotUpdated)
		return err
	} else {
		postUpdate(ctx, pclient.Event{
			Action:     pclient.PostUpdate,
			ObjectID:   model.ID,
			ObjectName: ModelName,
			Data:       updates,
		})
		// model.DefaultPostUpdate(ctx, 1, updates)
	}
	return nil
}
