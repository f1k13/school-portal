package services

import (
	"github.com/f1k13/school-portal/internal/repositories"
	"github.com/f1k13/school-portal/internal/storage/postgres/school-portal/public/model"
)

type UserService struct {
	UserRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{UserRepo: userRepo}
}

func (s *UserService) GetUserByEmail(email string) (*model.Users, error) {
	return s.UserRepo.GetUserByEmail(email)
}

func (s *UserService) GetUserByID(id string) (*model.Users, error) {
	return s.UserRepo.GetUserByID(id)
}
