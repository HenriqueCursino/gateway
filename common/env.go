package common

import (
	"os"
)

var CurrentMode = GetEnviroment("CURRENT_MODE")

func GetEnviroment(envName string) string {
	var CurrentMode = os.Getenv(envName)
	if CurrentMode == "" {
		panic(envName + " is empty!")
	}
	return CurrentMode
}
