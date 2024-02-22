package model

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// UserResponse struct to define custom response
type UserResponse struct {
	Status  uint32     `json:"status"`
	Message string     `json:"message"`
	Data    *fiber.Map `json:"data"`
}

// BadResponse response the error with 400 status code
func BadResponse(c *fiber.Ctx, err error) error {
	return c.Status(http.StatusBadRequest).JSON(UserResponse{
		Status:  http.StatusBadRequest,
		Message: "error",
		Data:    &fiber.Map{"data": err.Error()},
	})
}

// InternalServerErrorResponse response the error with 500 status code
func InternalServerErrorResponse(c *fiber.Ctx, err error) error {
	return c.Status(http.StatusInternalServerError).JSON(UserResponse{
		Status:  http.StatusInternalServerError,
		Message: "error",
		Data:    &fiber.Map{"data": err.Error()},
	})
}

// OkResponse response the data with 200 status code
func OkResponse(c *fiber.Ctx, data interface{}) error {
	return c.Status(http.StatusOK).JSON(UserResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    &fiber.Map{"data": data},
	})
}

// CreatedResponse response the data with 201 status code
func CreatedResponse(c *fiber.Ctx, data interface{}) error {
	return c.Status(http.StatusCreated).JSON(UserResponse{
		Status:  http.StatusCreated,
		Message: "success",
		Data:    &fiber.Map{"data": data},
	})
}
