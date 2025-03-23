package experienceMapper

import (
	experienceAdapter "github.com/f1k13/school-portal/internal/domain/adapter/experience"
	"github.com/f1k13/school-portal/internal/domain/models/experience"
	experienceDto "github.com/f1k13/school-portal/internal/dto/experience"
)

type ExperienceToModelMapper struct {
	adapter *experienceAdapter.ExperienceToModelAdapter
}

func NewExperienceToModelMapper(adapter *experienceAdapter.ExperienceToModelAdapter) *ExperienceToModelMapper {
	return &ExperienceToModelMapper{adapter: adapter}
}

func (d *ExperienceToModelMapper) CreateExperienceMapperToModel(dto []experienceDto.ExperienceDto) []experience.ExperienceModel {
	var models []experience.ExperienceModel

	for _, dto := range dto {
		models = append(models, d.adapter.CreateExperienceAdapter(&dto))
	}
	return models
}
