package usecase

import (
	"github.com/tigaweb/reversi-app/backend/model"
	"github.com/tigaweb/reversi-app/backend/repository"
)

type IGameUsecase interface {
	NewGame(game *model.Game) (*model.CreateGameResponse, error)
	GameHistory(user_id uint) ([]*model.GameResult, error)
	FindGameRecordByGameId(game_id uint) (model.Game, error)
}

type gameUsecase struct {
	gr repository.IGameRepository
	rr repository.IResultRepository
}

func NewGameUsecase(gr repository.IGameRepository, rr repository.IResultRepository) IGameUsecase {
	return &gameUsecase{gr, rr}
}

func (gu *gameUsecase) NewGame(game *model.Game) (*model.CreateGameResponse, error) {
	game_response := model.CreateGameResponse{}
	if err := gu.gr.CreateGame(game); err != nil {
		return nil, err
	}
	if err := gu.rr.CreateResutl(game.ID); err != nil {
		return nil, err
	}
	game_response.ID = game.ID
	return &game_response, nil
}

func (gu *gameUsecase) GameHistory(user_id uint) ([]*model.GameResult, error) {
	var game_response []*model.GameResult
	var err error
	if game_response, err = gu.gr.GetGameResultByUser(user_id); err != nil {
		return nil, err
	}
	return game_response, nil
}

func (gu *gameUsecase) FindGameRecordByGameId(game_id uint) (model.Game, error) {
	game, err := gu.gr.FindGameRecordByGameId(game_id)
	if err != nil {
		return game, err
	}
	return game, nil
}
