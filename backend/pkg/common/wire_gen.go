// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package common

import (
	"api-money-management/internal/handlers"
	"api-money-management/internal/repositories"
	"api-money-management/internal/services"
	"github.com/google/wire"
	"gorm.io/gorm"
)

// Injectors from injector.go:

func InjectDependencies(db *gorm.DB) (*Handler, error) {
	authRepository := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository)
	authHandler := handlers.NewAuthHandler(authService)
	walletRepository := repositories.NewWalletRepository(db)
	walletService := services.NewWalletService(walletRepository)
	walletHandler := handlers.NewWalletHandler(walletService)
	incomeRepository := repositories.NewIncomeRepository(db)
	categoryIncomeRepo := repositories.NewCategoryIncomeRepository(db)
	incomeService := services.NewIncomeService(incomeRepository, walletRepository, categoryIncomeRepo)
	incomeHandler := handlers.NewIncomeHandler(incomeService)
	categoryIncomeService := services.NewCategoryIncomeService(categoryIncomeRepo)
	categoryIncomeHandler := handlers.NewCategoryIncomeHandler(categoryIncomeService)
	categoryExpenseRepo := repositories.NewCategoryExpenseRepository(db)
	categoryExpenseService := services.NewCategoryExpenseService(categoryExpenseRepo)
	categoryExpenseHandler := handlers.NewCategoryExpenseHandler(categoryExpenseService)
	handler := &Handler{
		AuthHandler:            authHandler,
		WalletHandler:          walletHandler,
		IncomeHandler:          incomeHandler,
		CategoryIncomeHandler:  categoryIncomeHandler,
		CategoryExpenseHandler: categoryExpenseHandler,
	}
	return handler, nil
}

// injector.go:

var repositorySet = wire.NewSet(repositories.NewAuthRepository, repositories.NewWalletRepository, repositories.NewIncomeRepository, repositories.NewCategoryIncomeRepository, repositories.NewCategoryExpenseRepository)

var serviceSet = wire.NewSet(services.NewAuthService, services.NewWalletService, services.NewIncomeService, services.NewCategoryIncomeService, services.NewCategoryExpenseService)

var handlerSet = wire.NewSet(handlers.NewAuthHandler, handlers.NewWalletHandler, handlers.NewIncomeHandler, handlers.NewCategoryIncomeHandler, handlers.NewCategoryExpenseHandler)

type Handler struct {
	AuthHandler            *handlers.AuthHandler
	WalletHandler          *handlers.WalletHandler
	IncomeHandler          *handlers.IncomeHandler
	CategoryIncomeHandler  *handlers.CategoryIncomeHandler
	CategoryExpenseHandler *handlers.CategoryExpenseHandler
}
