package models

type Post struct {
	id      int    `json:"id"`
	title   string `json:"title"`
	content string `json:"content"`
}

// Get all posts
func (p *Post) All() Post {
	p = &Post{
		id:      1,
		title:   "Lorem ipsum dolor sit amet.",
		content: "Lonsectetur adipiscing elit. Donec facilisis dapibus suscipit. Duis dignissim eleifend justo, quis consequat lorem lacinia vel. Quisque scelerisque massa a pellentesque laoreet. In ac eleifend orci. Praesent ut nulla ac felis sagittis viverra ut in velit. Vivamus eu mauris fringilla, tempor mi et, mollis tellus. Vivamus in lorem at sem finibus laoreet in id orci. Morbi et consectetur nisl, quis ultrices ante. In malesuada sem a ex finibus, vel ullamcorper justo laoreet.",
	}
	return *p
}
