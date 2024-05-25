package config

import (
	"os"
)

type Config struct {
	DatabaseURL string
	UseInMemory bool
}

func LoadConfig() Config {
	useInMemory := os.Getenv("USE_IN_MEMORY") == "true"
	databaseURL := os.Getenv("DATABASE_URL")

	return Config{
		DatabaseURL: databaseURL,
		UseInMemory: useInMemory,
	}
}
