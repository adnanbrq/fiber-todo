package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestNotFound(t *testing.T) {
	app := fiber.New()
	app.Use(NotFound)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, res.StatusCode)
	assert.Equal(t, fiber.MIMEApplicationJSON, res.Header.Get("Content-Type"))
}
