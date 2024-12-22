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

func (o OrderService) GetOrderDetails(orderId int) (models.OrderDetails, error) {

	var order models.OrderDetails
	var orderDish models.OrderDish
	var userId int
	tx, err := o.db.Begin()

	if err != nil {
		return models.OrderDetails{}, err
	}
	err = tx.QueryRow("SELECT customer_id,order_price,order_status FROM orders WHERE id=$1", orderId).Scan(&userId, &order.OrderPrice, &order.OrderStatus)

	if err != nil {
		tx.Rollback()
		return models.OrderDetails{}, errors.New(fmt.Sprintf("failed to fetch order details: %s", err.Error()))
	}

	// getting dishes

	rows, err := tx.Query("SELECT dish_id, dish_quantity FROM order_dishes WHERE order_id = $1", orderId)

	if err != nil {
		tx.Rollback()
		return models.OrderDetails{}, errors.New(fmt.Sprintf("failed to fetch order details: %s", err.Error()))
	}
	if rows.Err() != nil {
		tx.Rollback()
		return models.OrderDetails{}, errors.New(fmt.Sprintf("failed to fetch order details: %s", err.Error()))
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&orderDish.DishId, &orderDish.Quantity)
		if err != nil {
			tx.Rollback()
			return models.OrderDetails{}, errors.New(fmt.Sprintf("failed to fetch order details: %s", err.Error()))
		}
		err = o.db.QueryRow("SELECT dish_name,dish_photo FROM dishes WHERE id=$1", orderDish.DishId).Scan(&orderDish.DishName, &orderDish.DishPhoto)
		if err != nil {
			tx.Rollback()
			return models.OrderDetails{}, errors.New(fmt.Sprintf("failed to fetch order details: %s", err.Error()))
		}
		order.Dishes = append(order.Dishes, orderDish)
	}

	if err := tx.Commit(); err != nil {
		return models.OrderDetails{}, errors.New(fmt.Sprintf("failed to fetch order details: %s", err.Error()))
	}

	return order, nil
}

func (o OrderService) CreateOrder(order models.CreateOrder) error {

	var orderId int
	tx, err := o.db.Begin()
	if err != nil {
		return errors.New("error creating order: " + err.Error())
	}

	err = tx.QueryRow("INSERT INTO orders (customer_id, order_price, order_status) VALUES ($1, $2, $3) RETURNING id", order.CustomerId, order.Price, "pending").Scan(&orderId)
	if err != nil {
		tx.Rollback()
		return errors.New("error creating order: " + err.Error())
	}

	for _, dish := range order.Dishes {
		_, err = tx.Exec("INSERT INTO order_dishes (order_id,dish_id,dish_quantity) VALUES ($1,$2,$3)", orderId, dish.DishId, dish.Quantity)
		if err != nil {
			tx.Rollback()
			return errors.New("error creating order_dishes: " + err.Error())
		}
	}

	if err = tx.Commit(); err != nil {
		return errors.New("error creating order: " + err.Error())
	}
	return nil
}

func (o OrderService) FinishOrder(orderId, workerId int) error {
	tx, err := o.db.Begin()
	if err != nil {
		return errors.New("error starting transaction: " + err.Error())
	}
	defer tx.Rollback()

	// Check if the order is assigned to the worker
	var assignedWorkerId int
	err = tx.QueryRow("SELECT worker_id FROM orders WHERE id = $1", orderId).Scan(&assignedWorkerId)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("order not found")
		}
		return errors.New("error checking worker_id: " + err.Error())
	}

	if assignedWorkerId != workerId {
		return errors.New("order is not assigned to this worker")
	}

	// Update the order status to finished
	_, err = tx.Exec("UPDATE orders SET order_status = 'finished' WHERE id = $1", orderId)
	if err != nil {
		return errors.New("error finishing order: " + err.Error())
	}

	// Commit the transaction
	if err = tx.Commit(); err != nil {
		return errors.New("error committing transaction: " + err.Error())
	}

	return nil
}

func (o OrderService) StartOrder(orderId int, workerId int) error {
	tx, err := o.db.Begin()
	if err != nil {
		return errors.New("error starting transaction: " + err.Error())
	}
	defer tx.Rollback()

	// Check if the order has a worker assigned
	var currentWorkerId *int
	err = tx.QueryRow("SELECT worker_id FROM orders WHERE id = $1", orderId).Scan(&currentWorkerId)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("order not found")
		}
		return errors.New("error checking worker_id: " + err.Error())
	}

	if currentWorkerId != nil {
		return errors.New("order already assigned to a worker")
	}

	// Update the worker_id for the order
	_, err = tx.Exec("UPDATE orders SET worker_id = $1, order_status = 'in_process' WHERE id = $2", workerId, orderId)
	if err != nil {
		return errors.New("error assigning worker to order: " + err.Error())
	}

	// Commit the transaction
	if err = tx.Commit(); err != nil {
		return errors.New("error committing transaction: " + err.Error())
	}

	return nil
}

func NewOrderService(db *sqlx.DB) *OrderService {
	return &OrderService{db: db}
}

func (o OrderService) GetOrders(workerId int) ([]models.OrderInfo, error) {
	var ordersInfo []models.OrderInfo

	// Query to fetch free orders and worker orders in a single query
	rows, err := o.db.Query(`
        SELECT id, customer_id,order_status 
        FROM orders 
        WHERE order_status = 'pending' OR (worker_id = $1 AND order_status = 'in_process')
    `, workerId)
	if err != nil {
		return nil, errors.New("error getting orders: " + err.Error())
	}
	defer rows.Close()

	// Iterate over the result set
	for rows.Next() {
		var orderInfo models.OrderInfo
		var userId int
		if err := rows.Scan(&orderInfo.OrderId, &userId, &orderInfo.OrderStatus); err != nil {
			return nil, errors.New("error scanning orders: " + err.Error())
		}

		// Retrieve user info for the customer
		userInfo, err := o.getUserInfo(userId)
		if err != nil {
			return nil, errors.New("error getting user info: " + err.Error())
		}
		orderInfo.UserLogin = userInfo.UserLogin
		orderInfo.UserPhoto = userInfo.UserLogin
		orderInfo.Address = userInfo.Address

		// Append the order info to the list
		ordersInfo = append(ordersInfo, orderInfo)
	}

	// Check for iteration errors
	if rows.Err() != nil {
		return nil, errors.New("error iterating orders: " + err.Error())
	}

	return ordersInfo, nil

}

func (o OrderService) getUserInfo(userId int) (models.UserOrderInfo, error) {
	var userInfo models.UserOrderInfo
	tx, err := o.db.Begin()
	if err != nil {
		return models.UserOrderInfo{}, err
	}

	err = tx.QueryRow("SELECT user_login FROM users WHERE id = $1", userId).Scan(&userInfo.UserLogin)
	if err != nil {
		tx.Rollback()
		return models.UserOrderInfo{}, err
	}

	err = tx.QueryRow("SELECT user_name, user_surname,user_photo,user_address FROM users_info WHERE user_id = $1", userId).Scan(&userInfo.Name, &userInfo.Surname, &userInfo.Photo, &userInfo.Address)
	if err != nil {
		tx.Rollback()
		return models.UserOrderInfo{}, err
	}

	if err := tx.Commit(); err != nil {
		return models.UserOrderInfo{}, err
	}
	return userInfo, nil
}
