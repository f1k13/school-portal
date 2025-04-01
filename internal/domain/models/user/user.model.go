package user

import (
	"time"

	"github.com/f1k13/school-portal/internal/controllers"
	handlers "github.com/f1k13/school-portal/internal/controllers"
	"github.com/f1k13/school-portal/internal/storage/postgres/school-portal/public/model"
	"github.com/google/uuid"
)

type UserModel = model.Users
type ProfileModel = model.Profiles
type UserWithToken struct {
	User  UserModel
	Token string
}

type User struct {
	ID           uuid.UUID `json:"id"`
	Email        string    `json:"email"`
	Role         string    `json:"role"`
	RefreshToken string    `json:"refreshToken"`
	CreatedAt    time.Time `json:"createdAt"`
	Verified     bool      `json:"verified"`
}

type Profile struct {
	FirstName   *string    `json:"firstName"`
	LastName    *string    `json:"lastName"`
	PhoneNumber *string    `json:"phoneNumber"`
	AvatarUrl   *string    `json:"avatarUrl"`
	Dob         *string    `json:"dob"`
	UserId      *uuid.UUID `json:"userId"`
	ID          uuid.UUID  `json:"id"`
}

type UserResponseAuth struct {
	handlers.Response
	User  User   `json:"user"`
	Token string `json:"token"`
}
type UserSelfRes struct {
	controllers.Response
	User User
}
type UserProfile struct {
	User    User    `json:"user"`
	Profile Profile `json:"profile"`
}
type UserProfileRes struct {
	controllers.Response
	UserProfile
}
type UserProfileModel struct {
	User    UserModel
	Profile ProfileModel
}
type UserWithTokenEntity struct {
	User
	Token string `json:"string"`
}
