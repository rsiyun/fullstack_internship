//go:build wireinject
// +build wireinject

package common

import (
	"api-money-management/internal/handlers"
	"api-money-management/internal/repositories"
	"api-money-management/internal/services"

	"github.com/google/wire"
	"gorm.io/gorm"
)

var repositorySet = wire.NewSet(
	repositories.NewAuthRepository,
	repositories.NewWalletRepository,
	repositories.NewIncomeRepository,
	repositories.NewCategoryIncomeRepository,
	repositories.NewCategoryExpenseRepository,
)

var serviceSet = wire.NewSet(
	services.NewAuthService,
	services.NewWalletService,
	services.NewIncomeService,
	services.NewCategoryIncomeService,
	services.NewCategoryExpenseService,
)

var handlerSet = wire.NewSet(
	handlers.NewAuthHandler,
	handlers.NewWalletHandler,
	handlers.NewIncomeHandler,
	handlers.NewCategoryIncomeHandler,
	handlers.NewCategoryExpenseHandler,
)

type Handler struct {
	AuthHandler            *handlers.AuthHandler
	WalletHandler          *handlers.WalletHandler
	IncomeHandler          *handlers.IncomeHandler
	CategoryIncomeHandler  *handlers.CategoryIncomeHandler
	CategoryExpenseHandler *handlers.CategoryExpenseHandler
}

func InjectDependencies(db *gorm.DB) (*Handler, error) {
	wire.Build(repositorySet, serviceSet, handlerSet, wire.Struct(new(Handler), "*"))
	return nil, nil
}
