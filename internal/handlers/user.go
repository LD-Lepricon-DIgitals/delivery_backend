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
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "invalid params"})
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

type LoginPayload struct {
	UserLogin    string `json:"user_login" binding:"required"`
	UserPassword string `json:"user_password" binding:"required"`
}

func (h *Handlers) LoginUser(c fiber.Ctx) error {
	userId := c.Locals("userId").(int)
	if userId != 0 {
		return c.SendStatus(fiber.StatusAccepted)
	}
	var payload LoginPayload
	err := c.Bind().Body(&payload)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid params"})
	}
	exists, err := h.services.IfUserExists(payload.UserLogin)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if !exists {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "user does not exists"})
	}
	ok, err := h.services.IsCorrectPassword(payload.UserLogin, payload.UserPassword)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "invalid user password"})
	}
	userId, err = h.services.GetUserId(payload.UserLogin)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	token, err := h.services.CreateToken(userId)
	cookie := fiber.Cookie{
		Name:  "token",
		Value: token,
	}
	cookie.Partitioned = true
	c.Cookie(&cookie)
	return c.SendStatus(fiber.StatusAccepted)
}

func (h *Handlers) ChangeUserCredentials(c fiber.Ctx) error {
	userId := c.Locals("userId").(int)
	var payload models.UserInfo
	err := c.Bind().Body(&payload)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	err = h.services.ChangeUserCredentials(userId, payload.UserLogin, payload.Name, payload.Surname, payload.Address, payload.Phone)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusOK)
}

type ChangePasswordPayload struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}

func (h *Handlers) ChangeUserPassword(c fiber.Ctx) error {
	userId := c.Locals("userId").(int)
	var payload ChangePasswordPayload
	err := c.Bind().Body(&payload)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if ok, err := h.services.IsCorrectPasswordId(userId, payload.OldPassword); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	} else if !ok {
		return (c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid user password"}))
	}
	err = h.services.ChangePassword(userId, payload.NewPassword)
	return nil
}

func (h *Handlers) LogoutUser(c fiber.Ctx) error {
	c.Locals("userId", 0)
	return c.SendStatus(fiber.StatusAccepted)
}

func (h *Handlers) DeleteUser(c fiber.Ctx) error {
	userId := c.Locals("userId").(int)
	err := h.services.DeleteUser(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	c.Locals("userId", 0)
	c.ClearCookie("token")
	return c.SendStatus(fiber.StatusOK)
}
