package usecase

import (
	"github.com/tigaweb/reversi-app/backend/model"
	"github.com/tigaweb/reversi-app/backend/repository"
)

type IGameUsecase interface {
	NewGame(game *model.Game) (*model.CreateGameResponse, error)
	GameHistory(user_id uint) ([]*model.GameResult, error)
}

type GameUsecase struct {
	gr repository.IGameRepository
}

func NewGameUsecase(gr repository.IGameRepository) IGameUsecase {
	return &GameUsecase{gr}
}

func (gu *GameUsecase) NewGame(game *model.Game) (*model.CreateGameResponse, error) {
	game_response := model.CreateGameResponse{}
	if err := gu.gr.CreateGame(game); err != nil {
		return nil, err
	}
	game_response.ID = game.ID
	return &game_response, nil
}

func (gu *GameUsecase) GameHistory(user_id uint) ([]*model.GameResult, error) {
	var game_response []*model.GameResult
	var err error
	if game_response, err = gu.gr.GetGameResultByUser(user_id); err != nil {
		return nil, err
	}
	return game_response, nil
}
