package educationDto

type EducationDto struct {
	UserID      string `json:"user_id"`
	Institution string `json:"institution"`
	Degree      string `json:"degree"`
	Year        int    `json:"year"`
	City        string `json:"city"`
}
