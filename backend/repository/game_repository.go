package repository

import (
	"github.com/tigaweb/reversi-app/backend/model"
	"gorm.io/gorm"
)

type IGameRepository interface {
	CreateGame(game *model.Game) error
	GetGameResultByUser(user_id uint) ([]*model.GameResult, error)
	FindGameRecordByGameId(game_id uint) (model.Game, error)
}

type gameRepository struct {
	db *gorm.DB
}

func NewGameRepository(db *gorm.DB) IGameRepository {
	return &gameRepository{db}
}

func (gr *gameRepository) CreateGame(game *model.Game) error {
	if err := gr.db.Create(&game).Error; err != nil {
		return err
	}
	return nil
}

func (gr *gameRepository) GetGameResultByUser(user_id uint) ([]*model.GameResult, error) {
	var gameResults []*model.GameResult
	gameIDs := gr.db.Model(&model.Game{}).
		Select("id").
		Where("join_by_id = ? OR created_by_id = ?", user_id, user_id)
	if err := gr.db.Where("game_id IN (?)", gameIDs).First(&gameResults).Error; err != nil {
		return nil, err
	}
	return gameResults, nil
}

func (gr *gameRepository) FindGameRecordByGameId(game_id uint) (model.Game, error) {
	var game model.Game
	if err := gr.db.Where("id=?", game_id).Preload("CreatedBy").Preload("JoinBy").First(&game).Error; err != nil {
		return game, err
	}
	return game, nil
}
