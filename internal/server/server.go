package server

import (
	"fincraft-telegram/internal/interfaces"
	"fmt"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type WebhookConfig struct {
	Host     string `env:"WEBHOOK_HOST"`
	Port     string `env:"WEBHOOK_PORT"`
	CertPath string `env:"WEBHOOK_CERT"`
	KeyPath  string `env:"WEBHOOK_CERT_KEY"`
}

// RunTelegramServer запускает сервер для обработки запросов от Telegram.
func RunTelegramServer(config WebhookConfig, handler interfaces.TelegramBotServer, bot *tgbotapi.BotAPI) error {
	webhookURL := fmt.Sprintf("https://%s:%s/%s",
		config.Host,
		config.Port,
		bot.Token,
	)

	wh, err := tgbotapi.NewWebhookWithCert(webhookURL, tgbotapi.FilePath(config.CertPath))
	if err != nil {
		return fmt.Errorf("failed to create webhook: %w", err)
	}

	_, err = bot.Request(wh)
	if err != nil {
		return fmt.Errorf("failed to set webhook: %w", err)
	}

	info, err := bot.GetWebhookInfo()
	if err != nil {
		return fmt.Errorf("failed to get webhook info: %w", err)
	}

	if info.LastErrorDate != 0 {
		return fmt.Errorf("telegram callback failed: %s", info.LastErrorMessage)
	}

	updates := bot.ListenForWebhook("/" + bot.Token)

	addr := fmt.Sprintf(":%s", config.Port)

	//noinspection GoUnhandledErrorResult
	go http.ListenAndServeTLS(addr, config.CertPath, config.KeyPath, nil)

	for update := range updates {
		handler.HandleUpdate(update)
	}

	return nil
}
