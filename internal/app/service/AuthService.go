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

func (s *AuthService) RegisterUser(req dto.RegistrationRequest) error {

	_, exsists, err := s.UserRepo.GetUserByEmail(req.Email)
	if exsists {
		return errors.New("User with this email already exsits")
	}
	if err != nil {
		return err
	}

	user := models.User{
		UserId:   uuid.New(),
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
		Created:  time.Now(),
	}
	user.HashPassword()

	return s.UserRepo.Create(user)

}

// Исправляем функцию LoginUser для сохранения UUID как строки в токене
func (s *AuthService) LoginUser(req dto.LoginRequest) (string, error) {
	user, exsists, err := s.UserRepo.GetUserByEmail(req.Email)

	if !exsists {
		return "", errors.New("There is no user with this email")
	}

	//переписать под валидацию данных
	if err != nil {
		return "", errors.New("Invalid email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", errors.New("Invalid password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.UserId.String(),
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
