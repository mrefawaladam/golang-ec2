package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title       string `json:"title"  form:"title"`
	Description string `json:"description"  form:"description"`
	Author      User   `json:"author" gorm:"foreignKey:AuthorID"`
	AuthorID    uint   `json:"author_id" form:"author_id"`
}
