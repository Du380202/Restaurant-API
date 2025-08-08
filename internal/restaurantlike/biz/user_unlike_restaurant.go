package biz

import (
	"context"
	"restaurant/internal/restaurantlike/model"
)

type UserUnlikeRestaurantStore interface {
	Delete(context context.Context, userId, restaurantId int) error
}

type userUnlikeRestaurantBiz struct {
	store UserUnlikeRestaurantStore
}

func NewUserUnLikeRestaurantBiz(store UserUnlikeRestaurantStore) *userUnlikeRestaurantBiz {
	return &userUnlikeRestaurantBiz{store: store}
}

func (biz *userUnlikeRestaurantBiz) UnlikeRestaurant(ctx context.Context, userId, restaurantId int) error {
	err := biz.store.Delete(ctx, userId, restaurantId)

	if err != nil {
		return model.ErrCannotUnlikeRestaurant(err)
	}

	return nil
}
