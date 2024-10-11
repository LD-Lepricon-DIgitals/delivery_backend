package db

import (
	"github.com/LD-Lepricon-DIgitals/delivery_backend/internal/models"
	"github.com/jmoiron/sqlx"
)

type UserServices interface {
	CreateUser(login, name, surname, address, phoneNumber, password string) (int, error)
	GetUserId(login string) (int, error)
	IsCorrectPassword(id int, password string) (bool, error)
	IfUserExists(login string) (bool, error)
	ChangeUserCredentials(id int, login, name, surname, address, phone string) error
	ChangePassword(id int, password string) error //14
	DeleteUser(id int) error
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
	GetDishById(id int) (models.Dish, error)
	//TODO: GetDishById, SearchByName
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
