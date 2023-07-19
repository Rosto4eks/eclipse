package handlers

import (
	"errors"
	"github.com/Rosto4eks/eclipse/internal/models"
	"github.com/labstack/echo/v4"
	"strconv"
)

func (h *handler) GetArticle(ctx echo.Context) error {
	headerName := h.authHeader(ctx)
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.Redirect(301, "/articles")
	}
	article, err := h.usecase.GetArticleById(id)
	if err != nil {
		return ctx.Redirect(301, "/articles")
	}
	comments, err := h.usecase.GetArticleComments(id)
	if err != nil {
		return ctx.Redirect(301, "/articles")
	}
	return ctx.Render(200, "article.html", map[string]interface{}{
		"header":   headerName,
		"article":  article,
		"comments": comments,
	})
}

func (h *handler) GetArticles(ctx echo.Context) error {
	headerName := h.authHeader(ctx)
	articles, err := h.usecase.GetAllArticles()
	if err != nil {
		return err
	}
	return ctx.Render(200, "allArticles.html", map[string]interface{}{
		"header":   headerName,
		"articles": articles,
	})
}

func (h *handler) GetNewArticle(ctx echo.Context) error {
	if err := h.auth(ctx, "author"); err != nil {
		return ctx.Redirect(301, "/")
	}
	headerName := h.authHeader(ctx)
	return ctx.Render(200, "newArticle.html", map[string]interface{}{
		"header": headerName,
	})
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
	if err := h.auth(ctx, "author"); err != nil {
		return ctx.Redirect(301, "/")
	}

	id, _ := strconv.Atoi(ctx.Param("id"))
	article, err := h.usecase.GetArticleById(id)
	if err != nil {
		return err
	}
	name := h.authHeader(ctx)

	if article.NameAuthor != name {
		return errors.New("invalid person")
	}
	err = h.usecase.DeleteArticle(id)
	if err != nil {
		return err
	}
	return ctx.Render(200, "allArticles.html", nil)
}
