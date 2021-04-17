package post

type CreateRequest struct {
	Title       string
	Description string
	Content     string
	Category_id int
}

func NewCreateRequest() *CreateRequest {
	return &CreateRequest{}
}
