package offerController

import (
	"encoding/json"
	"net/http"

	"github.com/f1k13/school-portal/internal/controllers"
	offerDto "github.com/f1k13/school-portal/internal/dto/offer"
	"github.com/f1k13/school-portal/internal/logger"
	"github.com/f1k13/school-portal/internal/models/offer"
	offerService "github.com/f1k13/school-portal/internal/services/offer"
)

type OfferController struct {
	offerService *offerService.OfferService
	controllers  *controllers.Controller
}

func NewOfferController(offerService *offerService.OfferService) *OfferController {
	return &OfferController{
		offerService: offerService,
		controllers:  &controllers.Controller{},
	}
}

func (h *OfferController) CreateOffer(w http.ResponseWriter, r *http.Request) {
	var req offerDto.OfferDto
	userID := h.controllers.GetUserIDCtx(r.Context())
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Log.Error("error decoding json", err)
		res := controllers.Response{Message: err.Error()}
		h.controllers.ResponseJson(w, http.StatusBadRequest, res)
	}

	o, err := h.offerService.CreateOffer(req, userID)
	if err != nil {
		logger.Log.Error("error in create offer handler", err)
		res := controllers.Response{Message: err.Error()}
		h.controllers.ResponseJson(w, http.StatusBadRequest, res)
		return
	}

	res := offer.OfferRes{Response: controllers.Response{Message: "Успешно"}, Offer: *o}
	h.controllers.ResponseJson(w, http.StatusCreated, res)

}

func (h *OfferController) GetOffer(w http.ResponseWriter, r *http.Request) {}

func (h *OfferController) CreateEducation(w http.ResponseWriter, r *http.Request) {}

func (h *OfferController) CreateExperiences(w http.ResponseWriter, r *http.Request) {}

func (h *OfferController) CreateSkills(w http.ResponseWriter, r *http.Request) {}
