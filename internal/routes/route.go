package routes

import (
	"database/sql"

	authController "github.com/f1k13/school-portal/internal/controllers/auth"
	offerController "github.com/f1k13/school-portal/internal/controllers/offer"
	userController "github.com/f1k13/school-portal/internal/controllers/user"
	"github.com/f1k13/school-portal/internal/infrastructure/email"
	"github.com/f1k13/school-portal/internal/middleware/auth"
	"github.com/f1k13/school-portal/internal/repositories/offer"
	userRepo "github.com/f1k13/school-portal/internal/repositories/user"
	authRoute "github.com/f1k13/school-portal/internal/routes/auth"
	offerRoute "github.com/f1k13/school-portal/internal/routes/offer"
	userRoute "github.com/f1k13/school-portal/internal/routes/user"
	authService "github.com/f1k13/school-portal/internal/services/auth"
	offerService "github.com/f1k13/school-portal/internal/services/offer"
	userService "github.com/f1k13/school-portal/internal/services/user"
	"github.com/go-chi/chi/v5"
)

func StartRouter(r *chi.Mux, db *sql.DB) {
	userRepo := userRepo.NewUserRepository(db)
	offerRepo := offerRepo.NewOfferRepository(db)

	emailService := email.NewEmailInfrastructure()

	authService := authService.NewAuthService(userRepo, emailService)
	userService := userService.NewUserService(userRepo)
	offerService := offerService.NewOfferService(offerRepo)

	authController := authController.NewAuthController(authService)
	userController := userController.NewUserController(userService)
	offerController := offerController.NewOfferController(offerService)
	authMiddleware := auth.NewAuthMiddleware()

	authRoute.AuthRouter(r, authController)
	userRoute.UserRoute(r, userController, authMiddleware)
	offerRoute.OfferRoute(r, offerController, authMiddleware)
}
