package rmodel

type Restaurant struct {
	Id     int    `gorm:"column:id" json:"id"`
	Name   string `gorm:"column:name" json:"name"`
	Addr   string `gorm:"column:addr" json:"addr"`
	Status int    `gorm:"column:status;default:true" json:"status"`
}

func (Restaurant) TableName() string { return "restaurants" }

type RestaurantCreate struct {
	Id   int    `gorm:"column:id" json:"id"`
	Name string `gorm:"column:name" json:"name"`
	Addr string `gorm:"column:addr" json:"addr"`
}

func (RestaurantCreate) TableName() string { return Restaurant{}.TableName() }

type RestaurantUpdate struct {
	Name *string `gorm:"column:name" json:"name"`
	Addr *string `gorm:"column:addr" json:"addr"`
}

func (RestaurantUpdate) TableName() string { return Restaurant{}.TableName() }
