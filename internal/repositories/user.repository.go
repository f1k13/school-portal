package repositories

import (
	"database/sql"
	"errors"

	"github.com/f1k13/school-portal/internal/dto"
	"github.com/f1k13/school-portal/internal/logger"
	"github.com/f1k13/school-portal/internal/storage/postgres/school-portal/public/model"
	"github.com/f1k13/school-portal/internal/storage/postgres/school-portal/public/table"
	"github.com/f1k13/school-portal/internal/utils"
	"github.com/go-jet/jet/v2/postgres"
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
		LastName:    userDto.SurName,
	}
	existUser, err := r.GetUserByEmail(userDto.Email)
	if err != nil && err.Error() != "user not found" {
		return nil, err
	}
	if existUser != nil {
		return nil, errors.New("user already exists")
	}
	stmt := table.Users.INSERT(table.Users.ID,
		table.Users.FirstName,
		table.Users.MiddleName,
		table.Users.PhoneNumber,
		table.Users.Email,
		table.Users.Role, table.Users.LastName, table.Users.RefreshToken).MODEL(u).RETURNING(table.Users.AllColumns)
	var dest []model.Users
	err = stmt.Query(r.DB, &dest)

	if err != nil {
		return nil, err
	}
	if len(dest) == 0 {
		return nil, errors.New("error in create user")
	}
	return &dest[0], nil
}

func (r *UserRepository) GetUserByEmail(email string) (*model.Users, error) {
	if email == "" {
		return nil, errors.New("email is empty")
	}
	stmt := table.Users.SELECT(table.Users.AllColumns).FROM(table.Users).WHERE(
		table.Users.Email.EQ(postgres.String(email)),
	)
	var dest model.Users
	err := stmt.Query(r.DB, &dest)

	if err != nil {
		if err.Error() == "qrm: no rows in result set" {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	logger.Log.Info("dest", dest)
	return &dest, nil
}
