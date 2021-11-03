package common

import "os"

func GetEnv(name string, fallback string) string {
	if env, ok := os.LookupEnv(name); ok {
		return env
	}
	return fallback
}
