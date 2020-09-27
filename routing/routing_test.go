package routing

import (
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestSetupRoutes(t *testing.T) {
	app := fiber.New()
	SetupRoutes(app)

	assert.NotEmpty(t, app.Stack())
}
