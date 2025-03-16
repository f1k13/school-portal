package routes

import (
	"github.com/f1k13/school-portal/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func AuthRouter(r *chi.Mux, authHandler *handlers.AuthHandler) {
	r.Post("/auth/sign-up", authHandler.SignUp)
	r.Post("/auth/init-sign-up", authHandler.InitAuthSignUp)
}
