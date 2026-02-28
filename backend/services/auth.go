package services

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"
	"golang.org/x/crypto/bcrypt"

	"github.com/TeeHaoBin/TicketLah/backend/models"
	"github.com/TeeHaoBin/TicketLah/backend/utils"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type AuthService struct {
	repository models.AuthRepository
}

func (s *AuthService) Login(ctx context.Context, loginData *models.AuthCredentials) (string, *models.User, error) {
	user, err := s.repository.GetUser(ctx, "email = ?", loginData.Email)
	
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", nil, fmt.Errorf("Invalid credentials")
		}
		return "", nil, err
	}

	if !models.MatchesHash(loginData.Password, user.Password) {
		return "", nil, fmt.Errorf("Invalid credentials")
	}

	claims := jwt.MapClaims{
		"id": user.ID,
		"role": user.Role,
		"exp": time.Now().Add(time.Hour * 168).Unix(),
	}

	token, err := utils.GenerateJWT(claims, jwt.SigningMethodHS256, os.Getenv("JWT_SECRET"))
	if err != nil {
		return "", nil, err
	}

	return token, user, nil
}

func (s *AuthService) Register(ctx context.Context, RegisterData *models.AuthCredentials) (string, *models.User, error) {
	if !models.IsValidEmail(RegisterData.Email) {
		return "", nil, fmt.Errorf("Please provide a VALID email to register")
	}

	if _, err := s.repository.GetUser(ctx, "email = ?", RegisterData.Email); !errors.Is(err, gorm.ErrRecordNotFound) {
		return "", nil, fmt.Errorf("Email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(RegisterData.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", nil, err
	}

	RegisterData.Password = string(hashedPassword)

	user, err := s.repository.RegisterUser(ctx, RegisterData)
	
	if err != nil {
		return "", nil, err
	}

	claims := jwt.MapClaims{
		"id": user.ID,
		"role": user.Role,
		"exp": time.Now().Add(time.Hour * 168).Unix(),
	}

	token, err := utils.GenerateJWT(claims, jwt.SigningMethodHS256, os.Getenv("JWT_SECRET"))
	
	if err != nil {
		return "", nil, err
	}

	return token, user, nil
}

func NewAuthService(repository models.AuthRepository) models.AuthService {
	return &AuthService{
		repository: repository,
	}
}