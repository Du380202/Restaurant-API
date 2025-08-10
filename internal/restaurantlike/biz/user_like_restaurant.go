package biz

import (
	"context"
	"log"
	"restaurant/common"
	"restaurant/internal/restaurantlike/model"
	"time"
)

type UserLikeRestaurantStore interface {
	Create(ctx context.Context, data *model.Like) error
}

type IncreaseLikeCountStore interface {
	IncreaseLikeCount(c context.Context, id int) error
}

type userLikeRestaurantBiz struct {
	store    UserLikeRestaurantStore
	increase IncreaseLikeCountStore
}

func NewUserLikeRestaurantBiz(store UserLikeRestaurantStore, increase IncreaseLikeCountStore) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{store: store, increase: increase}
}

func (biz *userLikeRestaurantBiz) LikeRestaurant(
	ctx context.Context,
	data *model.Like,
) error {
	err := biz.store.Create(ctx, data)

	if err != nil {
		return model.ErrCannotLikeRestaurant(err)
	}
	go func() {
		defer common.AppRecover()
		time.Sleep(time.Second * 3)
		if err := biz.increase.IncreaseLikeCount(ctx, data.RestaurantId); err != nil {
			log.Println(err)
		}
	}()

	return nil
}
