package utils

import (
	"regexp"
	"strings"
)

var (
	regNonAlnum   = regexp.MustCompile(`[^a-z0-9\s-]`)
	regMultispace = regexp.MustCompile(`[\s_-]+`)
)

func GenerateSlug(input string) string {
	slug := strings.ToLower(input)
	slug = regNonAlnum.ReplaceAllString(slug, "")
	slug = regMultispace.ReplaceAllString(slug, "-")
	slug = strings.Trim(slug, "-")

	return slug
}
