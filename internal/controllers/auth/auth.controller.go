package authController

import (
	"encoding/json"
	"net/http"

	controllers "github.com/f1k13/school-portal/internal/controllers"
	"github.com/f1k13/school-portal/internal/dto"
	"github.com/f1k13/school-portal/internal/logger"
	"github.com/f1k13/school-portal/internal/models/auth"
	"github.com/f1k13/school-portal/internal/models/user"
	authService "github.com/f1k13/school-portal/internal/services/auth"
	userService "github.com/f1k13/school-portal/internal/services/user"
)

type AuthController struct {
	AuthService *authService.AuthService
	UserService *userService.UserService
	controllers *controllers.Controller
}

func NewAuthController(authService *authService.AuthService, userService *userService.UserService) *AuthController {
	return &AuthController{
		AuthService: authService,
		UserService: userService,
		controllers: &controllers.Controller{},
	}
}

func (c *AuthController) SignUp(w http.ResponseWriter, r *http.Request) {
	var req auth.AuthCodeReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		res := controllers.Response{Message: err.Error()}
		c.controllers.ResponseJson(w, http.StatusBadRequest, res)
		logger.Log.Error("error decoding json", err)
		return
	}

	u, err := c.AuthService.SignUp(req.Code)
	if err != nil {
		res := controllers.Response{Message: err.Error()}
		c.controllers.ResponseJson(w, http.StatusBadRequest, res)
		logger.Log.Error("error sign up method")
		return
	}
	res := user.UserResponseAuth{
		Response: controllers.Response{Message: "Успешная регистрация"},
		User:     u.User,
		Token:    u.Token,
	}
	c.controllers.ResponseJson(w, http.StatusCreated, res)
}

func (c *AuthController) InitAuthSignUp(w http.ResponseWriter, r *http.Request) {
	var userDto dto.UserDto
	if err := json.NewDecoder(r.Body).Decode(&userDto); err != nil {
		res := controllers.Response{Message: err.Error()}
		c.controllers.ResponseJson(w, http.StatusBadRequest, res)
		logger.Log.Error("error decoding json", err)
		return
	}
	err := c.AuthService.InitSignUp(userDto)
	if err != nil {
		res := controllers.Response{Message: err.Error()}
		c.controllers.ResponseJson(w, http.StatusBadRequest, res)
		logger.Log.Error("error sign up method")
		return
	}
	res := controllers.Response{Message: "Код отправлен на почту"}
	c.controllers.ResponseJson(w, http.StatusCreated, res)
}

func (c *AuthController) InitAuthSignIn(w http.ResponseWriter, r *http.Request) {
	var req auth.SignInReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		res := controllers.Response{Message: err.Error()}
		c.controllers.ResponseJson(w, http.StatusBadRequest, res)
		logger.Log.Error("error decoding json", err)
		return
	}
	err := c.AuthService.InitSignIn(req.Email)
	if err != nil {
		res := controllers.Response{Message: err.Error()}
		c.controllers.ResponseJson(w, http.StatusBadRequest, res)
		logger.Log.Error("error sign in method")
		return
	}
	res := controllers.Response{Message: "Код отправлен на почту"}
	c.controllers.ResponseJson(w, http.StatusCreated, res)
}

func (c *AuthController) SignIn(w http.ResponseWriter, r *http.Request) {
	var req auth.AuthCodeReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		res := controllers.Response{Message: err.Error()}
		c.controllers.ResponseJson(w, http.StatusBadRequest, res)
		logger.Log.Error("error decoding json", err)
		return
	}

	u, err := c.AuthService.SignIn(req.Code)
	if err != nil {
		logger.Log.Error("error in sign in method")
		res := controllers.Response{Message: err.Error()}
		c.controllers.ResponseJson(w, http.StatusBadRequest, res)
		return
	}
	res := user.UserWithToken{
		User:  u.User,
		Token: u.Token,
	}
	c.controllers.ResponseJson(w, http.StatusOK, res)
}
