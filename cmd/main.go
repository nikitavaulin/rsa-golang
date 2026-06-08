package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/nikitavaulin/rsa-golang/internal/app"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func main() {
	app := app.NewApp()
	app.Run()
}
