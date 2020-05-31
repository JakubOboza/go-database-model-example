package utils

import "os"

//GetEnv gets Env with possible default fallback in case its empty
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
