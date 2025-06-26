package services

import (
	"api-money-management/internal/dtos"
	"api-money-management/internal/models"
	"api-money-management/internal/repositories"
	"api-money-management/pkg/auth"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	authRepo *repositories.AuthRepository
}

func NewAuthService(authRepo *repositories.AuthRepository) *AuthService {
	return &AuthService{authRepo: authRepo}
}

func (s *AuthService) Login(req *dtos.LoginRequest) (*dtos.LoginResponse, *dtos.ErrorResponse) {
	user, err := s.authRepo.FindByEmail(req.Email)
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, dtos.NewErrorResponse("Invalid email or password", 400, "user not found")
	}
	token, tokenErr := auth.GenerateJWT(user.ID, user.Email)
	if tokenErr != nil {
		return nil, dtos.NewErrorResponse(tokenErr.Error(), 500, "token generation error")
	}
	return &dtos.LoginResponse{
		Token:     token,
		ExpiresAt: 24 * 60 * 60,
		TokenType: "Bearer",
		UserID:    user.ID,
		UserName:  user.Name,
	}, nil
}

func (s *AuthService) Register(req *dtos.RegisterRequest) (*models.User, *dtos.ErrorResponse) {

	// periksa apakah email sudah terdaftar atau belum
	existinguser, err := s.authRepo.FindByEmail(req.Email)
	if existinguser != nil {
		return nil, err
	}

	// hash password
	hashedpassword, hashedError := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if hashedError != nil {
		return nil, dtos.NewErrorResponse("Failed to hash password", 500, "hashing error")
	}
	user := &models.User{
		Email:    req.Email,
		Name:     req.Name,
		Password: string(hashedpassword),
	}
	result, err := s.authRepo.Create(user)
	if err != nil {
		return nil, err
	}
	return result, nil
	// insert to database
}
