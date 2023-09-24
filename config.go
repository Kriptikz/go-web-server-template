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
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable not defined")
	}

	ProjectConfig = Config{
		Port: os.Getenv("PORT"),
	}
}
