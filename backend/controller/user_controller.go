package controller

import (
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/tigaweb/reversi-app/backend/model"
	"github.com/tigaweb/reversi-app/backend/usecase"
)

type IUserController interface {
	SignUp(c echo.Context) error
	Login(c echo.Context) error
	LogOut(c echo.Context) error
	CsrfToken(c echo.Context) error
	CheckAuth(c echo.Context) error
}

type userContoroller struct {
	uu usecase.IUserUsecase
}

func NewUserController(uu usecase.IUserUsecase) IUserController {
	return &userContoroller{uu}
}

func (uc userContoroller) SignUp(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	userRes, err := uc.uu.SignUp(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, userRes)
}

func (uc userContoroller) Login(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	tokenString, err := uc.uu.Login(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true // postman検証の時はコメントアウト
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode

	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}

func (uc userContoroller) LogOut(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Expires = time.Now()
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true // postman検証の時はコメントアウト
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode

	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}

func (uc userContoroller) CsrfToken(c echo.Context) error {
	token := c.Get("csrf").(string)
	return c.JSON(http.StatusOK, echo.Map{
		"csrf_token": token,
	})
}

func (uc userContoroller) CheckAuth(c echo.Context) error {
	// トークンが存在するか確認
	user := c.Get("user")
	if user == nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"is_login": false,
			"message":  "No token provided or token is invalid",
		})
	}
	// トークンの型アサーション
	token := user.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	user_id := claims["user_id"].(float64)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"is_login": true,
		"user_id":  user_id,
	})
}
