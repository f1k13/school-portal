package educationController

import (
	"encoding/json"
	"net/http"

	"github.com/f1k13/school-portal/internal/controllers"
	educationDto "github.com/f1k13/school-portal/internal/dto/education"
	"github.com/f1k13/school-portal/internal/logger"
	educationService "github.com/f1k13/school-portal/internal/services/education"
)

type EducationController struct {
	educationService *educationService.EducationService
	controllers      *controllers.Controller
}

func NewEducationController(educationService *educationService.EducationService) *EducationController {
	return &EducationController{
		educationService: educationService,
		controllers:      &controllers.Controller{},
	}
}

func (c *EducationController) CreateEducation(w http.ResponseWriter, r *http.Request) {
	var req educationDto.EducationDto

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Log.Error("error decoding json", err)
		res := controllers.Response{Message: err.Error()}
		c.controllers.ResponseJson(w, http.StatusBadRequest, res)
		return
	}

}
