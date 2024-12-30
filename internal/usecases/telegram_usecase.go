package usecases

// TelegramUseCase реализует бизнес-логику для работы с Telegram
type TelegramUseCase struct {
	repo TelegramRepository
}

// NewTelegramUseCase создает новый экземпляр TelegramUseCase
func NewTelegramUseCase(repo TelegramRepository) *TelegramUseCase {
	return &TelegramUseCase{
		repo: repo,
	}
}

// HandleStartCommand обрабатывает команду /start
func (u *TelegramUseCase) HandleStartCommand(chatID int64) error {
	return u.repo.SendMessage(chatID, "Не? Тағы жұмыс па?")
}
