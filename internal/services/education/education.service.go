package educationService

import (
	"errors"

	"github.com/f1k13/school-portal/internal/domain/models/education"
	educationDto "github.com/f1k13/school-portal/internal/dto/education"
	educationRepo "github.com/f1k13/school-portal/internal/repositories/education"
	"github.com/google/uuid"
)

type EducationService struct {
	educationRepo *educationRepo.EducationRepository
}

func NewEducationService(educationRepo *educationRepo.EducationRepository) *EducationService {
	return &EducationService{
		educationRepo: educationRepo,
	}
}

func (s *EducationService) CreateEducation(dto *[]educationDto.EducationDto, userID string) ([]education.EducationModel, error) {
	userIDUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, errors.New("invalid UUID format")
	}

	if len(*dto) == 0 {
		return nil, errors.New("education list is empty")
	}
	for i := range *dto {
		(*dto)[i].UserID = &userIDUUID
		if (*dto)[i].EndYear < (*dto)[i].StartYear {
			return nil, errors.New("end year must be greater than start year")
		}
	}
	e, err := s.educationRepo.CreateEducation(dto)

	if err != nil {
		return nil, err
	}
	return e, nil
}
