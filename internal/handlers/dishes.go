package handlers

import (
	"github.com/gofiber/fiber/v3"
	"log"
	"strconv"
)

// GetDishes godoc
// @Summary Get all dishes
// @Description Retrieve a list of all available dishes
// @Tags Dishes
// @Produce json
// @Success 200 {array} []Dish "List of dishes"
// @Failure 404 {object} fiber.Error "No dishes found"
// @Failure 500 {object} fiber.Error "Failed to retrieve dishes"
// @Router /dishes [get]
func (h *Handlers) GetDishes(ctx fiber.Ctx) error {
	dishes, err := h.services.GetDishes()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to get dishes")
	}
	if len(dishes) == 0 {
		return fiber.NewError(fiber.StatusNotFound, "Not found")
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

// AddDish godoc
// @Summary Add a new dish
// @Description Create a new dish with the provided details
// @Tags Dishes
// @Accept json
// @Produce json
// @Param dish body AddDishPayload true "Dish details"
// @Success 201 {object} map[string]int "Dish created successfully"
// @Failure 400 {object} fiber.Error "Invalid request body"
// @Failure 500 {object} fiber.Error "Failed to add dish"
// @Router /dishes [post]
func (h *Handlers) AddDish(ctx fiber.Ctx) error {
	var payload AddDishPayload
	err := ctx.Bind().Body(&payload)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}
	id, err := h.services.AddDish(payload.Name, payload.Price, payload.Weight, payload.Description, payload.Photo, payload.Category)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to add dish")
	}
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id": id,
	})
}

// DeleteDish godoc
// @Summary Delete a dish by ID
// @Description Remove a dish from the system by its ID
// @Tags Dishes
// @Param id path int true "Dish ID"
// @Success 200 "Dish deleted successfully"
// @Failure 400 {object} fiber.Error "Invalid dish ID"
// @Failure 500 {object} fiber.Error "Failed to delete dish"
// @Router /dishes/{id} [delete]
func (h *Handlers) DeleteDish(ctx fiber.Ctx) error {
	dishId := ctx.Params("id")
	id, err := strconv.Atoi(dishId)
	err = h.services.DeleteDish(id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to delete dish by id")
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

// ChangeDish godoc
// @Summary Update dish details
// @Description Update the details of an existing dish
// @Tags Dishes
// @Accept json
// @Produce json
// @Param dish body ChangeDishPayload true "Dish details"
// @Success 200 "Dish updated successfully"
// @Failure 400 {object} fiber.Error "Invalid request body"
// @Failure 500 {object} fiber.Error "Failed to update dish"
// @Router /dishes [put]
func (h *Handlers) ChangeDish(ctx fiber.Ctx) error {
	var payload ChangeDishPayload
	err := ctx.Bind().Body(&payload)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}
	err = h.services.ChangeDish(payload.Id, payload.Name, payload.Price, payload.Weight, payload.Description, payload.Photo, payload.Category)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to change dish")
	}
	return ctx.SendStatus(fiber.StatusOK)
}

type GetDishesByCategoryPayload struct {
	Category string `json:"dish_category" binding:"required"`
}

// GetDishesByCategory godoc
// @Summary Get dishes by category
// @Description Retrieve a list of dishes based on their category
// @Tags Dishes
// @Accept json
// @Produce json
// @Param category body GetDishesByCategoryPayload true "Category details"
// @Success 200 {array} []Dish "List of dishes"
// @Failure 400 {object} fiber.Error "Invalid request body"
// @Failure 404 {object} fiber.Error "No dishes found"
// @Failure 500 {object} fiber.Error "Failed to retrieve dishes"
// @Router /dishes/category [post]
func (h *Handlers) GetDishesByCategory(ctx fiber.Ctx) error {
	var payload GetDishesByCategoryPayload
	err := ctx.Bind().Body(&payload)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}
	dishes, err := h.services.GetDishesByCategory(payload.Category)
	if err != nil {
		log.Printf(err.Error())
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to get dishes by category")
	}
	if len(dishes) == 0 {
		return fiber.NewError(fiber.StatusNotFound, "Dishes not found")
	}
	return ctx.Status(fiber.StatusOK).JSON(dishes)
}

// GetDishById godoc
// @Summary Get a dish by ID
// @Description Retrieve a specific dish by its ID
// @Tags Dishes
// @Param dish_id path int true "Dish ID"
// @Success 200 {object} Dish "Dish details"
// @Failure 400 {object} fiber.Error "Invalid dish ID"
// @Failure 404 {object} fiber.Error "Dish not found"
// @Failure 500 {object} fiber.Error "Failed to retrieve dish"
// @Router /dishes/{dish_id} [get]
func (h *Handlers) GetDishById(ctx fiber.Ctx) error {
	dishId := ctx.Params("dish_id")
	id, err := strconv.Atoi(dishId)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid parameter 'dish_id'")
	}
	dish, err := h.services.GetDishById(id)
	if err != nil {
		log.Printf("Error getting dishes: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to get dish by id")
	}
	return ctx.Status(fiber.StatusOK).JSON(dish)
}

// SearchByName godoc
// @Summary Search dishes by name
// @Description Search for dishes by their name
// @Tags Dishes
// @Param name query string true "Dish name"
// @Success 200 {array} []Dish "List of matching dishes"
// @Failure 400 {object} fiber.Error "Query parameter 'name' is required"
// @Failure 404 {object} fiber.Error "No dishes found"
// @Failure 500 {object} fiber.Error "Failed to search dishes"
// @Router /dishes/search [get]
func (h *Handlers) SearchByName(ctx fiber.Ctx) error {
	name := ctx.Query("name")
	if name == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Query parameter 'name' is required")
	}
	dishes, err := h.services.SearchByName(name)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to search by name")
	}
	if dishes == nil {
		return fiber.NewError(fiber.StatusNotFound, "Dish not found")
	}
	return ctx.Status(fiber.StatusOK).JSON(dishes)
}
