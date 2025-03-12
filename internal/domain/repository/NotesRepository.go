package repository

import (
	"2/internal/domain/models"
	"github.com/google/uuid"
)

type NotesRepository interface {
	Create(note models.Note) (models.Note, error)
	GetNoteById(idStr string) (models.Note, error)
	GetAllByUserId(id uuid.UUID) ([]models.Note, error)
	Update(note models.Note) error
	Delete(id uuid.UUID) error
}
