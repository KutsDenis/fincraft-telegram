package config

import (
	"fincraft-telegram/internal/server"
	"fmt"
	"os"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

// Config содержит конфигурацию приложения.
type Config struct {
	server.WebhookConfig
	BotToken string `env:"BOT_TOKEN,required"`
}

// LoadConfig загружает конфигурацию из переменных окружения или файла .env.
func LoadConfig() (*Config, error) {
	if _, err := os.Stat(".env"); err == nil {
		if err := godotenv.Load(); err != nil {
			return nil, fmt.Errorf("failed to load .env file: %w", err)
		}
	}

	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("failed to parse environment variables: %w", err)
	}

	return cfg, nil
}
