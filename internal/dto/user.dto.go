package dto

type UserDto struct {
	Email string `json:"email"`
	Role  string `json:"role"`
}

type UserProfileDto struct {
	FirstName   *string `json:"firstName"`
	LastName    *string `json:"lastName"`
	PhoneNumber *string `json:"phoneNumber"`
	AvatarUrl   *string `json:"avatarUrl"`
	Dob         *string `json:"dob"`
}
