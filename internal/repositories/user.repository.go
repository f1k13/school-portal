package repositories

import (
	"database/sql"
	"errors"

	"github.com/f1k13/school-portal/internal/handlers/dto"
	"github.com/f1k13/school-portal/internal/storage/postgres/school-portal/public/model"
	"github.com/f1k13/school-portal/internal/storage/postgres/school-portal/public/table"
	"github.com/f1k13/school-portal/internal/utils"
	"github.com/google/uuid"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}
func (r *UserRepository) CreateUser(userDto dto.UserDto) (*model.Users, error) {
	u := model.Users{
		ID:          uuid.New(),
		FirstName:   userDto.FirstName,
		MiddleName:  userDto.MiddleName,
		PhoneNumber: utils.PtrToStr(userDto.PhoneNumber),
		Email:       userDto.Email,
		Role:        userDto.Role,
	}
	stmt := table.Users.INSERT(table.Users.AllColumns).MODEL(u).RETURNING(table.Users.AllColumns)
	var dest []model.Users
	err := stmt.Query(r.DB, &dest)

	if err != nil {
		return nil, err
	}
	if len(dest) == 0 {
		return nil, errors.New("Ошибка в репо")
	}
	return &dest[0], nil
}

func (r *UserRepository) GetUserByEmail(email string) (*model.Users, error) {
	u := model.Users{}
	if email == "" {
		return &model.Users{}, nil
	}
	// err := r.DB.Where("email = ?", email).First(&u).Error
	// if err != nil {
	// 	logger.Log.Error("Error getting user by email", err)
	// 	return user.User{}, nil
	// }
	return &u, nil
}
