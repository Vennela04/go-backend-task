package config

import "os"

type Config struct {
	DBUrl string
	Port  string
}

// Load loads configuration from environment variables or defaults
func Load() *Config {
	return &Config{
		DBUrl: getEnv("DATABASE_URL", "postgres://postgres:9514@localhost:5432/backenddb?sslmode=disable"),
		Port:  getEnv("PORT", "8080"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
