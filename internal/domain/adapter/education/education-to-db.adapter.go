package educationAdapter

import (
	"time"

	"github.com/f1k13/school-portal/internal/domain/models/education"
	educationDto "github.com/f1k13/school-portal/internal/dto/education"
	"github.com/google/uuid"
)

type EducationToModelAdapter struct{}

func NewEducationToModelAdapter() *EducationToModelAdapter {
	return &EducationToModelAdapter{}
}

func (a *EducationToModelAdapter) CreateEducationAdapter(dto *educationDto.EducationDto) *education.EducationModel {
	createdAt := time.Now()
	return &education.EducationModel{
		ID:          uuid.New(),
		UserID:      *dto.UserID,
		Institution: dto.Institution,
		Degree:      dto.Degree,
		EndYear:     dto.EndYear,
		City:        dto.City,
		StartYear:   dto.StartYear,
		CreatedAt:   &createdAt,
	}
}
