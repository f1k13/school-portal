package experienceMapper

import (
	experienceAdapter "github.com/f1k13/school-portal/internal/domain/adapter/experience"
	"github.com/f1k13/school-portal/internal/domain/models/experience"
)

type ExperienceToEntityMapper struct {
	adapter *experienceAdapter.ExperienceToEntityAdapter
}

func NewExperienceToEntityMapper(adapter *experienceAdapter.ExperienceToEntityAdapter) *ExperienceToEntityMapper {
	return &ExperienceToEntityMapper{
		adapter: adapter,
	}
}

func (d *ExperienceToEntityMapper) ExperienceMapper(e *[]experience.ExperienceModel) []experience.Experience {
	var models []experience.Experience

	for _, model := range *e {
		models = append(models, d.adapter.ExperienceAdapter(model))
	}
	return models
}
