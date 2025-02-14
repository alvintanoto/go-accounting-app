package service

import (
	"errors"
	"time"

	"vint.id/goaccounting/database"
	"vint.id/goaccounting/models"
)

func Register(user *models.User) error {
	user.CreatedAt = time.Now()
	user.HashPassword()

	err := database.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func Login(userLogin *models.Login) (registeredUser *models.User, err error) {
	registeredUser, err = database.GetUserByEmail(userLogin.Email)
	if err != nil {
		return nil, err
	}

	if !registeredUser.ComparePassword(userLogin.Password) {
		return nil, errors.New("user password error")
	}

	return registeredUser, nil
}
