package middlewares

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func JWTConfig() echojwt.Config {
	return echojwt.Config{
		SigningKey: []byte(viper.GetString("auth_jwt_secret")),
	}
}
func JWTMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(JWTConfig())
}
