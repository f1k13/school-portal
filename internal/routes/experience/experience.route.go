package experienceRoute

import (
	experienceController "github.com/f1k13/school-portal/internal/controllers/experience"
	"github.com/f1k13/school-portal/internal/middleware/auth"
	"github.com/go-chi/chi/v5"
)

type ExperienceRoute struct {
	experienceController *experienceController.ExperienceController
	router               *chi.Mux
	authMiddleware       *auth.AuthMiddleWare
}

func NewExperienceRouter(r *chi.Mux, experienceController *experienceController.ExperienceController, authMiddleware *auth.AuthMiddleWare) *ExperienceRoute {
	return &ExperienceRoute{
		experienceController: experienceController,
		router:               r,
		authMiddleware:       authMiddleware,
	}
}

func (r *ExperienceRoute) ExperienceRouter() {
	r.router.With(r.authMiddleware.Auth).Post("/experience/post", r.experienceController.CreateExperience)
}
