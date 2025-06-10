package rstorage

import (
	"gorm.io/gorm"
)

type restaurantStore struct {
	db *gorm.DB
}

func NewRestaurantStore(db *gorm.DB) *restaurantStore {
	return &restaurantStore{db: db}
}
