package services

import (
	"dot/app/models/domain"
	"dot/app/repositories"
	"dot/helpers"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AuthServices interface {
	Register(name, email, password string) (*domain.User, error)
	Login(email, password string) (*LoginResponse, error)
}

type LoginResponse struct {
	Token string `json:"token"`
}

type authServices struct {
	userRepo repositories.AuthRepository
}

func NewUserService(userRepo repositories.AuthRepository) AuthServices {
	return &authServices{userRepo: userRepo}
}

func (s *authServices) Register(name, email, password string) (*domain.User, error) {
	// Check if email already exists
	existingUser, _ := s.userRepo.FindByEmail(email)
	if existingUser != nil {
		return nil, errors.New("email already in use")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create user
	user := &domain.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}
	err = s.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
func (s *authServices) Login(email, password string) (*LoginResponse, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid email or password")
	}

	token, err := helpers.GenerateJWT(user.ID, user.Email)
	if err != nil {
		return nil, errors.New("cannot generate Token")
	}

	return &LoginResponse{Token: token}, nil
}
