package service

import (
	"github.com/LD-Lepricon-DIgitals/delivery_backend/internal/config"
	"github.com/LD-Lepricon-DIgitals/delivery_backend/internal/models"
	"github.com/LD-Lepricon-DIgitals/delivery_backend/pkg/db"
)

type UserServices interface {
	GetUserId(username, password string) (int, error)
	IsCorrectPassword(login, password string) (bool, error)
	IfUserExists(username string) (bool, error)
	CreateUser(login, name, surname, address, phoneNumber, password string) (int, error)
}

type WorkerServices interface {
}

type AdminServices interface {
}

type DishServices interface {
	GetDishes() map[int]models.Dish
	AddDish(name string, price, weight float64, description, photo string) error
	DeleteDish(id int) error
	ChangeDish(name string, price, weight float64, description, photo string) error
	GetDishesByCategory(category string) (map[int]models.Dish, error)
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
