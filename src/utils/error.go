package utils

// HandleError Handle basically every error with a panic
func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}
