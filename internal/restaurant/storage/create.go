package rstorage

import rmodel "restaurant/internal/restaurant/model"

func (s *restaurantStore) Create(r *rmodel.Restaurant) error {
	if err := s.db.Create(&r).Error; err != nil {
		return err
	}

	return nil
}
