package db

import "github.com/jmoiron/sqlx"

type UserServices interface {
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
	return &Repository{}
}
