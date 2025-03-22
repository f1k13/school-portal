package educationDataMapper

import (
	educationAdapter "github.com/f1k13/school-portal/internal/domain/adapter/education"
	"github.com/f1k13/school-portal/internal/domain/models/education"
	educationDto "github.com/f1k13/school-portal/internal/dto/education"
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
