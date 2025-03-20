package educationService

import educationRepo "github.com/f1k13/school-portal/internal/repositories/education"

type EducationService struct {
	educationRepo *educationRepo.EducationRepository
}

func NewEducationService(educationRepo *educationRepo.EducationRepository) *EducationService {
	return &EducationService{
		educationRepo: educationRepo,
	}
}
