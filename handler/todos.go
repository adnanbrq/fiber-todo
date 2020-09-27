package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/adnanbrq/fiber-todo/database"
	"github.com/adnanbrq/fiber-todo/dto"
	"github.com/adnanbrq/fiber-todo/errs"
	"github.com/adnanbrq/fiber-todo/model"
	"github.com/adnanbrq/validation"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// CreateTodo will try to create a new record in the database
// containing the received data
func CreateTodo(c *fiber.Ctx) error {
	input := dto.InputCreateTodo{}
	if err := c.BodyParser(&input); err != nil {
		log.Print(err)

		return c.SendStatus(http.StatusBadRequest)
	}

	if errors := validation.Validate(input); len(errors) > 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"errors": errors})
	}

	todo := input.ToTodo()
	if err := database.DB.Create(&todo).Error; err != nil {
		log.Print(err)

		return errs.ErrDatabaseOperationFailed
	}

	return c.JSON(todo)
}

// ReadTodos will fetch all todos from the database and return them or an empty array
func ReadTodos(c *fiber.Ctx) error {
	todos := model.Todos{}

	if err := database.DB.Order("created_at desc").Find(&todos).Error; err != nil {
		log.Print(err)

		return errs.ErrDatabaseOperationFailed
	}

	return c.JSON(todos)
}

// ReadTodo will try to find the todo given by the ID
func ReadTodo(c *fiber.Ctx) error {
	todoID, err := strconv.Atoi(c.Params("id", "-1"))
	if err != nil {
		return c.SendStatus(http.StatusNotFound)
	}

	todo := model.Todo{}
	if err := database.DB.First(&todo, todoID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.SendStatus(http.StatusNotFound)
		}

		log.Print(err)
		return errs.ErrDatabaseOperationFailed
	}

	log.Print(todo.ToString())
	return c.JSON(todo)
}

// UpdateTodo will try to find the todo tied to the given ID and update its content
func UpdateTodo(c *fiber.Ctx) error {
	todoID, err := strconv.Atoi(c.Params("id", "-1"))
	if err != nil {
		return c.SendStatus(http.StatusNotFound)
	}

	input := dto.InputUpdateTodo{}
	if err := c.BodyParser(&input); err != nil {
		log.Print(err)

		return errs.ErrBadRequest
	}

	if errs := validation.Validate(input); len(errs) > 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"errors": errs})
	}

	todo := model.Todo{}
	if err := database.DB.First(&todo, todoID).Error; err != nil {
		log.Print(err)

		if err == gorm.ErrRecordNotFound {
			return c.SendStatus(http.StatusNotFound)
		}

		return errs.ErrDatabaseOperationFailed
	}

	changes := input.ToMap()

	if len(changes) == 0 {
		return c.SendStatus(http.StatusNoContent)
	}

	if err := database.DB.Model(&todo).Updates(changes).Error; err != nil {
		log.Print(err)

		return errs.ErrDatabaseOperationFailed
	}

	return c.SendStatus(http.StatusNoContent)
}

// DeleteTodo will try to delete the todo tied to the given ID
func DeleteTodo(c *fiber.Ctx) error {
	todoID, err := strconv.Atoi(c.Params("id", "-1"))
	if err != nil {
		return c.SendStatus(http.StatusNotFound)
	}

	todo := model.Todo{}
	if err := database.DB.First(&todo, todoID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.SendStatus(http.StatusNotFound)
		}

		log.Print(err)
		return errs.ErrDatabaseOperationFailed
	}

	if err := database.DB.Delete(&todo, todoID).Error; err != nil {
		log.Print(err)

		return errs.ErrDatabaseOperationFailed
	}

	return c.SendStatus(http.StatusNoContent)
}
