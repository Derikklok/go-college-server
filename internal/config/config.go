package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var AppPort string

var DBHost string
var DBPort string
var DBUser string
var DBPassword string
var DBName string

func Load() {
	err := godotenv.Load()

	if err != nil {
		log.Println(".env file not found, using system environment")
	}

	AppPort = os.Getenv("APP_PORT")
	DBHost = os.Getenv("DB_HOST")
	DBPort = os.Getenv("DB_PORT")
	DBUser = os.Getenv("DB_USER")
	DBPassword = os.Getenv("DB_PASSWORD")
	DBName = os.Getenv("DB_NAME")
}
