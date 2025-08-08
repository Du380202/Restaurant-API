package storage

import (
	"context"
	"restaurant/common"
	"restaurant/internal/restaurantlike/model"
)

func (s *sqlStore) Delete(
	context context.Context,
	userId, restaurantId int,
) error {
	if err := s.db.Table(model.Like{}.TableName()).
		Where("user_id = ? and restaurant_id = ?", userId, restaurantId).
		Delete(nil).
		Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
