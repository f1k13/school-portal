package dto

type UserDto struct {
	FirstName   string  `json:"firstName"`
	MiddleName  *string `json:"middleName"`
	SurName     string  `json:"surName"`
	PhoneNumber *string `json:"phoneNumber"`
	Email       string  `json:"email"`
	Role        string  `json:"role"`
}
