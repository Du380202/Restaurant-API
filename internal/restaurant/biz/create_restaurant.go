package biz

import rmodel "restaurant/internal/restaurant/model"

type CreateRestaurantRepo interface {
	Create(r *rmodel.Restaurant) error
}

type createRestaurantBiz struct {
	repo CreateRestaurantRepo
}

func NewCreateRestaurantBiz(repo CreateRestaurantRepo) *createRestaurantBiz {
	return &createRestaurantBiz{repo: repo}
}

func (biz *createRestaurantBiz) CreateRestaurant(r *rmodel.Restaurant) error {
	if err := biz.repo.Create(r); err != nil {
		return err
	}
	return nil
}
