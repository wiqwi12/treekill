package repository

import (
	"2/internal/domain/models"
	"github.com/google/uuid"
)

type UserRepository interface {
	Create(user models.User) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	GetUserById(id uuid.UUID) (models.User, error)
}
