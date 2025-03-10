package userRepositories

import (
	"github.com/f1k13/school-portal/internal/handlers/dto"
	"github.com/f1k13/school-portal/internal/logger"
	db "github.com/f1k13/school-portal/internal/models"
	"github.com/f1k13/school-portal/internal/models/user"
	"github.com/f1k13/school-portal/internal/utils"
	"github.com/google/uuid"
)

func CreateUser(userDto dto.UserDto) (user.User, error) {

	u := user.User{
		ID:          uuid.New(),
		FirstName:   userDto.FirstName,
		MiddleName:  userDto.MiddleName,
		SurName:     userDto.SurName,
		PhoneNumber: utils.PtrToStr(userDto.PhoneNumber),
		Email:       userDto.Email,
		Role:        userDto.Role,
	}
	if err := db.DB.Create(&u).Error; err != nil {
		logger.Log.Error("Error creating user", err)
		return user.User{}, nil
	}
	return u, nil
}

func GetUserByEmail(email string) (user.User, error) {
	u := user.User{}
	if email == "" {
		return user.User{}, nil
	}
	err := db.DB.Where("email = ?", email).First(&u).Error
	if err != nil {
		logger.Log.Error("Error getting user by email", err)
		return user.User{}, nil
	}
	return u, nil
}
