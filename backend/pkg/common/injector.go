package common

import (
	"api-money-management/internal/handlers"

	"gorm.io/gorm"
)

// var repositorySet = wire.NewSet(
//     repositories.NewAuthRepository,
//     repositories.NewTransactionRepository,
// )

// var serviceSet = wire.NewSet(
//     services.NewAuthService,
//     services.NewTransactionService,
// )

// var handlerSet = wire.NewSet(
//     handlers.NewAuthHandler,
//     handlers.NewTransactionHandler,
// )

type Handler struct {
	authHandler *handlers.AuthHandler
}

func InjectDependencies(db *gorm.DB) (*Handler, error) {
	return &Handler{}, nil
}
