package offerRepo

import "database/sql"

type OfferRepository struct {
	db *sql.DB
}

func NewOfferRepository(db *sql.DB) *OfferRepository {
	return &OfferRepository{db: db}
}

func (r *OfferRepository) CreateOffer() {}

func (r *OfferRepository) GetOfferById() {}

func (r *OfferRepository) CreateOfferEducation() {}

func (r *OfferRepository) CreateOfferExperience() {}

func (r *OfferRepository) CreateOfferSkill() {}
