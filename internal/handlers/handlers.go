package handlers

import (
	"github.com/Rosto4eks/eclipse/internal/usecase"
	"github.com/labstack/echo/v4"
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
