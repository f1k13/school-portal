package educationRoute

import (
	educationController "github.com/f1k13/school-portal/internal/controllers/education"
	"github.com/f1k13/school-portal/internal/middleware/auth"
	"github.com/go-chi/chi/v5"
)

type EducationRoute struct {
	educationController *educationController.EducationController
	router              *chi.Mux
	authMiddleware      *auth.AuthMiddleWare
}

func NewEducationRouter(r *chi.Mux, educationController *educationController.EducationController, authMiddleware *auth.AuthMiddleWare) *EducationRoute {
	return &EducationRoute{
		educationController: educationController,
		router:              r,
		authMiddleware:      authMiddleware,
	}
}

func (r *EducationRoute) EducationRouter() {
	r.router.With(r.authMiddleware.Auth).Post("/education/post", r.educationController.CreateEducation)
	r.router.With(r.authMiddleware.Auth).Get("/education/get", r.educationController.GetEducations)
}
