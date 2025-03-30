package experienceRepo

import (
	"database/sql"
	"errors"

	"github.com/f1k13/school-portal/internal/domain/models/experience"
	experienceDto "github.com/f1k13/school-portal/internal/dto/experience"
	"github.com/f1k13/school-portal/internal/storage/postgres/school-portal/public/table"
	"github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"

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

func (r *ExperienceRepository) GetExperienceById(id uuid.UUID) (*experience.ExperienceModel, error) {
	stmt := table.Experiences.SELECT(table.Experiences.AllColumns).FROM(table.Experiences).WHERE(table.Experiences.ID.EQ(postgres.UUID(id)))
	var dest experience.ExperienceModel
	err := stmt.Query(r.db, &dest)
	if err != nil {
		if err.Error() == "qrm: no rows in result set" {
			return nil, errors.New("experience not found")
		}
		return nil, err
	}
	return &dest, nil
}

func (r *ExperienceRepository) GetExperiences() (*[]experience.ExperienceModel, error) {
	stmt := table.Experiences.SELECT(table.Experiences.AllColumns).FROM(table.Experiences)
	var dest []experience.ExperienceModel
	err := stmt.Query(r.db, &dest)
	if err != nil {
		return nil, err
	}
	if len(dest) == 0 {
		return nil, errors.New("error in get exp")
	}
	return &dest, nil
}
func (r *ExperienceRepository) GetExperiencesByIds(expIDS []uuid.UUID) (*[]experience.ExperienceModel, error) {

	ids := r.mapper.GetExperienceIds(expIDS)
	stmt := table.Experiences.SELECT(table.Experiences.AllColumns).FROM(table.Experiences).WHERE(table.Experiences.ID.IN(ids...))
	var dest []experience.ExperienceModel
	err := stmt.Query(r.db, &dest)
	if err != nil {
		return nil, err
	}
	if len(dest) == 0 {
		return nil, errors.New("error in get exp")
	}
	return &dest, nil
}
