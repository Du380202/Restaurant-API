package rstorage

import (
	"context"
	rmodel "restaurant/internal/restaurant/model"
)

func (s *restaurantStore) FindDataWithCondition(
	context context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*rmodel.Restaurant, error) {
	var data rmodel.Restaurant

	if err := s.db.Where(condition).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
