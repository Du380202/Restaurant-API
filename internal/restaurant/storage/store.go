package rstorage

import "gorm.io/gorm"

type RestaurantStore struct {
	db *gorm.DB
}
