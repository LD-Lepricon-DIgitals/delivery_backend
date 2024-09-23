package service

import (
	"github.com/LD-Lepricon-DIgitals/delivery_backend/internal/config"
	"github.com/LD-Lepricon-DIgitals/delivery_backend/pkg/db"
)

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
	}
}
