package middleware

import (
	"github.com/LD-Lepricon-DIgitals/delivery_backend/pkg/service"
	"github.com/gofiber/fiber/v3"
	"log"
	"strconv"
)

type Middleware struct {
	srv *service.Service
}

func NewMiddleware(srv *service.Service) *Middleware {
	return &Middleware{srv}
}

func (m *Middleware) AuthMiddleware(c fiber.Ctx) error {
	token := c.Cookies("token")
	if token == "" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	userId, err := m.srv.AuthServices.ParseToken(token)
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	log.Println(userId)
	c.Locals("userId", strconv.Itoa(userId))
	return c.Next()
}
