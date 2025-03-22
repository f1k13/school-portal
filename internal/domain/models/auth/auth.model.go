package auth

type AuthCodeReq struct {
	Code string `json:"code"`
}

type SignInReq struct {
	Email string `json:"email"`
}
