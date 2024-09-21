package utils

import (
	"os"
	"strconv"
)

// ParseBoolEnv parses a boolean value from the given environment variable name.
// It returns the parsed boolean value or the default value if parsing fails.
func ParseBoolEnv(envVar string, defaultValue bool) bool {
	value := os.Getenv(envVar)
	if value == "" {
		return defaultValue
	}
	parsedValue, err := strconv.ParseBool(value)
	if err != nil {
		return defaultValue
	}
	return parsedValue
}
