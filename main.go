package main

import (
	"dot/app/handler"
	"dot/app/repositories"
	"dot/app/services"
	"dot/config"
	"dot/routes"
	"log"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			config.Connection,
			repositories.NewAuthRepository,
			repositories.NewCategoryRepository,
			repositories.NewBookRepository,
			services.NewUserService,
			services.NewCategoryService,
			services.NewBookService,
			handler.NewAuthHandler,
			handler.NewCatgoryHandler,
			handler.NewBookHandler,
			routes.Api,
		),
		fx.Invoke(startServer),
	)
	app.Run()
}

func startServer(e *echo.Echo) {
	log.Println("starting server on :1323")
	e.Logger.Fatal(e.Start(":1323"))
}
