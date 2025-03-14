package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/f1k13/school-portal/internal/dto"
	"github.com/f1k13/school-portal/internal/logger"
	"github.com/f1k13/school-portal/internal/services"
	"github.com/f1k13/school-portal/internal/storage/postgres/school-portal/public/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AuthHandler struct {
	AuthService *services.AuthService
	UserService *services.UserService
}
type Users struct {
	ID          uuid.UUID `json:"id"`
	FirstName   string    `json:"first_name"`
	MiddleName  string    `json:"middle_name"`
	PhoneNumber *string   `json:"phone_number"`
	Email       string    `json:"email"`
	Role        string    `json:"role"`
	CreatedAt   string    `json:"created_at"`
}
type SignUpRes struct {
	Response
	User  model.Users `json:"user"`
	Token string      `json:"token"`
}

func NewAuthHandler(authService *services.AuthService, userService *services.UserService) *AuthHandler {
	return &AuthHandler{AuthService: authService, UserService: userService}
}
func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var userDto dto.UserDto
	if err := json.NewDecoder(r.Body).Decode(&userDto); err != nil {
		res := Response{Message: err.Error()}
		ResponseJson(w, http.StatusBadRequest, res)
		logger.Log.Error("error decoding json", err)
		return
	}

	u, err := h.AuthService.SignUp(userDto)
	if err != nil {
		res := Response{Message: err.Error()}
		ResponseJson(w, http.StatusBadRequest, res)
		logger.Log.Error("error sign up method")
		return
	}
	res := SignUpRes{
		Response: Response{Message: "Успешная регистрация"},
		User:     u.User,
		Token:    u.Token,
	}
	ResponseJson(w, http.StatusCreated, res)
}

func SignIn(c *gin.Context) {}
