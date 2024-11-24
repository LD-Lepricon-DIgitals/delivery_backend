package handlers

import (
	"errors"
	"github.com/gofiber/fiber/v3"
)

func CustomError(c fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}
	c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
	return c.Status(code).JSON(fiber.Map{
		"message": err.Error(),
	})
}
