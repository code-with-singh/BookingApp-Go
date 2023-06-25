package helper

func IsValidName(firstname string, lastname string) bool {
	if len(firstname) < 2 || len(lastname) < 2 {
		return false
	}

	return true
}
