package repository

import (
	"github.com/tigaweb/reversi-app/backend/model"
	"gorm.io/gorm"
)

type ITurnRepository interface {
	RegisterTurn(turn *model.Turn) error
	FindTurnByGameIdAndTurnCount(game_id uint, turn_count int) (model.Turn, error)
	FindMaxTurnCountByGameId(game_id uint) (model.Turn, error)
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

func (tr *turnRepository) FindMaxTurnCountByGameId(game_id uint) (model.Turn, error) {
	turn := &model.Turn{}
	if err := tr.db.Where("game_id = ?", game_id).Order("turn_count DESC").First(&turn).Error; err != nil {
		return *turn, err
	}
	return *turn, nil
}
