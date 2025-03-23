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

func (r *OfferRepository) CreateOfferEducation(dto offerDto.OfferEducationDto) (*offer.OfferEducationModel, error) {
	data := r.adapter.CreateOfferEducationAdapter(dto)
	stmt := table.OfferEducations.INSERT(table.OfferEducations.AllColumns).MODEL(data).RETURNING(table.OfferEducations.AllColumns)
	var dest []offer.OfferEducationModel
	err := stmt.Query(r.db, &dest)
	if err != nil {
		return nil, err
	}
	if len(dest) == 0 {
		return nil, errors.New("error in create offer education")
	}
	return &dest[0], nil
}

func (r *OfferRepository) CreateOfferExperience(dto *offerDto.OfferExperienceDto) (*offer.OfferExperienceModel, error) {
	data := r.adapter.CreateOfferExperienceAdapter(*dto)
	stmt := table.OfferExperiences.INSERT(table.OfferExperiences.AllColumns).MODEL(data).RETURNING(table.OfferExperiences.AllColumns)
	var dest []offer.OfferExperienceModel
	err := stmt.Query(r.db, &dest)
	if err != nil {
		return nil, err
	}
	if len(dest) == 0 {
		return nil, errors.New("error in create offer experience")
	}
	return &dest[0], nil
}

func (r *OfferRepository) CreateOfferSkill() {}
