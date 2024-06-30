package model

import "time"

type Turn struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	GameId      uint      `json:"game_id"`
	Game        Game      `json:"game"  gorm:"foreignkey:game_id"`
	TurnCount   int       `json:"turn_count"`
	NextDisc    int       `json:"next_disc"`
	EndAt       time.Time `json:"end_at"`
	CreatedByID uint      `json:"created_by_id"`
	CreatedBy   User      `json:"created_by" gorm:"foreignkey:created_by_id"`
}

type RegisterTurnRequest struct {
	GameID    uint `json:"game_id"`
	TurnCount int  `json:"turn_count"`
	Move      Move `json:"move"`
}
