package models

// User mapping users table
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (User) TableName() string {
	return "users"
}
