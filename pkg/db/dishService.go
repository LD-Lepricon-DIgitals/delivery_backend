package db

import (
	"errors"
	"fmt"
	"github.com/LD-Lepricon-DIgitals/delivery_backend/internal/models"
	"github.com/jmoiron/sqlx"
)

type DishService struct {
	db *sqlx.DB
}

func NewDishService(db *sqlx.DB) *DishService {
	return &DishService{db: db}
}

func (d *DishService) GetDishes() ([]models.Dish, error) {
	var dishes []models.Dish
	query := fmt.Sprintf("SELECT id, dish_name, dish_description, dish_price, dish_weight, dish_photo, dish_rating FROM dishes")
	rows, err := d.db.Query(query)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error getting dishes: %s", err))
	}
	defer rows.Close()

	for rows.Next() {
		var dish models.Dish
		err := rows.Scan(&dish.Id, &dish.Name, &dish.Description, &dish.Price, &dish.Weight, &dish.PhotoUrl, &dish.Rating)
		if err != nil {
			return nil, errors.New("Error getting dishes: " + err.Error())
		}
		dishes = append(dishes, dish)
	}
	if err := rows.Err(); err != nil {
		return nil, errors.New("Error iterating rows: " + err.Error())
	}
	return dishes, nil
}

func (d *DishService) AddDish(name string, price, weight float64, description, photo string) error {
	query := fmt.Sprintf("INSERT INTO dishes (dish_name, dish_price, dish_weight, dish_description, dish_photo) VALUES ($1, $2, $3, $4, $5);")
	_, err := d.db.Exec(query, name, price, weight, description, photo)
	if err != nil {
		return errors.New("Error adding dish: " + err.Error())
	}
	return nil
}

func (d *DishService) DeleteDish(id int) error {
	query := fmt.Sprintf("DELETE FROM dishes WHERE id=$1;")
	_, err := d.db.Exec(query, id)
	if err != nil {
		return errors.New("Error deleting dish: " + err.Error())
	}
	return nil
}

func (d *DishService) ChangeDish(id int, name string, price, weight float64, description, photo string) error {
	var res int
	checkQuery := fmt.Sprintf("SELECT COUNT(1) FROM dishes WHERE id=$1;")
	updateQuery := fmt.Sprintf("UPDATE dishes SET dish_name=$1, dish_price=$2, dish_weight=$3, dish_description=$4, dish_photo=$5 WHERE id=$6;")
	row := d.db.QueryRow(checkQuery, id)
	if err := row.Scan(&res); err != nil {
		return errors.New("Error checking dish count: " + err.Error())
	}
	if res == 1 {
		_, err := d.db.Exec(updateQuery, name, price, weight, description, photo, id)
		if err != nil {
			return errors.New("Error updating dish: " + err.Error())
		}
	}
	return nil
}

func (d *DishService) GetDishesByCategory(category string) ([]models.Dish, error) {
	var dishes []models.Dish
	query := fmt.Sprintf("SELECT dishes.dish_id, dishes.dish_name, dishes.dish_price, dishes.dish_weight, dishes.dish_description, dishes.dish_photo FROM dishes INNER JOIN dish_categories ON dishes.dish_category = dish_categories.id WHERE dish_categories.category_name=$1")
	rows, err := d.db.Query(query, category)
	if err != nil {
		return nil, errors.New("Error querying dishes: " + err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var dish models.Dish
		err := rows.Scan(&dish.Id, &dish.Name, &dish.Description, &dish.Price, &dish.Weight, &dish.PhotoUrl, &dish.Rating)
		if err != nil {
			return nil, errors.New("Error getting dishes: " + err.Error())
		}
		dishes = append(dishes, dish)
	}
	if err := rows.Err(); err != nil {
		return nil, errors.New("Error iterating rows: " + err.Error())
	}
	return dishes, nil
}

func (d *DishService) GetDishById(id int) (models.Dish, error) {
	var dish models.Dish
	query := fmt.Sprintf("SELECT * FROM dishes WHERE id=$1;")
	row := d.db.QueryRow(query, id)
	if err := row.Scan(&dish.Id, &dish.Name, &dish.Description, &dish.Price, &dish.Weight, &dish.PhotoUrl, &dish.Rating); err != nil {
		return dish, errors.New("Error getting dish by id: " + err.Error())
	}
	return dish, nil
}

func (d *DishService) SearchByName(name string) ([]models.Dish, error) {
	var dishes []models.Dish
	query := fmt.Sprintf("SELECT dishes.id, dishes.dish_name, dishes.dish_description, dishes.dish_price, dishes.dish_weight, dishes.dish_photo, dishes.dish_rating, dish_categories.category_name FROM dishes LEFT JOIN dish_categories ON dishes.dish_categories = dish_category.id WHERE dishes.dish_name ILIKE $1;")
	rows, err := d.db.Query(query, "%"+name+"%")
	if err != nil {
		return nil, errors.New("Error getting dishes by name: " + err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var dish models.Dish
		if err := rows.Scan(&dish.Id, &dish.Name, &dish.Description, &dish.Price, &dish.Weight, &dish.PhotoUrl, &dish.Rating); err != nil {
			return nil, errors.New("Error getting values from rows: " + err.Error())
		}
		dishes = append(dishes, dish)
	}
	if err := rows.Err(); err != nil {
		return nil, errors.New("Error iterating rows: " + err.Error())
	}
	return dishes, nil
}
