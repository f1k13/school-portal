package experience

import (
	"github.com/f1k13/school-portal/internal/controllers"
	"github.com/f1k13/school-portal/internal/storage/postgres/school-portal/public/model"
	"github.com/google/uuid"
	"time"
)

type ExperienceModel = model.Experiences

type Experience struct {
	Company   string     `json:"company"`
	Years     int32      `json:"year"`
	UserID    uuid.UUID  `json:"userId"`
	Role      string     `json:"role"`
	CreatedAt *time.Time `json:"createdAt"`
	ID        uuid.UUID  `json:"id"`
}

type ExperienceRes struct {
	controllers.Response
	Experience []Experience `json:"experience"`
}
