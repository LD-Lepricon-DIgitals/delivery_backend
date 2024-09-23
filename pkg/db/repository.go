package db

import "github.com/jmoiron/sqlx"

type UserServices interface {
	GetId(username, password string) (int, error)
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
