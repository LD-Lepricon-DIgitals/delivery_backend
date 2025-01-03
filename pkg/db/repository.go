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
	GetUserRole(userId int) (string, error)
}

type AdminServices interface {
}

type WorkerServices interface {
	ConfirmOrder(orderId int, workerId int) error
	DeclineOrder(orderId int) error
}

type OrderServices interface {
	CreateOrder(order models.CreateOrder) error
	GetOrders(workerId int) ([]models.OrderInfo, error) //includes private methods getFreeOrders and getWorkerOrders
	FinishOrder(orderId, workerId int) error
	StartOrder(orderId int, workerId int) error //includes isFree check
	GetOrderDetails(orderId int) (models.OrderDetails, error)
}

type DishServices interface {
	AddDish(models.Dish) (int, error)
	GetDishes() ([]models.Dish, error)
	DeleteDish(id int) error
	ChangeDish(dish models.ChangeDishPayload) error
	GetDishesByCategory(category string) ([]models.Dish, error)
	GetDishById(id int) (models.Dish, error)
	SearchByName(name string) ([]models.Dish, error)
	AddCategory(categoryName string) (int, error)
	GetCategories() ([]models.Category, error)
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
