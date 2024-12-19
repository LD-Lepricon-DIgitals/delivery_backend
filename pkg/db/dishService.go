package db

import (
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

// Get all dishes with their categories
func (d *DishService) GetDishes() ([]models.Dish, error) {
	var dishes []models.Dish
	err := d.db.Select(&dishes, `
		SELECT 
			d.id, d.dish_name, d.dish_description, d.dish_price, 
			d.dish_weight, d.dish_photo, d.dish_rating, c.category_name AS dish_category
		FROM 
			dishes d 
		LEFT JOIN 
			dish_categories c ON d.dish_category = c.id`)
	if err != nil {
		return nil, fmt.Errorf("error getting dishes: %s", err.Error())
	}
	return dishes, nil
}

// Add a new dish
func (d *DishService) AddDish(dish models.Dish) (int, error) {
	tx, err := d.db.Beginx()
	if err != nil {
		return 0, fmt.Errorf("failed to start transaction: %s", err)
	}
	var id int
	err = tx.QueryRowx(`
		INSERT INTO dishes 
			(dish_name, dish_description, dish_price, dish_weight, dish_photo, dish_rating, dish_category) 
		VALUES 
			(:dish_name, :dish_description, :dish_price, :dish_weight, :dish_photo, :dish_rating, 
			(SELECT id FROM dish_categories WHERE category_name = :dish_category))
		RETURNING id`, dish).Scan(&id)
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("error adding dish: %s", err.Error())
	}
	if err := tx.Commit(); err != nil {
		return 0, fmt.Errorf("error committing transaction: %s", err.Error())
	}
	return id, nil
}

// Delete a dish by ID
func (d *DishService) DeleteDish(id int) error {
	tx, err := d.db.Beginx()
	if err != nil {
		return fmt.Errorf("failed to start transaction: %s", err.Error())
	}
	res, err := tx.Exec(`DELETE FROM dishes WHERE id = $1`, id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error deleting dish: %s", err.Error())
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		tx.Rollback()
		return fmt.Errorf("dish not found")
	}
	return tx.Commit()
}

// Update an existing dish
func (d *DishService) ChangeDish(dish models.Dish) error {
	tx, err := d.db.Beginx()
	if err != nil {
		return fmt.Errorf("failed to start transaction: %s", err.Error())
	}
	res, err := tx.Exec(`
		UPDATE dishes 
		SET 
			dish_name = :dish_name, dish_price = :dish_price, dish_weight = :dish_weight, 
			dish_description = :dish_description, dish_photo = :dish_photo, 
			dish_category = (SELECT id FROM dish_categories WHERE category_name = :dish_category)
		WHERE id = :id`, dish)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error updating dish: %s", err.Error())
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		tx.Rollback()
		return fmt.Errorf("dish not found or no changes made")
	}
	return tx.Commit()
}

// Get dishes by category name
func (d *DishService) GetDishesByCategory(category string) ([]models.Dish, error) {
	var dishes []models.Dish
	err := d.db.Select(&dishes, `
		SELECT 
			d.id, d.dish_name, d.dish_description, d.dish_price, 
			d.dish_weight, d.dish_photo, d.dish_rating, c.category_name AS dish_category 
		FROM 
			dishes d
		INNER JOIN 
			dish_categories c ON d.dish_category = c.id 
		WHERE 
			c.category_name = $1`, category)
	if err != nil {
		return nil, fmt.Errorf("error getting dishes by category: %s", err.Error())
	}
	return dishes, nil
}

// Get a single dish by ID
func (d *DishService) GetDishById(id int) (models.Dish, error) {
	var dish models.Dish
	err := d.db.Get(&dish, `
		SELECT 
			d.id, d.dish_name, d.dish_description, d.dish_price, 
			d.dish_weight, d.dish_photo, d.dish_rating, c.category_name AS dish_category
		FROM 
			dishes d
		LEFT JOIN 
			dish_categories c ON d.dish_category = c.id 
		WHERE 
			d.id = $1`, id)
	if err != nil {
		return dish, fmt.Errorf("error getting dish by id: %s", err.Error())
	}
	return dish, nil
}

// Add a new category
func (d *DishService) AddCategory(categoryName string) (int, error) {
	var id int
	err := d.db.QueryRow(`
		INSERT INTO dish_categories (category_name) 
		VALUES ($1) RETURNING id`, categoryName).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("error adding category: %s", err.Error())
	}
	return id, nil
}

// Get all categories
func (d *DishService) GetCategories() ([]models.Category, error) {
	var categories []models.Category
	err := d.db.Select(&categories, `SELECT * FROM dish_categories`)
	if err != nil {
		return nil, fmt.Errorf("error getting categories: %s", err.Error())
	}
	return categories, nil
}
