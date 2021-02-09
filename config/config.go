package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	Environment string

	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string

	RabbitMQHost     string
	RabbitMQPort     int
	RabbitMQUser     string
	RabbitMQPassword string

	LogLevel string
	HttpPort string
}

func Load() Config {
	config := Config{}

	config.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))

	config.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost"))
	config.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))
	config.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "postgres"))
	config.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "postgres"))
	config.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "mb_corporate_service"))

	config.RabbitMQHost = cast.ToString(getOrReturnDefault("RABBITMQ_HOST", "localhost"))
	config.RabbitMQPort = cast.ToInt(getOrReturnDefault("RABBITMQ_PORT", 5672))
	config.RabbitMQUser = cast.ToString(getOrReturnDefault("RABBITMQ_USER", "admin"))
	config.RabbitMQPassword = cast.ToString(getOrReturnDefault("RABBITMQ_PASSWORD", "admin"))

	config.HttpPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":1234"))
	return config
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)

	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
