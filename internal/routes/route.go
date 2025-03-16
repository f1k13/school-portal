package routes

import (
	"database/sql"

	authHandler "github.com/f1k13/school-portal/internal/handlers/auth"
	userHandler "github.com/f1k13/school-portal/internal/handlers/user"
	"github.com/f1k13/school-portal/internal/middleware"
	repositories "github.com/f1k13/school-portal/internal/repositories/user"
	authRoute "github.com/f1k13/school-portal/internal/routes/auth"
	userRoute "github.com/f1k13/school-portal/internal/routes/user"
	authService "github.com/f1k13/school-portal/internal/services/auth"
	userService "github.com/f1k13/school-portal/internal/services/user"
	"github.com/go-chi/chi/v5"
)

func StartRouter(r *chi.Mux, db *sql.DB) {
	userRepo := repositories.NewUserRepository(db)
	authService := authService.NewAuthService(userRepo)
	userService := userService.NewUserService(userRepo)
	authHandler := authHandler.NewAuthHandler(authService, userService)
	userHandler := userHandler.NewUserHandler(userService)
	authMiddleware := middleware.NewAuthMiddleware()
	authRoute.AuthRouter(r, authHandler)
	userRoute.UserRoute(r, userHandler, authMiddleware)
}
