package service

import (
	"fmt"

	"github.com/tigaweb/reversi-app/backend/model"
	"github.com/tigaweb/reversi-app/backend/usecase"
)

type ITurnService interface {
	RegisterTurn(turn_count int, game_id uint, move model.Move) error
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
func (ts turnService) RegisterTurn(turn_count int, game_id uint, move model.Move) error {
	// 受け取ったgame_idのレコードが存在することの確認
	// turn_countが0だった場合 star_new_game_usecaseの実行
	game, err := ts.gu.FindGameRecordByGameId(game_id)
	if err != nil {
		return err
	}
	fmt.Println("RegisterTurnの処理")
	fmt.Println(game.ID)
	fmt.Println(turn_count)

	// 前回(このターンで石を置く前)のturn_countを取得する turn_count - 1
	// previous_turn_count := game.turn_count -1

	// 前回(このターンで石を置く前)の盤面を取得する GetLatestBoard(turn_count - 1) []Square

	// moveから石が置ける場所か判定する square_usecaseに実装 (Move, []Square) error
	// からのマス目ではない場合置くことはできない
	// ひっくり返せる石が存在することを判定する(ひっくり返せる位置のリスト)を作る (Move, []Square) x,yのリスト,error
	// ひっくり返せる点がない場合、置くことはできない
	// 次のターンの盤面を登録する Turnの登録、Moveの登録、Squareの登録

	// 動きを記録する
	return nil
}

// 最新(指定ターン)の盤面を取得して返す処理 GetLatestBoard
