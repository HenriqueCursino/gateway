package common

import (
	"os"
)

var CurrentMode = GetEnviroment("CURRENT_MODE")

func GetEnviroment(envName string) string {
	var valueEnv = os.Getenv(envName)
	if valueEnv == "" {
		panic(envName + " is empty!")
	}
	return valueEnv
}
