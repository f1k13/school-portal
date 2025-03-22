package education

import (
	"time"

	"github.com/f1k13/school-portal/internal/controllers"
	"github.com/f1k13/school-portal/internal/storage/postgres/school-portal/public/model"
	"github.com/google/uuid"
)

type EducationModel = model.Educations

type Education struct {
	ID          string     `json:"id"`
	UserID      uuid.UUID  `json:"userId"`
	Institution string     `json:"institution"`
	Degree      string     `json:"degree"`
	StartYear   int32      `json:"startYear"`
	EndYear     int32      `json:"endYear"`
	City        string     `json:"city"`
	CreatedAt   *time.Time `json:"created_at"`
}
type EducationRes struct {
	Education []Education          `json:"education"`
	Response  controllers.Response `json:"response"`
}
