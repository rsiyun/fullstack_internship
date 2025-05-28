package handler

import (
	"dot/app/models/dto"
	"dot/app/services"
	"dot/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authService services.AuthServices
}

func NewAuthHandler(authService services.AuthServices) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (ac *AuthHandler) Register(c echo.Context) error {
	var req dto.Register
	if err := helpers.BindAndValidate(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	_, err := ac.authService.Register(req.Name, req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "User registered successfully",
	})
}

func (ac *AuthHandler) Login(c echo.Context) error {
	var req dto.Login
	if err := helpers.BindAndValidate(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	resp, err := ac.authService.Login(req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Login successful",
		"token":   resp.Token,
	})
}
