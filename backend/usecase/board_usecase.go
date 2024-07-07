package usecase

import (
	"fmt"

	"github.com/tigaweb/reversi-app/backend/model"
)

type IBoardUsecase interface {
	GetFlipPoints(move model.Move, board model.Board) (model.ListFlipPoints, error)
}

type boardUsecase struct {
}

func NewBoardUsecase() IBoardUsecase {
	return &boardUsecase{}
}

func (bu *boardUsecase) GetFlipPoints(move model.Move, board model.Board) (model.ListFlipPoints, error) {
	// 番兵アルゴリズムによる判定を行う
	// 番兵を考慮したWalledDiscを作成する(board.DiscsをW(3)で囲んだ配列を作成)
	walledBoard, err := getWallDiscs(board)
	if err != nil {
		return nil, err
	}

	// ひっくり返せる位置の取得
	listFlipPoints, err := listUpCanFlipPoints(move, *walledBoard)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return *listFlipPoints, err

}

func getWallDiscs(board model.Board) (*model.Board, error) {
	// オリジナルの盤面のサイズを取得
	// 縦軸の数を取得(8)
	originalRows := len(board.Discs)
	if originalRows == 0 {
		return nil, fmt.Errorf("盤面の情報が存在しません")
	}
	// 横軸の数を取得(8)
	originalCols := len(board.Discs[0])

	// 番兵を考慮した盤面の定義(10*10)
	newRows := originalRows + 2
	newCols := originalCols + 2

	// 番兵用の盤面を作成
	wallDiscs := make(model.Discs, newRows)
	for i := range wallDiscs {
		wallDiscs[i] = make([]model.Disc, newCols)
	}

	// 番兵用の盤面のすべての位置にWallをセット
	for y := 0; y < newRows; y++ {
		for x := 0; x < newCols; x++ {
			wallDiscs[y][x] = model.W
		}
	}

	// オリジナルの盤面を新しい盤面の中央にコピー
	for y := 0; y < originalRows; y++ {
		for x := 0; x < originalCols; x++ {
			wallDiscs[y+1][x+1] = board.Discs[y][x]
		}
	}

	return &model.Board{Discs: wallDiscs}, nil
}

func listUpCanFlipPoints(move model.Move, walledBoard model.Board) (*model.ListFlipPoints, error) {
	// ひっくり返せる位置の配列の配列の初期化
	listFlipPoints := &model.ListFlipPoints{}

	walledX := move.X + 1
	walledY := move.Y + 1

	checkFlipPoints := func(move model.Move, xMove int, yMove int, listFlipPoints *model.ListFlipPoints) {

		cursorX := walledX + xMove
		cursorY := walledY + yMove

		// ひっくり返せる石の位置の配列
		flipPoints := model.ListFlipPoints{}

		// 隣のマスの色が違う色の場合に処理を継続する
		for {
			if !isOppositeDisc(model.Disc(move.Disc), walledBoard.Discs[cursorY][cursorX]) {
				// fmt.Println("終了")
				break
			}
			fmt.Println("処理が進む")

			// リストに位置を追加
			flipPoints = append(flipPoints, model.Point{
				"X": cursorX - 1,
				"Y": cursorY - 1,
			})

			// 位置を一つずらす
			cursorX += xMove
			cursorY += yMove

			// 隣のマスの色が置いた石と同じ色の場合、処理を確定する
			if move.Disc == int(walledBoard.Discs[cursorY][cursorX]) {
				*listFlipPoints = append(*listFlipPoints, flipPoints...)
				break
			}
		}
	}

	// 上
	checkFlipPoints(move, 0, -1, listFlipPoints)
	// 左上
	checkFlipPoints(move, -1, -1, listFlipPoints)
	// 左
	checkFlipPoints(move, -1, 0, listFlipPoints)
	// 左下
	checkFlipPoints(move, -1, 1, listFlipPoints)
	// 下
	checkFlipPoints(move, 0, 1, listFlipPoints)
	// 右下
	checkFlipPoints(move, 1, 1, listFlipPoints)
	// 右下
	checkFlipPoints(move, 1, 0, listFlipPoints)
	// 右上
	checkFlipPoints(move, 1, -1, listFlipPoints)

	if len(*listFlipPoints) == 0 {
		return nil, fmt.Errorf("ひっくり返せる場所がありません")
	}

	return listFlipPoints, nil
}

func isOppositeDisc(disc1 model.Disc, disc2 model.Disc) bool {
	if disc1 == model.Dark && disc2 == model.Light {
		return true
	}
	if disc1 == model.Light && disc2 == model.Dark {
		return true
	}
	return false
}
