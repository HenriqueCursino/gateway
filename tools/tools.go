package tools

import (
	"fmt"
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

func GetStringFromBody(body interface{}) string {
	return fmt.Sprintf("%v", body)
}
