package main

import (
	"github.com/henriquecursino/gateway/router"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	router.Router()
}
