package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// Health is used to check if the service is running
// it will just send a 200 Status code
func Health(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusOK)
}
