package db

import (
	"github.com/LD-Lepricon-DIgitals/delivery_backend/internal/models"
	"github.com/jmoiron/sqlx"
)

type UserServices interface {
	GetId(login, password string) (int, error)
	Create(email, login, password string) (int, error)
	CheckIfExists(login string) (bool, error)
	GetUserInfo(id int) (*models.UserInfo, error)
	ChangePassword(id int, password string) error
	DeleteUser(id int) error
	GetUserPass(username string) (string, error)
	AddUserAddress(id int, address string) error
	ChangeUserCredentials() error
	/*	AddUserInfo(id int, userPhone, userName, userSurname, userCity string) error*/
	/*	ChangeCity(id int, city string) error*/
	/*  ChangeLogin(id int, login string) error*/
	/*	ChangeEmail(id int, email string) error*/
	/*	ChangePhone(id int, phone string) error*/
}
type WorkerServices interface {
}

type AdminServices interface {
}

type DishServices interface {
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
	}
}
