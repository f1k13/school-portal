package educationRepo

import "database/sql"

type EducationRepository struct {
	db *sql.DB
}

func NewEducationRepository(db *sql.DB) *EducationRepository {
	return &EducationRepository{db: db}
}
