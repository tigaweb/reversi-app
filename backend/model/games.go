package model

import "time"

type Game struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	JoinById    uint      `json:"join_by_id"`
	JoinBy      User      `json:"join_by" gorm:"foreignkey:JoinById"`
	StartedAt   time.Time `json:"started_at"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedByID uint      `json:"created_by_id"`
	CreatedBy   User      `json:"created_by" gorm:"foreignkey:CreatedByID"`
	UpdatedAt   time.Time `json:"updated_at"`
}
