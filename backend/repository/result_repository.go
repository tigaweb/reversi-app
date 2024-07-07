package repository

import (
	"github.com/tigaweb/reversi-app/backend/model"
	"gorm.io/gorm"
)

type IResultRepository interface {
	RegisterResult(result model.GameResult) error
	FindResultByGameId(game_id uint) (model.GameResult, error)
}

type resultRepository struct {
	db *gorm.DB
}

func NewResultRepository(db *gorm.DB) IResultRepository {
	return &resultRepository{db}
}

func (rr *resultRepository) RegisterResult(result model.GameResult) error {
	if err := rr.db.Create(&result).Error; err != nil {
		return nil
	}
	return nil
}

func (rr *resultRepository) FindResultByGameId(game_id uint) (model.GameResult, error) {
	result := &model.GameResult{}
	if err := rr.db.Where("game_id = ?", game_id).Find(&result).Error; err != nil {
		return *result, err
	}
	return *result, nil
}
