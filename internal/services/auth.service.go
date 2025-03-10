package services

import (
	"github.com/f1k13/school-portal/internal/handlers/dto"
	"github.com/f1k13/school-portal/internal/logger"
	"github.com/f1k13/school-portal/internal/models/user"
	userRepositories "github.com/f1k13/school-portal/internal/repositories"
)

func SignUp(userDto dto.UserDto) (user.User, error) {
	u, err := userRepositories.CreateUser(userDto)
	if err != nil {
		logger.Log.Error("Error creating user", err)
		return user.User{}, nil
	}
	return u, nil
}
func sendEmail(email string) {

}

func verifyEmail(email string) {

}
