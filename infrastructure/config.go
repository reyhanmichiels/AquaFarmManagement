package infrastructure

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
		log.Fatal(err)
	}
}
