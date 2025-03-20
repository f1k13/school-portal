package educationRoute

import (
	educationController "github.com/f1k13/school-portal/internal/controllers/education"
	"github.com/f1k13/school-portal/internal/middleware/auth"
	"github.com/go-chi/chi/v5"
)

func EducationRoute(r *chi.Mux, educationController *educationController.EducationController, authMiddleware *auth.AuthMiddleWare) {

}
