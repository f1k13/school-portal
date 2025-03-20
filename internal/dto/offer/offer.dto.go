package offerDto

import "github.com/google/uuid"

type OfferDto struct {
	Price       int32      `json:"price"`
	DirectionId *uuid.UUID `json:"directionId"`
	UserId      *uuid.UUID `json:"userId"`
}
