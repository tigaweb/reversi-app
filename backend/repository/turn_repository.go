package repository

import (
	"github.com/tigaweb/reversi-app/backend/model"
	"gorm.io/gorm"
)

type ITurnRepository interface {
	RegisterTurn(turn *model.Turn) error
	// game_idとturn_idからTurnを取得する処理
	FindTurnByGameIdAndTurnCount(game_id uint, turn_count int) (model.Turn, error)
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

func (tr *turnRepository) FindTurnByGameIdAndTurnCount(game_id uint, turn_count int) (model.Turn, error) {
	turn := model.Turn{}
	if err := tr.db.Where("game_id = ? AND turn_count = ?", game_id, turn_count).Find(&turn).Error; err != nil {
		return turn, err
	}
	return turn, nil
}
