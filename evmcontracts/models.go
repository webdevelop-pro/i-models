package evmcontracts

import (
	"context"

	"github.com/pkg/errors"
	"github.com/webdevelop-pro/go-common/db"
	"github.com/webdevelop-pro/go-common/logger"
	"github.com/webdevelop-pro/go-common/queue/pclient"
	"github.com/webdevelop-pro/i-models/models"
	"github.com/webdevelop-pro/i-models/pgtype"
)

// Wallet is an object representing the database table.
type Contract struct {
	ID      int `db:"id" json:"id" yaml:"id"`
	UserID  int `db:"user_id" json:"user_id" yaml:"user_id"`
	OfferID int `db:"offer_id" json:"offer_id" yaml:"offer_id"`

	Status        string `db:"status" json:"status" yaml:"status"`
	Address       string `db:"address" json:"address" yaml:"address"`
	TransactionTX string `db:"transaction_tx" json:"transaction_tx" yaml:"transaction_tx"`

	CreatedAt pgtype.Timestamptz `db:"created_at" json:"created_at" yaml:"created_at"`
	UpdatedAt pgtype.Timestamptz `db:"updated_at" json:"updated_at" yaml:"updated_at"`

	updatedFields []string      `db:"-" json:"-"`
	db            db.Repository `db:"-" json:"-"`
}

func (model Contract) GetField(name string) any {
	switch name {
	case "ID":
		return model.ID
	case "UserID":
		return model.UserID
	case "OfferID":
		return model.OfferID
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

func (model Contract) GetValueByTag(name string) any {
	switch name {
	case "id":
		return model.ID
	case "user_id":
		return model.UserID
	case "offer_id":
		return model.OfferID
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

func (model Contract) ToJSON() map[string]any {
	res := map[string]any{}
	fields := model.Fields()
	for _, key := range fields {
		res[key] = model.GetField(key)
	}
	return res
}

func (model Contract) Fields() []string {
	return models.DefaultFields(&model)
}

func (model Contract) Table() string {
	return TableName
}

func (model Contract) GetID() any {
	return model.ID
}

func (model *Contract) SetID(id any) {
	model.ID = id.(int)
}

func (model Contract) Save(ctx context.Context, postUpdate func(ctx context.Context, msg pclient.Event) error) error {
	if model.ID == 0 {
		err := errors.Errorf("%s: Contract %d", models.ErrIDEmpty, model.ID)
		logger.FromCtx(ctx, pkgName).Error().Stack().Err(err).Msg(models.ErrIDEmpty)
		return err
	}

	updates := map[string]any{}
	for _, field := range model.updatedFields {
		updates[field] = model.GetValueByTag(field)
	}
	updated, err := models.Update[Contract](
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
		err := errors.Errorf("%s: Contract %d", models.ErrNotUpdated, model.ID)
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
