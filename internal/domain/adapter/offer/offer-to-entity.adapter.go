package offerAdapter

import "github.com/f1k13/school-portal/internal/domain/models/offer"

type OfferToEntityAdapter struct{}

func NewOfferToEntityAdapter() *OfferToEntityAdapter {
	return &OfferToEntityAdapter{}
}

func (a *OfferToEntityAdapter) OfferAdapter(o *offer.OfferModel) *offer.Offer {
	return &offer.Offer{ID: o.ID, Price: o.Price, DirectionID: o.DirectionID, UserID: o.UserID}
}
