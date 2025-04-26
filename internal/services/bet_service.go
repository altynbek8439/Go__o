package services

import (
	"betting-site/internal/models"
	"betting-site/internal/repository"

	"gorm.io/gorm"
)

type BetService struct {
	repo *repository.BetRepository
	db   *gorm.DB
}

func NewBetService(repo *repository.BetRepository, db *gorm.DB) *BetService {
	return &BetService{repo: repo, db: db}
}

func (s *BetService) Create(userID, eventID int, amount float32, outcome string) (*models.Bet, error) {
	// Начинаем транзакцию
	tx := s.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	// Находим пользователя
	var user models.User
	if err := tx.First(&user, userID).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Проверяем, достаточно ли средств
	if user.Balance < amount {
		tx.Rollback()
		return nil, gorm.ErrInvalidData // Можно заменить на кастомную ошибку
	}

	// Уменьшаем баланс
	user.Balance -= amount
	if err := tx.Save(&user).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Создаём ставку
	bet := &models.Bet{
		UserID:  userID,
		EventID: eventID,
		Amount:  amount,
		Outcome: outcome,
	}
	if err := s.repo.Create(bet); err != nil {
		tx.Rollback()
		return nil, err
	}

	// Подтверждаем транзакцию
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return bet, nil
}

func (s *BetService) GetBetsByUserID(userID int) ([]models.Bet, error) {
	return s.repo.GetBetsByUserID(userID)
}

func (s *BetService) GetByID(id int) (*models.Bet, error) {
	return s.repo.GetByID(id)
}

func (s *BetService) Delete(betID int) error {
	return s.repo.Delete(betID)
}
