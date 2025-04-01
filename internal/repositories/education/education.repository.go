package educationRepo

import (
	"database/sql"
	"errors"

	educationDataMapper "github.com/f1k13/school-portal/internal/domain/data-mapper/education"
	"github.com/f1k13/school-portal/internal/domain/models/education"
	educationDto "github.com/f1k13/school-portal/internal/dto/education"
	"github.com/f1k13/school-portal/internal/storage/postgres/school-portal/public/table"
	"github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
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
	err := stmt.Query(r.db, &dest)
	if err != nil {
		return nil, err
	}
	if len(dest) == 0 {
		return nil, errors.New("error in create education")
	}

	return dest, nil
}

func (r *EducationRepository) GetEducationById(id uuid.UUID) (*education.EducationModel, error) {
	stmt := table.Educations.SELECT(table.Educations.AllColumns).FROM(table.Educations).WHERE(table.Educations.ID.EQ(postgres.UUID(id)))
	var dest education.EducationModel
	err := stmt.Query(r.db, &dest)
	if err != nil {
		if err.Error() == "qrm: no rows in result set" {
			return nil, errors.New("education not found")
		}
		return nil, err
	}
	return &dest, nil
}

func (r *EducationRepository) GetEducations() (*[]education.EducationModel, error) {
	stmt := table.Educations.SELECT(table.Educations.AllColumns).FROM(table.Educations)
	var dest []education.EducationModel
	err := stmt.Query(r.db, &dest)
	if err != nil {
		return nil, err
	}
	if len(dest) == 0 {
		return nil, errors.New("error in get education")
	}
	return &dest, nil
}

func (r *EducationRepository) GetEducationsByIds(eduIDS []uuid.UUID) (*[]education.EducationModel, error) {
	ids := r.mapper.EducationIds(eduIDS)
	stmt := table.Educations.SELECT(table.Educations.AllColumns).FROM(table.Educations).WHERE(table.Educations.ID.IN(ids...))
	var dest []education.EducationModel
	err := stmt.Query(r.db, &dest)
	if err != nil {
		return nil, err
	}
	if len(dest) == 0 {
		return nil, errors.New("error in get education")
	}
	return &dest, nil
}

func (r *EducationRepository) GetEducationsByUserID(dto uuid.UUID) (*[]education.EducationModel, error) {
	stmt := table.Educations.SELECT(table.Educations.AllColumns).FROM(table.Educations).WHERE(table.Educations.UserID.EQ(postgres.UUID(dto)))

	var dest []education.EducationModel

	err := stmt.Query(r.db, &dest)

	if err != nil {
		return nil, err
	}
	return &dest, nil
}
