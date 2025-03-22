package offerController

import (
	"encoding/json"
	"net/http"

	"github.com/f1k13/school-portal/internal/controllers"
	offerAdapter "github.com/f1k13/school-portal/internal/domain/adapter/offer"
	"github.com/f1k13/school-portal/internal/domain/models/offer"
	offerDto "github.com/f1k13/school-portal/internal/dto/offer"
	"github.com/f1k13/school-portal/internal/logger"
	offerService "github.com/f1k13/school-portal/internal/services/offer"
)

type OfferController struct {
	offerService *offerService.OfferService
	controllers  *controllers.Controller
	adapter      *offerAdapter.OfferToEntityAdapter
}

func NewOfferController(offerService *offerService.OfferService, adapter *offerAdapter.OfferToEntityAdapter) *OfferController {
	return &OfferController{
		offerService: offerService,
		controllers:  &controllers.Controller{},
		adapter:      adapter,
	}
}

func (c *OfferController) CreateOffer(w http.ResponseWriter, r *http.Request) {
	var req offerDto.OfferDto
	userID := c.controllers.GetUserIDCtx(r.Context())
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Log.Error("error decoding json", err)
		res := controllers.Response{Message: err.Error()}
		c.controllers.ResponseJson(w, http.StatusBadRequest, res)
		return
	}

	o, err := c.offerService.CreateOffer(req, userID)
	if err != nil {
		logger.Log.Error("error in create offer handler", err)
		res := controllers.Response{Message: err.Error()}
		c.controllers.ResponseJson(w, http.StatusBadRequest, res)
		return
	}
	adapterOffer := c.adapter.OfferAdapter(o)
	res := offer.OfferRes{Response: controllers.Response{Message: "Успешно"}, Offer: *adapterOffer}
	c.controllers.ResponseJson(w, http.StatusCreated, res)

}

func (c *OfferController) GetOffer(w http.ResponseWriter, r *http.Request) {}

func (ch *OfferController) CreateEducation(w http.ResponseWriter, r *http.Request) {}

func (c *OfferController) CreateExperiences(w http.ResponseWriter, r *http.Request) {}

func (c *OfferController) CreateSkills(w http.ResponseWriter, r *http.Request) {}
