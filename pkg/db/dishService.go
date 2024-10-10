package db

import (
	"github.com/LD-Lepricon-DIgitals/delivery_backend/internal/models"
	"github.com/jmoiron/sqlx"
)

type DishService struct {
	db *sqlx.DB
}

func NewDishService(db *sqlx.DB) *DishService {
	return &DishService{db: db}
}

/*func (d *DishService) GetAll() ([]models.Dish, error) {
	query := fmt.Sprintf("SELECT id, dish_name, dish_description, dish_price, dish_weight, dish_photo, dish_rating FROM dishes")
	rows, err := d.db.Query(query)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error getting dishes: %s", err))
	}
	var dishes []models.Dish

	for rows.Next() {
		var dish models.Dish
		err := rows.Scan(&dish.Id, &dish.Name, &dish.Description, &dish.Price, &dish.Weight, &dish.PhotoUrl, &dish.Rating)
		if err != nil {
			return nil, errors.New("Error getting dishes: " + err.Error())
		}
		dishes = append(dishes, dish)
	}

	return dishes, nil
}*/

func (d *DishService) GetDishes() (map[int]models.Dish, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DishService) AddDish(name string, price, weight float64, description, photo string) error {
	return nil
}

func (d *DishService) DeleteDish(id int) error { return nil }

func (d *DishService) ChangeDish(name string, price, weight float64, description, photo string) error {
	return nil
}

func (d *DishService) GetDishesByCategory(category string) (map[int]models.Dish, error) {
	return nil, nil
}
