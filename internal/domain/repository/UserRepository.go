package repository

import (
	"2/internal/domain/models"
	"github.com/google/uuid"
)

type UserRepository interface {
	Create(user models.User) error
	GetUserByEmail(email string) (models.User, bool, error)
	GetUserById(id uuid.UUID) (models.User, error)
}
