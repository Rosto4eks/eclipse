package handlers

import (
	"github.com/Rosto4eks/eclipse/internal/models"
	"github.com/labstack/echo/v4"
)

func (h *handler) Albums(ctx echo.Context) error {
	return ctx.String(200, "Albums")
}

func (h *handler) Album(ctx echo.Context) error {
	return ctx.String(200, ctx.Param("id"))
}

func (h *handler) NewAlbum(ctx echo.Context) error {
	return ctx.Render(200, "newAlbum.html", nil)
}

func (h *handler) CreateNewAlbum(ctx echo.Context) error {
	form, err := ctx.MultipartForm()
	if err != nil {
		return ctx.Render(500, "newAlbum.html", map[string]interface{}{
			"error": "could not parse images",
		})
	}
	files := form.File["files"]

	album := models.Album{
		Name:        ctx.FormValue("name"),
		Date:        ctx.FormValue("date"),
		Author:      ctx.FormValue("author"),
		Description: ctx.FormValue("description"),
		Count:       len(files),
	}

	if err = h.usecase.NewAlbum(files, album); err != nil {
		return ctx.Render(500, "newAlbum.html", map[string]interface{}{
			"error": err.Error(),
		})
	}

	return ctx.Redirect(301, "/albums")
}
