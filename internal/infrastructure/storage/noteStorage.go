package storage

import (
	"2/internal/domain/models"
	"errors"
	"github.com/google/uuid"
	"sync"
)

type NotesRepository struct {
	notes     map[uuid.UUID]models.Note //Contains key - note id, val - note
	userIndex map[uuid.UUID][]uuid.UUID //key - userId, val - slice of user`s notes id`s
	mutex     sync.Mutex
}

func NewNotesRepository() *NotesRepository {
	return &NotesRepository{
		notes:     make(map[uuid.UUID]models.Note),
		userIndex: make(map[uuid.UUID][]uuid.UUID),
	}
}

func (s *NotesRepository) Create(note models.Note) (models.Note, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.notes[note.ID] = note
	s.userIndex[note.UserId] = append(s.userIndex[note.UserId], note.ID)
	return note, nil
}

func (s *NotesRepository) Get(id uuid.UUID) (models.Note, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	note, ok := s.notes[id]
	if !ok {
		return models.Note{}, errors.New("not found")
	}

	return note, nil
}

func (s *NotesRepository) GetAllByUserId(id uuid.UUID) ([]models.Note, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	noteIds, ok := s.userIndex[id]
	if !ok {
		return []models.Note{}, errors.New("User does not have any notes")
	}

	notes := make([]models.Note, len(noteIds))
	for _, noteId := range noteIds {
		notes = append(notes, s.notes[noteId])
	}
	return notes, nil

}

func (s *NotesRepository) Update(note models.Note) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	exsistingNote, ok := s.notes[note.ID]
	if !ok {
		return errors.New("not found")
	}

	if exsistingNote.UserId != note.UserId {
		return errors.New("note is not owned by this user")
	}

	s.notes[note.ID] = note

	return nil

}

func (s *NotesRepository) Delete(id uuid.UUID) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	note, ok := s.notes[id]
	if !ok {
		return errors.New("Note does not exist")
	}

	delete(s.notes, id)

	userNoteIds, _ := s.userIndex[id]
	updated := make([]uuid.UUID, 0, len(userNoteIds))

	for _, userNoteId := range userNoteIds {
		updated = append(updated, userNoteId)
	}

	if len(updated) == 0 {
		delete(s.userIndex, note.UserId)
	}

	s.userIndex[note.UserId] = updated

	return nil
}
