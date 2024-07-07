package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/tigaweb/reversi-app/backend/model"
	"github.com/tigaweb/reversi-app/backend/service"
)

type ITurnController interface {
	RegisterTurn(c echo.Context) error
	FindLatestTurn(c echo.Context) error
}

type turnController struct {
	ts service.ITurnService
}

func NewTurnController(ts service.ITurnService) ITurnController {
	return &turnController{ts}
}

func (tc turnController) RegisterTurn(c echo.Context) error {
	req := &model.RegisterTurnRequest{}
	if err := c.Bind(req); err != nil {
		return err
	}
	tc.ts.RegisterTurn(req.TurnCount, req.GameID, req.Move)
	return c.NoContent(http.StatusOK)
}

// game_idから最新の盤面を取得して返す処理
func (tc turnController) FindLatestTurn(c echo.Context) error {
	game_id := c.Param("game_id")
	gameIDUint, err := strconv.ParseUint(game_id, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid game_id"})
	}

	res := &model.FindLatestTurnResponse{}
	res.GameID = uint(gameIDUint)

	// ターンカウントの最大を取得
	latest_turn, err := tc.ts.FindMaxTurnCountByGameId(res.GameID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// 盤面を取得
	latest_board, err := tc.ts.GetBoardByTurnId(latest_turn.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	res.Board = latest_board.Discs

	// 次の石
	res.NextDisc = latest_turn.NextDisc

	// 勝者の石
	result, err := tc.ts.FindResultByGameId(res.GameID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	res.WinnerDisc = 0
	if result.ID != 0 {
		res.WinnerDisc = result.WinnerDisc
	}

	return c.JSON(http.StatusOK, res)
}
