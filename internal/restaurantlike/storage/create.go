package storage

import (
	"context"
	"restaurant/common"
	"restaurant/internal/restaurantlike/model"
)

func (s *sqlStore) Create(context context.Context, r *model.Like) error {
	if err := s.db.Create(&r).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
