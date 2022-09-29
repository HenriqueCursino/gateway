package common

const (
	DEVELOPMENT = "development"
)

var (
	CurrentMode = GetEnviroment("CURRENT_MODE")
)
