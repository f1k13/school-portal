package offerDataMapper

import (
	offerAdapter "github.com/f1k13/school-portal/internal/domain/adapter/offer"
	"github.com/f1k13/school-portal/internal/domain/models/offer"
)

type OfferToEntityDataMapper struct {
	adapter *offerAdapter.OfferToEntityAdapter
}

func NewOfferToEntityDataMapper(adapter *offerAdapter.OfferToEntityAdapter) *OfferToEntityDataMapper {
	return &OfferToEntityDataMapper{
		adapter: adapter,
	}
}

func (d *OfferToEntityDataMapper) OfferDataMapper(o *[]offer.OfferModel) *[]offer.Offer {
	var model []offer.Offer

	for _, v := range *o {
		model = append(model, *d.adapter.OfferAdapter(&v))
	}
	return &model
}
