package offerDto

import "github.com/google/uuid"

type OfferDto struct {
	Price         int32        `json:"price"`
	DirectionId   *uuid.UUID   `json:"directionId"`
	UserId        *uuid.UUID   `json:"userId"`
	IsOnline      bool         `json:"isOnline"`
	Title         string       `json:"title"`
	Description   string       `json:"description"`
	ExperienceIDS *[]uuid.UUID `json:"experienceIds"`
	EducationIDS  *[]uuid.UUID `json:"educationIds"`
	SkillIDS      *[]uuid.UUID `json:"skillIds"`
}

type OfferEducationDto struct {
	EducationIDS *[]uuid.UUID `json:"educationIds"`
	OfferId      *uuid.UUID   `json:"offerId"`
}

type OfferExperienceDto struct {
	ExperienceIDS *[]uuid.UUID `json:"experienceIds"`
	OfferId       *uuid.UUID   `json:"offerId"`
}

type OfferSkillDto struct {
	SkillIDS *[]uuid.UUID `json:"skillIds"`
	OfferId  *uuid.UUID   `json:"offerId"`
}
