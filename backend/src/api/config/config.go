package config

import "os"

var (
	AuthHostname = getEnv("AUTH_HOSTNAME", "localhost")
	AuthPort     = getEnv("AUTH_PORT", "8080")
)

func getEnv(key string, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if ok {
		return value
	}
	return defaultValue

}
