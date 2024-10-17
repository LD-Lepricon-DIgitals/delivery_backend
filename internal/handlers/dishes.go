package handlers

import (
	"github.com/gofiber/fiber/v3"
	"strconv"
)

func respondWithError(ctx fiber.Ctx, statusCode int, message string) error {
	return ctx.Status(statusCode).JSON(fiber.Map{"error": message})
}

func (h *Handlers) GetDishes(ctx fiber.Ctx) error {
	dishes, err := h.services.GetDishes()
	if err != nil {
		return respondWithError(ctx, fiber.StatusBadRequest, "Failed to get dishes")
	}
	return ctx.Status(fiber.StatusOK).JSON(dishes)
}

type AddDishPayload struct {
	Name        string  `json:"dish_name" binding:"required"`
	Price       float64 `json:"dish_price" binding:"required"`
	Weight      float64 `json:"dish_weight" binding:"required"`
	Description string  `json:"dish_description" binding:"required"`
	Photo       string  `json:"dish_photo" binding:"required"`
	Category    int     `json:"dish_category" binding:"required"`
}

func (h *Handlers) AddDish(ctx fiber.Ctx) error {
	var payload AddDishPayload
	err := ctx.Bind().Body(&payload)
	if err != nil {
		return respondWithError(ctx, fiber.StatusBadRequest, "Invalid request body")
	}
	id, err := h.services.AddDish(payload.Name, payload.Price, payload.Weight, payload.Description, payload.Photo, payload.Category)
	if err != nil {
		return respondWithError(ctx, fiber.StatusBadRequest, "Failed to add dish")
	}
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id": id,
	})
}

func (h *Handlers) DeleteDish(ctx fiber.Ctx) error {
	dishId := ctx.Params("id")
	id, err := strconv.Atoi(dishId)
	err = h.services.DeleteDish(id)
	if err != nil {
		return respondWithError(ctx, fiber.StatusInternalServerError, "Failed to delete dish by id")
	}
	return ctx.SendStatus(fiber.StatusOK)
}

type ChangeDishPayload struct {
	Id          int     `json:"id" binding:"required"`
	Name        string  `json:"dish_name" binding:"required"`
	Price       float64 `json:"dish_price" binding:"required"`
	Weight      float64 `json:"dish_weight" binding:"required"`
	Description string  `json:"dish_description" binding:"required"`
	Photo       string  `json:"dish_photo" binding:"required"`
	Category    int     `json:"dish_category" binding:"required"`
}

func (h *Handlers) ChangeDish(ctx fiber.Ctx) error {
	var payload ChangeDishPayload
	err := ctx.Bind().Body(&payload)
	if err != nil {
		return respondWithError(ctx, fiber.StatusBadRequest, "Invalid request body")
	}
	err = h.services.ChangeDish(payload.Id, payload.Name, payload.Price, payload.Weight, payload.Description, payload.Photo, payload.Category)
	if err != nil {
		return respondWithError(ctx, fiber.StatusInternalServerError, "Failed to change dish")
	}
	return ctx.SendStatus(fiber.StatusOK)
}

type GetDishesByCategoryPayload struct {
	Category string `json:"category" binding:"required"`
}

func (h *Handlers) GetDishesByCategory(ctx fiber.Ctx) error {
	var payload GetDishesByCategoryPayload
	err := ctx.Bind().Body(&payload)
	if err != nil {
		return respondWithError(ctx, fiber.StatusBadRequest, "Invalid request body")
	}
	dishes, err := h.services.GetDishesByCategory(payload.Category)
	if err != nil {
		return respondWithError(ctx, fiber.StatusInternalServerError, "Failed to get dishes by category")
	}
	return ctx.Status(fiber.StatusOK).JSON(dishes)
}

func (h *Handlers) GetDishById(ctx fiber.Ctx) error {
	dishId := ctx.Params("dish_id")
	id, err := strconv.Atoi(dishId)
	if err != nil {
		return respondWithError(ctx, fiber.StatusBadRequest, "Invalid parameter 'dish_id'")
	}
	dish, err := h.services.GetDishById(id)
	if err != nil {
		return respondWithError(ctx, fiber.StatusInternalServerError, "Failed to get dish by id")
	}
	return ctx.Status(fiber.StatusOK).JSON(dish)
}

func (h *Handlers) SearchByName(ctx fiber.Ctx) error {
	name := ctx.Query("name")
	if name == "" {
		return respondWithError(ctx, fiber.StatusBadRequest, "Query parameter 'name' is required")
	}
	dishes, err := h.services.SearchByName(name)
	if err != nil {
		return respondWithError(ctx, fiber.StatusInternalServerError, "Failed to search by name")
	}
	return ctx.Status(fiber.StatusOK).JSON(dishes)
}
