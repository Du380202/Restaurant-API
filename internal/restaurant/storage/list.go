package rstorage

import (
	"context"
	"restaurant/common"
	rmodel "restaurant/internal/restaurant/model"
)

func (s *restaurantStore) ListDataWithCondition(
	context context.Context,
	filter *rmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]rmodel.Restaurant, error) {
	var listData []rmodel.Restaurant

	db := s.db.Where("status in (1)")

	if err := db.Find(&listData).Error; err != nil {
		return nil, err
	}

	return listData, nil

}
