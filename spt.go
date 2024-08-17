package spt

import "os"

// FetchEnv returns the passed default value to a environment variable.
func FetchEnv(envVar string, defaultValue string) string {
	envValue, ok := os.LookupEnv(envVar)

	if !ok {
		envValue = defaultValue
	}

	return envValue
}
