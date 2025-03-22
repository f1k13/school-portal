package offer

import (
	"github.com/f1k13/school-portal/internal/controllers"
	"github.com/f1k13/school-portal/internal/storage/postgres/school-portal/public/model"
)

type Offer = model.Offers

type OfferRes struct {
	controllers.Response
	Offer Offer
}
