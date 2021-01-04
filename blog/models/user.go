package models

import kernel "blog/kernel/model"

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`

	BaseModel kernel.BaseModel
}
