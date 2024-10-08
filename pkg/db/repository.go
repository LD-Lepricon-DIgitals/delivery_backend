package db

import (
	"github.com/LD-Lepricon-DIgitals/delivery_backend/internal/models"
	"github.com/jmoiron/sqlx"
)

type UserServices interface {
	CreateUser(login, name, surname, address, phoneNumber, password string) (int, error)
	GetUserId(login string) (int, error)
	IsCorrectPassword(login, password string) (bool, error)
	IfUserExists(login string) (bool, error)
	ChangeUserCredentials(id int, login, name, surname, address string) error
	ChangePassword(id, password string) error //14
	//TODO: ChangePhoto
}
type WorkerServices interface {
}

type AdminServices interface {
}

type DishServices interface {
	AddDish(name string, price, weight float64, description, photo string) error
	GetDishes() ([]models.Dish, error)
	DeleteDish(id int) error
	ChangeDish(id int, name string, price, weight float64, description, photo string) error
	GetDishesByCategory(category string) ([]models.Dish, error)
}

type ReviewServices interface {
}
type Repository struct {
	UserServices
	WorkerServices
	AdminServices
	DishServices
	ReviewServices
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		UserServices: NewUserService(db),
		DishServices: NewDishService(db),
	}
}
