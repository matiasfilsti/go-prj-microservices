package config

import "os"

var (
	RabbitMQHostname = getEnv("MQ_HOSTNAME", "localhost")
	RabbitMQUser     = getEnv("MQ_USER", "backUser")
	RabbitMQPassword = getEnv("MQ_PASSWORD", "backPassword")
	RabbitMQPort     = getEnv("MQ_PORT", "5672")
)

func getEnv(key string, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if ok {
		return value
	}
	return defaultValue

}
