package repository

import (
	"github.com/tigaweb/reversi-app/backend/model"
	"gorm.io/gorm"
)

type ITurnRepository interface {
	RegisterTurn(turn *model.Turn) error
	// game_idとturn_idからTurnを取得する処理
	// GetTurnByGameIdAndTurnId(game_id uint, turn_id uint) (turn model.Turn, error error)
	// 最新のターン情報を取得する処理 Turnモデルにレスポンスの型を定義するか?
}

type turnRepository struct {
	db *gorm.DB
}

func NewTurnRepository(db *gorm.DB) ITurnRepository {
	return &turnRepository{db}
}

func (tr *turnRepository) RegisterTurn(turn *model.Turn) error {
	if err := tr.db.Create(&turn).Error; err != nil {
		return err
	}
	return nil
}
