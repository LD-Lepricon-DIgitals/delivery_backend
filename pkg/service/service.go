package service

import "github.com/LD-Lepricon-DIgitals/delivery_backend/pkg/db"

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
type Service struct {
	UserServices
	WorkerServices
	AdminServices
	DishServices
	ReviewServices
}

func NewService(repo *db.Repository) *Service {
	return &Service{}
}
