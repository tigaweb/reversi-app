package service

import (
	"fmt"

	"github.com/tigaweb/reversi-app/backend/model"
	"github.com/tigaweb/reversi-app/backend/usecase"
)

type ITurnService interface {
	RegisterTurn(turn_count int, game_id uint, move model.Move) error
	// game_idから、最新のTurn,Squareを取得して返す処理
}

type turnService struct {
	gu usecase.IGameUsecase
	su usecase.ISquareUsecase
	tu usecase.ITurnUsecase
	// move_usecaseを実装
}

func NewTurnService(gu usecase.IGameUsecase, su usecase.ISquareUsecase, tu usecase.ITurnUsecase) ITurnService {
	return &turnService{gu, su, tu}
}

// ターンを登録する処理 RegisterTurn ()
func (ts turnService) RegisterTurn(turn_count int, game_id uint, move model.Move) error {
	// 受け取ったgame_idのレコードが存在することの確認
	game, err := ts.gu.FindGameRecordByGameId(game_id)
	if err != nil {
		return err
	}
	fmt.Println("RegisterTurnの処理")
	fmt.Println(game.ID)
	fmt.Println(turn_count)

	// 前回(このターンで石を置く前)のturn_countを取得する turn_count - 1
	previous_turn_count := turn_count - 1

	// game_idとturn_countからturn_idを取得
	previous_turn_id, err := ts.tu.FindTurnIdByGameIdAndTurnCount(game.ID, previous_turn_count)
	if err != nil {
		return err
	}
	previous_board, err := ts.su.GetBoardByTurnId(previous_turn_id)
	if err != nil {
		return err
	}
	fmt.Println(previous_board.Discs)

	// 判定処理
	// 石をおこうとしている場所が空か
	// 	空ではない場合エラーを返す
	// moveと現在のboardから、位置が空か判定
	// ひっくり返せる場所があるか
	// 	moveと現在のboardから、ひっくり返せる位置のリストを取得
	// 		位置のリストが空の場合、エラーを返す
	// ひっくり返せる場所がある場合、現在のboardとリストから次の盤面を作成する
	// 次の盤面と今回配置した石と逆の色の石から、ひっくり返せる場所があるか確認する
	// 	存在する場合、次の色の石は今回の石の逆の色の石を指定する
	//  存在しない場合、かつ、さらに今回の石も配置する場所がない場合、盤面のそれぞれの石の色をカウントし、多い方が勝利者となる

	// Turnを登録
	// Moveを登録
	// Squareを登録

	return nil
}

// game_idから、最新のTurn,Squareを取得して返す処理
