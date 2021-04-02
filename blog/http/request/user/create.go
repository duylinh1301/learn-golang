package user

type UserRequest struct {
	Username string
	Email    string
	Password string
}

func NewUserRequest() *UserRequest {
	return &UserRequest{}
}
