package config

import (
    "os"
    "strconv"
)

type Config struct {
    DBURL      string
    RedisURL   string
    NatsURL    string
    GRPCPort   string
    HTTPPort   string
    LogLevel   string
}

func LoadConfig() (*Config, error) {
    // Получаем значения из переменных окружения или используем значения по умолчанию
    natsURL := os.Getenv("NATS_URL")
    if natsURL == "" {
        natsURL = "nats://localhost:4222" // Значение по умолчанию для NATS
    }

    return &Config{
        DBURL:    getEnv("DB_URL", "postgres://postgres:password@localhost:5432/orderdb?sslmode=disable"),
        RedisURL: getEnv("REDIS_URL", "localhost:6379"),
        NatsURL:  natsURL,
        GRPCPort: getEnv("GRPC_PORT", "50051"),
        HTTPPort: getEnv("HTTP_PORT", "9091"),
        LogLevel: getEnv("LOG_LEVEL", "info"),
    }, nil
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
    if value := os.Getenv(key); value != "" {
        if intVal, err := strconv.Atoi(value); err == nil {
            return intVal
        }
    }
    return defaultValue
}
