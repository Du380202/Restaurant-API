package rmodel

import (
	"restaurant/common"
)

type RestaurantType string

const TypeNormal RestaurantType = "normal"
const TypePremium RestaurantType = "premium"
const EntityName = "restaurants"

type Restaurant struct {
	common.SQLmodel
	Name      string             `gorm:"column:name" json:"name"`
	Addr      string             `gorm:"column:addr" json:"addr"`
	Logo      *common.Image      `gorm:"column:logo" json:"logo"`
	Cover     *common.Images     `gorm:"column:cover" json:"cover"`
	UserId    int                `gorm:"column:owner_id" json:"-"`
	User      *common.SimpleUser `gorm:"preload:false" json:"user"`
	LikeCount int                `gorm:"column:like_count" json:"like_count"`
}

func (Restaurant) TableName() string { return EntityName }

func (r *Restaurant) Mask(isAdmin bool) {
	r.GenUID(common.DbTypeRestaurant)

	if u := r.User; u != nil {
		u.Mask(isAdmin)
	}
}

type RestaurantCreate struct {
	common.SQLmodel
	Name   string         `gorm:"column:name" json:"name"`
	Addr   string         `gorm:"column:addr" json:"addr"`
	Type   RestaurantType `gomr:"column:type" json:"type"`
	UserId int            `gorm:"column:owner_id" json:"-"`
	Logo   *common.Image  `gorm:"column:logo" json:"logo"`
	Cover  *common.Images `gorm:"column:cover" json:"cover"`
}

func (RestaurantCreate) TableName() string { return Restaurant{}.TableName() }

type RestaurantUpdate struct {
	Name  *string        `gorm:"column:name" json:"name"`
	Addr  *string        `gorm:"column:addr" json:"addr"`
	Logo  *common.Image  `gorm:"column:logo" json:"logo"`
	Cover *common.Images `gorm:"column:cover" json:"cover"`
}

func (RestaurantUpdate) TableName() string { return Restaurant{}.TableName() }
