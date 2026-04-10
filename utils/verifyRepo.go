package utils

import "regexp"

var regRepo = regexp.MustCompile(`^[^/]+/[^/]+$`)

func IsValidRepo(repo string) bool {
	return regRepo.MatchString(repo)
}
