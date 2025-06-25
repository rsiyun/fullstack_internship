package repositories

import (
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

func (r *AuthRepository) FindByEmail(email string) (*models.User, error) {
	user := new(models.User)
	err := r.db.Where("email = ?", email).First(user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("not found")
	}
	return user, err
}
func (r *AuthRepository) Create(user *models.User) (*models.User, error) {
	result := r.db.Create(user)
	if result.RowsAffected > 0 {
		return user, nil
	}
	return nil, errors.New("failed to insert data")
}
