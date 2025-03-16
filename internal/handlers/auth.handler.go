package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/f1k13/school-portal/internal/dto"
	"github.com/f1k13/school-portal/internal/logger"
	"github.com/f1k13/school-portal/internal/services"
	"github.com/f1k13/school-portal/internal/storage/postgres/school-portal/public/model"
)

type AuthHandler struct {
	AuthService *services.AuthService
	UserService *services.UserService
}

type AuthRes struct {
	Response
	User  model.Users `json:"user"`
	Token string      `json:"token"`
}
type AuthCodeReq struct {
	Code string `json:"code"`
}
type SignInReq struct {
	Email string `json:"email"`
}

func NewAuthHandler(authService *services.AuthService, userService *services.UserService) *AuthHandler {
	return &AuthHandler{AuthService: authService, UserService: userService}
}
func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var req AuthCodeReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		res := Response{Message: err.Error()}
		ResponseJson(w, http.StatusBadRequest, res)
		logger.Log.Error("error decoding json", err)
		return
	}

	u, err := h.AuthService.SignUp(req.Code)
	if err != nil {
		res := Response{Message: err.Error()}
		ResponseJson(w, http.StatusBadRequest, res)
		logger.Log.Error("error sign up method")
		return
	}
	res := AuthRes{
		Response: Response{Message: "Успешная регистрация"},
		User:     u.User,
		Token:    u.Token,
	}
	ResponseJson(w, http.StatusCreated, res)
}
func (h *AuthHandler) InitAuthSignUp(w http.ResponseWriter, r *http.Request) {
	var userDto dto.UserDto
	if err := json.NewDecoder(r.Body).Decode(&userDto); err != nil {
		res := Response{Message: err.Error()}
		ResponseJson(w, http.StatusBadRequest, res)
		logger.Log.Error("error decoding json", err)
		return
	}
	err := h.AuthService.InitSignUp(userDto)
	if err != nil {
		res := Response{Message: err.Error()}
		ResponseJson(w, http.StatusBadRequest, res)
		logger.Log.Error("error sign up method")
		return
	}
	res := Response{Message: "Код отправлен на почту"}
	ResponseJson(w, http.StatusCreated, res)
}
func (h *AuthHandler) InitAuthSignIn(w http.ResponseWriter, r *http.Request) {
	var req SignInReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		res := Response{Message: err.Error()}
		ResponseJson(w, http.StatusBadRequest, res)
		logger.Log.Error("error decoding json", err)
		return
	}
	err := h.AuthService.InitSignIn(req.Email)
	if err != nil {
		res := Response{Message: err.Error()}
		ResponseJson(w, http.StatusBadRequest, res)
		logger.Log.Error("error sign up method")
		return
	}
	res := Response{Message: "Код отправлен на почту"}
	ResponseJson(w, http.StatusCreated, res)
}
func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	var req AuthCodeReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		res := Response{Message: err.Error()}
		ResponseJson(w, http.StatusBadRequest, res)
		logger.Log.Error("error decoding json", err)
		return
	}
	u, err := h.AuthService.SignIn(req.Code)
	if err != nil {
		logger.Log.Fatal("error in sign in method")
		res := Response{Message: err.Error()}
		ResponseJson(w, http.StatusBadRequest, res)
		return
	}
	res := dto.UserToken{
		User:  u.User,
		Token: u.Token,
	}
	ResponseJson(w, http.StatusOK, res)

}
