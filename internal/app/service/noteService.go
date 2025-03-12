package service

import (
	"2/internal/domain/models"
	"2/internal/infrastructure/storage"
	"2/internal/interface/http/dto"
	"errors"
	"github.com/google/uuid"
	"time"
)

type NoteService struct {
	noteRepo storage.NotesRepository
}

func NewNoteService(noteRepo storage.NotesRepository) *NoteService {
	return &NoteService{noteRepo: noteRepo}
}

func (s *NoteService) CreateNote(userId uuid.UUID, req dto.CreateNoteRequest) (models.Note, error) {
	if req.Title == "" {
		return models.Note{}, errors.New("title is required")
	}

	note := models.Note{
		ID:        uuid.New(),
		UserId:    userId,
		Title:     req.Title,
		Content:   req.Content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return s.noteRepo.Create(note)
}

func (s *NoteService) GetNote(userId uuid.UUID, noteId uuid.UUID) (models.Note, error) {

	note, err := s.noteRepo.Get(noteId)
	if err != nil {
		return models.Note{}, err
	}

	if note.UserId != userId {
		return models.Note{}, errors.New("Accsess denied")
	}

	return note, nil
}

func (s *NoteService) UpdateNote(userId uuid.UUID, noteId uuid.UUID, req dto.UpdateNoteRequest) error {

	note, err := s.noteRepo.Get(noteId)
	if err != nil {
		return err
	}

	if note.UserId != userId {
		return errors.New("Accsess denied")
	}

	if req.Title == "" {
		return errors.New("title is required")
	}

	if req.Content == "" {
		return errors.New("content is required")
	}

	note.Title = req.Title
	note.Content = req.Content
	note.UpdatedAt = time.Now()
	return s.noteRepo.Update(note)
}

func (s *NoteService) GetUserNotes(userID uuid.UUID) ([]models.Note, error) {
	return s.noteRepo.GetAllByUserId(userID)
}

func (s *NoteService) DeleteNote(userId uuid.UUID, noteId uuid.UUID) error {
	return s.noteRepo.Delete(noteId)
}
