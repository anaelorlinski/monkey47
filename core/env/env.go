package env

import (
	"os"
)

func EnvOrString(key, defaultValue string) string {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		return v
	}
	return defaultValue
}
