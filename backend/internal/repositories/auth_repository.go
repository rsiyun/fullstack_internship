package repositories

import (
	"api-money-management/internal/dtos"
	"api-money-management/internal/models"
	"errors"

	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) FindByEmail(email string) (*models.User, *dtos.ErrorResponse) {
	user := new(models.User)
	err := r.db.Where("email = ?", email).First(user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, dtos.NewErrorResponse("User not found", 404, "user not found")
	}
	return user, nil
}
func (r *AuthRepository) Create(user *models.User) (*models.User, *dtos.ErrorResponse) {
	result := r.db.Create(user)
	if result.RowsAffected > 0 {
		return user, nil
	}
	return nil, dtos.NewErrorResponse("Failed to create user", 500, "database error")
}
