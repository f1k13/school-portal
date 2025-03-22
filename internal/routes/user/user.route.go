package userRoute

import (
	userController "github.com/f1k13/school-portal/internal/controllers/user"
	"github.com/f1k13/school-portal/internal/middleware/auth"
	"github.com/go-chi/chi/v5"
)

type UserRoute struct {
	userController *userController.UserController
	router         *chi.Mux
	authMiddleware *auth.AuthMiddleWare
}

func NewUserRouter(r *chi.Mux, userController *userController.UserController, authMiddleware *auth.AuthMiddleWare) *UserRoute {
	return &UserRoute{
		userController: userController,
		router:         r,
		authMiddleware: authMiddleware,
	}
}

func (r *UserRoute) UserRouter() {

	r.router.With(r.authMiddleware.Auth).Get("/user/get-self", r.userController.GetSelf)
	r.router.With(r.authMiddleware.Auth).Post("/user/profile/post", r.userController.ProfilePost)
	r.router.With(r.authMiddleware.Auth).Get("/user/profile/get", r.userController.GetProfile)
}
