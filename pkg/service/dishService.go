package service

import (
	"github.com/LD-Lepricon-DIgitals/delivery_backend/internal/models"
	"github.com/LD-Lepricon-DIgitals/delivery_backend/pkg/db"
)

type DishService struct {
	repo *db.Repository
}

func NewDishService(repo *db.Repository) *DishService {
	return &DishService{
		repo: repo,
	}
}

func (d DishService) GetDishes() map[int]models.Dish {
	//TODO implement me
	panic("implement me")
}

func (d DishService) AddDish(name string, price, weight float64, description, photo string) error {
	//TODO implement me
	panic("implement me")
}

func (d DishService) DeleteDish(id int) error {
	//TODO implement me
	panic("implement me")
}

func (d DishService) ChangeDish(name string, price, weight float64, description, photo string) error {
	//TODO implement me
	panic("implement me")
}

func (d DishService) GetDishesByCategory(category string) (map[int]models.Dish, error) {
	//TODO implement me
	panic("implement me")
}
