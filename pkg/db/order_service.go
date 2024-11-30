package db

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/LD-Lepricon-DIgitals/delivery_backend/internal/models"
	"github.com/jmoiron/sqlx"
)

type OrderService struct {
	db *sqlx.DB
}

// NewOrderService creates a new instance of order service
func NewOrderService(db *sqlx.DB) *OrderService {
	return &OrderService{
		db: db,
	}
}

// CreateOrder adds a new order of type models.CreateOrder to database
func (o *OrderService) CreateOrder(order models.CreateOrder) (int, error) {
	orderId := 0

	tx, err := o.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("failed to begin transaction: %w", err)
	}

	err = tx.QueryRow("INSERT INTO orders (customer_id,order_price) VALUES ($1,$2) RETURNING order_id", order.CustomerId, order.Price).Scan(&orderId)
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("failed to insert into orders table: %w", err)
	}

	for _, dish := range order.Dishes {
		_, err = tx.Exec("INSERT INTO order_dishes (order_id,dish_id,dish_quantity) VALUES ($1,$2,$3)", orderId, dish.DishId, dish.Quantity)

		if err != nil {
			tx.Rollback()
			return 0, fmt.Errorf("failed to insert into orders table: %w", err)
		}
	}

	if err = tx.Commit(); err != nil {
		return 0, fmt.Errorf("failed to commit transaction: %w", err)
	}
	return orderId, nil

}

// GetOrder returns an order models.Order by id
func (o *OrderService) GetOrder(orderId int) (models.Order, error) {
	order := models.Order{}

	tx, err := o.db.Begin()
	if err != nil {
		return models.Order{}, fmt.Errorf("failed to begin transaction: %w", err)
	}

	err = tx.QueryRow("SELECT * FROM orders WHERE order_id = $1", orderId).Scan(&order)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Order{}, fmt.Errorf("order not found with id %d", orderId)
		}
		return models.Order{}, fmt.Errorf("failed to query orders table: %w", err)
	}

	rows, err := tx.Query("SELECT dish_id, dish_quantity FROM order_dishes WHERE order_id = $1", orderId)
	if err != nil {
		return models.Order{}, fmt.Errorf("failed to query order_dishes table: %w", err)
	}
	defer rows.Close()

	var dishes []models.OrderDish
	for rows.Next() {
		var dish models.OrderDish
		if err := rows.Scan(&dish.DishId, &dish.Quantity); err != nil {
			return models.Order{}, fmt.Errorf("failed to scan dish data: %w", err)
		}
		dishes = append(dishes, dish)
	}

	if err = rows.Err(); err != nil {
		return order, fmt.Errorf("error occurred while iterating over order_dishes rows: %w", err)
	}

	order.Dishes = dishes

	if err = tx.Commit(); err != nil {
		return models.Order{}, fmt.Errorf("failed to commit transaction: %w", err)
	}
	return order, nil
}

// GetUsersOrders returns all user`s orders
func (o *OrderService) GetUsersOrders(userId int) ([]models.Order, error) {

	var orders []models.Order
	tx, err := o.db.Begin()
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	rows, err := tx.Query("SELECT * FROM orders WHERE customer_id = $1", userId)

	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to get user orders: %w", err)
	}

	for rows.Next() {
		var order models.Order
		if err = rows.Scan(&order); err != nil {
			return nil, fmt.Errorf("failed to get order: %w", err)
		}
		orders = append(orders, order)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to get user orders: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return orders, nil
}

// DeleteOrder deletes order
func (o *OrderService) DeleteOrder(orderId int) error {
	tx, err := o.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	_, err = tx.Exec("DELETE FROM orders WHERE order_id = $1", orderId)
	if err != nil {
		return fmt.Errorf("failed to delete order: %w", err)
	}
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}
	return nil
}

// GetWorkerOrders returns all orders to worker
func (o *OrderService) GetWorkerOrders(workerId int) ([]models.Order, error) {
	var orders []models.Order
	tx, err := o.db.Begin()
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}

	rows, err := tx.Query("SELECT * FROM orders WHERE worker_id = $1", workerId)
	if err != nil {
		return nil, fmt.Errorf("failed to get worker orders: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var order models.Order
		if err = rows.Scan(&order); err != nil {
			return nil, fmt.Errorf("failed to get worker orders: %w", err)
		}

		orders = append(orders, order)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to get worker orders: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}
	return orders, nil
}

func (o *OrderService) GetOrderCustomer(orderId int) (int, error) {
	customerId := 0
	tx, err := o.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("failed to begin transaction: %w", err)
	}

	err = tx.QueryRow("SELECT customer_id FROM orders WHERE order_id = $1", orderId).Scan(&customerId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, fmt.Errorf("order not found with id %d", orderId)
		}
		return 0, fmt.Errorf("failed to query order customer table: %w", err)
	}
	if err = tx.Commit(); err != nil {
		return 0, fmt.Errorf("failed to commit transaction: %w", err)
	}
	return customerId, nil
}
