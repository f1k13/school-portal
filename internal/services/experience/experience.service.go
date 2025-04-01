package experienceService

import (
	"errors"
	"github.com/f1k13/school-portal/internal/domain/models/experience"
	experienceDto "github.com/f1k13/school-portal/internal/dto/experience"
	experienceRepo "github.com/f1k13/school-portal/internal/repositories/experience"
	"github.com/google/uuid"
)

type ExperienceService struct {
	experienceRepo *experienceRepo.ExperienceRepository
}

func NewExperienceService(experienceRepo *experienceRepo.ExperienceRepository) *ExperienceService {
	return &ExperienceService{
		experienceRepo: experienceRepo,
	}
}

func (s *ExperienceService) CreateExperience(dto *[]experienceDto.ExperienceDto, userID string) (*[]experience.ExperienceModel, error) {
	userIDUUID, err := uuid.Parse(userID)

	if err != nil {
		return nil, errors.New("invalid user id")
	}
	if len(*dto) == 0 {
		return nil, errors.New("experience is empty")
	}
	for i := range *dto {
		(*dto)[i].UserId = &userIDUUID
	}
	e, err := s.experienceRepo.CreateExperience(dto)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (s *ExperienceService) GetMyExperience(dto uuid.UUID) (*[]experience.ExperienceModel, error) {
	return s.experienceRepo.GetExperiencesByUserID(dto)
}
