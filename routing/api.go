package routing

import (
	"github.com/adnanbrq/fiber-todo/handler"
	"github.com/adnanbrq/fiber-todo/middleware"
	"github.com/gofiber/fiber/v2"
)

// attach api routes
func apiRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/", handler.Health)

	todos := api.Group("/todos")
	todos.Post("/", handler.CreateTodo)
	todos.Get("/", handler.ReadTodos)
	todos.Get("/:id", handler.ReadTodo)
	todos.Put("/:id", handler.UpdateTodo)
	todos.Delete("/:id", handler.DeleteTodo)

	api.Use(middleware.NotFound)
}
