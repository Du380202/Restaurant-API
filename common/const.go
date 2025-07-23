package common

import "log"

const (
	DbTypeRestaurant = 1
	DbTypeUser       = 2
	RecordNotFound   = "record not found"
)

const (
	CurentUser = "user"
)

func AppRecover() {
	if err := recover(); err != nil {
		log.Println("Recover error:", err)
	}
}
