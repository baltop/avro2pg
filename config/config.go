package config

import (
	"os"

	"github.com/joho/godotenv"
)

func Load() {

	if err := godotenv.Load(".env"); err != nil {
		panic("Error loading .env file")
	}
}

func TestLoad() {

	if err := godotenv.Load("../.env"); err != nil {
		panic("Error loading .env file")
	}
}

// Config func to get env value
func Config(key string) string {
	return os.Getenv(key)
}
