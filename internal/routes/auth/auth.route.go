package authRoute

import (
	authHandler "github.com/f1k13/school-portal/internal/handlers/auth"
	"github.com/go-chi/chi/v5"
)

func AuthRouter(r *chi.Mux, authHandler *authHandler.AuthHandler) {
	r.Post("/auth/sign-up", authHandler.SignUp)
	r.Post("/auth/init-sign-up", authHandler.InitAuthSignUp)
	r.Post("/auth/init-sign-in", authHandler.InitAuthSignIn)
	r.Post("/auth/sign-in", authHandler.SignIn)
}
