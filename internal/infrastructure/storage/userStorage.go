package storage

import (
	"2/internal/domain/models"
	"database/sql"
	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"time"
)

type UserRepository struct {
	Db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		Db: db,
	}
}

func (r *UserRepository) Create(user models.User) error {

	query, args, err := squirrel.Insert("users").
		Columns("user_id", "username", "email", "password", "created").
		Values(user.UserId, user.Username, user.Email, user.Password, time.Now()).
		PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return err
	}

	_, err = r.Db.Query(query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetUserByEmail(email string) (models.User, bool, error) {
	query, args, err := squirrel.Select("user_id", "username", "email", "password", "created").
		From("users").
		Where(squirrel.Eq{
			"email": email,
		}).PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return models.User{}, false, err
	}

	row := r.Db.QueryRow(query, args...)
	user := models.User{}
	err = row.Scan(
		&user.UserId,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Created,
	)

	// Проверяем, был ли найден пользователь
	if err != nil {
		if err == sql.ErrNoRows {
			// Пользователь не найден, но это не ошибка
			return models.User{}, false, nil
		}
		// Произошла реальная ошибка при выполнении запроса
		return models.User{}, false, err
	}

	// Пользователь найден
	return user, true, nil
}
func (r *UserRepository) GetUserById(id uuid.UUID) (models.User, error) {

	query, args, err := squirrel.Select("user_id", "username", "email", "password", "created").
		From("users").
		Where(squirrel.Eq{
			"user_id": id,
		}).PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return models.User{}, err
	}

	row := r.Db.QueryRow(query, args...)
	user := models.User{}
	row.Scan(
		&user.UserId,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Created,
	)

	return user, nil

}
