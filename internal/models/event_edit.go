package models

type EventEdit struct {
	Name     string  `json:"name"`
	Date     string  `json:"date"`
	OddsWin1 float32 `json:"odds_win1"`
	OddsDraw float32 `json:"odds_draw"`
	OddsWin2 float32 `json:"odds_win2"`
}
