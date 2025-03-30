package offerController

import (
	"encoding/json"
	"net/http"

	"github.com/f1k13/school-portal/internal/controllers"
	offerAdapter "github.com/f1k13/school-portal/internal/domain/adapter/offer"
	educationDataMapper "github.com/f1k13/school-portal/internal/domain/data-mapper/education"
	experienceMapper "github.com/f1k13/school-portal/internal/domain/data-mapper/experience"
	"github.com/f1k13/school-portal/internal/domain/models/offer"
	offerDto "github.com/f1k13/school-portal/internal/dto/offer"
	"github.com/f1k13/school-portal/internal/logger"
	offerService "github.com/f1k13/school-portal/internal/services/offer"
	"github.com/google/uuid"
)

type OfferController struct {
	offerService *offerService.OfferService
	controllers  *controllers.Controller
	adapter      *offerAdapter.OfferToEntityAdapter
	expMapper    *experienceMapper.ExperienceToEntityMapper
	eduMapper    *educationDataMapper.EducationToEntityDataMapper
}

func NewOfferController(offerService *offerService.OfferService, adapter *offerAdapter.OfferToEntityAdapter, expMapper *experienceMapper.ExperienceToEntityMapper, eduMapper *educationDataMapper.EducationToEntityDataMapper) *OfferController {
	return &OfferController{
		offerService: offerService,
		controllers:  &controllers.Controller{},
		adapter:      adapter,
		expMapper:    expMapper,
		eduMapper:    eduMapper,
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

func (c *OfferController) GetOfferById(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	req := query.Get("id")

	if req == "" {
		res := controllers.Response{Message: "id is empty"}
		c.controllers.ResponseJson(w, http.StatusBadRequest, res)
		return
	}
	id, err := uuid.Parse(req)
	if err != nil {
		logger.Log.Error("error in get offer by id handler", err)
		res := controllers.Response{Message: err.Error()}
		c.controllers.ResponseJson(w, http.StatusBadRequest, res)
		return
	}
	o, err := c.offerService.GetOfferById(id)
	if err != nil {
		logger.Log.Error("error in get offer by id handler", err)
		res := controllers.Response{Message: err.Error()}
		c.controllers.ResponseJson(w, http.StatusBadRequest, res)
		return
	}
	exp := c.expMapper.ExperienceMapper(&o.Experience)
	edu := c.eduMapper.EducationMapper(&o.Education)
	offerAdapter := c.adapter.OfferAdapter(
		&offer.OfferModel{
			ID:          o.ID,
			Price:       o.Price,
			DirectionID: o.DirectionID,
			UserID:      o.UserID,
			Title:       o.Title,
			Description: o.Description,
			IsOnline:    o.IsOnline,
		},
	)

	offerWithDetails := offer.OfferWithExpEdSkill{
		Offer:       *offerAdapter,
		Experiences: exp,
		Educations:  *edu,
	}

	res := offer.OfferWithExpEdSkillRes{
		Response: controllers.Response{Message: "Успешно"},
		Offer:    offerWithDetails,
	}
	c.controllers.ResponseJson(w, http.StatusOK, res)
}

func (c *OfferController) SearchOffers(w http.ResponseWriter, r *http.Request) {
	var req offerDto.SearchOfferDto

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Log.Error("error decoding json", err)
		res := controllers.Response{Message: err.Error()}
		c.controllers.ResponseJson(w, http.StatusBadRequest, res)
		return
	}
	o, err := c.offerService.SearchOffers(&req)
	if err != nil {
		logger.Log.Error("error in search offers handler", err)
		res := controllers.Response{Message: err.Error()}
		c.controllers.ResponseJson(w, http.StatusBadRequest, res)
		return
	}
	res := offer.OfferSearchWithExpEdSkillRes{Response: controllers.Response{Message: "Успешно"}, Offer: o}
	c.controllers.ResponseJson(w, http.StatusOK, res)

}
