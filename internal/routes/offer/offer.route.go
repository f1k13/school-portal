package offerRoute

import (
	offerController "github.com/f1k13/school-portal/internal/controllers/offer"
	"github.com/f1k13/school-portal/internal/middleware/auth"
	"github.com/go-chi/chi/v5"
)

func OfferRoute(r *chi.Mux, offerController *offerController.OfferController, authMiddleware *auth.AuthMiddleWare) {
	r.With(authMiddleware.Auth).Post("/offer/post", offerController.CreateOffer)
}
