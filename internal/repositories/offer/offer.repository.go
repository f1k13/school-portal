package offerRepo

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	offerAdapter "github.com/f1k13/school-portal/internal/domain/adapter/offer"
	"github.com/f1k13/school-portal/internal/domain/models/offer"
	offerDto "github.com/f1k13/school-portal/internal/dto/offer"
	"github.com/f1k13/school-portal/internal/logger"
	"github.com/f1k13/school-portal/internal/storage/postgres/school-portal/public/table"
	"github.com/go-jet/jet/v2/postgres"
	jet "github.com/go-jet/jet/v2/postgres"
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
func (r *OfferRepository) GetOfferByIdWithEducationExperienceSkills(id string) (*offer.OfferWithExpEdSkill, error) {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("invalid UUID format")
	}

	var offerExists bool
	err = table.Offers.
		SELECT(jet.BoolExp(jet.Raw("1"))).
		WHERE(table.Offers.ID.EQ(postgres.UUID(parsedID))).
		Query(r.db, &offerExists)

	if err != nil {
		return nil, fmt.Errorf("offer existence check failed: %v", err)
	}
	if !offerExists {
		return nil, errors.New("offer not found")
	}

	stmt := table.Offers.
		SELECT(
			table.Offers.AllColumns,
			jet.Raw(`COALESCE(
                json_agg(
                    DISTINCT jsonb_build_object(
                        'id', educations.id,
                        'institution', educations.institution,
                        'degree', educations.degree,
                        'start_year', educations.start_year,
                        'end_year', educations.end_year,
                        'city', educations.city
                    )
                ) FILTER (WHERE educations.id IS NOT NULL), '[]'
            )::jsonb`).AS("educations"),
			jet.Raw(`COALESCE(
                json_agg(
                    DISTINCT jsonb_build_object(
                        'id', experiences.id,
                        'company', experiences.company,
                        'role', experiences.role,
                        'years', experiences.years
                    )
                ) FILTER (WHERE experiences.id IS NOT NULL), '[]'
            )::jsonb`).AS("experiences"),
			jet.Raw(`COALESCE(
                json_agg(
                    DISTINCT jsonb_build_object(
                        'id', skills.id,
                        'name', skills.name,
                        'image', skills.image
                    )
                ) FILTER (WHERE skills.id IS NOT NULL), '[]'
            )::jsonb`).AS("skills"),
		).
		FROM(
			table.Offers.
				LEFT_JOIN(table.OfferEducations, table.OfferEducations.OfferID.EQ(table.Offers.ID)).
				LEFT_JOIN(table.Educations, table.Educations.ID.EQ(table.OfferEducations.EducationID)).
				LEFT_JOIN(table.OfferExperiences, table.OfferExperiences.OfferID.EQ(table.Offers.ID)).
				LEFT_JOIN(table.Experiences, table.Experiences.ID.EQ(table.OfferExperiences.ExperienceID)).
				LEFT_JOIN(table.OfferSkills, table.OfferSkills.OfferID.EQ(table.Offers.ID)).
				LEFT_JOIN(table.Skills, table.Skills.ID.EQ(table.OfferSkills.SkillID)),
		).
		WHERE(table.Offers.ID.EQ(postgres.UUID(parsedID))).
		GROUP_BY(table.Offers.ID)

	var rawData struct {
		offer.Offer
		Experiences []byte `db:"experiences"`
		Educations  []byte `db:"educations"`
		Skills      []byte `db:"skills"`
	}

	if err := stmt.Query(r.db, &rawData); err != nil {
		return nil, fmt.Errorf("query failed: %v", err)
	}

	if rawData.Offer.ID == uuid.Nil {
		return nil, errors.New("offer data not found (zero UUID)")
	}

	var result offer.OfferWithExpEdSkill
	result.Offer = rawData.Offer

	if err := json.Unmarshal(rawData.Experiences, &result.Experiences); err != nil {
		return nil, fmt.Errorf("failed to parse experiences: %v", err)
	}
	if err := json.Unmarshal(rawData.Educations, &result.Educations); err != nil {
		return nil, fmt.Errorf("failed to parse educations: %v", err)
	}
	if err := json.Unmarshal(rawData.Skills, &result.Skills); err != nil {
		return nil, fmt.Errorf("failed to parse skills: %v", err)
	}

	return &result, nil
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
