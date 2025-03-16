package services

import (
	"errors"
	"fmt"
	"math/rand"
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
func (s *AuthService) SignUp(code string) (*dto.UserToken, error) {
	u, err := s.UserRepo.GetUserByAuthCode(code)
	if err != nil {
		logger.Log.Error("Error getting user by auth code", err)
		return nil, err
	}
	t, err := generateJwt(u.ID.String(), time.Now().Add(time.Hour*72).Unix())
	if err != nil {
		logger.Log.Error("Error generating token", err)
		return nil, err
	}
	return &dto.UserToken{Token: t, User: *u}, nil
}
func generateJwt(id string, time int64) (string, error) {
	payload := jwt.MapClaims{
		"sub": id,
		"exp": time,
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

func (s *AuthService) InitSignUp(user dto.UserDto) error {
	userExist, err := s.UserRepo.GetUserByEmail(user.Email)
	if err != nil && err.Error() != "user not found" {
		logger.Log.Error("Error getting user by email", err)
		return err
	}
	if userExist != nil {
		return errors.New("user already exists")
	}

	u, err := s.UserRepo.CreateUser(user)
	refreshT, err := generateJwt(u.ID.String(), time.Now().Add(time.Hour*72).Unix())
	if err != nil {
		return err
	}
	err = s.UserRepo.SetRefreshToken(u, refreshT)
	if err != nil {
		logger.Log.Error("Error creating user", err)
		return err
	}

	code := fmt.Sprintf("%06d", rand.Intn(1000000))

	err = s.UserRepo.SetAuthCode(u, code)
	if err != nil {
		logger.Log.Error("Error saving auth code", err)
		return err
	}

	// err = sendEmail(user.Email, code)
	logger.Log.Info("code", code)
	if err != nil {
		logger.Log.Error("Error sending email", err)
		return err
	}

	return nil
}
func sendEmail(email string, code string) {

}

func verifyEmail(email string) {

}
