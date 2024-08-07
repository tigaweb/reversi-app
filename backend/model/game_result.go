package model

import "time"

type GameResult struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	GameId     uint      `json:"game_id"`
	Game       Game      `json:"game"  gorm:"foreignkey:game_id"`
	WinnerById uint      `json:"winner_by_id"`
	WinnerBy   User      `json:"winner" gorm:"foreignkey:winner_by_id"`
	WinnerDisc int       `json:"winner_disc"`
	EndAt      time.Time `json:"end_at"`
}

type ResponseGameHistory struct {
	GameId         uint      `json:"game_id"`
	GameState      int       `json:"game_state"`
	WinnerUserName string    `json:"winner_user_name"`
	WinnerDisc     Disc      `json:"winner_disc"`
	StartedAt      time.Time `json:"started_at"`
	EndAt          time.Time `json:"end_at"`
}
