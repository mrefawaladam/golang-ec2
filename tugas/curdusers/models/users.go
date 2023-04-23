package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email"  form:"email"`
	Password string `json:"password" form:"password"`
	Books    []Book `json:"books" gorm:"foreignKey:AuthorID"`
	Blogs    []Blog `json:"blogs" gorm:"foreignKey:AuthorID"`
}

type UserResponse struct {
	ID    uint   `json:"id" form:"name"`
	Name  string `json:"name" form:"name"`
	Email string `json:"email"  form:"email"`
	Token string `json:"token" form:"token"`
}
