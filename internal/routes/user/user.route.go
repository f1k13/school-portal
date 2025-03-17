package userRoute

import (
	userController "github.com/f1k13/school-portal/internal/controllers/user"
	"github.com/f1k13/school-portal/internal/middleware"
	"github.com/go-chi/chi/v5"
)

func UserRoute(r *chi.Mux, userController *userController.UserController, authMiddleware *middleware.AuthMiddleWare) {

	r.With(authMiddleware.Auth).Get("/user/get-self", userController.GetSelf)
	r.With(authMiddleware.Auth).Post("/user/profile/post", userController.ProfilePost)
}
