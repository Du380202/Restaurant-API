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

// type LikeRestaurantStore interface {
// 	GetRestaurantLikes(ctx context.Context, ids []int) (map[int]int, error)
// }

type listRestaurantBiz struct {
	store ListRestaurantStore
	// likeStore LikeRestaurantStore
}

func NewListRestaurantBiz(store ListRestaurantStore) *listRestaurantBiz {
	return &listRestaurantBiz{store: store}
}

func (biz *listRestaurantBiz) ListRestaurant(
	context context.Context,
	filter *rmodel.Filter,
	paging *common.Paging,
) ([]rmodel.Restaurant, error) {
	result, err := biz.store.ListDataWithCondition(context, filter, paging, "User")

	if err != nil {
		return nil, err
	}

	// ids := make([]int, len(result))

	// for i := range result {
	// 	ids[i] = result[i].Id
	// }

	// likeMap, err := biz.likeStore.GetRestaurantLikes(context, ids)

	// if err != nil {
	// 	log.Println(err)
	// 	return result, nil
	// }

	// for i, item := range result {
	// 	result[i].LikeCount = likeMap[item.Id]
	// }

	return result, nil
}
