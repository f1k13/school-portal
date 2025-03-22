package user

import (
	"github.com/f1k13/school-portal/internal/controllers"
	handlers "github.com/f1k13/school-portal/internal/controllers"
	"github.com/f1k13/school-portal/internal/storage/postgres/school-portal/public/model"
)

type UserModel = model.Users
type ProfileModel = model.Profiles
type UserWithToken struct {
	User  UserModel
	Token string
}
type UserResponseAuth struct {
	handlers.Response
	User  UserModel `json:"user"`
	Token string    `json:"token"`
}
type UserSelfRes struct {
	controllers.Response
	User *model.Users
}
type UserProfile struct {
	User    UserModel
	Profile ProfileModel
}
type UserProfileRes struct {
	controllers.Response
	UserProfile
}
