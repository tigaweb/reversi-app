package controller

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tigaweb/reversi-app/backend/model"
	"github.com/tigaweb/reversi-app/backend/service"
)

type ITurnController interface {
	RegisterTurn(c echo.Context) error
	// game_idから最新の盤面を取得して返す処理
}

type turnController struct {
	ts service.ITurnService
}

func NewTurnController(ts service.ITurnService) ITurnController {
	return &turnController{ts}
}

func (tc turnController) RegisterTurn(c echo.Context) error {
	req := model.RegisterTurnRequest{}
	if err := c.Bind(&req); err != nil {
		return err
	}
	tc.ts.RegisterTurn(req.TurnCount, req.GameID, req.Move)
	return c.NoContent(http.StatusOK)
}

// game_idから最新の盤面を取得して返す処理
