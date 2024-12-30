package interfaces

import (
	"fincraft-telegram/internal/usecases"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// TelegramBotServer обрабатывает входящие обновления на сервере
type TelegramBotServer interface {
	HandleUpdate(update tgbotapi.Update)
}

// TelegramHandler обрабатывает запросы от Telegram
type TelegramHandler struct {
	useCase *usecases.TelegramUseCase
}

// NewTelegramHandler создает новый экземпляр TelegramHandler
func NewTelegramHandler(useCase *usecases.TelegramUseCase) *TelegramHandler {
	return &TelegramHandler{
		useCase: useCase,
	}
}

// HandleUpdate обрабатывает входящие обновления от Telegram
func (h *TelegramHandler) HandleUpdate(update tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	if update.Message.Command() == "start" {
		h.useCase.HandleStartCommand(update.Message.Chat.ID)
	}
}
