package biz

import (
	"context"
	"errors"
	"restaurant/component/appctx"
	"restaurant/component/tokenprovider"
	"restaurant/internal/user/model"
)

type LoginStorage interface {
	FindUser(context context.Context, condition map[string]interface{}, moreInfo ...string) (*model.User, error)
}

type loginBusiness struct {
	appCtx        appctx.AppContext
	storeUser     LoginStorage
	tokenProvider tokenprovider.Provider
	hasher        Hasher
	expiry        int
}

func NewLoginBusiness(storeUser LoginStorage, tokenProvider tokenprovider.Provider, hasher Hasher,
	expiry int) *loginBusiness {
	return &loginBusiness{
		storeUser:     storeUser,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		expiry:        expiry,
	}
}

func (biz *loginBusiness) Login(ctx context.Context, data *model.UserLogin) (*tokenprovider.Token, error) {
	user, err := biz.storeUser.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if err != nil {
		return nil, err
	}

	passHashed := biz.hasher.Hash(data.Password + user.Salt)

	if user.Password != passHashed {
		return nil, errors.New("password is incorrect")
	}

	payload := tokenprovider.TokenPayload{
		UserId: user.Id,
		Role:   user.Role,
	}

	accessToken, err := biz.tokenProvider.Generate(payload, biz.expiry)

	if err != nil {
		return nil, err
	}

	/*
		refreshToken, err := biz.tokenProvider.Generate(payload, biz.tkCfg.GetRtExp())
		if err != nil {
			return nil
		}

		account := model.NewAccount(accessToken, refreshToken)
	*/

	return accessToken, nil
}
