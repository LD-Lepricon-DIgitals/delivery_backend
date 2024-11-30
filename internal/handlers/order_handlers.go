package handlers

import (
	"github.com/LD-Lepricon-DIgitals/delivery_backend/internal/models"
	"github.com/gofiber/fiber/v3"
	"strconv"
)

func (h *Handlers) CreateOrderHandler(ctx fiber.Ctx) error {
	userId, _, err := verifyUserToken(ctx)
	if err != nil {
		return err
	}
	var payload models.CreateOrder

	err = ctx.Bind().Body(&payload)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if payload.CustomerId != userId {
		return fiber.NewError(fiber.StatusForbidden, "Customer ID is not equal user id")
	}

	orderId, err := h.services.CreateOrder(payload)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"orderId": orderId,
	})
}

func (h *Handlers) GetOrderHandler(ctx fiber.Ctx) error {
	orderId := ctx.Params("orderId")

	id, err := strconv.Atoi(orderId)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if id <= 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid order id")
	}

	order, err := h.services.GetOrder(id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(order)
}

func (h *Handlers) DeleteOrder(ctx fiber.Ctx) error {
	orderId := ctx.Params("orderId")
	id, err := strconv.Atoi(orderId)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if id <= 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid order id")
	}
	userId, _, err := verifyUserToken(ctx)
	if err != nil {
		return err
	}
	customerId, err := h.services.GetOrderCustomer(id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	if userId != customerId {
		return fiber.NewError(fiber.StatusForbidden, "Customer ID is not equal user id")
	}

	err = h.services.DeleteOrder(id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.SendStatus(fiber.StatusOK)

}

func (h *Handlers) GetUserOrders(ctx fiber.Ctx) error {
	userId := ctx.Params("userId")
	uid, _, err := verifyUserToken(ctx)
	id, err := strconv.Atoi(userId)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if id <= 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid user id")
	}

	if uid != id {
		return fiber.NewError(fiber.StatusForbidden, "User ID in params is not equal user id in token")
	}

	orders, err := h.services.GetUsersOrders(id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.Status(fiber.StatusOK).JSON(orders)
}
