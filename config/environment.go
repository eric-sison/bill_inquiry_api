package config

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}

	log.Print("Loaded environment!")
}
