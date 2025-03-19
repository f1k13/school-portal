package offerService

import offerRepo "github.com/f1k13/school-portal/internal/repositories/offer"

type OfferService struct {
	offerRepo *offerRepo.OfferRepository
}

func NewOfferService(offerRepo *offerRepo.OfferRepository) *OfferService {
	return &OfferService{
		offerRepo: offerRepo,
	}
}

func (s *OfferService) CreateOffer() {}

func (s *OfferService) CreateEducationOffer() {}

func (s *OfferService) CreateExperienceOffer() {}

func (s *OfferService) CreateSkillOffer() {}

func (s *OfferService) GetOffer() {}
