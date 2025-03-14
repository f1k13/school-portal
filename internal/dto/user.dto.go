package dto

import "github.com/f1k13/school-portal/internal/storage/postgres/school-portal/public/model"

type UserDto struct {
	FirstName   string  `json:"firstName"`
	MiddleName  *string `json:"middleName"`
	SurName     string  `json:"surName"`
	PhoneNumber *string `json:"phoneNumber"`
	Email       string  `json:"email"`
	Role        string  `json:"role"`
}
type UserToken struct {
	Token string
	User  model.Users
}
