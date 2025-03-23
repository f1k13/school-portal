package experienceRepo

import (
	"database/sql"
	"errors"
	"github.com/f1k13/school-portal/internal/domain/models/experience"
	experienceDto "github.com/f1k13/school-portal/internal/dto/experience"
	"github.com/f1k13/school-portal/internal/storage/postgres/school-portal/public/table"

	experienceMapper "github.com/f1k13/school-portal/internal/domain/data-mapper/experience"
)

type ExperienceRepository struct {
	db     *sql.DB
	mapper *experienceMapper.ExperienceToModelMapper
}

func NewExperienceRepository(db *sql.DB, mapper *experienceMapper.ExperienceToModelMapper) *ExperienceRepository {
	return &ExperienceRepository{db: db, mapper: mapper}
}

func (r *ExperienceRepository) CreateExperience(dto *[]experienceDto.ExperienceDto) ([]experience.ExperienceModel, error) {
	data := r.mapper.CreateExperienceMapperToModel(*dto)
	var dest []experience.ExperienceModel

	stmt := table.Experiences.INSERT(table.Experiences.AllColumns)

	for _, v := range data {
		stmt = stmt.MODEL(v).RETURNING(table.Experiences.AllColumns)
	}
	err := stmt.Query(r.db, &dest)
	if err != nil {
		return nil, err
	}
	if len(dest) == 0 {
		return nil, errors.New("error in create exp")
	}
	return dest, nil
}
