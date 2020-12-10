package route

import (
	"blog/controllers/admin/dashboard"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() {

	r := mux.NewRouter()

	// Route Handler
	r.HandleFunc("/api/posts", getPosts).Methods("GET")
	// r.HandleFunc("/api/posts/{id}", getPost).Methods("GET")
	// r.HandleFunc("/api/posts", createPost).Methods("POST")
	// r.HandleFunc("/api/posts/{id}", updatePost).Methods("PUT")
	// r.HandleFunc("/api/posts/{id}", deletePost).Methods("DELETE")

	http.HandleFunc("/admin/dashboard", dashboard.Dashboard)
}
