package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	CurrentMode  string
	SecretKeyJWT string
)

func Load() {
	CurrentMode = GetEnviroment("CURRENT_MODE")
	SecretKeyJWT = GetEnviroment("SECRET_KEY_JWT")
}

func GetEnviroment(envName string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env")
	}

	var valueEnv = os.Getenv(envName)
	if valueEnv == "" {
		log.Fatal(envName, " is empty!")
	}
	return valueEnv
}
