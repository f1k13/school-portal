package offerService

import (
	"errors"

	"github.com/f1k13/school-portal/internal/domain/models/offer"
	offerDto "github.com/f1k13/school-portal/internal/dto/offer"
	"github.com/f1k13/school-portal/internal/logger"
	offerRepo "github.com/f1k13/school-portal/internal/repositories/offer"
	"github.com/google/uuid"
)

type OfferService struct {
	offerRepo *offerRepo.OfferRepository
}

func NewOfferService(offerRepo *offerRepo.OfferRepository) *OfferService {
	return &OfferService{
		offerRepo: offerRepo,
	}
}

func (s *OfferService) CreateOffer(dto offerDto.OfferDto, userID string) (*offer.Offer, error) {
	userIDUUID, err := uuid.Parse(userID)
	if err != nil {
		logger.Log.Error("error parsing uuid", err)
		return nil, errors.New("invalid UUID format")
	}
	directionID, err := uuid.Parse(dto.DirectionId.String())
	if err != nil {
		logger.Log.Error("error parsing uuid", err)
		return nil, errors.New("invalid UUID format")
	}

	offer := offerDto.OfferDto{UserId: &userIDUUID, Price: dto.Price, DirectionId: &directionID}

	o, err := s.offerRepo.CreateOffer(offer)
	if err != nil {
		logger.Log.Error("error in create offer service", err)
		return nil, err
	}
	return o, nil
}

func (s *OfferService) CreateEducationOffer() {}

func (s *OfferService) CreateExperienceOffer() {}

func (s *OfferService) CreateSkillOffer() {}

func (s *OfferService) GetOffer() {}
