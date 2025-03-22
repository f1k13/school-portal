package authService

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/f1k13/school-portal/internal/domain/models/user"
	userDto "github.com/f1k13/school-portal/internal/dto/user"
	"github.com/f1k13/school-portal/internal/infrastructure/email"
	"github.com/f1k13/school-portal/internal/logger"
	userRepo "github.com/f1k13/school-portal/internal/repositories/user"
	"github.com/golang-jwt/jwt/v4"
)

type AuthService struct {
	UserRepo     *userRepo.UserRepository
	EmailService email.EmailService
}

func NewAuthService(userRepo *userRepo.UserRepository, emailService *email.EmailService) *AuthService {
	return &AuthService{UserRepo: userRepo, EmailService: *emailService}
}
func (s *AuthService) SignUp(code string) (*user.UserWithToken, error) {
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
	go func() {
		err := s.UserRepo.SetIsAccess(u)
		if err != nil {
			logger.Log.Error("error setting user access", err)
		}
	}()

	return &user.UserWithToken{Token: t, User: *u}, nil
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

func (s *AuthService) InitSignUp(user userDto.UserDto) error {
	userExist, err := s.UserRepo.GetUserByEmail(user.Email)
	if err != nil && err.Error() != "user not found" {
		logger.Log.Error("Error getting user by email", err)
		return err
	}
	if userExist != nil && !userExist.Verified {
		s.InitSignIn(user.Email)
		return nil
	}

	u, err := s.UserRepo.CreateUser(user)
	if err != nil {
		logger.Log.Error("Error creating user", err)
		return err
	}
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

	err = s.EmailService.SendEmail(user.Email, "Verification code", fmt.Sprintf("Your verification code is %s", code))
	logger.Log.Info("code", code)
	if err != nil {
		logger.Log.Error("Error sending email", err)
		return err
	}

	return nil
}
func (s *AuthService) InitSignIn(email string) error {
	u, err := s.UserRepo.GetUserByEmail(email)

	if err != nil {
		logger.Log.Error("Error getting user by email", err)
		return err
	}
	if u == nil {
		return errors.New("user not found")
	}
	code := fmt.Sprintf("%06d", rand.Intn(1000000))
	err = s.UserRepo.SetAuthCode(u, code)
	if err != nil {
		logger.Log.Error("Error saving auth code", err)
		return err
	}
	return nil
}

func (s *AuthService) SignIn(code string) (*user.UserWithToken, error) {
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
	return &user.UserWithToken{
		User:  *u,
		Token: t,
	}, nil
}
