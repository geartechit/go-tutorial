package config

import (
	"os"
)

type PostgresConfig struct {
	URL            string
	MigrationPaths string
}

type ServerConfig struct {
	Port string
	Env  string
}

type Config struct {
	Postgres PostgresConfig
	Server   ServerConfig
}

func LoadConfig() *Config {
	return &Config{
		Postgres: PostgresConfig{
			URL:            getEnv("POSTGRES_URL", "postgres://postgres:password@localhost:5432/postgres?sslmode=disable"),
			MigrationPaths: "file://internal/database/postgres/migrations",
		},
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "8080"),
			Env:  getEnv("SERVER_ENV", "development"),
		},
	}
}

func getEnv(key, defaultValue string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultValue
}

func (s *ServerConfig) IsProduction() bool {
	return s.Env == "production"
}
