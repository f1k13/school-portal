package educationController

import (
	"encoding/json"
	"net/http"

	"github.com/f1k13/school-portal/internal/controllers"
	educationDataMapper "github.com/f1k13/school-portal/internal/domain/data-mapper/education"
	"github.com/f1k13/school-portal/internal/domain/models/education"
	educationDto "github.com/f1k13/school-portal/internal/dto/education"
	"github.com/f1k13/school-portal/internal/logger"
	educationService "github.com/f1k13/school-portal/internal/services/education"
)

type EducationController struct {
	educationService *educationService.EducationService
	controllers      *controllers.Controller
	mapper           *educationDataMapper.EducationToEntityDataMapper
}

func NewEducationController(educationService *educationService.EducationService, mapper *educationDataMapper.EducationToEntityDataMapper) *EducationController {
	return &EducationController{
		educationService: educationService,
		controllers:      &controllers.Controller{},
		mapper:           mapper,
	}
}

func (c *EducationController) CreateEducation(w http.ResponseWriter, r *http.Request) {
	var req []educationDto.EducationDto
	userID := c.controllers.GetUserIDCtx(r.Context())
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Log.Error("error decoding json", err)
		res := controllers.Response{Message: err.Error()}
		c.controllers.ResponseJson(w, http.StatusBadRequest, res)
		return
	}
	e, err := c.educationService.CreateEducation(&req, userID)
	logger.Log.Info(e)
	if err != nil {
		logger.Log.Error("error in create education handler", err)
		res := controllers.Response{Message: err.Error()}
		c.controllers.ResponseJson(w, http.StatusBadRequest, res)
		return
	}
	eMapper := c.mapper.EducationMapper(&e)
	res := education.EducationRes{Education: *eMapper, Response: controllers.Response{Message: "Успешно"}}
	c.controllers.ResponseJson(w, http.StatusCreated, res)
}
