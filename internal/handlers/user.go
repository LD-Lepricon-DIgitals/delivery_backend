package handlers

import (
	"errors"
	"fmt"
	"github.com/LD-Lepricon-DIgitals/delivery_backend/internal/models"
	"github.com/gofiber/fiber/v3"
	"log"
)

// RegisterUser registers a new user.
// @Summary Register a new user
// @Description Registers a new user and returns a token in a cookie
// @Tags auth
// @Accept json
// @Param request body models.UserReg true "User Registration"
// @Success 200 "Token in cookie"
// @Failure 400 {object models.APIError "Invalid request data"
// @Failure 409 {object} models.APIError "User already exists"
// @Router /auth/register [post]
func (h *Handlers) RegisterUser(c fiber.Ctx) error {
	var params models.UserReg
	err := c.Bind().JSON(&params)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request data")
	}
	exists, err := h.services.IfUserExists(params.UserLogin)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	if exists {
		return fiber.NewError(fiber.StatusConflict, "User already exists")
	}

	userId, err := h.services.CreateUser(params)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	token, err := h.services.CreateToken(userId, params.Role)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	cookie := &fiber.Cookie{
		Name:        "token",
		Value:       token,
		Partitioned: true,
		SameSite:    "None",
		Secure:      true,
	}

	c.Cookie(cookie)
	return nil
}

type LoginPayload struct {
	UserLogin    string `json:"user_login" validate:"required"`
	UserPassword string `json:"user_password" validate:"required"`
}

// LoginUser logs in an existing user.
// @Summary Log in a user
// @Description Logs in a user and returns a token in a cookie
// @Tags auth
// @Accept json
// @Produce json
// @Param request body LoginPayload true "User Login Credentials"
// @Success 200
// @Failure 400 {object} models.APIError "Invalid request data"
// @Failure 401 {object} models.APIError "Invalid credentials"
// @Router /auth/login [post]
func (h *Handlers) LoginUser(c fiber.Ctx) error {
	token := c.Cookies("token")
	if token != "" {
		userId, role, err := h.services.AuthServices.ParseToken(token)
		if err == nil {

			c.Locals("userId", userId)
			c.Locals("role", role)
			return c.SendStatus(fiber.StatusOK)
		}
		// Если токен не валиден, продолжаем с обычной авторизацией
	}
	var payload LoginPayload
	err := c.Bind().JSON(&payload)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request data")
	}
	exists, err := h.services.IfUserExists(payload.UserLogin)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	if !exists {
		return fiber.NewError(fiber.StatusNotFound, "User not exists")
	}
	ok, err := h.services.IsCorrectPassword(payload.UserLogin, payload.UserPassword)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	if !ok {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid user password")
	}
	userId, err := h.services.GetUserId(payload.UserLogin)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	userRole, err := h.services.GetUserRole(userId)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	token, err = h.services.CreateToken(userId, userRole)
	cookie := fiber.Cookie{
		Name:        "token",
		Value:       token,
		Partitioned: true,
		SameSite:    "None",
		Secure:      true,
	}
	cookie.Partitioned = true
	c.Cookie(&cookie)
	return c.SendStatus(fiber.StatusOK)
}

// ChangeUserCredentials updates a user's personal details.
// @Summary Update user credentials
// @Description Allows the logged-in user to update their details
// @Tags user
// @Accept json
// @Produce json
// @Param request body models.ChangeUserCredsPayload true "Updated User Creds"
// @Success 200
// @Failure 400 {object} models.APIError "Invalid request data"
// @Failure 401 {object} models.APIError "Unauthorized"
// @Router /api/user/change [patch]
func (h *Handlers) ChangeUserCredentials(c fiber.Ctx) error {
	userId, _, err := verifyUserToken(c)
	if err != nil {
		return err
	}
	var payload models.ChangeUserCredsPayload
	err = c.Bind().JSON(&payload)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request data")
	}
	err = h.services.ChangeUserCredentials(userId, payload)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.SendStatus(fiber.StatusOK)
}

type ChangePasswordPayload struct {
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required"`
}

