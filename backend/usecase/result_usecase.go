package usecase

import (
	"time"

	"github.com/tigaweb/reversi-app/backend/model"
	"github.com/tigaweb/reversi-app/backend/repository"
)

type IResultUsecase interface {
	RegisterResult(game_id uint, winner_by_id uint, winner_disc int) error
	FindResultByGameId(game_id uint) (model.GameResult, error)
	FindResultByUserId(user_id uint) ([]model.ResponseGameHistory, error)
}

type resultUsecase struct {
	rr repository.IResultRepository
}

func NewResultUsecase(rr repository.IResultRepository) IResultUsecase {
	return &resultUsecase{rr}
}

func (ru *resultUsecase) RegisterResult(game_id uint, winner_by_id uint, winner_disc int) error {
	result := &model.GameResult{}
	result.GameId = game_id
	result.WinnerById = winner_by_id
	result.WinnerDisc = winner_disc
	result.EndAt = time.Now()
	if err := ru.rr.RegisterResult(*result); err != nil {
		return err
	}
	return nil
}

func (ru *resultUsecase) FindResultByGameId(game_id uint) (model.GameResult, error) {
	result, err := ru.rr.FindResultByGameId(game_id)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (ru *resultUsecase) FindResultByUserId(user_id uint) ([]model.ResponseGameHistory, error) {
	gameHistory, err := ru.rr.FindResultByUserId(user_id)
	if err != nil {
		return nil, err
	}
	return gameHistory, err
}
