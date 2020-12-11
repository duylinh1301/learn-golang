package postcontroller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Post struct {
	id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func GetAllPosts(w http.ResponseWriter, r *http.Request) {

	post := Post{
		id:      1,
		Title:   "Lorem ipsum dolor sit amet.",
		Content: "Lonsectetur adipiscing elit. Donec facilisis dapibus suscipit. Duis dignissim eleifend justo, quis consequat lorem lacinia vel.",
	}

	var jsonData []byte
	jsonData, err := json.Marshal(post)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(jsonData))

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
