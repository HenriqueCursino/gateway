package common

import (
	"os"
)

func GetEnviroment(envName string) string {
	var valueEnv = os.Getenv(envName)
	return valueEnv
}
