package rstorage

import (
	"context"
	rmodel "restaurant/internal/restaurant/model"
)

func (s *restaurantStore) Delete(
	context context.Context,
	id int,
) error {
	if err := s.db.Table(rmodel.Restaurant{}.TableName()).
		Where("id = ?", id).
		Updates(map[string]interface{}{"status": 0}).Error; err != nil {
		return err
	}
	return nil
}
