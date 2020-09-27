package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/stretchr/testify/assert"
)

func TestErrorHandler(t *testing.T) {
	app := fiber.New(fiber.Config{
		ErrorHandler: ErrorHandler,
	})
	app.Use(recover.New())
	app.Get("/foo", func(c *fiber.Ctx) error {
		panic("Lets just panic, ok?")
	})

	req := httptest.NewRequest(http.MethodGet, "/foo", nil)
	res, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.Equal(t, fiber.MIMEApplicationJSON, res.Header.Get("Content-Type"))
}
