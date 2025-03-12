package main

import (
	_ "2/docs"
	"2/internal/app/service"
	"2/internal/infrastructure/storage"
	"2/internal/interface/http/handlers/httpHandlers"
	"2/internal/interface/http/middleware"
	"context"
	"github.com/swaggo/http-swagger"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	//nolint
	_ "2/docs"
)

// @title StudyNoteAPI
// @version 0.8.2
// @description API for notes management with JWT authentication

// @contact.name API Support
// @contact.email support@studynoteapi.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey JWTAuth
// @in header
// @name Authorization
// @schemes http
func main() {
	Run()
}

func Run() {

	secret := "12345"

	UserRepo := storage.NewUserRepository()
	NotesRepo := storage.NewNotesRepository()
	NotesService := service.NewNoteService(*NotesRepo)
	AuthService := service.NewAuthService(UserRepo, secret)

	AuthHandler := httpHandlers.NewAuthHandler(AuthService)
	NotesHandler := httpHandlers.NewNoteHandler(NotesService)

	mux := http.NewServeMux()

	mux.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), // URL к JSON документации
	))

	mux.HandleFunc("POST  /user/login", AuthHandler.Login)
	mux.HandleFunc("POST /user/register", AuthHandler.Register)
	mux.HandleFunc("GET /notes", NotesHandler.GetNotes)
	mux.HandleFunc("GET /notes/{id}", NotesHandler.GetNoteHandler)
	mux.HandleFunc("POST /notes", NotesHandler.CreateNote)
	mux.HandleFunc("PUT /notes/{id}", NotesHandler.UpdateNote)
	mux.HandleFunc("DELETE /notes/{id}", NotesHandler.DeleteNote)

	AuthMiddleware := middleware.NewAuthMiddleware(secret)

	authMux := AuthMiddleware.AuthMiddleware(mux)
	loggMux := middleware.Logger(authMux)

	server := &http.Server{
		Addr:    ":8080",
		Handler: loggMux,
	}

	go func() {
		log.Print("Server is runnig...")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Print("Server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown error: %s", err)
	}

	slog.AnyValue("Server gracefully stopped")
}
