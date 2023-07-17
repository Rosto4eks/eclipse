package handlers

import (
	"fmt"
	"github.com/Rosto4eks/eclipse/internal/models"
	"github.com/labstack/echo/v4"
	"strconv"
)

func (h *handler) GetAlbums(ctx echo.Context) error {
	headerName := h.authHeader(ctx)
	albums, err := h.usecase.GetAllAlbums()
	if err != nil {
		return err
	}
	return ctx.Render(301, "allAlbums.html", map[string]interface{}{
		"header": headerName,
		"albums": albums,
	})
}

func (h *handler) GetAlbum(ctx echo.Context) error {
	headerName := h.authHeader(ctx)
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
		"header": headerName,
		"album":  album,
		"paths":  paths,
	})
}

func (h *handler) GetNewAlbum(ctx echo.Context) error {
	headerName := h.authHeader(ctx)
	if err := h.auth(ctx, "author"); err != nil {
		return ctx.Redirect(301, "/")
	}
	return ctx.Render(200, "newAlbum.html", map[string]interface{}{
		"header": headerName,
	})
}

func (h *handler) PostNewAlbum(ctx echo.Context) error {
	headerName := h.authHeader(ctx)
	if err := h.auth(ctx, "author"); err != nil {
		return ctx.Redirect(301, "/")
	}
	form, err := ctx.MultipartForm()
	if err != nil {
		return ctx.Render(500, "newAlbum.html", map[string]interface{}{
			"header": headerName,
			"error":  "could not parse images",
		})
	}
	files := form.File["files"]
	usr, err := h.usecase.GetUserByName(ctx.FormValue("author"))
	if err != nil {
		return ctx.Render(401, "newAlbum.html", map[string]interface{}{
			"header": headerName,
			"error":  "author with given name does not exist",
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
			"header": headerName,
			"error":  err.Error(),
		})
	}

	return ctx.Redirect(301, "/albums")
}

func (h *handler) DeleteAlbum(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return err
	}
	err = h.usecase.DeleteAlbum(id)
	if err != nil {
		return err
	}

	return nil
}
