package tools

import (
	"regexp"
)

func RemoveMask(document string) string {
	re := regexp.MustCompile(`[^a-zA-Z0-9 ]+`)
	return re.ReplaceAllString(document, "")
}
