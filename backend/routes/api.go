package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New(enableCors bool) (*echo.Echo, error) {
	e := echo.New()
	// sambungin db dulu
	// initialize semua handlernya
	SetupMiddleware(e, enableCors)
	Routes(e)
	return e, nil

}

func Routes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
}

func SetupMiddleware(e *echo.Echo, enableCors bool) {
	if enableCors {
		e.Use(middleware.CORS())
	}
}
