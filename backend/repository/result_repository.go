package repository

import (
	"github.com/tigaweb/reversi-app/backend/model"
	"gorm.io/gorm"
)

type IResultRepository interface {
	CreateResutl(game_id uint) error
	RegisterResult(result model.GameResult) error
	FindResultByGameId(game_id uint) (model.GameResult, error)
	FindResultByUserId(user_id uint) ([]model.ResponseGameHistory, error)
}

type resultRepository struct {
	db *gorm.DB
}

func NewResultRepository(db *gorm.DB) IResultRepository {
	return &resultRepository{db}
}

func (rr *resultRepository) CreateResutl(game_id uint) error {
	query := `
		INSERT INTO game_results (game_id, winner_by_id, winner_disc, end_at)
		VALUES (?, NULL, NULL, NULL)
	`
	if err := rr.db.Exec(query, game_id).Error; err != nil {
		return err
	}
	return nil
}

func (rr *resultRepository) RegisterResult(result model.GameResult) error {
	if err := rr.db.Where("game_id = ?", result.GameId).Updates(&result).Error; err != nil {
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

func (rr *resultRepository) FindResultByUserId(user_id uint) ([]model.ResponseGameHistory, error) {
	var games []model.Game
	err := rr.db.Where("join_by_id = ? OR created_by_id = ?", user_id, user_id).Find(&games).Error
	if err != nil {
		return nil, err
	}

	var gameIds []uint
	for _, game := range games {
		gameIds = append(gameIds, game.ID)
	}

	var gameResults []model.GameResult
	err = rr.db.Where("game_id IN ?", gameIds).Preload("WinnerBy").Find(&gameResults).Error
	if err != nil {
		return nil, err
	}

	var response []model.ResponseGameHistory
	for _, result := range gameResults {
		game := findGameByID(games, result.GameId)
		startedAt := game.StartedAt
		if result.WinnerById != 0 {
			startedAt = game.CreatedAt
		}

		response = append(response, model.ResponseGameHistory{
			GameId:         result.GameId,
			GameState:      getGameState(result),
			WinnerUserName: result.WinnerBy.UserName,
			WinnerDisc:     model.Disc(result.WinnerDisc),
			StartedAt:      startedAt,
			EndAt:          result.EndAt,
		})
	}

	return response, nil
}

func findGameByID(games []model.Game, gameID uint) model.Game {
	for _, game := range games {
		if game.ID == gameID {
			return game
		}
	}
	return model.Game{}
}

func getGameState(result model.GameResult) int {
	if result.WinnerById != 0 {
		return 1 // ゲーム終了
	}
	return 0 // ゲーム進行中
}
