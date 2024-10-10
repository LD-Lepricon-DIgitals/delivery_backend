package service

import "github.com/LD-Lepricon-DIgitals/delivery_backend/pkg/db"

type DishService struct {
	repo *db.Repository
}

func NewDishService(repo *db.Repository) *DishService {
	return &DishService{
		repo: repo,
	}
}
