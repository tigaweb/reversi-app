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
}

type turnController struct {
	ts service.ITurnService
}

func NewTurnController(ts service.ITurnService) ITurnController {
	return &turnController{ts}
}

func (tc turnController) RegisterTurn(c echo.Context) error {
	fmt.Println("turnController.RegisterTurnの処理")
	req := model.RegisterTurnRequest{}
	if err := c.Bind(&req); err != nil {
		return err
	}
	fmt.Println(req.GameID)
	fmt.Println(req.TurnCount)
	fmt.Println(req.Move.Disc)
	fmt.Println(req.Move.X)
	fmt.Println(req.Move.Y)
	tc.ts.RegisterTurn(req.TurnCount, req.GameID, req.Move)
	return c.NoContent(http.StatusOK)
}
