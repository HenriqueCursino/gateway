package env

import (
	"log"
	"os"
)

var (
	CurrentMode string
)

func Load() {
	CurrentMode = GetEnviroment("CURRENT_MODE")
}

func GetEnviroment(envName string) string {
	var valueEnv = os.Getenv(envName)
	if valueEnv == "" {
		log.Fatal(envName, " is empty!")
	}
	return valueEnv
}
