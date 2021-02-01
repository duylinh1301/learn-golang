package models

// User mapping users table
type User struct {
	ID       int    `json:"id"`       // Unique
	Username string `json:"username"` // Unique
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (User) TableName() string {
	return "users"
}
