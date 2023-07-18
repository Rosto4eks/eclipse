package handlers

import (
	"fmt"
	"github.com/Rosto4eks/eclipse/internal/models"
	"github.com/labstack/echo/v4"
	"strconv"
)

func (h *handler) GetArticles(ctx echo.Context) error {
	articles, err := h.usecase.GetAllArticles()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return ctx.Render(301, "article.html", map[string]interface{}{
		"articles": articles,
	})
}

func (h *handler) GetNewArticle(ctx echo.Context) error {
	if err := h.auth(ctx, "author"); err != nil {
		return ctx.Redirect(301, "/")
	}
	return ctx.Render(200, "newArticle.html", nil)
}

func (h *handler) PostNewArticle(ctx echo.Context) error {
	if err := h.auth(ctx, "author"); err != nil {
		return ctx.Redirect(301, "/")
	}

	usr, err := h.usecase.GetUserByName(ctx.FormValue("NameAuthor"))
	if err != nil {
		return ctx.Render(401, "newArticle.html", map[string]interface{}{
			"error": "author with given name doesn't exist",
		})
	}

	images_count, _ := strconv.Atoi(ctx.FormValue("ImagesCount"))
	article := models.Article{
		Name:        ctx.FormValue("name"),
		Theme:       ctx.FormValue("theme"),
		AuthorID:    usr.ID,
		ImagesCount: images_count,
		Date:        ctx.FormValue("date"),
		Text:        ctx.FormValue("text"),
	}

	err = h.usecase.NewArticle(article)
	if err != nil {
		return ctx.Render(500, "newArticle.html", map[string]interface{}{
			"error": err.Error(),
		})
	}

	return ctx.Redirect(301, "/articles")
}

func (h *handler) DeleteArticle(ctx echo.Context) error {
	return nil
}
