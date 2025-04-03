package repository

import (
	"betting-site/internal/models"

	"gorm.io/gorm"
)

type BetRepository struct {
	db *gorm.DB
}

func NewBetRepository(db *gorm.DB) *BetRepository {
	return &BetRepository{db: db}
}

func (r *BetRepository) Create(bet *models.Bet) error {
	return r.db.Create(bet).Error
}

func (r *BetRepository) GetBetsByUserID(userID int) ([]models.Bet, error) {
	var bets []models.Bet
	err := r.db.Where("user_id = ?", userID).Find(&bets).Error
	return bets, err
}
