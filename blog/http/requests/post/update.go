package post

type UpdateRequest struct {
	Title       string
	Description string
	Content     string
}

func NewUpdateRequest() *UpdateRequest {
	return &UpdateRequest{}
}
