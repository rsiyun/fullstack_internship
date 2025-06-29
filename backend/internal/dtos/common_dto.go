package dtos

import (
	"net/http"

	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Message string      `json:"message"`
	Code    int         `json:"code,omitempty"`
	Details interface{} `json:"details,omitempty"`
}
type SuccessResponse struct {
	Message string      `json:"message"`
	Code    int         `json:"code,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type ValidationError struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
}

func (e *ErrorResponse) Error() string {
	return e.Message
}

func NewErrorResponse(message string, statusCode int, detail interface{}) *ErrorResponse {
	return &ErrorResponse{
		Message: message,
		Code:    statusCode,
		Details: detail,
	}
}
func NewValidationError(err error) *ErrorResponse {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		var errValidationDetails []ValidationError
		for _, fieldErr := range validationErrors {
			errValidationDetails = append(errValidationDetails, ValidationError{
				Field: fieldErr.Field(),
				Tag:   fieldErr.ActualTag(),
			})
		}
		return &ErrorResponse{
			Message: "Validation failed",
			Code:    http.StatusBadRequest,
			Details: errValidationDetails,
		}
	}
	return &ErrorResponse{
		Message: "Invalid request format",
		Code:    http.StatusBadRequest,
		Details: err.Error(),
	}
}
