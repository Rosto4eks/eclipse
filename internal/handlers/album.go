package handlers

import (
	"fmt"
	"strconv"

	"github.com/Rosto4eks/eclipse/internal/models"
	"github.com/labstack/echo/v4"
)

func (h *handler) Albums(ctx echo.Context) error {
	return ctx.String(200, "Albums")
}

func (h *handler) Album(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.Redirect(301, "/albums")
	}
	album, err := h.usecase.GetAlbumById(id)
	if err != nil {
		return ctx.Redirect(301, "/albums")
	}
	paths := make([]string, album.Images_count)
	for i := 0; i < album.Images_count; i++ {
		paths[i] = fmt.Sprintf("%s-%s/%d", album.Date[0:10], album.Name, i)
	}
	return ctx.Render(200, "album.html", map[string]interface{}{
		"album": album,
		"paths": paths,
	})
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
	usr, err := h.usecase.GetUserByName(ctx.FormValue("author"))
	if err != nil {
		return ctx.Render(401, "newAlbum.html", map[string]interface{}{
			"error": "author with given name does not exist",
		})
	}
	album := models.Album{
		Name:         ctx.FormValue("name"),
		Date:         ctx.FormValue("date"),
		Author_id:    usr.ID,
		Description:  ctx.FormValue("description"),
		Images_count: len(files),
	}

	if err = h.usecase.NewAlbum(files, album); err != nil {
		return ctx.Render(500, "newAlbum.html", map[string]interface{}{
			"error": err.Error(),
		})
	}

	return ctx.Redirect(301, "/albums")
}
