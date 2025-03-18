package userController

import (
	"encoding/json"
	"net/http"

	"github.com/f1k13/school-portal/internal/controllers"
	"github.com/f1k13/school-portal/internal/dto"
	"github.com/f1k13/school-portal/internal/logger"
	"github.com/f1k13/school-portal/internal/models/user"
	userService "github.com/f1k13/school-portal/internal/services/user"
)

type UserController struct {
	UserService *userService.UserService
	controllers *controllers.Controller
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
	res := user.UserSelfRes{User: u, Response: controllers.Response{Message: "Успешно"}}
	c.controllers.ResponseJson(w, http.StatusOK, res)
}
func (c *UserController) ProfilePost(w http.ResponseWriter, r *http.Request) {
	userID := c.controllers.GetUserIDCtx(r.Context())
	var req dto.UserProfileDto
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		res := controllers.Response{Message: err.Error()}
		c.controllers.ResponseJson(w, http.StatusBadRequest, res)
		logger.Log.Error("error decoding json", err)
		return
	}

	res, err := c.UserService.CreateProfile(&req, userID)
	if err != nil {
		logger.Log.Error("error in create profile handler", err)
		res := controllers.Response{Message: err.Error()}
		c.controllers.ResponseJson(w, http.StatusBadRequest, res)
		return
	}
	c.controllers.ResponseJson(w, http.StatusCreated, res)
}

func (c *UserController) GetProfile(w http.ResponseWriter, r *http.Request) {
	userID := c.controllers.GetUserIDCtx(r.Context())

	profile, err := c.UserService.GetProfile(userID)
	if err != nil {
		logger.Log.Error("error in get profile handler", err)
		res := controllers.Response{Message: err.Error()}
		c.controllers.ResponseJson(w, http.StatusBadRequest, res)
		return
	}

	res := user.UserProfileRes{
		Response: controllers.Response{Message: "Успешно"},
		UserProfile: user.UserProfile{
			User:    profile.User,
			Profile: profile.Profile,
		},
	}

	c.controllers.ResponseJson(w, http.StatusOK, res)
}
