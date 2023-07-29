package handlers

import (
	"fmt"
	"github.com/Rosto4eks/eclipse/internal/models"
	"github.com/labstack/echo/v4"
	"strconv"
)

func (h *handler) GetAlbums(ctx echo.Context) error {
	headerName := h.authHeader(ctx)
	var author string
	if err := h.auth(ctx, "author"); err == nil {
		author = "author"
	}
	albums, err := h.usecase.GetAlbums(0, 5)
	if err != nil {
		h.logger.Error("handlers", "GetAlbums", err)
		return err
	}
	return ctx.Render(200, "allAlbums.html", map[string]interface{}{
		"header": headerName,
		"albums": albums,
		"author": author,
	})
}

func (h *handler) LoadAlbums(ctx echo.Context) error {
	offset, err := strconv.Atoi(ctx.QueryParam("offset"))
	if err != nil {
		offset = 0
	}

	count, err := strconv.Atoi(ctx.QueryParam("count"))
	if err != nil {
		count = 5
	}

	albums, err := h.usecase.GetAlbums(offset, count)
	if err != nil {
		h.logger.Error("handlers", "GetAlbums", err)
		return ctx.JSON(400, map[string]interface{}{
			"error": err.Error(),
		})
	}
	return ctx.JSON(200, map[string]interface{}{
		"albums": albums,
	})
}

func (h *handler) GetAlbum(ctx echo.Context) error {
	headerName := h.authHeader(ctx)
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		h.logger.Error("handlers", "GetAlbum", err)
		return ctx.Redirect(302, "/albums")
	}
	album, err := h.usecase.GetAlbumById(id)
	if err != nil {
		h.logger.Error("handlers", "GetAlbum", err)
		return ctx.Redirect(302, "/albums")
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
		h.logger.Error("handlers", "GetNewAlbum", err)
		return ctx.Redirect(302, "/")
	}
	return ctx.Render(200, "newAlbum.html", map[string]interface{}{
		"header": headerName,
	})
}

func (h *handler) PostNewAlbum(ctx echo.Context) error {
	headerName := h.authHeader(ctx)
	if err := h.auth(ctx, "author"); err != nil {
		h.logger.Error("handlers", "PostNewAlbum", err)
		return ctx.Redirect(302, "/")
	}
	form, err := ctx.MultipartForm()
	if err != nil {
		h.logger.Error("handlers", "PostNewAlbum", err)
		return ctx.Render(500, "newAlbum.html", map[string]interface{}{
			"header": headerName,
			"error":  "could not parse images",
		})
	}
	files := form.File["files"]
	usr, err := h.usecase.GetUserByName(headerName)
	if err != nil {
		h.logger.Error("handlers", "PostNewAlbum", err)
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
	if album.Images_count == 0 {
		h.logger.Warning("handlers", "PostNewAlbum", "0 images uploaded")
		return ctx.Render(400, "newAlbum.html", map[string]interface{}{
			"header": headerName,
			"error":  "images not uploaded",
		})
	}
	if err = h.usecase.NewAlbum(files, album); err != nil {
		h.logger.Error("handlers", "PostNewAlbum", err)
		return ctx.Render(500, "newAlbum.html", map[string]interface{}{
			"header": headerName,
			"error":  err.Error(),
		})
	}

	return ctx.Redirect(302, "/albums")
}

func (h *handler) DeleteAlbum(ctx echo.Context) error {
	if err := h.auth(ctx, "author"); err != nil {
		h.logger.Error("handlers", "DeleteAlbum", err)
		return ctx.JSON(400, map[string]interface{}{
			"success": false,
			"message": "permisson denied",
		})
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		h.logger.Error("handlers", "DeleteAlbum", err)
		return ctx.JSON(500, map[string]interface{}{
			"success": false,
			"message": "cant parse id",
		})
	}

	name := h.authHeader(ctx)
	album, err := h.usecase.GetAlbumById(id)
	if err != nil {
		h.logger.Error("handlers", "DeleteAlbum", err)
		return ctx.JSON(404, map[string]interface{}{
			"success": false,
			"message": "cant find album",
		})
	}

	if album.Author != name {
		h.logger.Warning("handlers", "DeleteAlbum", fmt.Sprintf("permission denied album author: %s, name: %s", album.Author, name))
		return ctx.JSON(400, map[string]interface{}{
			"success": false,
			"message": "permisson denied",
		})
	}

	err = h.usecase.DeleteAlbum(id)
	if err != nil {
		h.logger.Error("handlers", "DeleteAlbum", err)
		return ctx.JSON(500, map[string]interface{}{
			"success": false,
			"message": "cant delete album",
		})
	}

	return ctx.JSON(200, map[string]interface{}{
		"success": true,
		"message": "",
	})
}
