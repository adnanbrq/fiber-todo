package errs

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

var (
	// ErrInternal is used as the default error message for the error handler
	ErrInternal = fiber.NewError(http.StatusInternalServerError, "Internal Server Error")

	// ErrDatabaseOperationFailed in case there was an error while performing a SQL task
	ErrDatabaseOperationFailed = fiber.NewError(http.StatusInternalServerError, "Request could not be processed. Please try again later.")

	// ErrBadRequest in case fiber cannot parse the given data
	ErrBadRequest = fiber.NewError(http.StatusBadRequest, "Request could not be Processed. Data is invalid.")
)
