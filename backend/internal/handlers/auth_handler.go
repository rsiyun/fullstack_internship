package handlers

import (
	"api-money-management/internal/dtos"
	"api-money-management/internal/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Login(c echo.Context) error {
	req := new(dtos.LoginRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, dtos.ErrorResponse{
			Message: "Invalid request",
			Code:    http.StatusBadRequest,
		})
	}

	response, err := h.authService.Login(req)
	if err != nil {
		return c.JSON(err.Code, dtos.ErrorResponse{
			Message: err.Message,
			Code:    err.Code,
			Details: err.Details,
		})
	}

	return c.JSON(http.StatusOK, response)
}
func (h *AuthHandler) Register(c echo.Context) error {
	req := new(dtos.RegisterRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, dtos.ErrorResponse{
			Message: "Invalid request",
			Code:    http.StatusBadRequest,
		})
	}
	result, err := h.authService.Register(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.ErrorResponse{
			Message: "Failed to register user",
			Code:    http.StatusInternalServerError,
			Details: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, dtos.UserResponse{
		ID:        result.ID,
		Name:      result.Name,
		Email:     result.Email,
		CreatedAt: result.CreatedAt.String(),
	})
}
