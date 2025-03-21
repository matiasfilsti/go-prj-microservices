package config

import "os"

var (
	AuthHostname = getEnv("AUTH_HOSTNAME", "localhost")
	AuthPort     = getEnv("AUTH_PORT", "8080")

	RabbitMQHostname = getEnv("MQ_HOSTNAME", "localhost")
	RabbitMQUser     = getEnv("MQ_USER", "backUser")
	RabbitMQPassword = getEnv("MQ_PASSWORD", "backPassword")
	RabbitMQPort     = getEnv("MQ_PORT", "5672")

	RedisHostname = getEnv("REDIS_HOSTNAME", "localhost")
	RedisPort     = getEnv("REDIS_PORT", "6389")
	RedisDb       = getEnv("REDIS_DB", "0")
	RedisPassword = getEnv("REDIS_PASSWORD", "testpassword")
	RedisUser     = getEnv("REDIS_USER", "testuser")
)

func getEnv(key string, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if ok {
		return value
	}
	return defaultValue

}
