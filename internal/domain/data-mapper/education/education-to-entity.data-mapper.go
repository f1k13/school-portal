package educationDataMapper

import (
	educationAdapter "github.com/f1k13/school-portal/internal/domain/adapter/education"
	"github.com/f1k13/school-portal/internal/domain/models/education"
)

type EducationToEntityDataMapper struct {
	adapter *educationAdapter.EducationToEntityAdapter
}

func NewEducationToEntityDataMapper(adapter *educationAdapter.EducationToEntityAdapter) *EducationToEntityDataMapper {
	return &EducationToEntityDataMapper{
		adapter: adapter,
	}
}

func (a *EducationToEntityDataMapper) EducationAdapter(e *[]education.EducationModel) *[]education.Education {
	var educations []education.Education
	for _, e := range *e {
		educations = append(educations, *a.adapter.EducationAdapter(&e))
	}
	return &educations
}
