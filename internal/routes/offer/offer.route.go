package offerRoute

import (
	offerController "github.com/f1k13/school-portal/internal/controllers/offer"
	"github.com/f1k13/school-portal/internal/middleware/auth"
	"github.com/go-chi/chi/v5"
)

type OfferRoute struct {
	offerController *offerController.OfferController
	router          *chi.Mux
	authMiddleware  *auth.AuthMiddleWare
}

func NewOfferRouter(r *chi.Mux, offerController *offerController.OfferController, authMiddleware *auth.AuthMiddleWare) *OfferRoute {
	return &OfferRoute{
		offerController: offerController,
		router:          r,
		authMiddleware:  authMiddleware,
	}
}

func (r *OfferRoute) OfferRouter() {
	r.router.With(r.authMiddleware.Auth).Post("/offer/post", r.offerController.CreateOffer)
	r.router.With(r.authMiddleware.Auth).Get("/offer/get", r.offerController.GetOfferById)
}
