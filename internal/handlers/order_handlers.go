package handlers

import (
	"fmt"
	"github.com/LD-Lepricon-DIgitals/delivery_backend/internal/models"
	"github.com/gofiber/fiber/v3"
	"strconv"
)

// CreateOrder creates a new order.
// @Summary Create a new order
// @Description Creates an order for the authenticated user or a specified customer.
// @Tags orders
// @Accept json
// @Produce json
// @Param order body models.CreateOrder true "Order payload"
// @Success 200 {string} string "Order created successfully"
// @Failure 400 {object} models.APIError "Bad Request - Invalid payload"
// @Failure 401 {object} models.APIError "Unauthorized - Invalid or missing token"
// @Failure 403 {object} models.APIError "Forbidden - User ID in token does not match customer ID in the order"
// @Failure 500 {object} models.APIError "Internal Server Error - Order creation failed"
// @Router /api/orders [post]
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

// GetOrders retrieves all orders for the authenticated user.
// @Summary Retrieve all orders
// @Description Fetches a list of all orders associated with the authenticated user with the "worker" role.
// @Tags orders
// @Accept json
// @Produce json
// @Success 200 {array} models.OrderInfo "List of orders"
// @Failure 401 {object} models.APIError "Unauthorized - Invalid or missing token"
// @Failure 403 {object} models.APIError "Forbidden - User does not have the required role"
// @Failure 500 {object} models.APIError "Internal Server Error - Failed to fetch orders"
// @Router /api/orders [get]
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

// GetOrderDetails retrieves the details of a specific order.
// @Summary Retrieve order details
// @Description Fetches detailed information about a specific order by its ID. Accessible only to users with the "worker" role.
// @Tags orders
// @Accept json
// @Produce json
// @Param order_id path int true "Order ID"
// @Success 200 {object} models.OrderDetails "Detailed information about the order"
// @Failure 400 {object} models.APIError "Bad Request - Invalid order ID"
// @Failure 401 {object} models.APIError "Unauthorized - Invalid or missing token"
// @Failure 403 {object} models.APIError "Forbidden - User does not have the required role"
// @Failure 500 {object} models.APIError "Internal Server Error - Failed to fetch order details"
// @Router /api/orders/{order_id} [get]
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

// ConfirmOrder confirms an order by its ID.
// @Summary Confirm an order
// @Description Confirms an order, marking it as accepted by a worker. Only accessible to users with the "worker" role.
// @Tags orders
// @Accept json
// @Produce json
// @Param order_id path int true "Order ID"
// @Success 200 {string} string "Order confirmed successfully"
// @Failure 400 {object} models.APIError "Bad Request - Invalid order ID"
// @Failure 401 {object} models.APIError "Unauthorized - Invalid or missing token"
// @Failure 403 {object} models.APIError "Forbidden - User does not have the required role"
// @Failure 500 {object} models.APIError "Internal Server Error - Failed to confirm order"
// @Router /api/orders/confirm/{order_id} [post]
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

	err = h.services.StartOrder(orderId, workerId)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.SendStatus(fiber.StatusOK)
}

// FinishOrder finishes an order by its ID.
// @Summary Finish an order
// @Description Marks an order as completed by the assigned worker. Only accessible to users with the "worker" role.
// @Tags orders
// @Accept json
// @Produce json
// @Param order_id path int true "Order ID"
// @Success 200 {string} string "Order finished successfully"
// @Failure 400 {object} models.APIError "Bad Request - Invalid order ID"
// @Failure 401 {object} models.APIError "Unauthorized - Invalid or missing token"
// @Failure 403 {object} models.APIError "Forbidden - User does not have the required role"
// @Failure 500 {object} models.APIError "Internal Server Error - Failed to finish order"
// @Router /api/orders/finish/{order_id} [post]
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
