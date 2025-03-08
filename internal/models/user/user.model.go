package user

import (
	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	FirstName    string    `json:"firstName" gorm:"not null"`
	MiddleName   string    `json:"middleName" gorm:"not null"`
	SurName      string    `json:"surName" gorm:"not null"`
	PhoneNumber  string    `json:"phoneNumber" gorm:"not null"`
	Email        string    `json:"email" gorm:"not null"`
	RefreshToken string    `json:"refreshToken" gorm:"not null"`
	Avatar string `json:"avatar"`
	Role string `json:"role" gorm:"not null"`
}


