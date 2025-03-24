package offer

import (
	"time"

	"github.com/f1k13/school-portal/internal/controllers"
	"github.com/f1k13/school-portal/internal/domain/models/education"
	"github.com/f1k13/school-portal/internal/domain/models/experience"
	"github.com/f1k13/school-portal/internal/storage/postgres/school-portal/public/model"
	"github.com/google/uuid"
)

type OfferModel = model.Offers

type OfferExperienceModel = model.OfferExperiences

type OfferEducationModel = model.OfferEducations

type OfferSkillModel = model.OfferSkills

type OfferWithExpEdSkill struct {
	OfferModel
	Experiences []experience.Experience `json:"experiences"`
	Educations  []education.Education   `json:"educations"`
}
type Offer struct {
	ID          uuid.UUID `json:"id"`
	Price       int32     `json:"price"`
	UserID      uuid.UUID `json:"userId"`
	DirectionID uuid.UUID `json:"directionId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	IsOnline    bool      `json:"isOnline"`
	CreatedAt   time.Time `json:"createdAt"`
}

type OfferRes struct {
	controllers.Response `json:"response"`
	Offer                Offer `json:"offer"`
}
type OfferWithExpEdSkillRes struct {
	controllers.Response `json:"response"`
	Offer                OfferWithExpEdSkill `json:"offer"`
}
