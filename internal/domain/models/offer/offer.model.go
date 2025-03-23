package offer

import (
	"github.com/f1k13/school-portal/internal/controllers"
	"github.com/f1k13/school-portal/internal/storage/postgres/school-portal/public/model"
	"github.com/google/uuid"
)

type OfferModel = model.Offers

type OfferExperienceModel = model.OfferExperiences

type OfferEducationModel = model.OfferEducations

type Offer struct {
	ID          uuid.UUID `json:"id"`
	Price       int32     `json:"price"`
	UserID      uuid.UUID `json:"userId"`
	DirectionID uuid.UUID `json:"directionId"`
}

type OfferRes struct {
	controllers.Response `json:"response"`
	Offer                Offer `json:"offer"`
}
