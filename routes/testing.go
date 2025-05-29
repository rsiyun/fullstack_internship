package routes

import (
	"dot/app/handler"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, authHandler *handler.AuthHandler) {
	e.Validator = &CustomValidator{validator: validator.New()}
	v1 := e.Group("/testing")

	// Auth routes
	v1.POST("/register", authHandler.Register)
	v1.POST("/login", authHandler.Login)
}
