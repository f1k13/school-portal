package routes

import (
	"database/sql"

	"github.com/f1k13/school-portal/internal/handlers"
	"github.com/f1k13/school-portal/internal/repositories"
	"github.com/f1k13/school-portal/internal/services"
	"github.com/go-chi/chi/v5"
)

func StartRouter(r *chi.Mux, db *sql.DB) {
	userRepo := repositories.NewUserRepository(db)
	authService := services.NewAuthService(userRepo)
	userService := services.NewUserService(userRepo)
	authHandler := handlers.NewAuthHandler(authService, userService)
	AuthRouter(r, authHandler)
}
