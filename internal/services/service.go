package services

import (
	authService "github.com/f1k13/school-portal/internal/services/auth"
	userService "github.com/f1k13/school-portal/internal/services/user"
)

type Services struct {
	AuthService *authService.AuthService
	UserService *userService.UserService
}

func NewServices(authService *authService.AuthService, userService *userService.UserService) *Services {
	return &Services{
		AuthService: authService,
		UserService: userService,
	}
}
