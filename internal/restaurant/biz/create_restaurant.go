package biz

import rmodel "restaurant/internal/restaurant/model"

type CreateRestaurantRepo interface {
	CreateRestaurant(r *rmodel.Restaurant) error
}

type createRestaurantBiz struct {
	repo CreateRestaurantRepo
}

func NewCreateRestaurantBiz(repo CreateRestaurantRepo) *createRestaurantBiz {
	return &createRestaurantBiz{repo: repo}
}

func (biz *createRestaurantBiz) CreateRestaurant(r *rmodel.Restaurant) error {
	if err := biz.repo.CreateRestaurant(r); err != nil {
		return err
	}
	return nil
}
