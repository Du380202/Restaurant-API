package biz

import (
	"context"
	"restaurant/common"
	"restaurant/internal/restaurantlike/model"
)

type ListUserLikeRestaurantStore interface {
	GetUserLikeRestaurant(
		c context.Context,
		conditions map[string]interface{},
		filter *model.Filter,
		paging *common.Paging,
	) ([]common.SimpleUser, error)
}

type listUserLikeRestaurantBiz struct {
	store ListUserLikeRestaurantStore
}

func NewListUserLikeRestaurant(store ListUserLikeRestaurantStore) *listUserLikeRestaurantBiz {
	return &listUserLikeRestaurantBiz{store: store}
}

func (biz *listUserLikeRestaurantBiz) ListUser(
	context context.Context,
	filter *model.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]common.SimpleUser, error) {
	users, err := biz.store.GetUserLikeRestaurant(context, nil, filter, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(model.EntityName, err)
	}

	return users, nil
}
