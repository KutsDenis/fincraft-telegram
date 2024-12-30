package usecases

// TelegramRepository определяет контракт репозитория для работы с Telegram
type TelegramRepository interface {
	SendMessage(chatID int64, text string) error
}
