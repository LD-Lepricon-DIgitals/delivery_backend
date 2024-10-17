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

func (d *DishService) GetDishes() ([]models.Dish, error) {
	dishes, err := d.repo.GetDishes()
	if err != nil {
		return nil, err
	}
	return dishes, nil
}

func (d *DishService) AddDish(name string, price, weight float64, description, photo string) (int, error) {
	id, err := d.repo.AddDish(name, price, weight, description, photo)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (d *DishService) DeleteDish(id int) error {
	err := d.repo.DeleteDish(id)
	if err != nil {
		return err
	}
	return nil
}

func (d *DishService) ChangeDish(id int, name string, price, weight float64, description, photo string) error {
	err := d.repo.ChangeDish(id, name, price, weight, description, photo)
	if err != nil {
		return err
	}
	return nil
}

func (d *DishService) GetDishesByCategory(category string) ([]models.Dish, error) {
	dishes, err := d.repo.GetDishesByCategory(category)
	if err != nil {
		return nil, err
	}
	return dishes, nil
}

func (d *DishService) GetDishById(id int) (models.Dish, error) {
	dish, err := d.repo.GetDishById(id)
	if err != nil {
		return dish, nil
	}
	return dish, err
}

func (d *DishService) SearchByName(name string) ([]models.Dish, error) {
	dishes, err := d.repo.SearchByName(name)
	if err != nil {
		return nil, err
	}
	return dishes, nil
}