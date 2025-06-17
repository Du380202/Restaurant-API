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

	db := s.db.Table(rmodel.Restaurant{}.TableName()).Where("status in (1)")

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	offset := (paging.Page - 1) * paging.Limit

	if err := db.Offset(offset).
		Limit(paging.Limit).
		Order("id desc").
		Find(&listData).Error; err != nil {
		return nil, err
	}

	return listData, nil

}
