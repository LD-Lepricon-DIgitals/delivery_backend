package handlers

import (
	"fmt"
	"github.com/LD-Lepricon-DIgitals/delivery_backend/internal/models"
	"github.com/gofiber/fiber/v3"
	"log"
	"strconv"
)

// GetDishes godoc
// @Summary Get all dishes
// @Description Retrieve a list of all available dishes
// @Tags Dishes
// @Produce json
// @Success 200 {array} models.Dish "List of dishes"
// @Failure 404 {object} models.APIError "No dishes found"
// @Failure 500 {object} models.APIError "Failed to retrieve dishes"
// @Router /api/dishes [get]
func (h *Handlers) GetDishes(ctx fiber.Ctx) error {
	log.Println("Handler: GetDishes - Start")
	dishes, err := h.services.GetDishes()
	if err != nil {
		log.Printf("Handler: GetDishes - Failed to get dishes: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to get dishes")
	}
	if len(dishes) == 0 {
		log.Println("Handler: GetDishes - No dishes found")
		return fiber.NewError(fiber.StatusNotFound, "Not found")
	}
	log.Println("Handler: GetDishes - Successfully retrieved dishes")
	return ctx.Status(fiber.StatusOK).JSON(dishes)
}

// AddDish godoc
// @Summary Add a new dish
// @Description Create a new dish with the provided details
// @Tags Dishes (Admin)
// @Accept json
// @Produce json
// @Param dish body AddDishPayload true "Dish details"
// @Success 201 {object} map[string]int "Dish created successfully"
// @Failure 400 {object} models.APIError "Invalid request body"
// @Failure 403 {object} models.APIError "Access forbidden"
// @Failure 500 {object} models.APIError "Failed to add dish"
// @Router /api/dishes/admin/add [post]
func (h *Handlers) AddDish(ctx fiber.Ctx) error {
	log.Println("Handler: AddDish - Start")
	if err := validateAdmin(ctx); err != nil {
		log.Printf("Handler: AddDish - Admin validation failed: %v", err)
		return err
	}

	var payload models.Dish
	if err := ctx.Bind().JSON(&payload); err != nil {
		log.Printf("Handler: AddDish - Failed to parse request body: %v", err)
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}
	id, err := h.services.AddDish(payload)
	if err != nil {
		log.Printf("Handler: AddDish - Failed to add dish: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to add dish")
	}

	log.Printf("Handler: AddDish - Dish added successfully with ID: %d", id)
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"id": id})
}

// DeleteDish godoc
// @Summary Delete a dish by ID
// @Description Remove a dish from the system by its ID
// @Tags Dishes (Admin)
// @Param id path int true "Dish ID"
// @Success 200 {string} string "Dish deleted successfully"
// @Failure 400 {object} models.APIError "Invalid dish ID"
// @Failure 403 {object} models.APIError "Access forbidden"
// @Failure 500 {object} models.APIError "Failed to delete dish"
// @Router /api/dishes/admin/delete/{id} [delete]
func (h *Handlers) DeleteDish(ctx fiber.Ctx) error {
	log.Println("Handler: DeleteDish - Start")
	if err := validateAdmin(ctx); err != nil {
		log.Printf("Handler: DeleteDish - Admin validation failed: %v", err)
		return err
	}

	dishId := ctx.Params("id")
	log.Printf("Handler: DeleteDish - Dish ID: %s", dishId)
	id, err := strconv.Atoi(dishId)
	if err != nil {
		log.Printf("Handler: DeleteDish - Invalid dish ID: %s", dishId)
		return fiber.NewError(fiber.StatusBadRequest, "Invalid dish ID")
	}

	if err := h.services.DeleteDish(id); err != nil {
		log.Printf("Handler: DeleteDish - Failed to delete dish: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to delete dish by id")
	}

	log.Printf("Handler: DeleteDish - Dish deleted successfully with ID: %d", id)
	return ctx.SendStatus(fiber.StatusOK)
}

