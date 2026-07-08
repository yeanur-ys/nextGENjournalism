package config

import (
	"errors"
	"os"
)

type Config struct {
	PostgresDSN   string
	Neo4jURI      string
	Neo4jUser     string
	Neo4jPassword string
	JWTSecret     string
	HTTPPort      string
}

func Load() (Config, error) {
	cfg := Config{
		PostgresDSN:   os.Getenv("POSTGRES_DSN"),
		Neo4jURI:      os.Getenv("NEO4J_URI"),
		Neo4jUser:     os.Getenv("NEO4J_USER"),
		Neo4jPassword: os.Getenv("NEO4J_PASSWORD"),
		JWTSecret:     os.Getenv("JWT_SECRET"),
		HTTPPort:      envOrDefault("HTTP_PORT", "8080"),
	}

	if cfg.PostgresDSN == "" || cfg.Neo4jURI == "" || cfg.Neo4jUser == "" || cfg.Neo4jPassword == "" || cfg.JWTSecret == "" {
		return Config{}, errors.New("missing required environment configuration")
	}

	return cfg, nil
}

func envOrDefault(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
