package usecase

import (
	"time"

	"github.com/tigaweb/reversi-app/backend/model"
	"github.com/tigaweb/reversi-app/backend/repository"
)

type ITurnUsecase interface {
	CreateFirstTurnAndBoardBySoloPlay(game model.Game) error
	FindTurnIdByGameIdAndTurnCount(game_id uint, turn_count int) (uint, error)
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
