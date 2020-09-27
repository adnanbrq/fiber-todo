package routing

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestApiRoutes(t *testing.T) {
	app := fiber.New()
	apiRoutes(app)

	req := httptest.NewRequest(http.MethodGet, "/api", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
