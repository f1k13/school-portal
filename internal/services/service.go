package services

import "github.com/f1k13/school-portal/internal/repositories"

type Service struct {
	Repo *repositories.Repository
}

func NewService(repo *repositories.Repository) *Service {
	return &Service{Repo: repo}
}
