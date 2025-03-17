package user

import (
	handlers "github.com/f1k13/school-portal/internal/controllers"
	"github.com/f1k13/school-portal/internal/storage/postgres/school-portal/public/model"
)

type User = model.Users

type UserWithToken struct {
	User  User
	Token string
}
type UserResponseAuth struct {
	handlers.Response
	User  User   `json:"user"`
	Token string `json:"token"`
}
