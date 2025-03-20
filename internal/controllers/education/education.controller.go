package educationController

import educationService "github.com/f1k13/school-portal/internal/services/education"

type EducationController struct {
	educationService *educationService.EducationService
}

func NewEducationController(educationService *educationService.EducationService) *EducationController {
	return &EducationController{
		educationService: educationService,
	}
}
