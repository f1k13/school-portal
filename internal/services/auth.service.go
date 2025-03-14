package services

import (
	"os"
	"time"

	"github.com/f1k13/school-portal/internal/dto"
	"github.com/f1k13/school-portal/internal/logger"
	"github.com/f1k13/school-portal/internal/repositories"
	"github.com/golang-jwt/jwt/v4"
)

type AuthService struct {
	UserRepo *repositories.UserRepository
}

func NewAuthService(userRepo *repositories.UserRepository) *AuthService {
	return &AuthService{UserRepo: userRepo}
}
func (s *AuthService) SignUp(userDto dto.UserDto) (*dto.UserToken, error) {
	u, err := s.UserRepo.CreateUser(userDto)

	if err != nil {
		logger.Log.Error("Error creating user", err)
		return nil, err
	}
	t, err := generateJwt(u.ID.String())
	if err != nil {
		logger.Log.Error("Error generating token", err)
		return nil, err
	}
	return &dto.UserToken{Token: t, User: *u}, nil
}
func generateJwt(id string) (string, error) {
	payload := jwt.MapClaims{
		"sub": id,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	secret := os.Getenv("JWT_SECRET_KEY")
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		logger.Log.Fatal("Error signing token", err)
		return "", err
	}
	return t, nil
}
func sendEmail(email string) {

}

func verifyEmail(email string) {

}
