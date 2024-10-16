package service

import (
	"github.com/LD-Lepricon-DIgitals/delivery_backend/internal/config"
	"github.com/LD-Lepricon-DIgitals/delivery_backend/internal/models"
	"github.com/LD-Lepricon-DIgitals/delivery_backend/pkg/db"
)

type UserServices interface {
	CreateUser(login, name, surname, address, phoneNumber, password string) (int, error)
	GetUserId(login string) (int, error)
	IsCorrectPassword(id int, password string) (bool, error)
	IfUserExists(login string) (bool, error)
	ChangeUserCredentials(id int, login, name, surname, address, phone string) error
	ChangePassword(id, password string) error //14
}

type WorkerServices interface {
}

type AdminServices interface {
}

type DishServices interface {
	GetDishes() ([]models.Dish, error)
	AddDish(name string, price, weight float64, description, photo string) (int, error)
	DeleteDish(id int) error
	ChangeDish(id int, name string, price, weight float64, description, photo string) error
	GetDishesByCategory(category string) ([]models.Dish, error)
	GetDishById(id int) (models.Dish, error)
	SearchByName(name string) ([]models.Dish, error)
}

type ReviewServices interface {
}

type AuthServices interface {
	CreateToken(login, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Service struct {
	UserServices
	WorkerServices
	AdminServices
	DishServices
	ReviewServices
	AuthServices
}

func NewService(repo *db.Repository, cfg *config.Config) *Service {
	return &Service{
		AuthServices: NewAuthService(cfg, repo),
		UserServices: NewUserService(repo),
		DishServices: NewDishService(repo),
	}
}
