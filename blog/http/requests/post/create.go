package post

type CreateRequest struct {
	Title       string
	Description string
	Content     string
}

func NewCreateRequest() *CreateRequest {
	return &CreateRequest{}
}
