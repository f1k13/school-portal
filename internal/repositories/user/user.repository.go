package userRepo

import (
	"database/sql"
	"errors"

	userDto "github.com/f1k13/school-portal/internal/dto/user"
	"github.com/f1k13/school-portal/internal/logger"
	"github.com/f1k13/school-portal/internal/models/user"

	"github.com/f1k13/school-portal/internal/storage/postgres/school-portal/public/model"
	"github.com/f1k13/school-portal/internal/storage/postgres/school-portal/public/table"
	"github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}
func (r *UserRepository) CreateUser(userDto userDto.UserDto) (*user.User, error) {
	u := model.Users{
		ID:    uuid.New(),
		Email: userDto.Email,
		Role:  userDto.Role,
	}
	existUser, err := r.GetUserByEmail(userDto.Email)
	if err != nil && err.Error() != "user not found" {
		return nil, err
	}
	if existUser != nil && existUser.Verified {
		return nil, errors.New("user already exists")
	}
	stmt := table.Users.INSERT(table.Users.ID,
		table.Users.Email,
		table.Users.Role, table.Users.RefreshToken).MODEL(u).RETURNING(table.Users.AllColumns)
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

func (r *UserRepository) GetUserByEmail(email string) (*user.User, error) {
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
func (r *UserRepository) SetAuthCode(user *user.User, authCode string) error {
	updatedUser := model.Users{
		AuthCode: authCode,
	}

	stmt := table.Users.
		UPDATE(table.Users.AuthCode).
		MODEL(updatedUser).
		WHERE(table.Users.ID.EQ(postgres.UUID(user.ID)))

	_, err := stmt.Exec(r.DB)
	if err != nil {
		return err
	}
	return nil
}
func (r *UserRepository) SetRefreshToken(user *user.User, refreshToken string) error {
	updatedUser := model.Users{
		RefreshToken: refreshToken,
	}

	stmt := table.Users.
		UPDATE(table.Users.RefreshToken).
		MODEL(updatedUser).
		WHERE(table.Users.ID.EQ(postgres.UUID(user.ID)))

	_, err := stmt.Exec(r.DB)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetUserByAuthCode(code string) (*user.User, error) {
	if code == "" {
		return nil, errors.New("code is empty")
	}
	stmt := table.Users.SELECT(table.Users.AllColumns).FROM(table.Users).WHERE(
		table.Users.AuthCode.EQ(postgres.String(code)),
	)
	var dest model.Users
	err := stmt.Query(r.DB, &dest)

	if err != nil {
		if err.Error() == "qrm: no rows in result set" {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &dest, nil
}

func (r *UserRepository) GetUserByID(id string) (*user.User, error) {
	if id == "" {
		return nil, errors.New("id is empty")
	}
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("invalid UUID format")
	}
	var stmt = table.Users.SELECT(table.Users.AllColumns).FROM(table.Users).WHERE(table.Users.ID.EQ(postgres.UUID(uuidID)))
	var dest model.Users
	err = stmt.Query(r.DB, &dest)

	if err != nil {
		if err.Error() == "qrm: no rows in result set" {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &dest, err
}

func (r *UserRepository) SetIsAccess(user *user.User) error {
	updatedUser := model.Users{
		Verified: true,
	}

	stmt := table.Users.
		UPDATE(table.Users.Verified).
		MODEL(updatedUser).
		WHERE(table.Users.ID.EQ(postgres.UUID(user.ID)))

	_, err := stmt.Exec(r.DB)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) CreateProfile(dto *userDto.UserProfileDto) (*user.Profile, error) {
	data := user.Profile{
		ID:          uuid.New(),
		FirstName:   dto.FirstName,
		LastName:    dto.LastName,
		PhoneNumber: dto.PhoneNumber,
		AvatarURL:   dto.AvatarUrl,
		Dob:         dto.Dob,
		UserID:      *dto.UserId,
	}
	stmt := table.Profiles.INSERT(table.Profiles.AllColumns).MODEL(data).RETURNING(table.Profiles.AllColumns)
	var dest []user.Profile

	err := stmt.Query(r.DB, &dest)
	if err != nil {
		logger.Log.Error(err)
		return nil, err
	}
	if len(dest) == 0 {
		return nil, errors.New("error in create profile")
	}

	return &dest[0], nil
}
func (r *UserRepository) GetProfileWithUser(userID uuid.UUID) (*user.UserProfile, error) {
	stmt := table.Profiles.SELECT(table.Profiles.AllColumns, table.Users.AllColumns).FROM(table.Profiles.LEFT_JOIN(table.Users, table.Users.ID.EQ(table.Profiles.UserID))).WHERE(table.Profiles.UserID.EQ(postgres.UUID(userID)))

	var dest user.UserProfile
	err := stmt.Query(r.DB, &dest)
	if err != nil {
		if err.Error() == "qrm: no rows in result set" {
			return nil, errors.New("profile not found")
		}
		return nil, err
	}
	return &dest, nil
}
func (r *UserRepository) GetProfile(userID uuid.UUID) (*user.Profile, error) {
	stmt := table.Profiles.SELECT(table.Profiles.AllColumns).FROM(table.Profiles).WHERE(table.Profiles.UserID.EQ(postgres.UUID(userID)))

	var dest user.Profile
	err := stmt.Query(r.DB, &dest)
	if err != nil {
		if err.Error() == "qrm: no rows in result set" {
			return nil, errors.New("profile not found")
		}
		return nil, err
	}
	return &dest, nil
}
