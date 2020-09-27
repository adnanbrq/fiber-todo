package handler

import (
	"github.com/adnanbrq/fiber-todo/errs"
	"github.com/gofiber/fiber/v2"
)

// ErrorHandler replaces the default fiber error handler
func ErrorHandler(c *fiber.Ctx, err error) error {
	if e, ok := err.(*fiber.Error); ok {
		return c.Status(e.Code).JSON(fiber.Map{"message": e.Message})
	}

	return c.Status(errs.ErrInternal.Code).JSON(fiber.Map{"message": errs.ErrInternal.Message})
}
