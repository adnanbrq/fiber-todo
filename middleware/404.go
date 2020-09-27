package middleware

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// NotFound will be the last called middleware which just sends a 404
func NotFound(c *fiber.Ctx) error {
	return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Resource not Found"})
}
