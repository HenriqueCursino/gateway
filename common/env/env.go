package env

import (
	"log"
	"os"
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
	var valueEnv = os.Getenv(envName)
	if valueEnv == "" {
		log.Fatal(envName, " is empty!")
	}
	return valueEnv
}
