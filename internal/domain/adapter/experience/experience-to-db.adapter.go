package experienceAdapter

import (
	"github.com/f1k13/school-portal/internal/domain/models/experience"
	experienceDto "github.com/f1k13/school-portal/internal/dto/experience"
	"github.com/google/uuid"
)

type ExperienceToModelAdapter struct{}

func NewExperienceToModelAdapter() *ExperienceToModelAdapter {
	return &ExperienceToModelAdapter{}
}

func (a *ExperienceToModelAdapter) CreateExperienceAdapter(dto *experienceDto.ExperienceDto) experience.ExperienceModel {
	return experience.ExperienceModel{
		ID:      uuid.New(),
		Company: dto.Company,
		UserID:  *dto.UserId,
		Role:    dto.Role,
		Years:   dto.Years,
	}
}
