package handlers

import "github.com/gofiber/fiber/v3"

type GetUser struct {
	id int `json:"id" binding:"required"`
}

func (h *Handlers) GetUser(ctx fiber.Ctx) error {
	payload := new(GetUser)
	err := ctx.Bind().Body(payload)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	user := h.services.UserServices.
}

func (h *Handlers) RegUser(ctx fiber.Ctx) error {
	return nil
}

func (h *Handlers) LoginUser(ctx fiber.Ctx) error {
	return nil
}

func (h *Handlers) ChangeUserCity(ctx fiber.Ctx) error { return nil }

func (h *Handlers) ChangeUserLogin(ctx fiber.Ctx) error { return nil }

func (h *Handlers) ChangeUserEmail(ctx fiber.Ctx) error { return nil }

func (h *Handlers) ChangeUserPhone(ctx fiber.Ctx) error { return nil }

func (h *Handlers) ChangeUserPassword(ctx fiber.Ctx) error { return nil }

func (h *Handlers) ChangeUserCredentials(ctx fiber.Ctx) error { return nil }
