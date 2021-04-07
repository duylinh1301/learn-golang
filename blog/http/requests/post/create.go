package post

type PostRequest struct {
	Title       string
	Description string
	Content     string
	Category_id int
}

func NewPostRequest() *PostRequest {
	return &PostRequest{}
}
