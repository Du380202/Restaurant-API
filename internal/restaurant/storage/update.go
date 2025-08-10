package rstorage

import (
	"context"
	"fmt"
	"restaurant/common"
	rmodel "restaurant/internal/restaurant/model"

	"gorm.io/gorm"
)

func (s *restaurantStore) IncreaseLikeCount(c context.Context, id int) error {
	db := s.db
	fmt.Println("Done")
	if err := db.Table(rmodel.Restaurant{}.TableName()).Where("id = ?", id).
		Update("like_count", gorm.Expr("like_count + ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *restaurantStore) DecreaseLikeCount(c context.Context, id int) error {
	db := s.db

	if err := db.Table(rmodel.Restaurant{}.TableName()).Where("id = ?", id).
		Update("like_count", gorm.Expr("like_count - ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
