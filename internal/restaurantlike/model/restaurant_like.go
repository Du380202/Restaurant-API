package model

import (
	"fmt"
	"restaurant/common"
	"time"
)

const EntityName = "UserLikeRestaurant"

type Like struct {
	RestaurantId int                `gorm:"column:restaurant_id" json:"restaurant_id"`
	UserId       int                `gorm:"column:user_id" json:"user_id"`
	CreatedAt    *time.Time         `gorm:"columg:created_at" json:"created_at"`
	User         *common.SimpleUser `gorm:"preload:false" json:"user"`
}

func (Like) TableName() string { return "restaurant_likes" }

func (l *Like) GetRestaurantId() int {
	return l.RestaurantId
}

func ErrCannotLikeRestaurant(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("Cannot like this restaurant"),
		fmt.Sprintf("ErrCannotLikeRestaurant"),
	)
}

func ErrCannotUnlikeRestaurant(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("Cannot unlike this restaurant"),
		fmt.Sprintf("ErrCannotUnlikeRestaurant"),
	)
}
