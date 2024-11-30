package db

import (
	"github.com/LD-Lepricon-DIgitals/delivery_backend/internal/models"
	"github.com/jmoiron/sqlx"
)

type UserServices interface {
	CreateUser(user models.UserReg) (int, error)
	GetUserId(login string) (int, error)
	IsCorrectPassword(login string, password string) (bool, error)
	IfUserExists(login string) (bool, error)
	ChangeUserCredentials(id int, info models.ChangeUserCredsPayload) error
	ChangePassword(id int, password string) error //14
	DeleteUser(id int) error
	IsCorrectPasswordId(id int, passwordToCheck string) (bool, error)
	GetUserInfo(id int) (models.UserInfo, error)
	UpdatePhoto(photoString string, userId int) error
}

type AdminServices interface {
}

type WorkerServices interface {
	ConfirmOrder(orderId int, workerId int) error
	DeclineOrder(orderId int) error
}

type OrderServices interface {
	GetWorkerOrders(workerId int) ([]models.Order, error)
	CreateOrder(order models.CreateOrder) (int, error)
	GetOrder(orderId int) (models.Order, error)
	DeleteOrder(orderId int) error
	GetUsersOrders(int) ([]models.Order, error)
	GetOrderCustomer(orderId int) (int, error)
}

type DishServices interface {
	AddDish(name string, price, weight float64, description, photo string, category int) (int, error)
	GetDishes() ([]models.Dish, error)
	DeleteDish(id int) error
	ChangeDish(id int, name string, price, weight float64, description, photo string, category int) error
	GetDishesByCategory(category string) ([]models.Dish, error)
	GetDishById(id int) (models.Dish, error)
	SearchByName(name string) ([]models.Dish, error)

	//TODO: Dish Categories return
}

type ReviewServices interface {
	PostReview(int, string) error
}
type Repository struct {
	UserServices
	AdminServices
	DishServices
	ReviewServices
	WorkerServices
	OrderServices
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		UserServices:   NewUserService(db),
		DishServices:   NewDishService(db),
		WorkerServices: NewWorkerService(db),
		OrderServices:  NewOrderService(db),
	}
}
