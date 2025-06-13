package biz

import (
	"context"
	"restaurant/common"
	rmodel "restaurant/internal/restaurant/model"
)

type ListRestaurantStore interface {
	ListDataWithCondition(
		context context.Context,
		filter *rmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]rmodel.Restaurant, error)
}

type listRestaurantBiz struct {
	store ListRestaurantStore
}

func NewListRestaurantBiz(store ListRestaurantStore) *listRestaurantBiz {
	return &listRestaurantBiz{store: store}
}

func (biz *listRestaurantBiz) ListRestaurant(
	context context.Context,
	filter *rmodel.Filter,
	paging *common.Paging,
) ([]rmodel.Restaurant, error) {
	result, err := biz.store.ListDataWithCondition(context, filter, paging)

	if err != nil {
		return nil, err
	}

	return result, nil
}
