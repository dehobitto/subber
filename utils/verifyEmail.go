package utils

import "regexp"

var regEmail = regexp.MustCompile(`^[^/]+@[^/]+.[^/]+$`)

func IsValidEmail(email string) bool {
	return regEmail.MatchString(email)
}
