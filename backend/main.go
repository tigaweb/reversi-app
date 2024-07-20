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
	// Repository
	userRepository := repository.NewUserRepository(db)
	gameRepository := repository.NewGameRepository(db)
	squareRepository := repository.NewSquareRepository(db)
	turnRepository := repository.NewTurnRepository(db)
	moveRepository := repository.NewMoveRepository(db)
	resultRepository := repository.NewResultRepository(db)
	// Validator
	userValidator := validator.NewUserValidator()
	// Usecase
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	gameUsecase := usecase.NewGameUsecase(gameRepository, resultRepository)
	turnUsecase := usecase.NewTurnUsecase(turnRepository, squareRepository)
	moveUsecase := usecase.NewMoveUsecase(moveRepository)
	squareUsecase := usecase.NewSquareUsecase(squareRepository)
	boardUsecase := usecase.NewBoardUsecase()
	reslutUsecase := usecase.NewResultUsecase(resultRepository)
	// Service
	turnService := service.NewTurnService(gameUsecase, squareUsecase, turnUsecase, boardUsecase, moveUsecase, reslutUsecase)
	// Controller
	userController := controller.NewUserController(userUsecase)
	gameController := controller.NewGameController(gameUsecase, turnUsecase, reslutUsecase)
	turnController := controller.NewTurnController(turnService)
	// router
	e := router.NewRouter(userController, gameController, turnController)
	e.Logger.Fatal(e.Start(":8080"))
}
