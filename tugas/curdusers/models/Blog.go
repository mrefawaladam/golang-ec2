package models

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	Title    string `json:"title"`
	Content  string `json:"content"`
	Author   User   `json:"author" gorm:"foreignKey:AuthorID"`
	AuthorID uint   `json:"author_id"`
}
