package tools

import (
	"regexp"
	"strconv"
)

func TreatDoc(doc string) int {
	documentUnmasked := removeMask(&doc)
	documentInt, _ := convertStrToInt(documentUnmasked)
	return documentInt
}

func removeMask(document *string) string {
	re := regexp.MustCompile(`[^a-zA-Z0-9 ]+`)
	return re.ReplaceAllString(*document, "")
}

func convertStrToInt(str string) (int, error) {
	num, err := strconv.Atoi(str)
	return num, err
}