// ChangeDish godoc
// @Summary Update dish details
// @Description Update the details of an existing dish
// @Tags Dishes (Admin)
// @Accept json
// @Produce json
// @Param dish body ChangeDishPayload true "Dish details"
// @Success 200 {string} string "Dish updated successfully"
// @Failure 400 {object} models.APIError "Invalid request body"
// @Failure 403 {object} models.APIError "Access forbidden"
// @Failure 500 {object} models.APIError "Failed to update dish"
// @Router /api/dishes/admin/update [put]
func (h *Handlers) ChangeDish(ctx fiber.Ctx) error {
	log.Println("Handler: ChangeDish - Start")
	if err := validateAdmin(ctx); err != nil {
		log.Printf("Handler: ChangeDish - Admin validation failed: %v", err)
		return err
	}

	var payload models.ChangeDishPayload
	if err := ctx.Bind().JSON(&payload); err != nil {
		log.Printf("Handler: ChangeDish - Failed to parse request body: %v", err)
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	if payload.Id <= 0 {
		log.Printf("Handler: ChangeDish - Invalid dish ID in payload: %d", payload.Id)
		return fiber.NewError(fiber.StatusBadRequest, "Invalid id")
	}

	if err := h.services.ChangeDish(payload); err != nil {
		log.Printf("Handler: ChangeDish - Failed to change dish: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to change dish")
	}

	log.Printf("Handler: ChangeDish - Dish updated successfully with ID: %d", payload.Id)
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
// @Success 200 {array} models.Dish "List of dishes"
// @Failure 400 {object} models.APIError "Invalid request body"
// @Failure 404 {object} models.APIError "No dishes found"
// @Failure 500 {object} models.APIError "Failed to retrieve dishes"
// @Router /api/dishes/by_category [post]
func (h *Handlers) GetDishesByCategory(ctx fiber.Ctx) error {
	log.Println("Handler: GetDishesByCategory - Start")
	var payload GetDishesByCategoryPayload
	if err := ctx.Bind().JSON(&payload); err != nil {
		log.Printf("Handler: GetDishesByCategory - Failed to parse request body: %v", err)
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	log.Printf("Handler: GetDishesByCategory - Category: %s", payload.Category)
	dishes, err := h.services.GetDishesByCategory(payload.Category)
	if err != nil {
		log.Printf("Handler: GetDishesByCategory - Failed to get dishes: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to get dishes by category")
	}

	if len(dishes) == 0 {
		log.Printf("Handler: GetDishesByCategory - No dishes found for category: %s", payload.Category)
		return fiber.NewError(fiber.StatusNotFound, "Dishes not found")
	}

	log.Printf("Handler: GetDishesByCategory - Successfully retrieved dishes for category: %s", payload.Category)
	return ctx.Status(fiber.StatusOK).JSON(dishes)
}

// GetDishById godoc
// @Summary Get a dish by ID
// @Description Retrieve a specific dish by its ID
// @Tags Dishes
// @Param dish_id path int true "Dish ID"
// @Success 200 {object} models.Dish "Dish details"
// @Failure 400 {object} models.APIError "Invalid dish ID"
// @Failure 404 {object} models.APIError "Dish not found"
// @Failure 500 {object} models.APIError "Failed to retrieve dish"
// @Router /api/dishes/by_id/{dish_id} [get]
func (h *Handlers) GetDishById(ctx fiber.Ctx) error {
	log.Println("Handler: GetDishById - Start")
	dishId := ctx.Params("dish_id")
	id, err := strconv.Atoi(dishId)
	if err != nil {
		log.Printf("Handler: GetDishById - Invalid parameter 'dish_id': %s", dishId)
		return fiber.NewError(fiber.StatusBadRequest, "Invalid parameter 'dish_id'")
	}

	log.Printf("Handler: GetDishById - Fetching dish with ID: %d", id)
	dish, err := h.services.GetDishById(id)
	if err != nil {
		log.Printf("Handler: GetDishById - Error getting dish: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to get dish by id")
	}

	log.Printf("Handler: GetDishById - Successfully retrieved dish with ID: %d", id)
	return ctx.Status(fiber.StatusOK).JSON(dish)
}

// SearchByName godoc
// @Summary Search dishes by name
// @Description Search for dishes by their name
// @Tags Dishes
// @Param name query string true "Dish name"
// @Success 200 {array} models.Dish "List of matching dishes"
// @Failure 400 {object} models.APIError "Query parameter 'name' is required"
// @Failure 404 {object} models.APIError "No dishes found"
// @Failure 500 {object} models.APIError "Failed to search dishes"
// @Router /api/dishes/search [get]
func (h *Handlers) SearchByName(ctx fiber.Ctx) error {
	log.Println("Handler: SearchByName - Start")
	name := ctx.Query("name")
	if name == "" {
		log.Println("Handler: SearchByName - Query parameter 'name' is missing")
		return fiber.NewError(fiber.StatusBadRequest, "Query parameter 'name' is required")
	}

	log.Printf("Handler: SearchByName - Searching for dishes with name containing: %s", name)
	dishes, err := h.services.SearchByName(name)
	if err != nil {
		log.Printf("Handler: SearchByName - Error searching by name: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to search by name")
	}

	if dishes == nil || len(dishes) == 0 {
		log.Printf("Handler: SearchByName - No dishes found for name: %s", name)
		return fiber.NewError(fiber.StatusNotFound, "Dish not found")
	}

	log.Printf("Handler: SearchByName - Successfully found %d dishes", len(dishes))
	return ctx.Status(fiber.StatusOK).JSON(dishes)
}

func validateAdmin(ctx fiber.Ctx) error {
	userRole, ok := ctx.Locals("userRole").(string)
	if !ok {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid token: can not parse role from locals")
	}
	if userRole != "admin" {
		return fiber.NewError(fiber.StatusForbidden, fmt.Sprintf("Invalid token: invalid role, expected worker got %s", userRole))
	}
	return nil
}

type AddDishCategoryPayload struct {
	CategoryName string `json:"category_name" bindings:"required"`
}

// AddCategory godoc
// @Summary Add a new dish category
// @Description Add a new dish category by providing its name
// @Tags Categories (Admin)
// @Accept json
// @Produce json
// @Param payload body AddDishCategoryPayload true "Category data"
// @Success 201 {object} map[string]int "ID of the created category"
// @Failure 400 {object} models.APIError "Invalid request body"
// @Failure 403 {object} models.APIError "Access forbidden"
// @Failure 500 {object} models.APIError "Failed to create category"
// @Router /api/categories [post]
func (h *Handlers) AddCategory(ctx fiber.Ctx) error {
	var payload AddDishCategoryPayload
	if err := ctx.Bind().JSON(&payload); err != nil {
		log.Printf("Handler: AddCategory - Invalid request body: %v", err)
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	log.Printf("Handler: AddCategory - Adding category: %s", payload.CategoryName)
	id, err := h.services.AddCategory(payload.CategoryName)
	if err != nil {
		log.Printf("Handler: AddCategory - Failed to create category: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create category")
	}

	log.Printf("Handler: AddCategory - Successfully created category with ID: %d", id)
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"id": id})
}

// GetCategories godoc
// @Summary Get all dish categories
// @Description Retrieve a list of all dish categories
// @Tags Categories
// @Produce json
// @Success 200 {array} string "List of categories"
// @Failure 500 {object} models.APIError "Failed to retrieve categories"
// @Router /api/categories [get]
func (h *Handlers) GetCategories(ctx fiber.Ctx) error {
	log.Println("Handler: GetCategories - Start")
	categories, err := h.services.GetCategories()
	if err != nil {
		log.Printf("Handler: GetCategories - Failed to retrieve categories: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to retrieve categories")
	}

	log.Printf("Handler: GetCategories - Successfully retrieved %d categories", len(categories))
	return ctx.Status(fiber.StatusOK).JSON(categories)
}
