package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)

func CreateConnection() (*gorm.DB, error) {

	// Get the database credentials
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	DBName := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")

	return gorm.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s user=%s dbname=%s sslmode=disable password=%s",
			host, user, DBName, password,
		),
	)
}
