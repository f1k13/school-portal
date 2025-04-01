package experienceController

import (
	"encoding/json"
	"net/http"

	experienceMapper "github.com/f1k13/school-portal/internal/domain/data-mapper/experience"
	"github.com/f1k13/school-portal/internal/domain/models/experience"
	"github.com/google/uuid"

	"github.com/f1k13/school-portal/internal/controllers"
	experienceDto "github.com/f1k13/school-portal/internal/dto/experience"
	"github.com/f1k13/school-portal/internal/logger"
	experienceService "github.com/f1k13/school-portal/internal/services/experience"
)

type ExperienceController struct {
	experienceService *experienceService.ExperienceService
	controllers       *controllers.Controller
	mapper            *experienceMapper.ExperienceToEntityMapper
}

func NewExperienceController(experienceService *experienceService.ExperienceService, mapper *experienceMapper.ExperienceToEntityMapper) *ExperienceController {
	return &ExperienceController{experienceService: experienceService, controllers: &controllers.Controller{}, mapper: mapper}
}

func (c *ExperienceController) CreateExperience(w http.ResponseWriter, r *http.Request) {
	var req []experienceDto.ExperienceDto
	userID := c.controllers.GetUserIDCtx(r.Context())
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		res := controllers.Response{Message: err.Error()}
		logger.Log.Error("error decoding json", err)
		c.controllers.ResponseJson(w, http.StatusBadRequest, res)
		return
	}
	e, err := c.experienceService.CreateExperience(&req, userID)
	if err != nil {
		res := controllers.Response{Message: err.Error()}
		c.controllers.ResponseJson(w, http.StatusBadRequest, res)
		return
	}
	eMapper := c.mapper.ExperienceMapper(e)
	res := experience.ExperienceRes{Experience: eMapper, Response: controllers.Response{Message: "Успешно"}}
	c.controllers.ResponseJson(w, http.StatusCreated, res)
}

func (c *ExperienceController) GetExperiences(w http.ResponseWriter, r *http.Request) {
	userID := c.controllers.GetUserIDCtx(r.Context())

	userIDParse, err := uuid.Parse(userID)

	if err != nil {
		res := controllers.Response{Message: err.Error()}
		logger.Log.Error("error in parse id", err)
		c.controllers.ResponseJson(w, http.StatusUnauthorized, res)
		return
	}

	e, err := c.experienceService.GetMyExperience(userIDParse)

	if err != nil {
		res := controllers.Response{Message: err.Error()}
		c.controllers.ResponseJson(w, http.StatusInternalServerError, res)
		return
	}

	eMapper := c.mapper.ExperienceMapper(e)

	res := experience.ExperienceRes{Experience: eMapper, Response: controllers.Response{Message: "Успешно"}}
	c.controllers.ResponseJson(w, http.StatusOK, res)

}
