package userController

import (
	"net/http"

	"github.com/f1k13/school-portal/internal/controllers"
	"github.com/f1k13/school-portal/internal/logger"
	userService "github.com/f1k13/school-portal/internal/services/user"
	"github.com/f1k13/school-portal/internal/storage/postgres/school-portal/public/model"
)

type UserController struct {
	UserService *userService.UserService
	controllers *controllers.Controller
}

type UserSelfRes struct {
	controllers.Response
	User *model.Users
}

func NewUserController(userService *userService.UserService) *UserController {
	return &UserController{
		UserService: userService,
		controllers: &controllers.Controller{},
	}
}

func (c *UserController) GetSelf(w http.ResponseWriter, r *http.Request) {
	userID := c.controllers.GetUserIDCtx(r.Context())
	u, err := c.UserService.GetUserByID(userID)
	if err != nil {
		logger.Log.Error("error in get self handler", err)
		res := controllers.Response{Message: err.Error()}
		c.controllers.ResponseJson(w, http.StatusBadRequest, res)
		return
	}
	res := UserSelfRes{User: u, Response: controllers.Response{Message: "Успешно"}}
	c.controllers.ResponseJson(w, http.StatusOK, res)
}
