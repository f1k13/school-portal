package educationDto

import "github.com/google/uuid"

type EducationDto struct {
	UserID      *uuid.UUID `json:"user_id"`
	Institution string     `json:"institution"`
	Degree      string     `json:"degree"`
	EndYear     int32      `json:"endYear"`
	City        string     `json:"city"`
	StartYear   int32      `json:"startYear"`
}
