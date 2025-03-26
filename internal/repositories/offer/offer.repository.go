package offerRepo

import (
	"database/sql"
	"errors"

	offerAdapter "github.com/f1k13/school-portal/internal/domain/adapter/offer"
	"github.com/f1k13/school-portal/internal/domain/models/offer"
	offerDto "github.com/f1k13/school-portal/internal/dto/offer"
	"github.com/f1k13/school-portal/internal/logger"
	"github.com/f1k13/school-portal/internal/storage/postgres/school-portal/public/table"
	"github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
)

type OfferRepository struct {
	db      *sql.DB
	adapter *offerAdapter.OfferToModelAdapter
}

func NewOfferRepository(db *sql.DB, adapter *offerAdapter.OfferToModelAdapter) *OfferRepository {
	return &OfferRepository{db: db, adapter: adapter}
}

func (r *OfferRepository) CreateOffer(dto offerDto.OfferDto) (*offer.OfferModel, error) {
	data := r.adapter.CreateOfferAdapter(dto)
	stmt := table.Offers.INSERT(table.Offers.AllColumns).MODEL(data).RETURNING(table.Offers.AllColumns)

	var dest []offer.OfferModel
	err := stmt.Query(r.db, &dest)
	if err != nil {
		logger.Log.Error("error in create offer", err)
		return nil, err
	}
	if len(dest) == 0 {
		return nil, errors.New("error in create offer")
	}
	return &dest[0], nil
}

func (r *OfferRepository) GetOfferById(id uuid.UUID) (*offer.Offer, error) {
	stmt := table.Offers.SELECT(table.Offers.AllColumns).FROM(table.Offers).WHERE(table.Offers.ID.EQ(postgres.UUID(id)))
	var dest offer.Offer
	err := stmt.Query(r.db, &dest)
	if err != nil {
		if err.Error() == "qrm: no rows in result set" {
			return nil, errors.New("offer not found")
		}
		return nil, err
	}
	return &dest, nil
}

func (r *OfferRepository) GetOfferByIdWithEducationExperienceSkills(id uuid.UUID) (*offer.OfferWithExpEdSkill, error) {
	stmt := table.Offers.
		SELECT(
			table.Offers.AllColumns,
			table.Direction.AllColumns,
		).
		FROM(table.Offers.LEFT_JOIN(table.Direction, table.Direction.ID.EQ(table.Offers.DirectionID))).
		WHERE(table.Offers.ID.EQ(postgres.UUID(id)))

	var dest offer.OfferWithExpEdSkill
	stmt.Query(r.db, &dest)

	eduStmt := table.OfferEducations.
		SELECT(
			table.OfferEducations.AllColumns,
			table.Educations.AllColumns,
		).
		FROM(table.OfferEducations.LEFT_JOIN(table.Educations, table.Educations.ID.EQ(table.OfferEducations.EducationID))).
		WHERE(table.OfferEducations.OfferID.EQ(postgres.UUID(id)))

	if err := eduStmt.Query(r.db, &dest.Educations); err != nil {
		return nil, err
	}

	expStmt := table.OfferExperiences.
		SELECT(
			table.OfferExperiences.AllColumns,
			table.Experiences.AllColumns,
		).
		FROM(table.OfferExperiences.LEFT_JOIN(table.Experiences, table.Experiences.ID.EQ(table.OfferExperiences.ExperienceID))).
		WHERE(table.OfferExperiences.OfferID.EQ(postgres.UUID(id)))

	if err := expStmt.Query(r.db, &dest.Experiences); err != nil {
		return nil, err
	}

	skillStmt := table.OfferSkills.
		SELECT(
			table.OfferSkills.AllColumns,
			table.Skills.AllColumns,
		).
		FROM(table.OfferSkills.LEFT_JOIN(table.Skills, table.Skills.ID.EQ(table.OfferSkills.SkillID))).
		WHERE(table.OfferSkills.OfferID.EQ(postgres.UUID(id)))

	if err := skillStmt.Query(r.db, &dest.Skills); err != nil {
		return nil, err
	}

	return &dest, nil
}

func (r *OfferRepository) CreateOfferEducation(dto offerDto.OfferEducationDto) error {

	data := r.adapter.CreateOfferEducationAdapter(dto)
	var dest []offer.OfferEducationModel
	stmt := table.OfferEducations.INSERT(table.OfferEducations.AllColumns).RETURNING(table.OfferEducations.AllColumns)

	for _, v := range data {
		stmt = stmt.MODEL(v).RETURNING(table.OfferEducations.AllColumns)
	}
	err := stmt.Query(r.db, &dest)
	if err != nil {
		return err
	}
	if len(dest) == 0 {
		return errors.New("error in create offer education")
	}
	return nil
}

func (r *OfferRepository) CreateOfferExperience(dto offerDto.OfferExperienceDto) error {

	data := r.adapter.CreateOfferExperienceAdapter(dto)
	var dest []offer.OfferEducationModel
	stmt := table.OfferExperiences.INSERT(table.OfferExperiences.AllColumns).RETURNING(table.OfferExperiences.AllColumns)

	for _, v := range data {
		stmt = stmt.MODEL(v).RETURNING(table.OfferExperiences.AllColumns)
	}
	err := stmt.Query(r.db, &dest)
	if err != nil {
		return err
	}
	if len(dest) == 0 {
		return errors.New("error in create offer education")
	}
	return nil
}

func (r *OfferRepository) CreateOfferSkill(dto offerDto.OfferSkillDto) error {

	data := r.adapter.CreateOfferSkillAdapter(dto)
	var dest []offer.OfferEducationModel
	stmt := table.OfferSkills.INSERT(table.OfferSkills.AllColumns).RETURNING(table.OfferSkills.AllColumns)

	for _, v := range data {
		stmt = stmt.MODEL(v).RETURNING(table.OfferSkills.AllColumns)
	}
	err := stmt.Query(r.db, &dest)
	if err != nil {
		return err
	}
	if len(dest) == 0 {
		return errors.New("error in create offer education")
	}
	return nil
}
