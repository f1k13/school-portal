package userAdapter

import (
	"github.com/f1k13/school-portal/internal/domain/models/user"
	userDto "github.com/f1k13/school-portal/internal/dto/user"
	"github.com/google/uuid"
)

type UserToModelAdapter struct{}

func NewUserToModelAdapter() *UserToModelAdapter {
	return &UserToModelAdapter{}
}

func (a *UserToModelAdapter) CreateUserAdapter(dto *userDto.UserDto) *user.UserModel {
	return &user.UserModel{
		ID:    uuid.New(),
		Email: dto.Email,
		Role:  dto.Role,
	}
}

func (a *UserToModelAdapter) CreateProfileAdapter(dto *userDto.UserProfileDto) *user.ProfileModel {
	return &user.ProfileModel{
		ID:          uuid.New(),
		FirstName:   dto.FirstName,
		LastName:    dto.LastName,
		PhoneNumber: dto.PhoneNumber,
		AvatarURL:   dto.AvatarUrl,
		Dob:         dto.Dob,
		UserID:      *dto.UserId,
	}
}
