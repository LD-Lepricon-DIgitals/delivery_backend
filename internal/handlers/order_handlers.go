package handlers

import (
	"fmt"
	"github.com/LD-Lepricon-DIgitals/delivery_backend/internal/models"
	"github.com/gofiber/fiber/v3"
	"strconv"
)

func (h *Handlers) CreateOrder(ctx fiber.Ctx) error {
	userId, _, err := verifyUserToken(ctx)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}
	var payload models.CreateOrder

	err = ctx.Bind().Body(&payload)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if payload.CustomerId != 0 && userId != payload.CustomerId {
		return fiber.NewError(fiber.StatusForbidden, "userId in token isn`t equal to customerId in order")
	}
	if payload.CustomerId == 0 {
		payload.CustomerId = userId
	}
	err = h.services.CreateOrder(payload)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func (h *Handlers) GetOrders(ctx fiber.Ctx) error {
	userId, role, err := verifyUserToken(ctx)

	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}

	if role != "worker" {
		return fiber.NewError(fiber.StatusForbidden, fmt.Sprintf("got %s role, expected worker", role))
	}

	ordersInfo, err := h.services.GetOrders(userId)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(ordersInfo)
}

func (h *Handlers) GetOrderDetails(ctx fiber.Ctx) error {
	param := ctx.Params("order_id")
	orderId, err := strconv.Atoi(param)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if orderId <= 0 {
		return fiber.NewError(fiber.StatusBadRequest, "invalid orderId")
	}
	_, role, err := verifyUserToken(ctx)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}
	if role != "worker" {
		return fiber.NewError(fiber.StatusForbidden, fmt.Sprintf("got %s role, expected worker", role))
	}

	orderDetails, err := h.services.GetOrderDetails(orderId)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.Status(fiber.StatusOK).JSON(orderDetails)
}

func (h *Handlers) ConfirmOrder(ctx fiber.Ctx) error {
	param := ctx.Params("order_id")
	orderId, err := strconv.Atoi(param)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if orderId <= 0 {
		return fiber.NewError(fiber.StatusBadRequest, "invalid orderId")
	}

	workerId, role, err := verifyUserToken(ctx)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}
	if role != "worker" {
		return fiber.NewError(fiber.StatusForbidden, fmt.Sprintf("got %s role, expected worker", role))
	}

	err = h.services.ConfirmOrder(orderId, workerId)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.SendStatus(fiber.StatusOK)
}

func (h *Handlers) FinishOrder(ctx fiber.Ctx) error {
	param := ctx.Params("order_id")
	orderId, err := strconv.Atoi(param)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if orderId <= 0 {
		return fiber.NewError(fiber.StatusBadRequest, "invalid orderId")
	}

	workerId, role, err := verifyUserToken(ctx)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}
	if role != "worker" {
		return fiber.NewError(fiber.StatusForbidden, fmt.Sprintf("got %s role, expected worker", role))
	}

	err = h.services.FinishOrder(orderId, workerId)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return ctx.SendStatus(fiber.StatusOK)
}
