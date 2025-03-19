package offerController

import (
	"net/http"

	offerService "github.com/f1k13/school-portal/internal/services/offer"
)

type OfferController struct {
	offerService *offerService.OfferService
}

func NewOfferController(offerService *offerService.OfferService) *OfferController {
	return &OfferController{
		offerService: offerService,
	}
}

func (h *OfferController) CreateOffer(w http.ResponseWriter, r *http.Request) {}

func (h *OfferController) GetOffer(w http.ResponseWriter, r *http.Request) {}

func (h *OfferController) CreateEducation(w http.ResponseWriter, r *http.Request) {}

func (h *OfferController) CreateExperiences(w http.ResponseWriter, r *http.Request) {}

func (h *OfferController) CreateSkills(w http.ResponseWriter, r *http.Request) {}
