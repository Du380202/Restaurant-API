package biz

import (
	"context"
	"restaurant/internal/restaurantlike/model"
)

type UserLikeRestaurantStore interface {
	Create(ctx context.Context, data *model.Like) error
}
type userLikeRestaurantBiz struct {
	store UserLikeRestaurantStore
}

func NewUserLikeRestaurantBiz(store UserLikeRestaurantStore) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{store: store}
}

func (biz *userLikeRestaurantBiz) LikeRestaurant(
	ctx context.Context,
	data *model.Like,
) error {
	err := biz.store.Create(ctx, data)

	if err != nil {
		return model.ErrCannotLikeRestaurant(err)
	}

	return nil
}
