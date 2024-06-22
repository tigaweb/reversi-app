package controller

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/tigaweb/reversi-app/backend/model"
	"github.com/tigaweb/reversi-app/backend/usecase"
)

type IGameController interface {
	CreateGame(c echo.Context) error
	// GetGameResult(c echo.Context) error
}

type gameController struct {
	gu usecase.IGameUsecase
}

func NewGameController(gu usecase.IGameUsecase) IGameController {
	return &gameController{gu}
}

func (gc *gameController) CreateGame(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	game := model.Game{}
	game.StartedAt = time.Now()
	game.CreatedByID = uint(userId.(float64))
	game.JoinById = uint(userId.(float64))
	gameRes, err := gc.gu.NewGame(&game)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, gameRes)
}
