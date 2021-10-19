package utils

import "strings"

// Capitalize returns a capitalized version of a string
func Capitalize(s string) string {
	return strings.ToUpper(string(s[0])) + s[1:]
}
