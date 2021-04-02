package post

type PostRequest struct {
	Title       string
	Description string
	Content     string
}

func NewPostRequest() *PostRequest {
	return &PostRequest{}
}
