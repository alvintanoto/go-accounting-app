package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id                 string
	Email              string  `form:"email" binding:"required"`
	Name               string  `form:"name" binding:"required"`
	Password           string  `form:"password" binding:"required"`
	CurrencyPreference *string `form:"currency_preference"`
	CreatedAt          time.Time
}

func (u *User) HashPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return nil
}

func (u *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

type Login struct {
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}
