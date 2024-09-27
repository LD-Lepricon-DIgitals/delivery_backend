package middleware

import (
	"github.com/LD-Lepricon-DIgitals/delivery_backend/pkg/service"
	"github.com/gofiber/fiber/v3"
	"log"
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
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	headerParts := strings.Split(headers["Authorization"][0], " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	userId, err := m.srv.AuthServices.ParseToken(headerParts[1])
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	log.Println(userId)
	c.Locals("userId", strconv.Itoa(userId))
	return c.Next()
}
