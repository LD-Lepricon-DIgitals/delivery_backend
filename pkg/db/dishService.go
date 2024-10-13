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

//TODO: filter by category name

func (d *DishService) GetDishes() ([]models.Dish, error) {
	var dishes []models.Dish
	query := "SELECT dishes.id, dishes.dish_name, dishes.dish_description, dishes.dish_price, dishes.dish_weight, dishes.dish_photo, dishes.dish_rating, dish_categories.category_name FROM dishes LEFT JOIN dish_categories ON dishes.dish_category=dish_categories.id;"
	rows, err := d.db.Query(query)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error getting dishes: %s", err))
	}
	defer rows.Close()

	for rows.Next() {
		var dish models.Dish
		err := rows.Scan(&dish.Id, &dish.Name, &dish.Description, &dish.Price, &dish.Weight, &dish.PhotoUrl, &dish.Rating, &dish.Category)
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

func (d *DishService) AddDish(name string, price, weight float64, description, photo string) (int, error) {
	var id int
	query := "INSERT INTO dishes (dish_name, dish_price, dish_weight, dish_description, dish_photo) VALUES ($1, $2, $3, $4, $5) RETURNING id;"
	row := d.db.QueryRow(query, name, price, weight, description, photo)
	if err := row.Scan(&id); err != nil {
		return 0, errors.New("Error adding dish: " + err.Error())
	}
	return id, nil
}

func (d *DishService) DeleteDish(id int) error {
	query := "DELETE FROM dishes WHERE id=$1;"
	res, err := d.db.Exec(query, id)
	if err != nil {
		return errors.New("Error deleting dish: " + err.Error())
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return errors.New("Error after affecting rows: " + err.Error())
	}
	if rowsAffected == 0 {
		return errors.New("no rows affected")
	}
	return nil
}

func (d *DishService) ChangeDish(id int, name string, price, weight float64, description, photo string) error {
	var res int
	checkQuery := "SELECT COUNT(1) FROM dishes WHERE id=$1;"
	updateQuery := "UPDATE dishes SET dish_name=$1, dish_price=$2, dish_weight=$3, dish_description=$4, dish_photo=$5 WHERE id=$6;"
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
	query := "SELECT dishes.dish_id, dishes.dish_name, dishes.dish_price, dishes.dish_weight, dishes.dish_description, dishes.dish_photo FROM dishes INNER JOIN dish_categories ON dishes.dish_category = dish_categories.id WHERE dish_categories.category_name=$1;"
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
	query := "SELECT dishes.id, dishes.dish_name, dishes.dish_description, dishes.dish_price, dishes.dish_weight, dishes.dish_photo, dishes.dish_rating, dish_categories.category_name FROM dishes LEFT JOIN dish_categories ON dishes.dish_categories = dish_category.id WHERE dishes.id=$1;"
	row := d.db.QueryRow(query, id)
	if err := row.Scan(&dish.Id, &dish.Name, &dish.Description, &dish.Price, &dish.Weight, &dish.PhotoUrl, &dish.Rating, &dish.Category); err != nil {
		return dish, errors.New("Error getting dish by id: " + err.Error())
	}
	return dish, nil
}

func (d *DishService) SearchByName(name string) ([]models.Dish, error) {
	var dishes []models.Dish
	query := "SELECT dishes.id, dishes.dish_name, dishes.dish_description, dishes.dish_price, dishes.dish_weight, dishes.dish_photo, dishes.dish_rating, dish_categories.category_name FROM dishes LEFT JOIN dish_categories ON dishes.dish_categories = dish_category.id WHERE dishes.dish_name ILIKE $1;"
	rows, err := d.db.Query(query, "%"+name+"%")
	if err != nil {
		return nil, errors.New("Error getting dishes by name: " + err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var dish models.Dish
		if err := rows.Scan(&dish.Id, &dish.Name, &dish.Description, &dish.Price, &dish.Weight, &dish.PhotoUrl, &dish.Rating, &dish.Category); err != nil {
			return nil, errors.New("Error getting values from rows: " + err.Error())
		}
		dishes = append(dishes, dish)
	}
	if err := rows.Err(); err != nil {
		return nil, errors.New("Error iterating rows: " + err.Error())
	}
	return dishes, nil
}
