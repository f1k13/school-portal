package offerService

import (
	"errors"
	"sync"

	"github.com/f1k13/school-portal/internal/domain/models/education"
	"github.com/f1k13/school-portal/internal/domain/models/experience"
	"github.com/f1k13/school-portal/internal/domain/models/offer"
	offerDto "github.com/f1k13/school-portal/internal/dto/offer"
	"github.com/f1k13/school-portal/internal/logger"
	educationRepo "github.com/f1k13/school-portal/internal/repositories/education"
	experienceRepo "github.com/f1k13/school-portal/internal/repositories/experience"
	offerRepo "github.com/f1k13/school-portal/internal/repositories/offer"
	"github.com/google/uuid"
)

type OfferService struct {
	offerRepo *offerRepo.OfferRepository
	expRepo   *experienceRepo.ExperienceRepository
	eduRepo   *educationRepo.EducationRepository
}

func NewOfferService(offerRepo *offerRepo.OfferRepository, expRepo *experienceRepo.ExperienceRepository, eduRepo *educationRepo.EducationRepository) *OfferService {
	return &OfferService{
		offerRepo: offerRepo,
		expRepo:   expRepo,
		eduRepo:   eduRepo,
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

func (s *OfferService) SearchOffers(dto *offerDto.SearchOfferDto) ([]offer.OfferWithExpEdSkill, error) {
	// o, err := s.offerRepo.GetOfferWithFilters(dto)
	// if err != nil {
	// 	logger.Log.Error("error in search offers", err)
	// 	return nil, err
	// }
	// return o, nil
	return nil, nil
}

func (s *OfferService) GetOfferById(id uuid.UUID) (*offer.OfferWithExpEduModel, error) {
	o, err := s.offerRepo.GetOfferById(id)
	if err != nil {
		return nil, err
	}
	offerExp, err := s.offerRepo.GetOfferExperience(o.ID)

	if err != nil {
		return nil, err
	}
	offerEdu, err := s.offerRepo.GetOfferEducation(o.ID)
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	errCh := make(chan error, 2)

	var exp []experience.ExperienceModel
	var edu []education.EducationModel

	if offerExp != nil && len(*offerExp) > 0 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ids := make([]uuid.UUID, len(*offerExp))
			for i, v := range *offerExp {
				ids[i] = v.ExperienceID
			}
			data, err := s.expRepo.GetExperiencesByIds(ids)
			if err != nil {
				errCh <- err
				return
			}
			exp = *data
		}()
	}

	if offerEdu != nil && len(*offerEdu) > 0 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ids := make([]uuid.UUID, len(*offerEdu))
			for i, v := range *offerEdu {
				ids[i] = v.EducationID
			}
			data, err := s.eduRepo.GetEducationsByIds(ids)
			if err != nil {
				errCh <- err
				return
			}
			edu = *data
		}()
	}

	go func() {
		wg.Wait()
		close(errCh)
	}()

	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	return &offer.OfferWithExpEduModel{
		OfferModel: *o,
		Experience: exp,
		Education:  edu,
	}, nil
}
