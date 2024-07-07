package usecase

import (
	"github.com/tigaweb/reversi-app/backend/model"
	"github.com/tigaweb/reversi-app/backend/repository"
)

type ISquareUsecase interface {
	CreateSquares(turn_id uint, board model.Board) error
	GetBoardByTurnId(turn_id uint) (*model.Board, error)
	FindWinnerDiscByTurnId(turn_id uint) (model.Disc, error)
}

type squareUsecase struct {
	sr repository.ISquareRepository
}

func NewSquareUsecase(sr repository.ISquareRepository) ISquareUsecase {
	return &squareUsecase{sr}
}

func (su *squareUsecase) CreateSquares(turn_id uint, board model.Board) error {
	if err := su.sr.CreateSquares(turn_id, board); err != nil {
		return nil
	}
	return nil
}

func (su *squareUsecase) GetBoardByTurnId(turn_id uint) (*model.Board, error) {
	squares, err := su.sr.GetSquaresByTurnId(turn_id)
	if err != nil {
		return nil, err
	}

	// Boardの初期化 8*8の配列を初期化
	board := &model.Board{
		Discs: make([][]model.Disc, 8), // 縦軸(Y=8)の配列を作成
	}

	for i := range board.Discs {
		board.Discs[i] = make([]model.Disc, 8) // 配列の中に横軸(X=8)の配列を作成
	}

	for _, square := range squares {
		board.Discs[square.Y][square.X] = model.Disc(square.Disc)
	}

	return board, nil
}

func (su *squareUsecase) FindWinnerDiscByTurnId(turn_id uint) (model.Disc, error) {
	winner_disc, err := su.sr.FindWinnerDiscByTurnId(turn_id)
	if err != nil {
		return 0, err
	}
	return winner_disc, err
}
