package educationRepo

import (
	"database/sql"
	"errors"

	educationDataMapper "github.com/f1k13/school-portal/internal/domain/data-mapper/education"
	"github.com/f1k13/school-portal/internal/domain/models/education"
	educationDto "github.com/f1k13/school-portal/internal/dto/education"
	"github.com/f1k13/school-portal/internal/storage/postgres/school-portal/public/table"
)

type EducationRepository struct {
	db     *sql.DB
	mapper *educationDataMapper.EducationDataMapper
}

func NewEducationRepository(db *sql.DB, mapper *educationDataMapper.EducationDataMapper) *EducationRepository {
	return &EducationRepository{db: db, mapper: mapper}
}

func (r *EducationRepository) CreateEducation(dto *[]educationDto.EducationDto) ([]education.EducationModel, error) {
	data := r.mapper.CreateEducationMapperToModel(*dto)
	var dest []education.EducationModel
	stmt := table.Educations.INSERT(table.Educations.AllColumns)
	for _, v := range data {
		stmt = stmt.MODEL(v).RETURNING(table.Educations.AllColumns)

	}
	err := stmt.RETURNING(table.Educations.AllColumns).Query(r.db, &dest)
	if err != nil {
		return nil, err
	}
	if len(dest) == 0 {
		return nil, errors.New("error in create education")
	}

	return dest, nil
}
