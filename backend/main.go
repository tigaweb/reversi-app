package main

import (
	"github.com/tigaweb/reversi-app/backend/controller"
	"github.com/tigaweb/reversi-app/backend/db"
	"github.com/tigaweb/reversi-app/backend/repository"
	"github.com/tigaweb/reversi-app/backend/router"
	"github.com/tigaweb/reversi-app/backend/service"
	"github.com/tigaweb/reversi-app/backend/usecase"
	"github.com/tigaweb/reversi-app/backend/validator"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	userValidator := validator.NewUserValidator()
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	userController := controller.NewUserController(userUsecase)
	gameRepository := repository.NewGameRepository(db)
	gameUsecase := usecase.NewGameUsecase(gameRepository)
	turnRepository := repository.NewTurnRepository(db)
	squareRepository := repository.NewSquareRepository(db)
	turnUsecase := usecase.NewTurnUsecase(turnRepository, squareRepository)
	squareUsecase := usecase.NewSquareRepository(squareRepository)
	gameController := controller.NewGameController(gameUsecase, turnUsecase)
	boardUsecase := usecase.NewBoardUsecase()
	turnService := service.NewTurnService(gameUsecase, squareUsecase, turnUsecase, boardUsecase)
	turnController := controller.NewTurnController(turnService)
	e := router.NewRouter(userController, gameController, turnController)
	e.Logger.Fatal(e.Start(":8080"))
}
