package config

import (
	"os"
	"strconv"
)

var (
	AuthHostname = getEnv("AUTH_HOSTNAME", "localhost")
	AuthPort     = getEnv("AUTH_PORT", "8080")

	RabbitMQHostname = getEnv("MQ_HOSTNAME", "localhost")
	RabbitMQUser     = getEnv("MQ_USER", "backUser")
	RabbitMQPassword = getEnv("MQ_PASSWORD", "backPassword")
	RabbitMQPort     = getEnv("MQ_PORT", "5672")

	RedisHostname = getEnv("REDIS_HOSTNAME", "localhost")
	RedisPort     = getEnv("REDIS_PORT", "6379")
	RedisPassword = getEnv("REDIS_PASSWORD", "testpassword")
	RedisUser     = getEnv("REDIS_USER", "testuser")
	RedisTTL      = getEnvInt("REDIS_TTL", 1800)
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
