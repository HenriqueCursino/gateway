package errors

func IsEmptyError(err error) bool {
	return err == nil
}
