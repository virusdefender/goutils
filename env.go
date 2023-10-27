package goutils

import (
	"golang.org/x/exp/constraints"
	"os"
)

func EnvOrDefaultNumber[T constraints.Integer | constraints.Float](key string, defaultValue T) T {
	s := os.Getenv(key)
	if s == "" {
		return defaultValue
	}

	val, err := ParseNumber[T](s)
	if err != nil {
		return defaultValue
	}
	return val
}

func EnvOrDefaultString(key string, defaultValue string) string {
	s := os.Getenv(key)
	if s == "" {
		return defaultValue
	}
	return s
}
