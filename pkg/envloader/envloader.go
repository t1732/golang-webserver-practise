package envloader

import (
	"fmt"
	"os"
	"strconv"
)

func Get(key string, fallback string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}

	return fallback
}

func GetInt(key string, fallback int) int {
	v, ok := os.LookupEnv(key)
	if ok {
		if n, err := strconv.Atoi(v); err == nil {
			return n
		}

		fmt.Printf("Failed to parse int %s: %s \n", key, v)
	}

	return fallback
}
