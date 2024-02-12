package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoURI() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("failed to load .env file")
	}
	return os.Getenv("MONGOURI")
}
func EnvDatabase() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load .env file")
	}
	return os.Getenv("DATABASE")
}
func EnvCollection() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load .env file")
	}
	return os.Getenv("COLLECTION")
}
