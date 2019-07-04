package utils

import (
	"log"
	"os"
)

func MustGetEnv(name string, logger *log.Logger) string {
	value, found := os.LookupEnv(name)
	if !found {
		logger.Fatalf("environment variable %s must be set", name)
	}
	return value
}
