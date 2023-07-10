package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadENvVariables() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}
}