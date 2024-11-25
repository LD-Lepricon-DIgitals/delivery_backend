package middleware

import (
	"github.com/LD-Lepricon-DIgitals/delivery_backend/pkg/service"
	"github.com/gofiber/fiber/v3"
	"log"
)

type Middleware struct {
	srv *service.Service
}

func NewMiddleware(srv *service.Service) *Middleware {
	return &Middleware{srv}
}

func (m *Middleware) UserAuthMiddleware(c fiber.Ctx) error {
	token := c.Cookies("token")
	if token == "" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	userId, role, err := m.srv.AuthServices.ParseToken(token)
	if role != "user" && role != "worker" && role != "admin" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	log.Println(userId)
	c.Locals("userId", userId)
	c.Locals("role", role)
	return c.Next()
}