// ChangeUserPassword changes a user's password.
// @Summary Change user password
// @Description Allows the logged-in user to change their password
// @Tags user
// @Accept json
// @Produce json
// @Param request body ChangePasswordPayload true "Old and New Password"
// @Success 200 {string} string "Password changed successfully"
// @Failure 400 {object} models.APIError "Invalid request data"
// @Failure 401 {object} models.APIError "Invalid old password"
// @Router /api/user/change_password [patch]
func (h *Handlers) ChangeUserPassword(c fiber.Ctx) error {
	userId, _, err := verifyUserToken(c)
	if err != nil {
		return err
	}
	var payload ChangePasswordPayload
	err = c.Bind().JSON(&payload)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request data")
	}
	if ok, err := h.services.IsCorrectPasswordId(userId, payload.OldPassword); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	} else if !ok {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid password")
	}
	err = h.services.ChangePassword(userId, payload.NewPassword)
	return nil
}

// LogoutUser logs out the currently logged-in user.
// @Summary Logout user
// @Description Logs out the currently logged-in user by clearing the authentication token cookie.
// @Tags auth
// @Produce json
// @Success 200
// @Router /api/user/logout [post]
func (h *Handlers) LogoutUser(c fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:   "token",
		Value:  "",
		MaxAge: -1,
	}
	c.Cookie(&cookie)
	c.Locals("userId", nil)
	c.Locals("role", nil)
	return c.SendStatus(fiber.StatusOK)
}

// DeleteUser deletes a user's account.
// @Summary Delete user account
// @Description Deletes the logged-in user's account
// @Tags user
// @Success 200
// @Failure 401 {object} models.APIError "Unauthorized"
// @Failure 500 {object} models.APIError "Internal server error"
// @Router /api/user/delete [delete]
func (h *Handlers) DeleteUser(c fiber.Ctx) error {
	userId, _, err := verifyUserToken(c)
	if err != nil {
		return err
	}
	err = h.services.DeleteUser(userId)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	cookie := fiber.Cookie{
		Name:   "token",
		Value:  "",
		MaxAge: -1,
	}
	c.Locals("userId", nil)
	c.Locals("role", nil)
	c.Cookie(&cookie)
	return c.SendStatus(fiber.StatusOK)
}

// GetUserInfo retrieves information about the logged-in user.
// @Summary Get user info
// @Description Retrieves the details of the logged-in user
// @Tags user
// @Produce json
// @Success 200 {object} models.UserInfo "User information"
// @Failure 401 {object} models.APIError "Unauthorized"
// @Router /api/user/info [get]
func (h *Handlers) GetUserInfo(c fiber.Ctx) error {
	userId, _, err := verifyUserToken(c)
	if err != nil {
		return err
	}
	user, err := h.services.GetUserInfo(userId)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

type ChangePhotoPayload struct {
	PhotoString string `json:"photo" validate:"required"`
}

// UpdatePhoto updates a user's profile photo.
// @Summary Update user photo
// @Description Allows the logged-in user to update their profile photo
// @Tags user
// @Accept json
// @Produce json
// @Param request body ChangePhotoPayload true "Photo Data"
// @Success 200
// @Failure 400 {object} models.APIError "Invalid request data"
// @Failure 401 {object} models.APIError "Unauthorized"
// @Router /api/user/photo [patch]
func (h *Handlers) UpdatePhoto(c fiber.Ctx) error {
	userId, _, err := verifyUserToken(c)
	if err != nil {
		return err
	}
	var payload ChangePhotoPayload
	err = c.Bind().JSON(&payload)
	if err != nil {
		log.Println(fmt.Sprintf("error: %s", err.Error()))
		return fiber.NewError(fiber.StatusBadRequest, errors.New("invalid request body").Error())
	}
	err = h.services.UpdatePhoto(payload.PhotoString, userId)
	if err != nil {
		log.Println(fmt.Sprintf("error: %s", err.Error()))
		return fiber.NewError(fiber.StatusInternalServerError, errors.New("failed to update photo").Error())
	}
	return c.SendStatus(fiber.StatusOK)
}

func verifyUserToken(c fiber.Ctx) (int, string, error) {
	userId, ok := c.Locals("userId").(int)
	if userId <= 0 || !ok {
		return 0, "", fiber.NewError(fiber.StatusUnauthorized, errors.New("invalid user id").Error())
	}

	userRole, ok := c.Locals("userRole").(string)
	if (userRole != "admin" && userRole != "user" && userRole != "worker") || !ok {
		return 0, "", fiber.NewError(fiber.StatusUnauthorized, errors.New("invalid user role").Error())
	}

	return userId, "", nil
}
