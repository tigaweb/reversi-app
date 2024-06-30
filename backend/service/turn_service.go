package service

import (
	"fmt"

	"github.com/tigaweb/reversi-app/backend/model"
	"github.com/tigaweb/reversi-app/backend/usecase"
)

type ITurnService interface {
	RegisterTurn(game_id uint, move model.Move) error
}

type turnService struct {
	// game_usecase,turn_usecase,move_usecase,square_usecaseのユースケースを実装
	gu usecase.IGameUsecase
}

func NewTurnService(gu usecase.IGameUsecase) ITurnService {
	return &turnService{gu}
}

// 初回アクセス時(0ターン目)の盤面を登録して返す処理 GetNewBoard (game_id) {game_id, Board}

// ターンを登録する処理 RegisterTurn ()
func (ts turnService) RegisterTurn(game_id uint, move model.Move) error {
	// 受け取ったgame_idのレコードが存在することの確認
	game, err := ts.gu.FindGameRecordByGameId(game_id)
	if err != nil {
		return err
	}
	fmt.Println("RegisterTurnの処理")
	fmt.Println(game.ID)

	// 前回のturn_countを取得する

	// moveから石が置ける場所か判定する
	// ひっくり返せる石が存在することを判定する(ひっくり返せる位置のリスト)を作る
	// 次のターンの盤面を登録する

	// 動きを記録する
	return nil
}

// 最新の盤面を取得して返す処理 GetLatestBoard
