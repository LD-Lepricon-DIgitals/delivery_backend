package service

import (
	"github.com/LD-Lepricon-DIgitals/delivery_backend/internal/config"
	"github.com/LD-Lepricon-DIgitals/delivery_backend/pkg/db"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAuthService_CreateToken(t *testing.T) {
	cfg := &config.Config{SigningKey: "test_signing_key"}
	repo := &db.Repository{}
	authService := NewAuthService(cfg, repo)

	userId := 1
	token, err := authService.CreateToken(userId)
	assert.NoError(t, err, "No errors are expected")
	assert.NotEmpty(t, token, "token cant be empty")

	// Тестуємо парсинг токена
	parsedId, err := authService.ParseToken(token)
	assert.NoError(t, err, "No errors are expected")
	assert.Equal(t, userId, parsedId, "Id has to be equal")
}

func TestAuthService_ParseToken_InvalidToken(t *testing.T) {
	cfg := &config.Config{SigningKey: "test_signing_key"}
	repo := &db.Repository{}
	authService := NewAuthService(cfg, repo)

	invalidToken := "invalid.token.string"
	parsedId, err := authService.ParseToken(invalidToken)
	assert.Error(t, err, "Error expected")
	assert.Equal(t, 0, parsedId, "Id has to be 0")
}

func TestAuthService_ParseToken_ExpiredToken(t *testing.T) {
	cfg := &config.Config{SigningKey: "test_signing_key"}
	repo := &db.Repository{}
	authService := NewAuthService(cfg, repo)

	claims := &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(-time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		Id: 1,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(cfg.SigningKey))
	assert.NoError(t, err, "No errors expected")

	parsedId, err := authService.ParseToken(signedToken)
	assert.Error(t, err, "Error expected while parsing outdated token")
	assert.Equal(t, 0, parsedId, "Id has to be 0")
}
