package utils

import (
	"os"
	"strconv"
)

// ReadEnvBool read and parse boolean environment variables.
// Will return `false` if the variable is not a `bool`
func ReadEnvBool(key string) bool {
	value := os.Getenv(key)
	parsed, err := strconv.ParseBool(value)

	if err != nil {
		return false
	}

	return parsed
}

// ReadEnvString read string environment variables.
// This function is a wrapper function to 'os.Getenv()'
func ReadEnvString(key string) string {
	return os.Getenv(key)
}

// ReadEnvInt read and integer environment variables.
// Will return `0` if the variable is not a `bool`
func ReadEnvInt(key string) int {
	value := os.Getenv(key)
	parsed, err := strconv.Atoi(value)

	if err != nil {
		return 0
	}

	return parsed
}
