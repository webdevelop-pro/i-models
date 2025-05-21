package evmwallets

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
type Wallet struct {
	ID     int `db:"id" json:"id" yaml:"id"`
	UserID int `db:"user_id" json:"user_id" yaml:"user_id"`

	PublicKey  string `db:"public_key" json:"public_key" yaml:"public_key"`
	PrivateKey string `db:"private_key" json:"private_key" yaml:"private_key"`

	Balance    float64            `db:"balance" json:"balance" yaml:"balance"`
	IncBalance float64            `db:"inc_balance" json:"inc_balance" yaml:"inc_balance"`
	OutBalance float64            `db:"out_balance" json:"out_balance" yaml:"out_balance"`
	Status     WalletStatusT      `db:"status" json:"status" yaml:"status"`
	CreatedAt  pgtype.Timestamptz `db:"created_at" json:"created_at" yaml:"created_at"`
	UpdatedAt  pgtype.Timestamptz `db:"updated_at" json:"updated_at" yaml:"updated_at"`

	ObjectID      string `db:"-" json:"object_id" yaml:"object_id"`
	ContentTypeID int    `db:"-" json:"content_type_id" yaml:"content_type_id"`

	updatedFields []string      `db:"-" json:"-"`
	db            db.Repository `db:"-" json:"-"`
}

func (model Wallet) GetField(name string) any {
	switch name {
	case "ID":
		return model.ID
	case "UserID":
		return model.UserID
	case "PublicKey":
		return model.PublicKey
	case "PrivateKey":
		return model.PrivateKey
	case "Balance":
		return model.Balance
	case "IncBalance":
		return model.IncBalance
	case "OutBalance":
		return model.OutBalance
	case "Status":
		return model.Status
	case "CreatedAt":
		return model.CreatedAt
	case "UpdatedAt":
		return model.UpdatedAt
	}
	return nil
}

func (model Wallet) GetValueByTag(name string) any {
	switch name {
	case "id":
		return model.ID
	case "user_id":
		return model.UserID
	case "public_key":
		return model.PublicKey
	case "private_key":
		return model.PrivateKey
	case "balance":
		return model.Balance
	case "inc_balance":
		return model.IncBalance
	case "out_balance":
		return model.OutBalance
	case "status":
		return model.Status
	case "created_at":
		return model.CreatedAt
	case "updated_at":
		return model.UpdatedAt
	}
	return nil
}

func (model Wallet) ToJSON() map[string]any {
	res := map[string]any{}
	fields := model.Fields()
	for _, key := range fields {
		res[key] = model.GetField(key)
	}
	return res
}

func (model Wallet) Fields() []string {
	return models.DefaultFields(&model)
}

func (model Wallet) Table() string {
	return TableName
}

func (model Wallet) GetID() any {
	return model.ID
}

func (model *Wallet) SetID(id any) {
	model.ID = id.(int)
}

func (model Wallet) Save(ctx context.Context, postUpdate func(ctx context.Context, msg pclient.Event) error) error {
	if model.ID == 0 {
		err := errors.Errorf("%s: Wallet %d", models.ErrIDEmpty, model.ID)
		logger.FromCtx(ctx, pkgName).Error().Stack().Err(err).Msg(models.ErrIDEmpty)
		return err
	}

	updates := map[string]any{}
	for _, field := range model.updatedFields {
		updates[field] = model.GetValueByTag(field)
	}
	updated, err := models.Update[Wallet](
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
		err := errors.Errorf("%s: Wallet %d", models.ErrNotUpdated, model.ID)
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
