package middleware

import (
	"fmt"
	"github.com/LD-Lepricon-DIgitals/delivery_backend/pkg/service"
	"github.com/gofiber/fiber/v3"
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
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid token: empty token")
	}
	userId, userRole, err := m.srv.AuthServices.ParseToken(token)

	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid token: cant parse token")
	}
	if userId <= 0 {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid token: invalid user id")
	}
	if userRole != "user" && userRole != "worker" && userRole != "admin" {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid token: invalid role")
	}

	c.Locals("userId", userId)
	c.Locals("userRole", userRole)
	return c.Next()
}

func (m *Middleware) AdminAuthMiddleware(c fiber.Ctx) error {
	token := c.Cookies("token")
	if token == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid token: empty token")
	}
	userId, userRole, err := m.srv.AuthServices.ParseToken(token)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid token: cant parse token")
	}
	if userId <= 0 {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid token: invalid user id")
	}
	if userRole != "admin" {
		return fiber.NewError(fiber.StatusUnauthorized, fmt.Sprintf("Invalid token: invalid role, expected admin got %s", userRole))
	}

	c.Locals("userId", userId)
	c.Locals("userRole", userRole)
	return c.Next()
}

func (m *Middleware) WorkerAuthMiddleware(c fiber.Ctx) error {
	token := c.Cookies("token")
	if token == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid token: empty token")
	}
	userId, userRole, err := m.srv.AuthServices.ParseToken(token)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid token: cant parse token")
	}
	if userId <= 0 {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid token: invalid user id")
	}
	if userRole != "worker" {
		return fiber.NewError(fiber.StatusUnauthorized, fmt.Sprintf("Invalid token: invalid role, expected worker got %s", userRole))
	}

	c.Locals("userId", userId)
	c.Locals("userRole", userRole)
	return c.Next()
}
