package main

import (
	"regexp"
	"strings"
)

func ErrorPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func Slugify(input string) string {
	// Trim spaces
	slug := strings.TrimSpace(input)

	// Convert all characters to lowercase
	slug = strings.ToLower(slug)

	// Replace spaces, punctuation, and special characters with a hyphen
	reg := regexp.MustCompile("[^a-z0-9]+")
	slug = reg.ReplaceAllString(slug, "-")

	// Remove leading and trailing hyphens
	slug = strings.Trim(slug, "-")

	return slug
}
