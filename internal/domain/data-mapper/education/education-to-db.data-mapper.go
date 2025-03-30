package educationDataMapper

import (
	educationAdapter "github.com/f1k13/school-portal/internal/domain/adapter/education"
	"github.com/f1k13/school-portal/internal/domain/models/education"
	educationDto "github.com/f1k13/school-portal/internal/dto/education"
	jet "github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
)

type EducationDataMapper struct {
	adapter *educationAdapter.EducationToModelAdapter
}

func NewEducationDataMapper(adapter *educationAdapter.EducationToModelAdapter) *EducationDataMapper {
	return &EducationDataMapper{
		adapter: adapter,
	}
}

func (d *EducationDataMapper) CreateEducationMapperToModel(dto []educationDto.EducationDto) []education.EducationModel {
	var models []education.EducationModel
	for _, dto := range dto {
		models = append(models, *d.adapter.CreateEducationAdapter(&dto))
	}
	return models
}
func (d *EducationDataMapper) EducationIds(ids []uuid.UUID) []jet.Expression {
	var eduIDExprs []jet.Expression

	for _, id := range ids {
		eduIDExprs = append(eduIDExprs, jet.UUID(id))
	}
	return eduIDExprs
}
