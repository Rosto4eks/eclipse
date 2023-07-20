package handlers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/Rosto4eks/eclipse/internal/models"
	"github.com/labstack/echo/v4"
)

func (h *handler) GetArticle(ctx echo.Context) error {
	headerName := h.authHeader(ctx)
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.Redirect(302, "/articles")
	}
	article, err := h.usecase.GetArticleById(id)
	if err != nil {
		return ctx.Redirect(302, "/articles")
	}
	comments, err := h.usecase.GetArticleComments(id)
	if err != nil {
		return ctx.Redirect(302, "/articles")
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
		return ctx.Redirect(302, "/")
	}
	headerName := h.authHeader(ctx)
	return ctx.Render(200, "newArticle.html", map[string]interface{}{
		"header": headerName,
	})
}

func (h *handler) PostNewArticle(ctx echo.Context) error {
	if err := h.auth(ctx, "author"); err != nil {
		return ctx.Redirect(302, "/")
	}
	author := h.authHeader(ctx)

	req := make(map[string]interface{})
	err := json.NewDecoder(ctx.Request().Body).Decode(&req)
	if err != nil {
		return err
	}
	user, err := h.usecase.GetUserByName(author)

	article := models.Article{
		Name:        req["title"].(string),
		Theme:       req["theme"].(string),
		AuthorID:    user.ID,
		ImagesCount: 0,
		Date:        time.Now().Format("2006-01-02"),
		Text:        req["text"].(string),
	}

	err = h.usecase.NewArticle(article)
	if err != nil {
		return ctx.JSON(500, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return ctx.JSON(201, map[string]interface{}{
		"error": nil,
	})
}

func (h *handler) DeleteArticle(ctx echo.Context) error {
	if err := h.auth(ctx, "author"); err != nil {
		return ctx.JSON(400, map[string]interface{}{
			"success": false,
			"message": "permission denied",
		})
	}
	articleId, _ := strconv.Atoi(ctx.Param("article_id"))
	article, err := h.usecase.GetArticleById(articleId)
	if err != nil {
		return ctx.JSON(400, map[string]interface{}{
			"success": false,
			"message": "can't find article",
		})
	}
	name := h.authHeader(ctx)

	if article.NameAuthor != name {
		return ctx.JSON(400, map[string]interface{}{
			"success": false,
			"message": "invalid person",
		})
	}
	err = h.usecase.DeleteArticle(articleId)
	if err != nil {
		fmt.Println(err)
		return ctx.JSON(500, map[string]interface{}{
			"success": false,
			"message": "can't delete article",
		})
	}
	return ctx.JSON(200, map[string]interface{}{
		"success": true,
		"message": "",
	})
}

func (h *handler) ChangeArticle(ctx echo.Context) error {
	return ctx.JSON(200, map[string]interface{}{
		"success": true,
		"message": "",
	})
}
