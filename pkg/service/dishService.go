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

func (d *DishService) AddDish(dish models.Dish) (int, error) {
	id, err := d.repo.AddDish(dish)
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

func (d *DishService) ChangeDish(dish models.ChangeDishPayload) error {
	err := d.repo.ChangeDish(dish)
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
		return dish, err
	}
	return dish, nil
}

func (d *DishService) SearchByName(name string) ([]models.Dish, error) {
	dishes, err := d.repo.SearchByName(name)
	if err != nil {
		return nil, err
	}
	return dishes, nil
}

func (d *DishService) AddCategory(categoryName string) (int, error) {
	id, err := d.repo.AddCategory(categoryName)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (d *DishService) GetCategories() ([]models.Category, error) {
	categories, err := d.repo.GetCategories()
	if err != nil {
		return nil, err
	}
	return categories, nil
}
