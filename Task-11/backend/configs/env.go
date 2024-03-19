package configs

import (
	"log"

	"github.com/joho/godotenv"
)

// ReadEnv will read .env file
func ReadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("ERROR: While Readning .env file", err.Error())
	}
}
