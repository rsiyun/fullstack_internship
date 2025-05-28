package repositories

import (
	"dot/app/models/domain"
	"errors"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Create(user *domain.User) error
	FindByEmail(email string) (*domain.User, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) Create(user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *authRepository) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("not found")
	}
	return &user, err
}
