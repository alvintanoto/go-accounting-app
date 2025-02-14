package database

import (
	"database/sql"
	"os"

	"github.com/joho/godotenv"
)

var DB *sql.DB

func init() {
	godotenv.Load(".env")
	var err error
	DB, err = sql.Open("postgres", os.Getenv("POSTGRES_URI"))
	if err != nil {
		panic(err)
	}
	data, err := os.ReadFile("database/init.sql")
	if err != nil {
		panic(err)
	}
	script := string(data)
	if _, err := DB.Exec(script); err != nil {
		panic(err)
	}
}
