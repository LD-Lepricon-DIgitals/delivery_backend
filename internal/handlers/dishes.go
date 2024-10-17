package handlers

import (
	"github.com/LD-Lepricon-DIgitals/delivery_backend/internal/models"
	"github.com/gofiber/fiber/v3"
)

func (h *Handlers) GetDishes(ctx fiber.Ctx) error {
	dishes, err := h.services.GetDishes()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(dishes)
}

type AddDishPayload struct {
	Name        string  `json:"dish_name" binding:"required"`
	Price       float64 `json:"dish_price" binding:"required"`
	Weight      float64 `json:"dish_weight" binding:"required"`
	Description string  `json:"dish_description" binding:"required"`
	Photo       string  `json:"dish_photo" binding:"required"`
}

func (h *Handlers) AddDish(ctx fiber.Ctx) error {
	var payload AddDishPayload
	err := ctx.Bind().Body(&payload)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	id, err := h.services.AddDish(payload.Name, payload.Price, payload.Weight, payload.Description, payload.Photo)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id": id,
	})
}

func (h *Handlers) DeleteDish(ctx fiber.Ctx) error {
	var id int
	err := ctx.Bind().Body(&id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	err = h.services.DeleteDish(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.SendStatus(fiber.StatusOK)
}

func (h *Handlers) ChangeDish(id int, name string, price, weight float64, description, photo string) error {
	return nil
}

func (h *Handlers) GetDishesByCategory(category string) ([]models.Dish, error) {
	return nil, nil
}

func (h *Handlers) GetDishById(id int) (models.Dish, error) {
	return models.Dish{}, nil
}

func (h *Handlers) SearchByName(name string) ([]models.Dish, error) {
	return nil, nil
}
