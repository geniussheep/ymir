package pkg

import (
	"os"
	"strconv"
	"strings"
	"time"
)

func GetEnvDuration(key string, fallback time.Duration) time.Duration {
	if value, ok := os.LookupEnv(key); ok {
		if n, err := time.ParseDuration(value); err == nil {
			return n
		}
	}
	return fallback
}

func GetEnvArray(key string, sep string, fallback []string) []string {
	if value, ok := os.LookupEnv(key); ok {
		return strings.Split(value, sep)
	}
	return fallback
}

func GetEnvBool(key string, fallback bool) bool {
	if value, ok := os.LookupEnv(key); ok {
		if b, err := strconv.ParseBool(value); err == nil {
			return b
		}
	}
	return fallback
}

func GetEnvInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		if n, err := strconv.ParseInt(value, 10, 64); err == nil {
			return n
		}
	}
	return fallback
}

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
