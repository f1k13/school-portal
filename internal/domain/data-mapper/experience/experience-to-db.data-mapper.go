package experienceMapper

import (
	experienceAdapter "github.com/f1k13/school-portal/internal/domain/adapter/experience"
	"github.com/f1k13/school-portal/internal/domain/models/experience"
	experienceDto "github.com/f1k13/school-portal/internal/dto/experience"
	jet "github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
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

func (d *ExperienceToModelMapper) GetExperienceIds(ids []uuid.UUID) []jet.Expression {
	var expIDExprs []jet.Expression

	for _, id := range ids {
		expIDExprs = append(expIDExprs, jet.UUID(id))
	}
	return expIDExprs
}
func (d *ExperienceToModelMapper) GetExperienceYears(years []int32) []jet.Expression {
	var expYears []jet.Expression

	for _, year := range years {
		expYears = append(expYears, jet.Int32(year))
	}
	return expYears
}
