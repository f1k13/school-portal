package offerDto

import "github.com/google/uuid"

type OfferDto struct {
	Price       int32      `json:"price"`
	DirectionId *uuid.UUID `json:"directionId"`
	UserId      *uuid.UUID `json:"userId"`
}

type OfferEducationDto struct {
	EducationId *uuid.UUID `json:"educationId"`
	OfferId     *uuid.UUID `json:"offerId"`
}

type OfferExperienceDto struct {
	ExperienceId *uuid.UUID `json:"experienceId"`
	OfferId      *uuid.UUID `json:"offerId"`
}

type OfferSkillDto struct {
	SkillId *uuid.UUID `json:"skillId"`
	OfferId *uuid.UUID `json:"offerId"`
}
