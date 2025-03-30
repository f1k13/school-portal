package offerAdapter

import (
	"time"

	"github.com/f1k13/school-portal/internal/domain/models/offer"
	offerDto "github.com/f1k13/school-portal/internal/dto/offer"
	"github.com/google/uuid"
)

type OfferToModelAdapter struct{}

func NewOfferToModelAdapter() *OfferToModelAdapter {
	return &OfferToModelAdapter{}
}

func (a *OfferToModelAdapter) CreateOfferAdapter(dto offerDto.OfferDto) offer.OfferModel {
	createdAt := time.Now()
	return offer.OfferModel{ID: uuid.New(), Price: dto.Price, DirectionID: *dto.DirectionId, UserID: *dto.UserId, Title: dto.Title, Description: dto.Description, IsOnline: dto.IsOnline, CreatedAt: &createdAt}
}

func (a *OfferToModelAdapter) CreateOfferEducationAdapter(dto offerDto.OfferEducationDto) []offer.OfferEducationModel {
	var models []offer.OfferEducationModel
	for _, educationID := range *dto.EducationIDS {
		models = append(models, offer.OfferEducationModel{
			ID:          uuid.New(),
			EducationID: educationID,
			OfferID:     *dto.OfferId,
		})
	}
	return models
}

func (a *OfferToModelAdapter) CreateOfferExperienceAdapter(dto offerDto.OfferExperienceDto) []offer.OfferExperienceModel {
	var models []offer.OfferExperienceModel
	for _, experienceID := range *dto.ExperienceIDS {
		models = append(models, offer.OfferExperienceModel{
			ID:           uuid.New(),
			ExperienceID: experienceID,
			OfferID:      *dto.OfferId,
		})
	}
	return models
}

func (a *OfferToModelAdapter) CreateOfferSkillAdapter(dto offerDto.OfferSkillDto) []offer.OfferSkillModel {
	var models []offer.OfferSkillModel
	for _, skillID := range *dto.SkillIDS {
		models = append(models, offer.OfferSkillModel{
			ID:      uuid.New(),
			SkillID: skillID,
			OfferID: *dto.OfferId,
		})
	}
	return models
}
