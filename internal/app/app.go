package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	educationAdapter "github.com/f1k13/school-portal/internal/domain/adapter/education"
	experienceAdapter "github.com/f1k13/school-portal/internal/domain/adapter/experience"
	offerAdapter "github.com/f1k13/school-portal/internal/domain/adapter/offer"
	userAdapter "github.com/f1k13/school-portal/internal/domain/adapter/user"
	educationDataMapper "github.com/f1k13/school-portal/internal/domain/data-mapper/education"
	experienceMapper "github.com/f1k13/school-portal/internal/domain/data-mapper/experience"
	"github.com/f1k13/school-portal/internal/logger"
	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"

	authController "github.com/f1k13/school-portal/internal/controllers/auth"
	educationController "github.com/f1k13/school-portal/internal/controllers/education"
	experienceController "github.com/f1k13/school-portal/internal/controllers/experience"
	offerController "github.com/f1k13/school-portal/internal/controllers/offer"
	userController "github.com/f1k13/school-portal/internal/controllers/user"
	"github.com/f1k13/school-portal/internal/infrastructure/email"
	"github.com/f1k13/school-portal/internal/middleware/auth"
	educationRepo "github.com/f1k13/school-portal/internal/repositories/education"
	experienceRepo "github.com/f1k13/school-portal/internal/repositories/experience"
	offerRepo "github.com/f1k13/school-portal/internal/repositories/offer"
	userRepo "github.com/f1k13/school-portal/internal/repositories/user"
	authRoute "github.com/f1k13/school-portal/internal/routes/auth"
	educationRoute "github.com/f1k13/school-portal/internal/routes/education"
	experienceRoute "github.com/f1k13/school-portal/internal/routes/experience"
	offerRoute "github.com/f1k13/school-portal/internal/routes/offer"
	userRoute "github.com/f1k13/school-portal/internal/routes/user"
	authService "github.com/f1k13/school-portal/internal/services/auth"
	educationService "github.com/f1k13/school-portal/internal/services/education"
	experienceService "github.com/f1k13/school-portal/internal/services/experience"
	offerService "github.com/f1k13/school-portal/internal/services/offer"
	userService "github.com/f1k13/school-portal/internal/services/user"
)

func App() {
	ConnectDB()
	StartApp()
}

var DB *sql.DB

func ConnectDB() error {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return fmt.Errorf("DATABASE_URL is not set")
	}

	var err error
	DB, err = sql.Open("postgres", dbURL)
	if err != nil {
		logger.Log.Fatalf("Не удалось подключиться к базе данных: %v", err)
		return err
	}

	if err := DB.Ping(); err != nil {
		log.Fatalf("Не удалось пинговать базу данных: %v", err)
		return err
	}

	log.Println("Успешно подключено к базе данных")
	return nil
}

func StartApp() {

	r := chi.NewRouter()
	logger.Log.Info("SERVER START ON PORT", 3000)

	userToModelAdapter := userAdapter.NewUserToModelAdapter()
	offerToModelAdapter := offerAdapter.NewOfferToModelAdapter()
	educationToModelAdapter := educationAdapter.NewEducationToModelAdapter()
	experienceToModelAdapter := experienceAdapter.NewExperienceToModelAdapter()
	experienceToEntityAdapter := experienceAdapter.NewExperienceToEntityAdapter()
	offerToEntityAdapter := offerAdapter.NewOfferToEntityAdapter()
	educationToEntityAdapter := educationAdapter.NewEducationToEntityAdapter()

	educationToModelDataMapper := educationDataMapper.NewEducationDataMapper(educationToModelAdapter)
	experienceToModelDataMapper := experienceMapper.NewExperienceToModelMapper(experienceToModelAdapter)

	educationToEntityDataMapper := educationDataMapper.NewEducationToEntityDataMapper(educationToEntityAdapter)
	experienceToEntityDataMapper := experienceMapper.NewExperienceToEntityMapper(experienceToEntityAdapter)

	userRepo := userRepo.NewUserRepository(DB, userToModelAdapter)
	offerRepo := offerRepo.NewOfferRepository(DB, offerToModelAdapter)
	educationRepo := educationRepo.NewEducationRepository(DB, educationToModelDataMapper)
	experienceRepo := experienceRepo.NewExperienceRepository(DB, experienceToModelDataMapper)

	emailService := email.NewEmailInfrastructure()

	authService := authService.NewAuthService(userRepo, emailService)
	userService := userService.NewUserService(userRepo)
	offerService := offerService.NewOfferService(offerRepo)
	educationService := educationService.NewEducationService(educationRepo)
	experienceService := experienceService.NewExperienceService(experienceRepo)

	authController := authController.NewAuthController(authService)
	userController := userController.NewUserController(userService)
	offerController := offerController.NewOfferController(offerService, offerToEntityAdapter)
	educationController := educationController.NewEducationController(educationService, educationToEntityDataMapper)
	experienceController := experienceController.NewExperienceController(experienceService, experienceToEntityDataMapper)

	authMiddleware := auth.NewAuthMiddleware()

	authRoutes := authRoute.NewAuthRouter(r, authController)
	authRoutes.AuthRouter()
	userRoutes := userRoute.NewUserRouter(r, userController, authMiddleware)
	userRoutes.UserRouter()
	offerRoutes := offerRoute.NewOfferRouter(r, offerController, authMiddleware)
	offerRoutes.OfferRouter()
	educationRoutes := educationRoute.NewEducationRouter(r, educationController, authMiddleware)
	educationRoutes.EducationRouter()
	experienceRoutes := experienceRoute.NewExperienceRouter(r, experienceController, authMiddleware)
	experienceRoutes.ExperienceRouter()
	http.ListenAndServe(":3000", r)

}
