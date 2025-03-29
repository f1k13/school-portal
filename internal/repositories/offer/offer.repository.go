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

func (r *OfferRepository) GetOfferById(id string) (*offer.Offer, error) {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("invalid UUID format")
	}
	var offer offer.Offer
	err = r.db.QueryRow("SELECT id, price, user_id, direction_id, title, description, is_online, created_at FROM public.offers WHERE id = $1", parsedID).
		Scan(&offer.ID, &offer.Price, &offer.UserID, &offer.DirectionID, &offer.Title, &offer.Description, &offer.IsOnline, &offer.CreatedAt)
	if err != nil {
		if err.Error() == "qrm: no rows in result set" {
			return nil, errors.New("offer not found")
		}
		return nil, err
	}
	return &offer, nil
}

func (r *OfferRepository) GetOfferByIdWithEducationExperienceSkills(id string) (*offer.OfferWithExpEdSkill, error) {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("invalid UUID format")
	}

	var rawData offer.OfferWithExpEdSkillRaw
	query := `
		SELECT
			offers.id,
			offers.price,
			offers.user_id,
			offers.direction_id,
			offers.title,
			offers.description,
			offers.is_online,
			offers.created_at,
			COALESCE(
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
			) AS educations,
			COALESCE(
				json_agg(
					DISTINCT jsonb_build_object(
						'id', experiences.id,
						'company', experiences.company,
						'role', experiences.role,
						'years', experiences.years
					)
				) FILTER (WHERE experiences.id IS NOT NULL), '[]'
			) AS experiences,
			COALESCE(
				json_agg(
					DISTINCT jsonb_build_object(
						'id', skills.id,
						'name', skills.name,
						'image', skills.image
					)
				) FILTER (WHERE skills.id IS NOT NULL), '[]'
			) AS skills
		FROM public.offers
		LEFT JOIN public.offer_educations AS offer_educations ON offer_educations.offer_id = offers.id
		LEFT JOIN public.educations AS educations ON educations.id = offer_educations.education_id
		LEFT JOIN public.offer_experiences AS offer_experiences ON offer_experiences.offer_id = offers.id
		LEFT JOIN public.experiences AS experiences ON experiences.id = offer_experiences.experience_id
		LEFT JOIN public.offer_skills AS offer_skills ON offer_skills.offer_id = offers.id
		LEFT JOIN public.skills AS skills ON skills.id = offer_skills.skill_id
		WHERE offers.id = $1
		GROUP BY
			offers.id,
			offers.price,
			offers.user_id,
			offers.direction_id,
			offers.title,
			offers.description,
			offers.is_online,
			offers.created_at
	`

	if err := r.db.QueryRow(query, parsedID).Scan(
		&rawData.ID,
		&rawData.Price,
		&rawData.UserID,
		&rawData.DirectionID,
		&rawData.Title,
		&rawData.Description,
		&rawData.IsOnline,
		&rawData.CreatedAt,
		&rawData.Experiences,
		&rawData.Educations,
		&rawData.Skills,
	); err != nil {
		return nil, fmt.Errorf("query failed: %v", err)
	}

	var result offer.OfferWithExpEdSkill
	result.Offer = r.adapter.OfferWithExpEduSkillAdapter(&rawData)

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
