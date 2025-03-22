package offerAdapter

import (
	"github.com/f1k13/school-portal/internal/domain/models/offer"
	offerDto "github.com/f1k13/school-portal/internal/dto/offer"
	"github.com/google/uuid"
)

type OfferToModelAdapter struct{}

func NewOfferToModelAdapter() *OfferToModelAdapter {
	return &OfferToModelAdapter{}
}

func (a *OfferToModelAdapter) CreateOfferAdapter(dto *offerDto.OfferDto) *offer.OfferModel {
	return &offer.OfferModel{ID: uuid.New(), Price: dto.Price, DirectionID: *dto.DirectionId, UserID: *dto.UserId}
}
