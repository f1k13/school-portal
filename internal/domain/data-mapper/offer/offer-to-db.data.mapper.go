package offerDataMapper

import (
	jet "github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
)

type OfferToModelDataMapper struct{}

func NewOfferToModelDataMapper() *OfferToModelDataMapper {
	return &OfferToModelDataMapper{}
}

func (d *OfferToModelDataMapper) MapToIdsToDb(ids []uuid.UUID) []jet.Expression {
	var offerIDExprs []jet.Expression

	for _, id := range ids {
		offerIDExprs = append(offerIDExprs, jet.UUID(id))
	}
	return offerIDExprs
}
