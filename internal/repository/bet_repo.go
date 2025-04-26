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

func (r *BetRepository) GetByID(id int) (*models.Bet, error) {
	var bet models.Bet
	err := r.db.First(&bet, id).Error
	return &bet, err
}

func (r *BetRepository) Delete(betID int) error {
	return r.db.Delete(&models.Bet{}, betID).Error
}
