package experienceAdapter

import "github.com/f1k13/school-portal/internal/domain/models/experience"

type ExperienceToEntityAdapter struct{}

func NewExperienceToEntityAdapter() *ExperienceToEntityAdapter {
	return &ExperienceToEntityAdapter{}
}

func (a *ExperienceToEntityAdapter) ExperienceAdapter(e experience.ExperienceModel) experience.Experience {
	return experience.Experience{
		ID:        e.ID,
		UserID:    e.UserID,
		Company:   e.Company,
		Years:     e.Years,
		Role:      e.Role,
		CreatedAt: e.CreatedAt,
	}
}
