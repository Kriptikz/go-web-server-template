package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var ProjectConfig Config

type Config struct {
	Port string
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ProjectConfig = Config{
		Port: os.Getenv("PORT"),
	}
}
