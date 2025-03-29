package userAdapter

import "github.com/f1k13/school-portal/internal/domain/models/user"

type UserToEntityAdapter struct{}

func NewUserToEntityAdapter() *UserToEntityAdapter {
	return &UserToEntityAdapter{}
}

func (a *UserToEntityAdapter) UserAdapter(u *user.UserModel) *user.User {
	return &user.User{
		ID:           u.ID,
		Email:        u.Email,
		Role:         u.Role,
		Verified:     u.Verified,
		CreatedAt:    *u.CreatedAt,
		RefreshToken: u.RefreshToken,
	}
}

func (a *UserToEntityAdapter) UserProfileAdapter(u *user.UserModel, p *user.ProfileModel) *user.UserProfile {
	return &user.UserProfile{
		User: *a.UserAdapter(u),
		Profile: user.Profile{
			FirstName:   p.FirstName,
			LastName:    p.LastName,
			PhoneNumber: p.PhoneNumber,
			AvatarUrl:   p.AvatarURL,
			Dob:         p.Dob,
			UserId:      &p.UserID,
			ID:          p.ID,
		},
	}
}

func (a *UserToEntityAdapter) ProfileAdapter(p *user.ProfileModel) *user.Profile {
	return &user.Profile{
		FirstName:   p.FirstName,
		LastName:    p.LastName,
		PhoneNumber: p.PhoneNumber,
		AvatarUrl:   p.AvatarURL,
		Dob:         p.Dob,
		UserId:      &p.UserID,
		ID:          p.ID,
	}
}
