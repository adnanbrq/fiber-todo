package routing

import (
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes will apply the routes to the given app instance
func SetupRoutes(app *fiber.App) {
	apiRoutes(app)
}
