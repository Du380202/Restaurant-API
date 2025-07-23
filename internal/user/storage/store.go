package storage

import "gorm.io/gorm"

type userStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) *userStore {
	return &userStore{db: db}
}
