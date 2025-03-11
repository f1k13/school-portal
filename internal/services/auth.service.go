package services

import (
	"github.com/f1k13/school-portal/internal/handlers/dto"
	"github.com/f1k13/school-portal/internal/logger"
	"github.com/f1k13/school-portal/internal/models/user"
	"github.com/f1k13/school-portal/internal/repositories"
)

type AuthService struct {
	UserRepo *repositories.UserRepository
}

func NewAuthService(userRepo *repositories.UserRepository) *AuthService {
	return &AuthService{UserRepo: userRepo}
}
func (s *AuthService) SignUp(userDto dto.UserDto) (user.User, error) {
	u, err := s.UserRepo.CreateUser(userDto)
	if err != nil {
		logger.Log.Error("Error creating user", err)
		return user.User{}, err
	}
	return u, nil
}
func sendEmail(email string) {

}

func verifyEmail(email string) {

}
