package routes

import (
	"github.com/f1k13/school-portal/internal/handlers"
	db "github.com/f1k13/school-portal/internal/models"
	"github.com/f1k13/school-portal/internal/repositories"
	"github.com/f1k13/school-portal/internal/services"
	"github.com/go-chi/chi/v5"
)

func StartRouter(r *chi.Mux) {
	userRepo := repositories.NewUserRepository(db.DB)
	authService := services.NewAuthService(userRepo)
	authHandler := handlers.NewAuthHandler(authService)
	AuthRouter(r, authHandler)
}
