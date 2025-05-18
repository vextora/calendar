package ratelimiter

import "strings"

var excludedPaths = []string{
	"/swagger/",
	"/metrics",
	"/health",
	"/ping",
}

func IsExcludePath(path string) bool {
	for _, p := range excludedPaths {
		if strings.HasPrefix(path, p) {
			return true
		}
	}
	return false
}
