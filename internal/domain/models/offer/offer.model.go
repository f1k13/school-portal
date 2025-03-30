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
	Offer
	Experiences []experience.Experience `json:"experiences"`
	Educations  []education.Education   `json:"educations"`
}
type Offer struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Price       int32     `json:"price" db:"price"`
	UserID      uuid.UUID `json:"userId" db:"user_id"`
	DirectionID uuid.UUID `json:"directionId" db:"direction_id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	IsOnline    bool      `json:"isOnline" db:"is_online"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
}

type OfferRes struct {
	controllers.Response `json:"response"`
	Offer                Offer `json:"offer"`
}
type OfferWithExpEdSkillRes struct {
	controllers.Response `json:"response"`
	Offer                OfferWithExpEdSkill `json:"offer"`
}
type OfferSearchWithExpEdSkillRes struct {
	controllers.Response `json:"response"`
	Offer                []OfferWithExpEdSkill `json:"offer"`
}

type OfferWithExpEduModel struct {
	OfferModel
	Experience []experience.ExperienceModel
	Education  []education.EducationModel
}
