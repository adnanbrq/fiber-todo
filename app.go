package main

import (
	"github.com/adnanbrq/fiber-todo/handler"
	"github.com/adnanbrq/fiber-todo/routing"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// creates a new fiber app, configures it and attaches the routes
func newApp() *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: handler.ErrorHandler,
	})

	app.Use(recover.New())
	routing.SetupRoutes(app)

	return app
}
