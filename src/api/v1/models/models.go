package models

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Connect() error {
	var err error
	host := os.Getenv("SQL_HOST")
	username := os.Getenv("SQL_USER")
	password := os.Getenv("SQL_PASSWORD")
	databaseName := os.Getenv("SQL_DATABASE")
	port := os.Getenv("SQL_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, password, databaseName, port)
	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	} else {
		fmt.Println("Successfully connected to the database")
		return nil
	}
}
