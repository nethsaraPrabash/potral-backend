package repository

import (
	"github.com/nethsaraPrabash/potral-backend/src/config"
	"github.com/nethsaraPrabash/potral-backend/src/model"
)

func Register(user *model.User) error {
	return config.DB.Create(user).Error
}

func Login(user *model.User) error {
	return config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(user).Error
}

