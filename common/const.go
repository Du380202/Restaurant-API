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

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}

func AppRecover() {
	if err := recover(); err != nil {
		log.Println("Recover error:", err)
	}
}
