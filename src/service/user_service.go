package service

import (
	"github.com/nethsaraPrabash/potral-backend/src/model"
	"github.com/nethsaraPrabash/potral-backend/src/repository"
)

func RegisterUser(user *model.User) error {

	return repository.Register(user)

}

func LoginUser(user *model.User) error {

	return repository.Login(user)

}



