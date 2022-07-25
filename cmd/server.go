package main

import (
	"log"

	webserver "github.com/MaximillianoNico/go-clean-architecture/app/infrastructure"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file")
	}

	App := webserver.NewApp("v1.0.0")

	App.RunServer()
}
