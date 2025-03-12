package service

import (
	"2/internal/domain/models"
	"2/internal/infrastructure/storage"
	"2/internal/interface/http/dto"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthService struct {
	UserRepo *storage.UserRepository
	Secret   string
}

func NewAuthService(userRepo *storage.UserRepository, secret string) *AuthService {
	return &AuthService{UserRepo: userRepo, Secret: secret}
}

func (s *AuthService) RegisterUser(req dto.RegistrationRequest) (models.User, error) {

	_, err := s.UserRepo.GetUserByEmail(req.Email)
	if err == nil {
		return models.User{}, errors.New("User with this email already exsits")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	user := models.User{
		UserId:   uuid.New(),
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
		Created:  time.Now(),
	}

	return s.UserRepo.Create(user)

}

// Исправляем функцию LoginUser для сохранения UUID как строки в токене
func (s *AuthService) LoginUser(req dto.LoginRequest) (string, error) {
	user, err := s.UserRepo.GetUserByEmail(req.Email)
	if err != nil {
		return "", errors.New("Invalid email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", errors.New("Invalid password")
	}

	// Исправленный код - сохраняем UUID как строку
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.UserId.String(), // Преобразуем UUID в строку
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
