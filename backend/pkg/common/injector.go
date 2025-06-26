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
)

var serviceSet = wire.NewSet(
	services.NewAuthService,
)

var handlerSet = wire.NewSet(
	handlers.NewAuthHandler,
)

type Handler struct {
	AuthHandler *handlers.AuthHandler
}

func InjectDependencies(db *gorm.DB) (*Handler, error) {
	wire.Build(repositorySet, serviceSet, handlerSet, wire.Struct(new(Handler), "authHandler"))
	return nil, nil
}
