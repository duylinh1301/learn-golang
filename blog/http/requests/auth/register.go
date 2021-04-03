package auth

type RegisterRequest struct {
	Username string
	Email    string
	Password string
}

func NewRegisterRequest() *RegisterRequest {
	return &RegisterRequest{}
}
