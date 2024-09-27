package service

import (
	"github.com/LD-Lepricon-DIgitals/delivery_backend/internal/config"
	"github.com/LD-Lepricon-DIgitals/delivery_backend/internal/models"
	"github.com/LD-Lepricon-DIgitals/delivery_backend/pkg/db"
)

type UserServices interface {
	CreateUser(email, login, password string) (int, error)
	CheckIfExists(email string) (bool, error)
	DeleteUser(id int) error
	AddUserInfo(id int, userPhone, userName, userSurname, userCity string) error
	ChangeCity(id int, city string) error
	ChangeLogin(id int, login string) error
	ChangePassword(id int, oldPassword, newPassword string) error
	ChangeEmail(id int, email string) error
	GetById(id int) (*models.User, error)
	ChangePhone(id int, phone string) error
}

type WorkerServices interface {
}

type AdminServices interface {
}

type DishServices interface {
}

type ReviewServices interface {
}

type AuthServices interface {
	CreateToken(username, password string) (string, error)
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
	}
}
