package route

import (
	"blog/controllers/admin/dashboard"
	postcontroller "blog/controllers/blog"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() {

	r := mux.NewRouter().StrictSlash(true)

	// Route Handler
	r.HandleFunc("/api/posts", postcontroller.GetAllPosts).Methods("GET")
	// http.HandleFunc("/api/posts", postcontroller.GetAllPosts)
	// r.HandleFunc("/api/posts/{id}", getPost).Methods("GET")
	// r.HandleFunc("/api/posts", createPost).Methods("POST")
	// r.HandleFunc("/api/posts/{id}", updatePost).Methods("PUT")
	// r.HandleFunc("/api/posts/{id}", deletePost).Methods("DELETE")

	http.HandleFunc("/admin/dashboard", dashboard.Dashboard)

	http.Handle("/", r)
}
