package middleware

import (
	"github.com/LD-Lepricon-DIgitals/delivery_backend/pkg/service"
	"github.com/gofiber/fiber/v3"
	"strconv"
	"strings"
)

type Middleware struct {
	srv *service.Service
}

func NewMiddleware(srv *service.Service) *Middleware {
	return &Middleware{srv}
}

func (m *Middleware) AuthMiddleware(c fiber.Ctx) error {
	headers := c.GetReqHeaders()

	if headers["Authorization"] == nil {
		c.Status(fiber.StatusUnauthorized)
		return c.Redirect().To("/login")
	}
	headerParts := strings.Split(headers["Authorization"][0], " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		c.Status(fiber.StatusUnauthorized)
		return c.Redirect().To("/login")
	}

	userId, err := m.srv.AuthServices.ParseToken(headerParts[0])
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.Redirect().To("/login")
	}
	c.Locals("userId", strconv.Itoa(userId))
	return c.Next()
}
