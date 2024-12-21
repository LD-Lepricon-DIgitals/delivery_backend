package db

import (
	"database/sql"
	"fmt"
	"github.com/LD-Lepricon-DIgitals/delivery_backend/internal/models"
	"github.com/jmoiron/sqlx"
)

type DishService struct {
	db *sqlx.DB
}

func (d *DishService) SearchByName(name string) ([]models.Dish, error) {
	const query = `
		SELECT 
			d.id, d.dish_name, d.dish_description, d.dish_price, 
			d.dish_weight, d.dish_photo, d.dish_rating, c.category_name AS dish_category 
		FROM 
			dishes d
		INNER JOIN 
			dish_categories c ON d.dish_category = c.id 
		WHERE 
			d.dish_name ILIKE '%' || $1 || '%'`

	// Execute the query
	rows, err := d.db.Query(query, name)
	if err != nil {
		return nil, fmt.Errorf("failed to search dishes by name '%s': %w", name, err)
	}
	defer rows.Close()

	// Iterate over the result set
	var dishes []models.Dish
	for rows.Next() {
		var dish models.Dish
		var categoryName string
		if err := rows.Scan(&dish.Id, &dish.Name, &dish.Description, &dish.Price, &dish.Weight, &dish.Photo, &dish.Rating, &categoryName); err != nil {
			return nil, fmt.Errorf("error scanning dishes: %w", err)
		}
		dish.Category = categoryName
		dishes = append(dishes, dish)
	}

	// Check for errors during iteration
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over dishes: %w", err)
	}

	return dishes, nil
}

func NewDishService(db *sqlx.DB) *DishService {
	return &DishService{db: db}
}

// Get all dishes with their categories
func (d *DishService) GetDishes() ([]models.Dish, error) {
	const query = `
		SELECT 
			d.id, d.dish_name, d.dish_description, d.dish_price, 
			d.dish_weight, d.dish_photo, d.dish_rating, c.category_name AS dish_category
		FROM 
			dishes d 
		LEFT JOIN 
			dish_categories c ON d.dish_category = c.id`

	// Execute the query
	rows, err := d.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch dishes: %w", err)
	}
	defer rows.Close()

	// Iterate over the result set
	var dishes []models.Dish
	for rows.Next() {
		var dish models.Dish
		var categoryName sql.NullString // Handles the possibility of NULL for LEFT JOIN
		if err := rows.Scan(
			&dish.Id, &dish.Name, &dish.Description, &dish.Price,
			&dish.Weight, &dish.Photo, &dish.Rating, &categoryName,
		); err != nil {
			return nil, fmt.Errorf("error scanning dishes: %w", err)
		}

		// Handle NULL category_name gracefully
		if categoryName.Valid {
			dish.Category = categoryName.String
		} else {
			dish.Category = "Uncategorized" // Default value if category is NULL
		}

		dishes = append(dishes, dish)
	}

	// Check for errors during iteration
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over dishes: %w", err)
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
func (d *DishService) ChangeDish(dish models.ChangeDishPayload) error {
	// Begin the transaction
	tx, err := d.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	// Update the dish
	res, err := tx.Exec(`
		UPDATE dishes 
		SET 
			dish_name = $1, dish_price = $2, dish_weight = $3, 
			dish_description = $4, dish_photo = $5, 
			dish_category = (SELECT id FROM dish_categories WHERE category_name = $6)
		WHERE id = $7`,
		dish.Name, dish.Price, dish.Weight, dish.Description, dish.Photo, dish.Category, dish.Id,
	)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update dish: %w", err)
	}

	// Check affected rows
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		tx.Rollback()
		return fmt.Errorf("dish with id %d not found or no changes made", dish.Id)
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (d *DishService) GetDishesByCategory(category string) ([]models.Dish, error) {
	const query = `
		SELECT 
			d.id, d.dish_name, d.dish_description, d.dish_price, 
			d.dish_weight, d.dish_photo, d.dish_rating, c.category_name AS dish_category 
		FROM 
			dishes d
		INNER JOIN 
			dish_categories c ON d.dish_category = c.id 
		WHERE 
			c.category_name = $1`

	// Execute the query
	rows, err := d.db.Query(query, category)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch dishes for category '%s': %w", category, err)
	}
	defer rows.Close()

	// Iterate over the result set
	var dishes []models.Dish
	for rows.Next() {
		var dish models.Dish
		var categoryName string
		if err := rows.Scan(&dish.Id, &dish.Name, &dish.Description, &dish.Price, &dish.Weight, &dish.Photo, &dish.Rating, &categoryName); err != nil {
			return nil, fmt.Errorf("error scanning dishes: %w", err)
		}
		dish.Category = categoryName
		dishes = append(dishes, dish)
	}

	// Check for errors during iteration
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over dishes: %w", err)
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
