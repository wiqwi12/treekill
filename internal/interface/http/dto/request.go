package dto

// RegistrationRequest represents user registration data
type RegistrationRequest struct {
	Email    string `json:"email" example:"user@example.com"`
	Username string `json:"username" example:"john_doe"`
	Password string `json:"password" example:"P@ssw0rd!"`
}

// LoginRequest represents user login credentials
type LoginRequest struct {
	Email    string `json:"email" example:"user@example.com"`
	Password string `json:"password" example:"P@ssw0rd!"`
}

// CreateNoteRequest represents note creation data
type CreateNoteRequest struct {
	Title   string `json:"title" example:"My First Note"`
	Content string `json:"content" example:"Note content here"`
}

// UpdateNoteRequest represents note update data
type UpdateNoteRequest struct {
	Title   string `json:"title" example:"Updated Note Title"`
	Content string `json:"content" example:"Updated note content"`
}
