package errors

func IsEmptyError(err error) bool {
	if err == nil {
		return true
	}
	return false
}
