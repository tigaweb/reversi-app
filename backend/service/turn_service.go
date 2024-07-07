package service

import (
	"fmt"
	"time"

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
	bu usecase.IBoardUsecase
	mu usecase.IMoveUsecase
	ru usecase.IResultUsecase
}

func NewTurnService(
	gu usecase.IGameUsecase,
	su usecase.ISquareUsecase,
	tu usecase.ITurnUsecase,
	bu usecase.IBoardUsecase,
	mu usecase.IMoveUsecase,
	ru usecase.IResultUsecase,
) ITurnService {
	return &turnService{gu, su, tu, bu, mu, ru}
}

// ターンを登録する処理
func (ts turnService) RegisterTurn(turn_count int, game_id uint, move model.Move) error {
	// 受け取ったgame_idのレコードが存在することの確認
	game, err := ts.gu.FindGameRecordByGameId(game_id)
	if err != nil {
		return err
	}

	// 前回(このターンで石を置く前)のturn_countを取得する
	previous_turn_count := turn_count - 1

	// game_idとturn_countからturn_idを取得
	previous_turn_id, err := ts.tu.FindTurnIdByGameIdAndTurnCount(game.ID, previous_turn_count)
	if err != nil {
		return err
	}
	current_board, err := ts.su.GetBoardByTurnId(previous_turn_id)
	if err != nil {
		return err
	}

	// 判定処理(空の場所ではない場合エラー)
	if current_board.Discs[move.Y][move.X] != model.E {
		return fmt.Errorf("すでに石がある場所です")
	}

	// ひっくり返せる場所のリストを取得(ない場合エラー)
	listFlipPoints, err := ts.bu.GetFlipPoints(move, *current_board)
	if err != nil {
		return err
	}

	// ひっくり返せる場所がある場合、次の盤面を作成
	if err := ts.bu.FlipDiscs(current_board, listFlipPoints, move); err != nil {
		return err
	}

	turn := &model.Turn{}

	// 次の盤面に逆の色の石を配置できるか確認
	reversedColor := model.ReverseColor(model.Disc(move.Disc))
	result_flag := 0
	if ts.bu.SerchCanFlipPointsByDisc(reversedColor, *current_board) {
		// 存在する場合、次ターンの石は逆の色
		turn.NextDisc = int(reversedColor)
	} else if ts.bu.SerchCanFlipPointsByDisc(model.Disc(move.Disc), *current_board) {
		// 存在せず、今回の石がまた置ける場合は連続して石を打つ
		turn.NextDisc = move.Disc
	} else {
		// 何もおけない場合、勝敗がついたと判断
		result_flag = 1

	}

	// Turnを登録
	turn.GameId = game.ID
	turn.TurnCount = turn_count
	turn.CreatedByID = game.CreatedBy.ID
	turn.EndAt = time.Now()
	if err := ts.tu.RegisterTurn(turn); err != nil {
		return err
	}

	// Moveを登録
	move.TurnId = turn.ID
	if err := ts.mu.RegisterMove(&move); err != nil {
		return err
	}

	// Squareを登録
	if err := ts.su.CreateSquares(turn.ID, *current_board); err != nil {
		return err
	}

	// 勝敗がついている場合、resultを記録
	if result_flag == 1 {
		// 盤面から多い方の石を計算する
		winner_disc, err := ts.su.FindWinnerDiscByTurnId(turn.ID)
		if err != nil {
			return err
		}
		// game_resultに書き込む
		if err := ts.ru.RegisterResult(
			game_id,
			game.CreatedByID,
			int(winner_disc),
		); err != nil {
			return err
		}
	}

	return nil
}

// game_idから、最新のTurn,Squareを取得して返す処理
