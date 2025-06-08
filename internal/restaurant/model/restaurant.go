package rmodel

type Restaurant struct {
	Id   int    `gorm:"column:id" json:"id"`
	Name string `gorm:"column:name" json:"name"`
	Addr string `gorm:"column:addr" json:"addr"`
}

func (Restaurant) TableName() string { return "restaurants" }

type RestaurantUpdate struct {
	Name *string `gorm:"column:name" json:"name"`
	Addr *string `gorm:"column:addr" json:"addr"`
}

func (RestaurantUpdate) TableName() string { return Restaurant{}.TableName() }
