package storage

import (
	"context"
	"restaurant/common"
	"restaurant/internal/user/model"
)

func (s *userStore) Create(context context.Context, data *model.UserCreate) error {
	db := s.db.Begin()

	if err := db.Table(data.TableName()).Create(data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil

}