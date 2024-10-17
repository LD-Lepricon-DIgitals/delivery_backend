package handlers

import (
	"github.com/LD-Lepricon-DIgitals/delivery_backend/internal/models"
	"github.com/gofiber/fiber/v3"
)

func (h *Handlers) RegisterUser(c fiber.Ctx) error {
	var params models.UserReg
	err := c.Bind().Body(&params)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{})
	}
	exists, err := h.services.IfUserExists(params.UserLogin)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if exists {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "user already exists"})
	}

	userId, err := h.services.CreateUser(params.UserLogin, params.UserName, params.UserSurname, params.UserAddress, params.UserPhone, params.UserPass)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	token, err := h.services.CreateToken(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	cookie := fiber.Cookie{
		Name:  "token",
		Value: token,
	}
	cookie.Partitioned = true
	c.Cookie(&cookie)
	return nil
}

func (h *Handlers) LoginUser(c fiber.Ctx) error {
	return nil
}
