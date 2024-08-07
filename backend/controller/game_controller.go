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
	GetGameResult(c echo.Context) error
}

type gameController struct {
	gu usecase.IGameUsecase
	tu usecase.ITurnUsecase
	ru usecase.IResultUsecase
}

func NewGameController(gu usecase.IGameUsecase, tu usecase.ITurnUsecase, ru usecase.IResultUsecase) IGameController {
	return &gameController{gu, tu, ru}
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
	// Game作成時、最初の盤面(Turn,Squareを作成する処理)
	gc.tu.CreateFirstTurnAndBoardBySoloPlay(game)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, gameRes)
}

func (gc *gameController) GetGameResult(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	userIdUint := uint(userId.(float64))
	gameHistory, err := gc.ru.FindResultByUserId(userIdUint)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, gameHistory)
}
