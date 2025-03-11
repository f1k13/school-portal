package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/f1k13/school-portal/internal/handlers/dto"
	"github.com/f1k13/school-portal/internal/logger"
	"github.com/f1k13/school-portal/internal/models/user"
	"github.com/f1k13/school-portal/internal/services"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	AuthService *services.AuthService
}
type SignUpRes struct {
	Message string    `json:"message"`
	U       user.User `json:"user"`
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{AuthService: authService}
}
func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var userDto dto.UserDto
	if err := json.NewDecoder(r.Body).Decode(&userDto); err != nil {
		http.Error(w, "invalid req body", http.StatusBadRequest)
		logger.Log.Error("error decoding json", err)
		return
	}
	u, err := h.AuthService.SignUp(userDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logger.Log.Error("error sign up method")
		return
	}
	res := SignUpRes{
		Message: "Успешная регистрация",
		U:       u,
	}
	ResponseJson(w, http.StatusCreated, res)
}

func SignIn(c *gin.Context) {}
