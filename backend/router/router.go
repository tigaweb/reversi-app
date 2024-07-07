package router

import (
	"net/http"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tigaweb/reversi-app/backend/controller"
)

func NewRouter(
	uc controller.IUserController,
	gc controller.IGameController,
	tc controller.ITurnController,
) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:3000",
			os.Getenv("FE_URL"),
		},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders,
			echo.HeaderXCSRFToken,
		},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
	}))
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath:     "/",
		CookieDomain:   os.Getenv("API_DOMAIN"),
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteNoneMode,
		// CookieSameSite: http.SameSiteDefaultMode, // POSTMAN動作確認用
		CookieMaxAge: 3600, // CSRFの有効期限を秒で設定
	}))

	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.Login)
	e.POST("/logout", uc.LogOut)
	e.GET("/csrf", uc.CsrfToken)

	c := e.Group("/check-auth")
	c.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
		ErrorHandler: func(c echo.Context, err error) error {
			return echo.NewHTTPError(http.StatusCreated, map[string]interface{}{
				"is_login": false,
				"message":  "first access",
			})
		},
	}))
	c.GET("", uc.CheckAuth)

	g := e.Group("/games")
	g.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	g.POST("", gc.CreateGame)
	t := g.Group("/latest/turns/")
	t.GET(":game_id", tc.FindLatestTurn)
	t.POST("", tc.RegisterTurn)

	return e
}
