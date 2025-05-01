package config

import "os"

var (
	DbUsername   = getEnv("DB_USERNAME", "authapi")
	DbPassword   = getEnv("DB_PASSWORD", "authapipass")
	DbDatabase   = getEnv("DB_DATABASE", "auth-api")
	DbHostname   = getEnv("DB_HOSTNAME", "localhost")
	DbPort       = getEnv("DB_PORT", "5432")
	SecretJwtKey = getEnv("JWT_SECRET", "secretkey")
)

func getEnv(key string, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if ok {
		return value
	}
	return defaultValue

}
