package main

import (
	"github.com/tigaweb/reversi-app/backend/controller"
	"github.com/tigaweb/reversi-app/backend/db"
	"github.com/tigaweb/reversi-app/backend/repository"
	"github.com/tigaweb/reversi-app/backend/router"
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
	gameController := controller.NewGameController(gameUsecase)
	e := router.NewRouter(userController, gameController)
	e.Logger.Fatal(e.Start(":8080"))
}
