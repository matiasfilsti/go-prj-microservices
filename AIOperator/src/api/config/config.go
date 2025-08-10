package config

import (
	"os"
	"strconv"
	"strings"
)

var (
	AllowedOrigin = getEnvSlice("ALLOWED_ORIGIN", "http://localhost:8085,http://localhost:8080,http://frontend-service:8085")
	RedisHost     = getEnv("REDIS_HOST", "localhost")
	RedisPort     = getEnvInt("REDIS_PORT", 6379)
	RedisUser     = getEnv("REDIS_USER", "admin")
	RedisPassword = getEnv("REDIS_PASSWORD", "AdminPass123")
	RedisDB       = getEnv("REDIS_DB", "airedis_vectorstore")
	LlmUrl        = getEnv("LLM_URL", "http://127.0.0.1:11435")
)

func getEnv(key string, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if ok {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	value, ok := os.LookupEnv(key)
	if ok {
		num, _ := strconv.Atoi(value)
		return num
	}
	return defaultValue
}

func getEnvSlice(key string, defaultValue string) []string {
	value, ok := os.LookupEnv(key)
	if ok {
		return strings.Split(value, ",")
	}
	return strings.Split(defaultValue, ",")
}
