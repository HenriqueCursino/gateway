package main

import (
	"log"

	"github.com/henriquecursino/gateway/common/env"
	"github.com/henriquecursino/gateway/router"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env")
	}

	env.Load()
	router.Router()
}
