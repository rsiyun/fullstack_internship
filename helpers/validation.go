package helpers

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func BindAndValidate(c echo.Context, req interface{}) echo.Map {
	if err := c.Bind(req); err != nil {
		return echo.Map{
			"message": "Invalid request",
		}
	}
	if err := c.Validate(req); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errors := make(map[string]string)

		for _, e := range validationErrors {
			errors[e.Field()] = e.Error()
		}

		return echo.Map{
			"message": "Validation failed",
			"errors":  errors,
		}
	}

	return nil
}
