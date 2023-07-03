package handlers

import (
	"github.com/labstack/echo/v4"
)

// first layer, handles incoming http requests
type Handler struct {
	//usecase
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Home(ctx echo.Context) error {
	return ctx.String(200, "HOME")
}

func (h *Handler) Albums(ctx echo.Context) error {
	return ctx.String(200, "Albums")
}

func (h *Handler) Album(ctx echo.Context) error {
	return ctx.String(200, ctx.Param("id"))
}

func (h *Handler) NewAlbum(ctx echo.Context) error {
	return ctx.Render(201, "newAlbum.html", nil)
}

func (h *Handler) CreateNewAlbum(ctx echo.Context) error {
	return ctx.String(201, "created")
}
