package handlers

import (
	"github.com/Rosto4eks/eclipse/internal/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type Ihandler interface {
	Home(echo.Context) error
	Albums(echo.Context) error
	Album(echo.Context) error
	NewAlbum(echo.Context) error
	CreateNewAlbum(echo.Context) error
	SignIn(echo.Context) error
	SignUp(echo.Context) error
	NewUser(echo.Context) error
	Authorise(ctx echo.Context) error
	WriteCookie(token string, ctx echo.Context) error
	ReadCookie(ctx echo.Context) (string, error)
}

// first layer, handles incoming http requests
type handler struct {
	usecase usecase.Iusecase
}

func New(usecase usecase.Iusecase) *handler {
	return &handler{
		usecase,
	}
}

func (h *handler) Home(ctx echo.Context) error {
	return ctx.String(200, "HOME")
}

func (h *handler) WriteCookie(token string, ctx echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "jwt_token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	ctx.SetCookie(cookie)
	return ctx.String(http.StatusOK, "write a cookie")
}

func (h *handler) ReadCookie(ctx echo.Context) (string, error) {
	cookie, err := ctx.Cookie("jwt_token")
	if err != nil {
		return "", err
	}
	return cookie.Value, ctx.String(http.StatusOK, "read a cookie")
}
