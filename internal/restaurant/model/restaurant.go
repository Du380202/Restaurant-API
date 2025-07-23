package rmodel

import "restaurant/common"

type Restaurant struct {
	common.SQLmodel
	Name  string         `gorm:"column:name" json:"name"`
	Addr  string         `gorm:"column:addr" json:"addr"`
	Logo  *common.Image  `gorm:"column:logo" json:"logo"`
	Cover *common.Images `gorm:"column:cover" json:"cover"`
}

func (Restaurant) TableName() string { return "restaurants" }

func (r *Restaurant) Mask(isAdmin bool) {
	r.GenUID(common.DbTypeRestaurant)
}

type RestaurantCreate struct {
	common.SQLmodel
	Name  string         `gorm:"column:name" json:"name"`
	Addr  string         `gorm:"column:addr" json:"addr"`
	Logo  *common.Image  `gorm:"column:logo" json:"logo"`
	Cover *common.Images `gorm:"column:cover" json:"cover"`
}

func (RestaurantCreate) TableName() string { return Restaurant{}.TableName() }

type RestaurantUpdate struct {
	Name  *string        `gorm:"column:name" json:"name"`
	Addr  *string        `gorm:"column:addr" json:"addr"`
	Logo  *common.Image  `gorm:"column:logo" json:"logo"`
	Cover *common.Images `gorm:"column:cover" json:"cover"`
}

func (RestaurantUpdate) TableName() string { return Restaurant{}.TableName() }
