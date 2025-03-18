package userService

import (
	"errors"

	"github.com/f1k13/school-portal/internal/dto"
	"github.com/f1k13/school-portal/internal/models/user"
	repositories "github.com/f1k13/school-portal/internal/repositories/user"
	"github.com/f1k13/school-portal/internal/storage/postgres/school-portal/public/model"
	"github.com/google/uuid"
)

type UserService struct {
	UserRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{UserRepo: userRepo}
}

func (s *UserService) GetUserByEmail(email string) (*user.User, error) {
	return s.UserRepo.GetUserByEmail(email)
}

func (s *UserService) GetUserByID(id string) (*user.User, error) {
	return s.UserRepo.GetUserByID(id)
}

func (s *UserService) CreateProfile(dto *dto.UserProfileDto, userID string) (*model.Profiles, error) {
	uuidID, err := uuid.Parse(userID)
	if err != nil {
		return nil, errors.New("invalid UUID format")
	}

	p, err := s.UserRepo.CreateProfile(dto, uuidID)

	if err != nil {
		return nil, err
	}

	return p, nil
}
func (s *UserService) GetProfile(userID string) (*user.UserProfile, error) {
	uuid, err := uuid.Parse(userID)
	if err != nil {
		return nil, errors.New("invalid UUID format")
	}
	p, err := s.UserRepo.GetProfileWithUser(uuid)
	if err != nil {
		return nil, err
	}
	return p, nil
}
