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
	repositories.NewCategoryIncome,
)

var serviceSet = wire.NewSet(
	services.NewAuthService,
	services.NewWalletService,
	services.NewCategoryIncomeService,
)

var handlerSet = wire.NewSet(
	handlers.NewAuthHandler,
	handlers.NewWalletHandler,
	handlers.NewCategoryIncomeHandler,
)

type Handler struct {
	AuthHandler           *handlers.AuthHandler
	WalletHandler         *handlers.WalletHandler
	CategoryIncomeHandler *handlers.CategoryIncomeHandler
}

func InjectDependencies(db *gorm.DB) (*Handler, error) {
	wire.Build(repositorySet, serviceSet, handlerSet, wire.Struct(new(Handler), "*"))
	return nil, nil
}
