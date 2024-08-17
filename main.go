package spt

import "os"


func FetchEnv(setVar string, defaultVar string) string {
	envVar, ok := os.LookupEnv(setVar)

	if ok {
		envVar = defaultVar
	}

	return envVar
}
