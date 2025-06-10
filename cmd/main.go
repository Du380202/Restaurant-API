package main

import "restaurant/config"

func main() {
	config.LoadConfig()
	config.ConnectPostgres()

	db := config.GetDB()

	sqlDB, err := db.DB()
	if err != nil {
		panic("Không thể lấy sql.DB từ GORM: " + err.Error())
	}
	defer sqlDB.Close()
}
