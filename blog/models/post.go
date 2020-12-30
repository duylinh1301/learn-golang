package models

import "gorm.io/gorm"

type Post struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	db      *gorm.DB
}
