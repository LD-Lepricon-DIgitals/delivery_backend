package handlers

import (
	"github.com/gofiber/fiber/v3"
)

type getUser struct {
	Id int `json:"id" binding:"required"`
}

func (h *Handlers) GetUser(ctx fiber.Ctx) error {
	payload := new(getUser)
	err := ctx.Bind().Body(payload)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	user, err := h.services.UserServices.GetById(payload.Id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(user)
}

type regPayload struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Login    string `json:"login" binding:"required"`
}

func (h *Handlers) RegUser(ctx fiber.Ctx) error {
	payload := new(regPayload)
	err := ctx.Bind().Body(payload)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	id, err := h.services.UserServices.CreateUser(payload.Email, payload.Login, payload.Password)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(fiber.Map{"id": id})
}

type signInInput struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handlers) LoginUser(ctx fiber.Ctx) error {
	payload := new(signInInput)

	err := ctx.Bind().Body(payload)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	token, err := h.services.AuthServices.CreateToken(payload.Login, payload.Password)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(fiber.Map{"token": token})
}

type changeCity struct {
	City string `json:"city" binding:"required"` //TODO: test the concept
}

func (h *Handlers) ChangeUserCity(ctx fiber.Ctx) error {
	userId := ctx.Locals("userId").(int)
	payload := new(changeCity)
	err := ctx.Bind().Body(payload)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err = h.services.ChangeCity(userId, payload.City)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.SendStatus(fiber.StatusOK)

}

type changeLogin struct {
	Id    int    `json:"id" binding:"required"`
	Login string `json:"login" binding:"required"`
}

func (h *Handlers) ChangeUserLogin(ctx fiber.Ctx) error {
	payload := new(changeLogin)
	err := ctx.Bind().Body(payload)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	err = h.services.ChangeLogin(payload.Id, payload.Login)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.SendStatus(fiber.StatusOK)
}

type changeEmail struct {
	Id    int    `json:"id" binding:"required"`
	Email string `json:"email" binding:"required"`
}

func (h *Handlers) ChangeUserEmail(ctx fiber.Ctx) error {
	payload := new(changeEmail)
	err := ctx.Bind().Body(payload)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	err = h.services.ChangeEmail(payload.Id, payload.Email)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.SendStatus(fiber.StatusOK)
}

type changePhone struct {
	Id    int    `json:"id" binding:"required"`
	Phone string `json:"phone" binding:"required"`
}

func (h *Handlers) ChangeUserPhone(ctx fiber.Ctx) error {
	payload := new(changePhone)
	err := ctx.Bind().Body(payload)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err = h.services.UserServices.ChangePhone(payload.Id, payload.Phone)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.SendStatus(fiber.StatusOK)
}

type changePassword struct {
	Id  int    `json:"id" binding:"required"`
	Old string `json:"old" binding:"required"`
	New string `json:"new" binding:"required"`
}

func (h *Handlers) ChangeUserPassword(ctx fiber.Ctx) error {
	payload := new(changePassword)
	err := ctx.Bind().Body(payload)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	err = h.services.UserServices.ChangePassword(payload.Id, payload.Old, payload.New)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.SendStatus(fiber.StatusOK)
}

type DeletePayload struct {
	Id int `json:"id" binding:"required"`
}

func (h *Handlers) DeleteUser(ctx fiber.Ctx) error {
	payload := new(DeletePayload)
	err := ctx.Bind().Body(payload)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err = h.services.DeleteUser(payload.Id)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.SendStatus(fiber.StatusOK)
}
