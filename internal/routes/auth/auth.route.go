package authRoute

import (
	authController "github.com/f1k13/school-portal/internal/controllers/auth"
	"github.com/go-chi/chi/v5"
)

func AuthRouter(r *chi.Mux, authController *authController.AuthController) {
	r.Post("/auth/sign-up", authController.SignUp)
	r.Post("/auth/init-sign-up", authController.InitAuthSignUp)
	r.Post("/auth/init-sign-in", authController.InitAuthSignIn)
	r.Post("/auth/sign-in", authController.SignIn)
}
