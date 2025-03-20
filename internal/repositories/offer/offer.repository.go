package offerRepo

import (
	"database/sql"
	"errors"

	offerDto "github.com/f1k13/school-portal/internal/dto/offer"
	"github.com/f1k13/school-portal/internal/logger"
	"github.com/f1k13/school-portal/internal/models/offer"
	"github.com/f1k13/school-portal/internal/storage/postgres/school-portal/public/table"
	"github.com/google/uuid"
)

type OfferRepository struct {
	db *sql.DB
}

func NewOfferRepository(db *sql.DB) *OfferRepository {
	return &OfferRepository{db: db}
}

func (r *OfferRepository) CreateOffer(dto offerDto.OfferDto) (*offer.Offer, error) {
	data := offer.Offer{
		ID:          uuid.New(),
		Price:       dto.Price,
		DirectionID: *dto.DirectionId,
		UserID:      *dto.UserId,
	}
	stmt := table.Offers.INSERT(table.Offers.AllColumns).MODEL(data).RETURNING(table.Offers.AllColumns)

	var dest []offer.Offer
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

func (r *OfferRepository) GetOfferById() {}

func (r *OfferRepository) CreateOfferEducation() {}

func (r *OfferRepository) CreateOfferExperience() {}

func (r *OfferRepository) CreateOfferSkill() {}
