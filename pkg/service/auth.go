package service

import (
	"fmt"
	"github.com/LD-Lepricon-DIgitals/delivery_backend/internal/config"
	"github.com/LD-Lepricon-DIgitals/delivery_backend/pkg/db"
	"github.com/golang-jwt/jwt"
	"time"
)

type AuthService struct {
	cfg  *config.Config
	repo *db.Repository
}

type tokenClaims struct {
	jwt.StandardClaims
	Id int `json:"id"`
}

const (
	tokenTll = 12 * time.Hour
)

func NewAuthService(cfg *config.Config, repository *db.Repository) *AuthService {
	return &AuthService{cfg: cfg, repo: repository}
}

func (a *AuthService) CreateToken(id int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTll).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		Id: id,
	})
	return token.SignedString([]byte(a.cfg.SigningKey))
}

func (a *AuthService) ParseToken(tokenString string) (int, error) {
	token, err := jwt.ParseWithClaims(tokenString, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(a.cfg.SigningKey), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, fmt.Errorf("invalid token")
	}
	return claims.Id, nil
}
