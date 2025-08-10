package biz

import (
	"context"
	"log"
	"restaurant/common"
	"restaurant/internal/restaurantlike/model"
	"time"
)

type UserUnlikeRestaurantStore interface {
	Delete(context context.Context, userId, restaurantId int) error
}

type DecreaseLikeCountStore interface {
	IncreaseLikeCount(c context.Context, id int) error
}

type userUnlikeRestaurantBiz struct {
	store    UserUnlikeRestaurantStore
	decrease DecreaseLikeCountStore
}

func NewUserUnLikeRestaurantBiz(store UserUnlikeRestaurantStore, decrease DecreaseLikeCountStore) *userUnlikeRestaurantBiz {
	return &userUnlikeRestaurantBiz{store: store, decrease: decrease}
}

func (biz *userUnlikeRestaurantBiz) UnlikeRestaurant(ctx context.Context, userId, restaurantId int) error {
	err := biz.store.Delete(ctx, userId, restaurantId)

	if err != nil {
		return model.ErrCannotUnlikeRestaurant(err)
	}

	go func() {
		defer common.AppRecover()
		time.Sleep(time.Second * 3)
		if err := biz.decrease.IncreaseLikeCount(ctx, restaurantId); err != nil {
			log.Println(err)
		}
	}()

	return nil
}
