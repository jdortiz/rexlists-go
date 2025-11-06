// Package config provides configuration functions for the application.
package config

import "os"

func GetFromEnvOrDefault(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
