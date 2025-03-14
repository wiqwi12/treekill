package dto

import "github.com/google/uuid"

// AuthResponse represents authentication token response
type AuthResponse struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}

// UserResponse represents user data response
type UserResponse struct {
	Email    string    `json:"email" example:"user@example.com"`
	Username string    `json:"username" example:"john_doe"`
	UserId   uuid.UUID `json:"userid" example:"550e8400-e29b-41d4-a716-446655440000"`
}

// StandartResponse represents standart response with message
type StandartResponse struct {
	Message string `json:"message" example:"Hello World"`
}
