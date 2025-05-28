package handler

import (
	"dot/app/models/dto"
	"dot/app/services"
	_ "dot/docs"
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

// Register godoc
// @Summary Register Account
// @Description Mendaftar akun
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.Register true "Register form"
// @Failure 400 {object} dto.ValidationError
// @Failure 400 {object} dto.ErrorResponse
// @Success 201 {object} dto.ErrorResponse
// @Router /register [post]
func (ac *AuthHandler) Register(c echo.Context) error {
	var req dto.Register
	if err := helpers.BindAndValidate(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	_, err := ac.authService.Register(req.Name, req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, dto.RegisterResponse{
		Message: "User registered successfully",
	})
}

// Register godoc
// @Summary Login Account
// @Description Login ke akun yang sudah ada
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.Login true "Login"
// @Failure 400 {object} dto.ValidationError
// @Failure 400 {object} dto.ErrorResponse
// @Success 200 {object} dto.LoginResponse
// @Router /login [post]
func (ac *AuthHandler) Login(c echo.Context) error {
	var req dto.Login
	if err := helpers.BindAndValidate(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	resp, err := ac.authService.Login(req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dto.LoginResponse{
		Message: "Login successful",
		Token:   resp.Token,
	})
}
