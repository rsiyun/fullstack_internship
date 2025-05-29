package helpers

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ErrorValidation struct {
	Message string   `json:"message"`
	Errors  []string `json:"errors"`
}

func BindAndValidate(c echo.Context, req interface{}) *ErrorValidation {
	if err := c.Bind(req); err != nil {
		return &ErrorValidation{
			Message: "Invalid request format",
			Errors:  []string{err.Error()},
		}
	}
	if err := c.Validate(req); err != nil {
		validationErrors, ok := err.(validator.ValidationErrors)
		if !ok {
			return &ErrorValidation{
				Message: "Validation failed",
				Errors:  []string{err.Error()},
			}
		}
		var errors []string
		for _, e := range validationErrors {
			errors = append(errors, fmt.Sprintf("Field '%s': %s", e.Field(), e.Error()))
		}

		return &ErrorValidation{
			Message: "Validation failed",
			Errors:  errors,
		}
	}

	// Jika tidak ada error, kembalikan nil
	return nil
}
