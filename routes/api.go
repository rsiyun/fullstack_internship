package routes

import (
	"dot/app/handler"

	"github.com/go-playground/validator/v10"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
func Api(authHandler *handler.AuthHandler, categoryHandler *handler.CategoryHandler, bookHandler *handler.BookHandler) *echo.Echo {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	gJwt := e.Group("/admin")
	gJwt.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte("secret"),
	}))

	e.POST("/register", authHandler.Register)
	e.POST("/login", authHandler.Login)

	gJwt.GET("/category", categoryHandler.GetAllCategory)
	gJwt.GET("/category/:id", categoryHandler.ShowCategory)
	gJwt.POST("/category", categoryHandler.CreateCategory)
	gJwt.PUT("/category/:id", categoryHandler.UpdateCategory)
	gJwt.DELETE("/category/:id", categoryHandler.DeleteCategory)

	gJwt.GET("/book", bookHandler.GetAllBook)
	gJwt.GET("/book/:id", bookHandler.ShowBook)
	gJwt.POST("/book", bookHandler.CreateBook)
	gJwt.PUT("/book/:id", bookHandler.UpdateBook)
	gJwt.DELETE("/book/:id", bookHandler.DeleteBook)
	gJwt.GET("/book/category/:category_id", bookHandler.GetBooksByCategory)
	gJwt.GET("/book/category/:category_id/count", bookHandler.CountBooksByCategory)
	return e
}
