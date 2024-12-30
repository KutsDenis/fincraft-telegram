package main

import (
	"fincraft-telegram/internal/infrastructure"
	"fincraft-telegram/internal/interfaces"
	"fincraft-telegram/internal/server"
	"fincraft-telegram/internal/usecases"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"os"

	"github.com/KutsDenis/logzap"
	"go.uber.org/zap"

	"fincraft-telegram/internal/config"
)

func main() {
	// Инициализация логгера
	appEnv := os.Getenv("APP_ENV")
	logzap.Init(appEnv)
	defer logzap.Sync()
	log := logzap.Logger

	// Загрузка конфигурации
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load configuration", zap.Error(err))
		os.Exit(1)
	}
	log.Info("Configuration loaded successfully")

	bot, err := tgbotapi.NewBotAPI(cfg.BotToken)
	if err != nil {
		log.Fatal("Failed to create bot", zap.Error(err))
		os.Exit(1)
	}
	log.Info("Bot created successfully", zap.String("username", bot.Self.UserName))

	telegramRepo := infrastructure.NewTelegramRepository(bot)
	telegramUseCase := usecases.NewTelegramUseCase(telegramRepo)
	telegramHandler := interfaces.NewTelegramHandler(telegramUseCase)

	if err := server.RunTelegramServer(cfg.WebhookConfig, telegramHandler, bot); err != nil {
		log.Fatal("Failed to start server", zap.Error(err))
		os.Exit(1)
	}
}
