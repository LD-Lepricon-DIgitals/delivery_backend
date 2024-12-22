package db

import (
	"fmt"
	"github.com/LD-Lepricon-DIgitals/delivery_backend/internal/models"
	"github.com/jmoiron/sqlx"
)

type DishService struct {
	db *sqlx.DB
}

func (d *DishService) GetDishes() ([]models.Dish, error) {
	rows, err := d.db.Queryx("SELECT * FROM dishes")
	if err != nil {
		return nil, err
	}
	if rows.Err() != nil || rows == nil {
		return nil, fmt.Errorf("error getting dishes: %s", rows.Err())
	}
	defer rows.Close()

	var dishes []models.Dish
	for rows.Next() {
		var dish models.Dish
		if err := rows.StructScan(&dish); err != nil {
			return nil, fmt.Errorf("error getting dishes: %s", err.Error())
		}
		dishes = append(dishes, dish)
	}
	if rows.Err() != nil || rows == nil {
		return nil, fmt.Errorf("error getting dishes: %s", rows.Err())
	}
	return dishes, nil
}

func (d *DishService) DeleteDish(id int) error {
	tx, err := d.db.Begin()
	if err != nil {
		return fmt.Errorf("error creating transaction: %s", err.Error())
	}
	rows, err := tx.Exec("DELETE FROM dishes WHERE id=$1", id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error deleting dishes: %s", err.Error())
	}
	r, err := rows.RowsAffected()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error deleting dishes: %s", err.Error())
	}
	if r == 0 {
		tx.Rollback()
		return fmt.Errorf("dishes with id %d not found", id)
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("error committing dishes: %s", err.Error())
	}
	return nil
}

func (d *DishService) ChangeDish(dish models.ChangeDishPayload) error {
	tx, err := d.db.Begin()
	if err != nil {
		return fmt.Errorf("error creating transaction: %s", err.Error())
	}

	rows, err := tx.Exec("UPDATE dishes SET dish_name = $1, dish_description = $2, dish_price = $3, dish_weight = $4, dish_photo = $5, dish_rating = $6, dish_category = $7 WHERE id = $8", dish.Name, dish.Description, dish.Price, dish.Weight, dish.Photo, dish.Rating, dish.Category, dish.Id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error updating dish: %s", err.Error())
	}
	if r, _ := rows.RowsAffected(); r == 0 {
		tx.Rollback()
		return fmt.Errorf("dish with id %d not found", dish.Id)
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("error committing dishes: %s", err.Error())
	}
	return nil
}

func (d *DishService) GetDishesByCategory(category string) ([]models.Dish, error) {
	var dishes []models.Dish
	query := `
		SELECT 
			dishes.id, 
			dishes.dish_name, 
			dishes.dish_description, 
			dishes.dish_price, 
			dishes.dish_weight, 
			dishes.dish_photo, 
			dishes.dish_rating, 
			dishes.dish_category 
		FROM dishes 
		INNER JOIN dish_categories ON dishes.dish_category = dish_categories.category_name 
		WHERE dish_categories.name = $1`
	rows, err := d.db.Queryx(query, category)
	if err != nil {
		return nil, fmt.Errorf("error getting dishes: %s", err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var dish models.Dish
		if err := rows.StructScan(&dish); err != nil {
			return nil, fmt.Errorf("error getting dishes: %s", err.Error())
		}
		dishes = append(dishes, dish)
	}
	if rows.Err() != nil || rows == nil {
		return nil, fmt.Errorf("error getting dishes: %s", rows.Err())
	}
	return dishes, nil
}

func (d *DishService) GetDishById(id int) (models.Dish, error) {
	var model models.Dish
	rows, err := d.db.Queryx("SELECT * FROM dishes WHERE id=$1", id)
	if err != nil {
		return model, fmt.Errorf("error getting dish: %s", err.Error())
	}
	if rows == nil {
		return model, fmt.Errorf("dish with id %d not found", id)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.StructScan(&model)
		if err != nil {
			return model, fmt.Errorf("error getting dish: %s", err.Error())
		}
	}
	return model, nil
}

func (d *DishService) SearchByName(name string) ([]models.Dish, error) {
	var dishes []models.Dish
	query := `
		SELECT 
			d.id, d.dish_name, d.dish_description, d.dish_price, 
			d.dish_weight, d.dish_photo, d.dish_rating, c.category_name AS dish_category 
		FROM 
			dishes d
		INNER JOIN 
			dish_categories c ON d.dish_category = c.id 
		WHERE 
			d.dish_name ILIKE '%' || $1 || '%'`
	rows, err := d.db.Queryx(query, name)
	if err != nil {
		return nil, fmt.Errorf("error getting dishes: %s", err.Error())
	}
	for rows.Next() {
		var dish models.Dish
		if err := rows.StructScan(&dish); err != nil {
			return nil, fmt.Errorf("error getting dishes: %s", err.Error())
		}
		dishes = append(dishes, dish)
	}
	if rows.Err() != nil || rows == nil {
		return nil, fmt.Errorf("error getting dishes: %s", rows.Err().Error())
	}
	return dishes, nil
}

func (d *DishService) AddCategory(categoryName string) (int, error) {
	var id int
	tx, err := d.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("error creating transaction: %s", err.Error())
	}
	err = tx.QueryRow("INSERT INTO dish_categories (category_name) VALUES ($1) RETURNING id", categoryName).Scan(&id)
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("error adding category: %s", err.Error())
	}
	if err := tx.Commit(); err != nil {
		return 0, fmt.Errorf("error committing dishes: %s", err.Error())
	}
	return id, nil
}

func (d *DishService) GetCategories() ([]models.Category, error) {
	var categories []models.Category
	rows, err := d.db.Queryx("SELECT * FROM dish_categories")
	if err != nil {
		return nil, fmt.Errorf("error getting dish_categories: %s", err.Error())
	}
	for rows.Next() {
		var category models.Category
		if err := rows.StructScan(&category); err != nil {
			return nil, fmt.Errorf("error getting dish_categories: %s", err.Error())
		}
		categories = append(categories, category)
	}
	if rows.Err() != nil || rows == nil {
		return nil, fmt.Errorf("error getting categories: %s", rows.Err().Error())
	}
	return categories, nil
}

func NewDishService(db *sqlx.DB) *DishService {
	return &DishService{db: db}
}

func (d *DishService) AddDish(dish models.Dish) (int, error) {
	tx, err := d.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("Error starting transaction: %s", err)
	}
	var id int

	err = tx.QueryRow("INSERT INTO dishes (dish_name,dish_description,dish_price,dish_weight,dish_photo,dish_rating,dish_category) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id", dish.Name, dish.Description, dish.Price, dish.Weight, dish.Photo, dish.Rating, dish.Category).Scan(&id)
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("Error adding dish: %s", err)
	}
	if err = tx.Commit(); err != nil {
		return 0, fmt.Errorf("Error committing dish: %s", err)
	}
	return id, nil
}
