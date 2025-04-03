package services

import (
	"betting-site/internal/models"
	"betting-site/internal/repository" // Убедись, что этот импорт есть
)

type BetService struct {
	repo *repository.BetRepository
}

func NewBetService(repo *repository.BetRepository) *BetService {
	return &BetService{repo: repo}
}

func (s *BetService) Create(userID, eventID int, amount float32, outcome string) (*models.Bet, error) {
	bet := &models.Bet{
		UserID:  userID,
		EventID: eventID,
		Amount:  amount,
		Outcome: outcome,
	}
	err := s.repo.Create(bet)
	return bet, err
}

func (s *BetService) GetBetsByUserID(userID int) ([]models.Bet, error) {
	return s.repo.GetBetsByUserID(userID)
}
