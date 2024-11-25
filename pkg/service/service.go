package service

import (
	"github.com/LD-Lepricon-DIgitals/delivery_backend/internal/config"
	"github.com/LD-Lepricon-DIgitals/delivery_backend/internal/handlers"
	"github.com/LD-Lepricon-DIgitals/delivery_backend/internal/models"
	"github.com/LD-Lepricon-DIgitals/delivery_backend/pkg/db"
)

type UserServices interface {
	CreateUser(user models.UserReg) (int, error)
	GetUserId(login string) (int, error)
	IsCorrectPassword(login string, password string) (bool, error)
	IfUserExists(login string) (bool, error)
	ChangeUserCredentials(id int, info models.ChangeUserCredsPayload) error
	ChangePassword(id int, password string) error
	IsCorrectPasswordId(id int, passwordToCheck string) (bool, error)
	DeleteUser(id int) error //14
	GetUserInfo(id int) (models.UserInfo, error)
	UpdatePhoto(photo string, userId int) error
}

type WorkerServices interface {
	ConfirmOrder(orderId int, workerId int) error
	DeclineOrder(orderId int) error
}

type AdminServices interface {
}

type DishServices interface {
	GetDishes() ([]models.Dish, error)
	AddDish(name string, price, weight float64, description, photo string, category int) (int, error)
	DeleteDish(id int) error
	ChangeDish(id int, name string, price, weight float64, description, photo string, category int) error
	GetDishesByCategory(category string) ([]models.Dish, error)
	GetDishById(id int) (models.Dish, error)
	SearchByName(name string) ([]models.Dish, error)
}

type ReviewServices interface {
}

type AuthServices interface {
	CreateToken(id int, role string) (string, error)
	ParseToken(token string) (int, string, error)
}

type Service struct {
	UserServices
	AdminServices
	WorkerServices
	DishServices
	ReviewServices
	AuthServices
}

func NewService(repo *db.Repository, cfg *config.Config) *Service {
	return &Service{
		AuthServices:   NewAuthService(cfg, repo),
		UserServices:   NewUserService(repo),
		DishServices:   NewDishService(repo),
		WorkerServices: NewWorkerService(repo),
	}
}
