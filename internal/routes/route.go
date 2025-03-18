package routes

import (
	"database/sql"

	authController "github.com/f1k13/school-portal/internal/controllers/auth"
	userController "github.com/f1k13/school-portal/internal/controllers/user"
	"github.com/f1k13/school-portal/internal/infrastructure/email"
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

	emailService := email.NewEmailInfrastructure()

	authService := authService.NewAuthService(userRepo, emailService)
	userService := userService.NewUserService(userRepo)

	authController := authController.NewAuthController(authService)
	userController := userController.NewUserController(userService)

	authMiddleware := middleware.NewAuthMiddleware()

	authRoute.AuthRouter(r, authController)
	userRoute.UserRoute(r, userController, authMiddleware)
}
