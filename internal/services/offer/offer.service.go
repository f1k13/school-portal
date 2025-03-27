package offerService

import (
	"errors"
	"sync"

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

func (s *OfferService) CreateOffer(dto offerDto.OfferDto, userID string) (*offer.OfferModel, error) {
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

	offer := offerDto.OfferDto{UserId: &userIDUUID, Price: dto.Price, DirectionId: &directionID, IsOnline: dto.IsOnline, Title: dto.Title, Description: dto.Description}

	o, err := s.offerRepo.CreateOffer(offer)

	if err != nil {
		logger.Log.Error("error in create offer service", err)
		return nil, err
	}

	var wg sync.WaitGroup
	errCh := make(chan error, 3)

	wg.Add(3)
	if dto.ExperienceIDS != nil && len(*dto.ExperienceIDS) > 0 {
		expDto := offerDto.OfferExperienceDto{ExperienceIDS: dto.ExperienceIDS, OfferId: &o.ID}
		go func() {
			defer wg.Done()
			errCh <- s.offerRepo.CreateOfferExperience(expDto)
		}()
	} else {
		wg.Done()
	}
	if dto.EducationIDS != nil && len(*dto.EducationIDS) > 0 {
		educationDto := offerDto.OfferEducationDto{EducationIDS: dto.EducationIDS, OfferId: &o.ID}
		go func() {
			defer wg.Done()
			errCh <- s.offerRepo.CreateOfferEducation(educationDto)
		}()
	} else {
		wg.Done()
	}

	if dto.SkillIDS != nil && len(*dto.SkillIDS) > 0 {
		skillDto := offerDto.OfferSkillDto{SkillIDS: dto.SkillIDS, OfferId: &o.ID}
		go func() {
			defer wg.Done()
			errCh <- s.offerRepo.CreateOfferSkill(skillDto)
		}()
	} else {
		wg.Done()
	}

	go func() {
		wg.Wait()
		close(errCh)
	}()

	for err := range errCh {
		if err != nil {
			logger.Log.Error("error in create offer relation", err)
		}
	}

	return o, nil
}

func (s *OfferService) CreateEducationOffer() {}

func (s *OfferService) CreateExperienceOffer() {}

func (s *OfferService) CreateSkillOffer() {}

func (s *OfferService) GetOfferById(id string) (*offer.OfferWithExpEdSkill, error) {
	o, err := s.offerRepo.GetOfferByIdWithEducationExperienceSkills(id)
	if err != nil {
		logger.Log.Error("error in get offer by id", err)
		return nil, err
	}
	return o, nil
}
