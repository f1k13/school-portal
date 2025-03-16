package userHandler

import (
	"net/http"

	"github.com/f1k13/school-portal/internal/handlers"
	"github.com/f1k13/school-portal/internal/logger"
	userService "github.com/f1k13/school-portal/internal/services/user"
	"github.com/f1k13/school-portal/internal/storage/postgres/school-portal/public/model"
)

type UserHandler struct {
	UserService *userService.UserService
	handlers    *handlers.Handlers
}

type UserSelfRes struct {
	handlers.Response
	User *model.Users
}

func NewUserHandler(userService *userService.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
		handlers:    &handlers.Handlers{},
	}
}

func (h *UserHandler) GetSelf(w http.ResponseWriter, r *http.Request) {
	userID := h.handlers.GetUserIDCtx(r.Context()) // Теперь вызываем через h.handlers
	u, err := h.UserService.GetUserByID(userID)
	if err != nil {
		logger.Log.Error("error in get self handler", err)
		res := handlers.Response{Message: err.Error()}
		h.handlers.ResponseJson(w, http.StatusBadRequest, res)
		return
	}
	res := UserSelfRes{User: u, Response: handlers.Response{Message: "Успешно"}}
	h.handlers.ResponseJson(w, http.StatusOK, res)
}
