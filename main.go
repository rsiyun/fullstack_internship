package main

import (
	"dot/app/handler"
	"dot/app/repositories"
	"dot/app/services"
	"dot/config"
	"dot/routes"
	"log"

	_ "dot/docs"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
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

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:1323
// @BasePath /
func startServer(e *echo.Echo) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	log.Println("starting server on :1323")
	e.Logger.Fatal(e.Start(":1323"))
}
