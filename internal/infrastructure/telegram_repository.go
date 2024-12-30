package infrastructure

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// TelegramRepository реализует интерфейс для работы с Telegram
type TelegramRepository struct {
	bot *tgbotapi.BotAPI
}

// NewTelegramRepository создает новый экземпляр TelegramRepository
func NewTelegramRepository(bot *tgbotapi.BotAPI) *TelegramRepository {
	return &TelegramRepository{
		bot: bot,
	}
}

// SendMessage отправляет сообщение в чат
func (r *TelegramRepository) SendMessage(chatID int64, text string) error {
	msg := tgbotapi.NewMessage(chatID, text)
	_, err := r.bot.Send(msg)
	return err
}
