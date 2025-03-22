package educationAdapter

import "github.com/f1k13/school-portal/internal/domain/models/education"

type EducationToEntityAdapter struct{}

func NewEducationToEntityAdapter() *EducationToEntityAdapter {
	return &EducationToEntityAdapter{}
}

func (a *EducationToEntityAdapter) EducationAdapter(e *education.EducationModel) *education.Education {
	return &education.Education{
		ID:          e.ID.String(),
		UserID:      e.UserID,
		Institution: e.Institution,
		Degree:      e.Degree,
		StartYear:   e.EndYear,
		EndYear:     e.EndYear,
		City:        e.City,
		CreatedAt:   e.CreatedAt,
	}
}
