package routes

import (
	"api-money-management/internal/middlewares"
	"api-money-management/pkg/common"
	"api-money-management/pkg/database"
	"api-money-management/pkg/utils"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New(enableCors bool) (*echo.Echo, error) {
	e := echo.New()
	e.Validator = utils.NewGlobalValidator()
	// sambungin db dulu
	db, err := database.DBConn()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	// initialize semua handlernya
	handler, err := common.InjectDependencies(db)
	if err != nil {
		log.Fatalf("Failed to inject dependencies: %v", err)

	}
	SetupMiddleware(e, enableCors)
	Routes(e, handler)
	return e, nil

}

func Routes(e *echo.Echo, allHandler *common.Handler) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/login", allHandler.AuthHandler.Login)
	e.POST("/register", allHandler.AuthHandler.Register)
	api := e.Group("/api")
	api.Use(middlewares.JWTMiddleware())
	api.GET("/user", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "user",
		})
	})
	api.POST("/wallet", allHandler.WalletHandler.CreateWallet)
	api.GET("/wallet", allHandler.WalletHandler.GetWalletUser)
	api.GET("/wallet/:id", allHandler.WalletHandler.ShowWallet)
	api.PUT("/wallet/:id", allHandler.WalletHandler.UpdateWallet)
	api.POST("/category-income", allHandler.CategoryIncomeHandler.CreateCategoryIncome)
	api.GET("/category-income", allHandler.CategoryIncomeHandler.GetCategoryIncome)
	api.GET("/category-income/:id", allHandler.CategoryIncomeHandler.ShowCategoryIncome)
	api.PUT("/category-income/:id", allHandler.CategoryIncomeHandler.UpdateCategoryIncome)
	api.POST("/category-expense", allHandler.CategoryExpenseHandler.CreateCategoryExpense)
	api.GET("/category-expense", allHandler.CategoryExpenseHandler.GetCategoryExpense)
	api.GET("/category-expense/:id", allHandler.CategoryExpenseHandler.ShowCategoryExpense)
	api.PUT("/category-expense/:id", allHandler.CategoryExpenseHandler.UpdateCategoryExpense)
	api.POST("/income", allHandler.IncomeHandler.CreateIncome)
	api.GET("/income", allHandler.IncomeHandler.GetIncomes)
	api.GET("/income/:id", allHandler.IncomeHandler.ShowIncome)
	api.PUT("/income/:id", allHandler.IncomeHandler.UpdateIncome)
	// api.PUT("/category-inco")
	// api.GET("")
}

func SetupMiddleware(e *echo.Echo, enableCors bool) {
	if enableCors {
		e.Use(middleware.CORS())
	}
}
