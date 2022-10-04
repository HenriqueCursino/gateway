package tools

import (
	"regexp"

	"github.com/google/uuid"
)

func RemoveMask(document string) string {
	re := regexp.MustCompile(`[^a-zA-Z0-9 ]+`)
	return re.ReplaceAllString(document, "")
}

func GenerateHash() string {
	return uuid.New().String()
}
