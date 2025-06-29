package config

import "os"

var (
	BackEndUrl  = getEnv("BACKEND_URL", "localhost")
	BackEndPort = getEnv("BACKEND_PORT", "8081")
)

func getEnv(key string, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if ok {
		return value
	}
	return defaultValue

}
