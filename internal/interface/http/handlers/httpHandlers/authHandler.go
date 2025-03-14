package httpHandlers

import (
	"2/internal/app/service"
	"2/internal/errors"
	"2/internal/interface/http/dto"
	"encoding/json"
	"net/http"
)

type AuthHandler struct {
	AuthService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{AuthService: authService}
}

// Register godoc
// @Summary User registration
// @Description Create new user account
// @Tags Auth
// @Accept json
// @Produce json
// @Param input body dto.RegistrationRequest true "Registration data"
// @Success 201 {object} dto.StandartResponse
// @Failure 400 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Router /user/register [post]
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {

	var req dto.RegistrationRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	if req.Username == "" || req.Email == "" || req.Password == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		resp := errors.ErrorResponse{
			Error: "Email, username and password are required",
		}
		json.NewEncoder(w).Encode(resp)
		return
	}

	err = h.AuthService.RegisterUser(req)
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
	response := dto.StandartResponse{
		Message: "succsess",
	}

	json.NewEncoder(w).Encode(response)

}

// Login godoc
// @Summary User authentication
// @Description Get JWT access token
// @Tags Auth
// @Accept json
// @Produce json
// @Param input body dto.LoginRequest true "Login credentials"
// @Success 200 {object} dto.AuthResponse
// @Failure 400 {object} errors.ErrorResponse
// @Failure 401 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Router /user/login [post]
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {

	var req dto.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

	}

	token, err := h.AuthService.LoginUser(req)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		resp := errors.ErrorResponse{
			Error: err.Error(),
		}
		json.NewEncoder(w).Encode(resp)
		return
	}

	response := dto.AuthResponse{
		Token: token,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}
