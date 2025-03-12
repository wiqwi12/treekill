package httpHandlers

import (
	"2/internal/app/service"
	"2/internal/errors"
	"2/internal/interface/http/dto"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

type NoteHandler struct {
	noteService *service.NoteService
}

func NewNoteHandler(service *service.NoteService) *NoteHandler {
	return &NoteHandler{
		noteService: service,
	}
}

// Вспомогательная функция для безопасного получения UUID из контекста
func getUserIDFromContext(r *http.Request) (uuid.UUID, error) {
	// Печатаем все ключи в контексте для отладки
	fmt.Println("Context keys:")
	fmt.Printf("  userId exists: %v\n", r.Context().Value("userId") != nil)
	fmt.Printf("  user_id exists: %v\n", r.Context().Value("user_id") != nil)

	// Пробуем получить UUID с ключом "userId"
	userIDVal := r.Context().Value("userId")
	if userIDVal == nil {
		// Если не нашли, пробуем с ключом "user_id"
		userIDVal = r.Context().Value("user_id")
		if userIDVal == nil {
			return uuid.UUID{}, fmt.Errorf("user ID not found in context")
		}
	}

	// Печатаем тип и значение для отладки
	fmt.Printf("User ID from context: (%T) %v\n", userIDVal, userIDVal)

	// Если это уже UUID, просто возвращаем его
	if userID, ok := userIDVal.(uuid.UUID); ok {
		return userID, nil
	}

	// Если это строка, парсим ее в UUID
	if userIDStr, ok := userIDVal.(string); ok {
		return uuid.Parse(userIDStr)
	}

	// Если ничего не подошло, возвращаем ошибку
	return uuid.UUID{}, fmt.Errorf("user ID has unexpected type: %T", userIDVal)
}

// GetNoteHandler godoc
// @Summary Get note by ID
// @Description Get single note by its ID
// @Tags Notes
// @Security JWTAuth
// @Produce json
// @Param id path string true "Note ID"
// @Success 200 {object} models.Note
// @Failure 400 {object} errors.ErrorResponse
// @Failure 401 {object} errors.ErrorResponse
// @Failure 404 {object} errors.ErrorResponse
// @Router /notes/{id} [get]
func (h *NoteHandler) GetNoteHandler(w http.ResponseWriter, r *http.Request) {
	// Безопасное получение UUID
	userId, err := getUserIDFromContext(r)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		resp := errors.ErrorResponse{
			Error: fmt.Sprintf("Auth error: %v", err),
		}
		json.NewEncoder(w).Encode(resp)
		return
	}

	noteIdStr := r.PathValue("id")
	noteId, err := uuid.Parse(noteIdStr)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		resp := errors.ErrorResponse{
			Error: fmt.Sprintf("Note ID %s is invalid", noteIdStr),
		}
		json.NewEncoder(w).Encode(resp)
		return
	}

	note, err := h.noteService.GetNote(userId, noteId)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		resp := errors.ErrorResponse{
			Error: err.Error(),
		}
		json.NewEncoder(w).Encode(resp)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(note)
}

// GetNotes godoc
// @Summary Get all notes
// @Description Get list of all user's notes
// @Tags Notes
// @Security JWTAuth
// @Produce json
// @Success 200 {array} models.Note
// @Failure 401 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Router /notes [get]
func (h *NoteHandler) GetNotes(w http.ResponseWriter, r *http.Request) {
	// Безопасное получение UUID
	userId, err := getUserIDFromContext(r)
	if err != nil {
		http.Error(w, fmt.Sprintf("Auth error: %v", err), http.StatusUnauthorized)
		return
	}

	notes, err := h.noteService.GetUserNotes(userId)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		resp := errors.ErrorResponse{
			Error: err.Error(),
		}
		json.NewEncoder(w).Encode(resp)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(notes)
}

// CreateNote godoc
// @Summary Create new note
// @Description Create new note for authenticated user
// @Tags Notes
// @Security JWTAuth
// @Accept json
// @Produce json
// @Param input body dto.CreateNoteRequest true "Note data"
// @Success 201 {object} models.Note
// @Failure 400 {object} errors.ErrorResponse
// @Failure 401 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Router /notes [post]
func (h *NoteHandler) CreateNote(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateNoteRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		resp := errors.ErrorResponse{
			Error: err.Error(),
		}
		json.NewEncoder(w).Encode(resp)
		return
	}

	// Безопасное получение UUID
	userId, err := getUserIDFromContext(r)
	if err != nil {
		http.Error(w, fmt.Sprintf("Auth error: %v", err), http.StatusUnauthorized)
		return
	}

	note, err := h.noteService.CreateNote(userId, req)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		resp := errors.ErrorResponse{
			Error: err.Error(),
		}
		json.NewEncoder(w).Encode(resp)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(note)
}

// DeleteNote godoc
// @Summary Delete note
// @Description Delete note by ID
// @Tags Notes
// @Security JWTAuth
// @Param id path string true "Note ID"
// @Success 204
// @Failure 400 {object} errors.ErrorResponse
// @Failure 401 {object} errors.ErrorResponse
// @Failure 404 {object} errors.ErrorResponse
// @Router /notes/{id} [delete]
func (h *NoteHandler) DeleteNote(w http.ResponseWriter, r *http.Request) {
	// Безопасное получение UUID
	userId, err := getUserIDFromContext(r)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		resp := errors.ErrorResponse{
			Error: err.Error(),
		}
		json.NewEncoder(w).Encode(resp)
		return
	}

	noteIdStr := r.PathValue("id")
	noteId, err := uuid.Parse(noteIdStr)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		resp := errors.ErrorResponse{
			Error: "somethin wrong with noteId",
		}
		json.NewEncoder(w).Encode(resp)
		return
	}

	err = h.noteService.DeleteNote(userId, noteId)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		resp := errors.ErrorResponse{
			Error: err.Error(),
		}
		json.NewEncoder(w).Encode(resp)
		return
	}
}

// UpdateNote godoc
// @Summary Update note
// @Description Update existing note
// @Tags Notes
// @Security JWTAuth
// @Accept json
// @Produce json
// @Param id path string true "Note ID"
// @Param input body dto.UpdateNoteRequest true "Updated note data"
// @Success 200 {object} models.Note
// @Failure 400 {object} errors.ErrorResponse
// @Failure 401 {object} errors.ErrorResponse
// @Failure 404 {object} errors.ErrorResponse
// @Router /notes/{id} [put]
func (h *NoteHandler) UpdateNote(w http.ResponseWriter, r *http.Request) {
	// Безопасное получение UUID
	userId, err := getUserIDFromContext(r)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		resp := errors.ErrorResponse{
			Error: err.Error(),
		}
		json.NewEncoder(w).Encode(resp)
		return
	}

	noteIdStr := r.PathValue("id")
	noteId, err := uuid.Parse(noteIdStr)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		resp := errors.ErrorResponse{
			Error: "somethin wrong with noteId",
		}
		json.NewEncoder(w).Encode(resp)
		return
	}

	var req dto.UpdateNoteRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		resp := errors.ErrorResponse{
			Error: err.Error(),
		}
		json.NewEncoder(w).Encode(resp)
		return
	}

	err = h.noteService.UpdateNote(userId, noteId, req)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		resp := errors.ErrorResponse{
			Error: err.Error(),
		}
		json.NewEncoder(w).Encode(resp)
		return

	}
}
