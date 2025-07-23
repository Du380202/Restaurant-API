package biz

import (
	"context"
	"restaurant/common"
	"restaurant/internal/user/model"
)

type RegisterStorage interface {
	FindUser(context context.Context, condition map[string]interface{}, moreInfo ...string) (*model.User, error)
	Create(context context.Context, data *model.UserCreate) error
}

type Hasher interface {
	Hash(data string) string
}

type registerBussiness struct {
	registerStorage RegisterStorage
	hasher          Hasher
}

func NewRegisterBussiness(registerStorage RegisterStorage, hasher Hasher) *registerBussiness {
	return &registerBussiness{
		registerStorage: registerStorage,
		hasher:          hasher,
	}
}

func (biz *registerBussiness) Register(ctx context.Context, data *model.UserCreate) error {
	user, _ := biz.registerStorage.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if user != nil {
		return model.ErrEmailExisted
	}

	salt := common.GenSalt(50)

	data.Password = biz.hasher.Hash(data.Password + salt)
	data.Salt = salt
	data.Role = "user"
	data.Status = 1

	if err := biz.registerStorage.Create(ctx, data); err != nil {
		return err

	}

	return nil
}
