package models

import "time"

// User mapping users table
type User struct {
	ID         int       `json:"id"`       // Unique
	Username   string    `json:"username"` // Unique
	Password   string    `json:"-"`
	Email      string    `json:"email"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

func NewUser() *User {
	return &User{
		Created_at: time.Now(),
		Updated_at: time.Now(),
	}
}

func (User) TableName() string {
	return "users"
}
