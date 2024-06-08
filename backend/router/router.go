package router

import (
	"github.com/labstack/echo"
	"github.com/tigaweb/reversi-app/backend/controller"
)

func NewRouter(uc controller.IUserController) *echo.Echo {
	e := echo.New()

	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.Login)
	e.POST("/logout", uc.LogOut)

	return e
}
