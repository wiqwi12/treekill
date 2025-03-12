package storage

import (
	"2/internal/domain/models"
	"errors"
	"github.com/google/uuid"
	"sync"
)

type UserRepository struct {
	storage map[uuid.UUID]models.User
	sync.Mutex
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		storage: make(map[uuid.UUID]models.User),
	}
}

func (r *UserRepository) Create(user models.User) (models.User, error) {
	r.Lock()
	defer r.Unlock()

	r.storage[user.UserId] = user
	return user, nil
}

func (r *UserRepository) GetUserByEmail(email string) (models.User, error) {
	r.Lock()
	defer r.Unlock()

	for _, user := range r.storage {
		if user.Email == email {
			return user, nil
		}
	}
	return models.User{}, errors.New("user not found")
}

func (r *UserRepository) GetUserById(id uuid.UUID) (models.User, error) {
	r.Lock()
	defer r.Unlock()

	user, ok := r.storage[id]
	if !ok {
		return models.User{}, errors.New("user not found")
	}
	return user, nil
}
