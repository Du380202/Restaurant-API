package rstorage

import rmodel "restaurant/internal/restaurant/model"

func (s *restaurantStore) CreateRestaurant(r *rmodel.Restaurant) error {
	if err := s.db.Create(&r).Error; err != nil {
		return err
	}

	return nil
}
