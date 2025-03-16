package handlers

import (
	"github.com/f1k13/school-portal/internal/logger"
	"github.com/f1k13/school-portal/internal/services"
	"github.com/f1k13/school-portal/internal/storage/postgres/school-portal/public/model"
	"net/http"
)

type UserHandler struct {
	UserService *services.UserService
}
type UserSelfRes struct {
	Response
	User *model.Users
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

func (h *UserHandler) GetSelf(w http.ResponseWriter, r *http.Request) {
	userID := GetUserIDCtx(r.Context())
	u, err := h.UserService.GetUserByID(userID)
	if err != nil {
		logger.Log.Error("error in get self handler", err)
		res := Response{Message: err.Error()}
		ResponseJson(w, http.StatusBadRequest, res)
		return
	}
	res := UserSelfRes{User: u, Response: Response{Message: "Успешно"}}
	ResponseJson(w, http.StatusOK, res)
}
