package userController

import (
	"encoding/json"
	"net/http"

	"github.com/f1k13/school-portal/internal/controllers"
	userAdapter "github.com/f1k13/school-portal/internal/domain/adapter/user"
	"github.com/f1k13/school-portal/internal/domain/models/user"
	userDto "github.com/f1k13/school-portal/internal/dto/user"
	"github.com/f1k13/school-portal/internal/logger"
	userService "github.com/f1k13/school-portal/internal/services/user"
)

type UserController struct {
	UserService *userService.UserService
	controllers *controllers.Controller
	adapter     *userAdapter.UserToEntityAdapter
}

func NewUserController(userService *userService.UserService, adapter *userAdapter.UserToEntityAdapter) *UserController {
	return &UserController{
		UserService: userService,
		controllers: &controllers.Controller{},
		adapter:     adapter,
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
	userAdapter := c.adapter.UserAdapter(u)
	res := user.UserSelfRes{User: *userAdapter, Response: controllers.Response{Message: "Успешно"}}
	c.controllers.ResponseJson(w, http.StatusOK, res)
}
func (c *UserController) ProfilePost(w http.ResponseWriter, r *http.Request) {
	userID := c.controllers.GetUserIDCtx(r.Context())
	var req userDto.UserProfileDto
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
	profile := c.adapter.ProfileAdapter(res)
	c.controllers.ResponseJson(w, http.StatusCreated, profile)
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
	result := c.adapter.UserProfileAdapter(&profile.User, &profile.Profile)
	res := user.UserProfileRes{
		Response: controllers.Response{Message: "Успешно"},
		UserProfile: user.UserProfile{
			User:    result.User,
			Profile: result.Profile,
		},
	}

	c.controllers.ResponseJson(w, http.StatusOK, res)
}
