package model

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
