package config

import (
    "os"
)

type Config struct {
    DatabaseURL string
    RedisURL    string
}

func LoadConfig() (*Config, error) {
    config := &Config{
        DatabaseURL: os.Getenv("DATABASE_URL"),
        RedisURL:    os.Getenv("REDIS_URL"),
    }

    return config, nil
}
