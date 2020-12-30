package postcontroller

import (
	// db "blog/bootstrap"

	"blog/http/response"
	models "blog/models"
	repository "blog/repositories"
	"net/http"
)

// GetAllPosts return all posts
func GetAll(w http.ResponseWriter, r *http.Request) {

	// post := models.Post{
	// 	Id:      1,
	// 	Title:   "Lorem ipsum dolor sit amet.",
	// 	Content: "Lonsectetur adipiscing elit. Donec facilisis dapibus suscipit. Duis dignissim eleifend justo, quis consequat lorem lacinia vel.",
	// }

	// db := db.ConnectDB()

	// env := config.Env

	post := []models.Post{}

	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", env["DB_USERNAME"], env["DB_PASSWORD"], env["DB_HOST"], env["DB_PORT"], env["DB_DATABASE"])

	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// if err != nil {
	// 	log.Println("Connection Failed to Open")
	// } else {
	// 	log.Println("Connection to " + env["DB_CONNECTION"] + " established")
	// }

	repository.All(post)
	// db.Find(&post)

	response.ReturnJSON(w, http.StatusOK, "", post)
}
