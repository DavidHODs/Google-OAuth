package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

var OAuthgolang *oauth2.Config

// LoadFile returns the value of a specified key from .env file
func LoadFile(key string) string {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("error loading .env file")
	}

	return os.Getenv(key)
}
