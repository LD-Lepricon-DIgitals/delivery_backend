package db

import (
	"github.com/LD-Lepricon-DIgitals/delivery_backend/internal/models"
	"github.com/jmoiron/sqlx"
)

type UserServices interface {
	GetId(login, password string) (int, error)
	Create(email, login, password string) (int, error)
	CheckIfExists(email string) (bool, error)
	GetById(id int) (*models.User, error)
	ChangeCity(id int, city string) error
	ChangeLogin(id int, login string) error
	ChangePassword(id int, password string) error
	ChangeEmail(id int, email string) error
	DeleteUser(id int) error
	ChangePhone(id int, phone string) error
	GetUserPass(username string) (string, error)
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
