package auth

import (
	"api-money-management/internal/dtos"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func GetUserIDFromToken(c echo.Context) (uint, *dtos.ErrorResponse) {
	token, ok := c.Get("user").(*jwt.Token)
	if !ok || token == nil {
		return 0, dtos.NewErrorResponse("Unauthorized", http.StatusUnauthorized, "User token not found")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, dtos.NewErrorResponse("Unauthorized", http.StatusUnauthorized, "User token not found")
	}
	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		return 0, dtos.NewErrorResponse("Unauthorized: User ID not found in token claims", http.StatusUnauthorized, "User ID claim missing or invalid type.")
	}
	return uint(userIDFloat), nil
}
