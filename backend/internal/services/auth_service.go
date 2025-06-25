package services

import (
	"api-money-management/internal/dtos"
	"api-money-management/internal/models"
	"api-money-management/internal/repositories"
	"api-money-management/pkg/auth"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	authRepo *repositories.AuthRepository
}

func NewAuthService(authRepo *repositories.AuthRepository) *AuthService {
	return &AuthService{authRepo: authRepo}
}

func (s *AuthService) Login(req *dtos.LoginRequest) (*dtos.LoginResponse, error) {
	user, err := s.authRepo.FindByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid email or password")
	}
	token, err := auth.GenerateJWT(user.ID, user.Email)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}
	return &dtos.LoginResponse{
		Token:     token,
		ExpiresAt: 24 * 60 * 60,
		TokenType: "Bearer",
		UserID:    user.ID,
		UserName:  user.Name,
	}, nil
}

func (s *AuthService) Register(req *dtos.RegisterRequest) (*models.User, error) {

	// periksa apakah email sudah terdaftar atau belum
	existinguser, err := s.authRepo.FindByEmail(req.Email)
	if existinguser != nil {
		return nil, errors.New("email have been used")
	}

	// hash password
	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to create hash password")
	}
	user := &models.User{
		Email:    req.Email,
		Name:     req.Name,
		Password: string(hashedpassword),
	}
	result, err := s.authRepo.Create(user)
	if err != nil {
		return nil, errors.New("failed to create user")
	}
	return result, nil
	// insert to database
}
