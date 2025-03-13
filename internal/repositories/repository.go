package repositories

import (
	"database/sql"
)

type Repository struct {
	UserRepo *UserRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		UserRepo: NewUserRepository(db),
	}
}
