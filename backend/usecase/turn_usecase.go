package usecase

import (
	"time"

	"github.com/tigaweb/reversi-app/backend/model"
	"github.com/tigaweb/reversi-app/backend/repository"
)

type ITurnUsecase interface {
	CreateFirstTurnAndBoardBySoloPlay(game model.Game) error
	FindTurnIdByGameIdAndTurnCount(game_id uint, turn_count int) (uint, error)
	RegisterTurn(turn *model.Turn) error
	FindMaxTurnCountByGameId(game_id uint) (model.Turn, error)
}

type turnUsecase struct {
	tr repository.ITurnRepository
	sr repository.ISquareRepository
}

func NewTurnUsecase(tr repository.ITurnRepository, sr repository.ISquareRepository) ITurnUsecase {
	return &turnUsecase{tr, sr}
}

func (tu *turnUsecase) CreateFirstTurnAndBoardBySoloPlay(game model.Game) error {
	// turnを作成 0ターン目 次の石；0(黒が先行)
	firstTurn := model.Turn{
		GameId:      game.ID,
		TurnCount:   0,
		NextDisc:    int(model.D),
		EndAt:       time.Now(),
		CreatedByID: game.CreatedByID,
	}
	if err := tu.tr.RegisterTurn(&firstTurn); err != nil {
		return err
	}
	// NOTE Gameの作成とGameIDを受け取って盤面を登録する処理を素直に分ければよかった..?
	// 初期盤面の登録
	boad := model.NewBoard() // 初期盤面の取得
	if err := tu.sr.CreateSquares(firstTurn.ID, *boad); err != nil {
		return err
	}
	return nil
}

func (tu *turnUsecase) FindTurnIdByGameIdAndTurnCount(game_id uint, turn_count int) (uint, error) {
	turn, err := tu.tr.FindTurnByGameIdAndTurnCount(game_id, turn_count)
	if err != nil {
		return 0, err
	}
	return turn.ID, nil
}

func (tu *turnUsecase) RegisterTurn(turn *model.Turn) error {
	if err := tu.tr.RegisterTurn(turn); err != nil {
		return err
	}
	return nil
}

func (tu *turnUsecase) FindMaxTurnCountByGameId(game_id uint) (model.Turn, error) {
	turn, err := tu.tr.FindMaxTurnCountByGameId(game_id)
	if err != nil {
		return turn, err
	}
	return turn, nil
}
