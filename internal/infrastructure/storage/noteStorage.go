package storage

import (
	"2/internal/domain/models"
	"database/sql"
	"errors"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"time"
)

type NotesRepository struct {
	Db *sql.DB
}

func NewNotesRepository(db *sql.DB) *NotesRepository {
	return &NotesRepository{
		Db: db,
	}
}

func (s *NotesRepository) Create(note models.Note) error {

	query, args, err := squirrel.Insert("notes").
		Columns("id", "user_id", "title", "content", "created_at", "updated_at").
		Values(note.ID, note.UserId, note.Title, note.Content, note.CreatedAt, note.UpdatedAt).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return err
	}
	_, err = s.Db.Query(query, args...)
	if err != nil {
		return errors.New(fmt.Sprint("Error inserting note into database: ", err))
	}

	return nil
}

func (s *NotesRepository) Get(id uuid.UUID) (models.Note, error) {

	query, args, err := squirrel.Select("id", "user_id", "title", "content", "created_at", "updated_at").
		From("notes").Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return models.Note{}, err
	}

	var note models.Note
	row := s.Db.QueryRow(query, args...)
	err = row.Scan(
		&note.ID,
		&note.UserId,
		&note.Title,
		&note.Content,
		&note.CreatedAt,
		&note.UpdatedAt)

	if err != nil {
		return models.Note{}, err
	}

	return note, nil
}

func (s *NotesRepository) GetAllByUserId(id uuid.UUID) ([]models.Note, error) {

	query, args, err := squirrel.Select("*").
		From("notes").
		Where(squirrel.Eq{
			"user_id": id,
		}).PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return []models.Note{}, err
	}

	rows, err := s.Db.Query(query, args...)
	if err != nil {
		return []models.Note{}, err
	}
	defer rows.Close()

	var notes []models.Note
	for rows.Next() {
		var note models.Note
		err = rows.Scan(
			&note.ID,
			&note.UserId,
			&note.Title,
			&note.Content,
			&note.CreatedAt,
			&note.UpdatedAt)
		if err != nil {
			return []models.Note{}, err
		}
		notes = append(notes, note)
	}

	err = rows.Err()
	if err != nil {
		return []models.Note{}, err
	}
	return notes, nil
}

func (s *NotesRepository) Update(note models.Note) error {
	prev, err := s.Get(note.ID)
	if err != nil {
		return err
	}

	if prev.Content == note.Content && prev.Title == note.Title {
		return errors.New("There is no updates")
	}

	query, args, err := squirrel.Update("notes").
		Set("title", note.Title).
		Set("content", note.Content).
		Set("updated_at", time.Now()).
		Where(squirrel.Eq{"id": note.ID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return err
	}

	_, err = s.Db.Exec(query, args...)
	return err
}

func (s *NotesRepository) Delete(id uuid.UUID) error {

	query, args, err := squirrel.Delete("notes").Where(squirrel.Eq{
		"user_id": id,
	}).PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return err
	}

	_, err = s.Db.Exec(query, args...)
	return err

}
