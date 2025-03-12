package errors

// ErrorResponse represents standard error format
type ErrorResponse struct {
	Error string `json:"error" example:"error description"`
}
