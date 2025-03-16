package authHandler

import (
	"encoding/json"
	"net/http"

	"github.com/f1k13/school-portal/internal/dto"
	"github.com/f1k13/school-portal/internal/handlers"
	"github.com/f1k13/school-portal/internal/logger"
	authService "github.com/f1k13/school-portal/internal/services/auth"
	userService "github.com/f1k13/school-portal/internal/services/user"
	"github.com/f1k13/school-portal/internal/storage/postgres/school-portal/public/model"
)

type AuthHandler struct {
	AuthService *authService.AuthService
	UserService *userService.UserService
	handlers    *handlers.Handlers
}

type AuthRes struct {
	handlers.Response
	User  model.Users `json:"user"`
	Token string      `json:"token"`
}

type AuthCodeReq struct {
	Code string `json:"code"`
}

type SignInReq struct {
	Email string `json:"email"`
}

func NewAuthHandler(authService *authService.AuthService, userService *userService.UserService) *AuthHandler {
	return &AuthHandler{
		AuthService: authService,
		UserService: userService,
		handlers:    &handlers.Handlers{},
	}
}

func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var req AuthCodeReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		res := handlers.Response{Message: err.Error()}
		h.handlers.ResponseJson(w, http.StatusBadRequest, res)
		logger.Log.Error("error decoding json", err)
		return
	}

	u, err := h.AuthService.SignUp(req.Code)
	if err != nil {
		res := handlers.Response{Message: err.Error()}
		h.handlers.ResponseJson(w, http.StatusBadRequest, res)
		logger.Log.Error("error sign up method")
		return
	}
	res := AuthRes{
		Response: handlers.Response{Message: "Успешная регистрация"},
		User:     u.User,
		Token:    u.Token,
	}
	h.handlers.ResponseJson(w, http.StatusCreated, res)
}

func (h *AuthHandler) InitAuthSignUp(w http.ResponseWriter, r *http.Request) {
	var userDto dto.UserDto
	if err := json.NewDecoder(r.Body).Decode(&userDto); err != nil {
		res := handlers.Response{Message: err.Error()}
		h.handlers.ResponseJson(w, http.StatusBadRequest, res)
		logger.Log.Error("error decoding json", err)
		return
	}
	err := h.AuthService.InitSignUp(userDto)
	if err != nil {
		res := handlers.Response{Message: err.Error()}
		h.handlers.ResponseJson(w, http.StatusBadRequest, res)
		logger.Log.Error("error sign up method")
		return
	}
	res := handlers.Response{Message: "Код отправлен на почту"}
	h.handlers.ResponseJson(w, http.StatusCreated, res)
}

func (h *AuthHandler) InitAuthSignIn(w http.ResponseWriter, r *http.Request) {
	var req SignInReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		res := handlers.Response{Message: err.Error()}
		h.handlers.ResponseJson(w, http.StatusBadRequest, res)
		logger.Log.Error("error decoding json", err)
		return
	}
	err := h.AuthService.InitSignIn(req.Email)
	if err != nil {
		res := handlers.Response{Message: err.Error()}
		h.handlers.ResponseJson(w, http.StatusBadRequest, res)
		logger.Log.Error("error sign in method")
		return
	}
	res := handlers.Response{Message: "Код отправлен на почту"}
	h.handlers.ResponseJson(w, http.StatusCreated, res)
}

func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	var req AuthCodeReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		res := handlers.Response{Message: err.Error()}
		h.handlers.ResponseJson(w, http.StatusBadRequest, res)
		logger.Log.Error("error decoding json", err)
		return
	}
	u, err := h.AuthService.SignIn(req.Code)
	if err != nil {
		logger.Log.Error("error in sign in method")
		res := handlers.Response{Message: err.Error()}
		h.handlers.ResponseJson(w, http.StatusBadRequest, res)
		return
	}
	res := dto.UserToken{
		User:  u.User,
		Token: u.Token,
	}
	h.handlers.ResponseJson(w, http.StatusOK, res)
}
