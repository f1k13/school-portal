package experienceDto

import "github.com/google/uuid"

type ExperienceDto struct {
	UserId  *uuid.UUID `json:"userId"`
	Company string     `json:"company"`
	Role    string     `json:"role"`
	Years   int32      `json:"years"`
}
