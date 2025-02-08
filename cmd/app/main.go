package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/nordew/tatl-test/internal/app"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading environment variables: %v", err)
	}
}

func main() {
	app.MustRun()
}
