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
)

var serviceSet = wire.NewSet(
	services.NewAuthService,
	services.NewWalletService,
)

var handlerSet = wire.NewSet(
	handlers.NewAuthHandler,
	handlers.NewWalletHandler,
)

type Handler struct {
	AuthHandler   *handlers.AuthHandler
	WalletHandler *handlers.WalletHandler
}

func InjectDependencies(db *gorm.DB) (*Handler, error) {
	wire.Build(repositorySet, serviceSet, handlerSet, wire.Struct(new(Handler), "*"))
	return nil, nil
}
