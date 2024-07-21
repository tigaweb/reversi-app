package controller

import (
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
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
	if err := tc.ts.RegisterTurn(req.TurnCount, req.GameID, req.Move); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

// game_idから最新の盤面を取得して返す処理
func (tc turnController) FindLatestTurn(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	userIdUint := uint(userId.(float64))

	game_id := c.Param("game_id")
	gameIDUint64, err := strconv.ParseUint(game_id, 10, 64)
	gameIDUint := uint(gameIDUint64)
	// 不正なgame_idの検出
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid game_id")
	}

	// 対戦に参加していないユーザーからのアクセスを不正として検出する
	if err := tc.ts.CheckParticipationByUserId(gameIDUint, userIdUint); err != nil {
		return c.JSON(http.StatusBadRequest, "not your business")
	}

	res := &model.FindLatestTurnResponse{}
	res.GameID = gameIDUint

	// ターンカウントの最大を取得
	latest_turn, err := tc.ts.FindMaxTurnCountByGameId(res.GameID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	res.TurnCount = latest_turn.TurnCount

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
