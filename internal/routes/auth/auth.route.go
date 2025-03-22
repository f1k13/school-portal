package authRoute

import (
	authController "github.com/f1k13/school-portal/internal/controllers/auth"
	"github.com/go-chi/chi/v5"
)

type AuthRoute struct {
	authController *authController.AuthController
	router         *chi.Mux
}

func NewAuthRouter(r *chi.Mux, authController *authController.AuthController) *AuthRoute {
	return &AuthRoute{
		authController: authController,
		router:         r,
	}
}

func (r *AuthRoute) AuthRouter() {
	r.router.Post("/auth/sign-up", r.authController.SignUp)
	r.router.Post("/auth/init-sign-up", r.authController.InitAuthSignUp)
	r.router.Post("/auth/init-sign-in", r.authController.InitAuthSignIn)
	r.router.Post("/auth/sign-in", r.authController.SignIn)
}
