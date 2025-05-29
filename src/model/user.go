package model

import (
	"gorm.io/gorm"
)


type User struct {
	gorm.Model

	UserName string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
	Role string `json:"role"`
	ProfilePic string `json:"propic"`
}