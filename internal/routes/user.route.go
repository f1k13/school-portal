package routes

import (
	"github.com/f1k13/school-portal/internal/handlers"
	"github.com/f1k13/school-portal/internal/middleware"
	"github.com/go-chi/chi/v5"
)

func UserRoute(r *chi.Mux, userHandler *handlers.UserHandler, authMiddleware *middleware.AuthMiddleWare) {

	r.With(authMiddleware.Auth).Get("/user/get-self", userHandler.GetSelf)
}
