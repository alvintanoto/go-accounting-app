package database

import (
	"log"

	"vint.id/goaccounting/models"
)

func CreateUser(user *models.User) error {
	err := DB.QueryRow(
		`INSERT INTO users(email, name, password_hash)
		VALUES ($1, $2, $3) RETURNING ID`,
		user.Email,
		user.Name,
		user.Password,
	).Scan(&user.Id)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func GetUserByEmail(username string) (*models.User, error) {
	var user models.User
	if err := DB.QueryRow(`SELECT id, email, name, password_hash, currency_preference, created_at FROM users WHERE email = $1`, username).Scan(
		&user.Id,
		&user.Email,
		&user.Name,
		&user.Password,
		&user.CurrencyPreference,
		&user.CreatedAt,
	); err != nil {
		log.Println(err)
		return nil, err
	}
	return &user, nil
}
