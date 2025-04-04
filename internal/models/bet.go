package models

import "gorm.io/gorm"

type Bet struct {
	gorm.Model
	ID      uint    `json:"id"`
	UserID  int     `json:"user_id"`
	EventID int     `json:"event_id"`
	Amount  float32 `json:"amount"`
	Outcome string  `json:"outcome"`
}
