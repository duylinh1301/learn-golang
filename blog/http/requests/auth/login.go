package auth

type LoginRequest struct {
	Email    string
	Password string
}

func NewLoginRequest() *LoginRequest {
	return &LoginRequest{}
}
