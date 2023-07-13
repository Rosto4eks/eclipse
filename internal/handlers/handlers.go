package handlers

import (
	"github.com/Rosto4eks/eclipse/internal/usecase"
	"github.com/labstack/echo/v4"
)

type Ihandler interface {
	GetHome(echo.Context) error
	GetAlbums(echo.Context) error
	GetAlbum(echo.Context) error
	GetNewAlbum(echo.Context) error
	PostNewAlbum(echo.Context) error
	GetSignIn(echo.Context) error
	GetSignUp(echo.Context) error
	PostSignUp(echo.Context) error
	PostSignIn(ctx echo.Context) error
	auth(echo.Context, string) error
	writeJWT(token string, ctx echo.Context)
	readJWT(echo.Context) (string, error)
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

func (h *handler) GetHome(ctx echo.Context) error {
	return ctx.String(200, "HOME")
}
