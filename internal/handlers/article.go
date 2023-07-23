package handlers

import (
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
	var author string
	if err := h.auth(ctx, "author"); err == nil {
		author = "author"
	}
	articles, err := h.usecase.GetAllArticles()
	if err != nil {
		return err
	}
	return ctx.Render(200, "allArticles.html", map[string]interface{}{
		"header":   headerName,
		"articles": articles,
		"author":   author,
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

	form, err := ctx.MultipartForm()
	if err != nil {
		fmt.Println(err)
		return ctx.JSON(400, map[string]interface{}{
			"error": err.Error(),
		})
	}
	images := form.File["images"]
	user, err := h.usecase.GetUserByName(author)

	article := models.Article{
		Name:        ctx.FormValue("title"),
		Theme:       ctx.FormValue("theme"),
		AuthorID:    user.ID,
		ImagesCount: len(images),
		Date:        time.Now().Format("2006-01-02"),
		Text:        ctx.FormValue("text"),
		TitleColor:  ctx.FormValue("color"),
	}
	if article.ImagesCount == 0 {
		return ctx.JSON(400, map[string]interface{}{
			"error": "images not uploaded",
		})
	}
	if err = h.usecase.NewArticle(images, article); err != nil {
		return ctx.JSON(500, map[string]interface{}{
			"error": err.Error(),
		})
	}
	fmt.Println("beb")
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
	if err := h.auth(ctx, "author"); err != nil {
		return ctx.JSON(403, map[string]interface{}{
			"success": false,
			"message": "forbidden",
		})
	}
	headerName := h.authHeader(ctx)
	jsonBody := make(map[string]interface{}, 0)
	err := json.NewDecoder(ctx.Request().Body).Decode(&jsonBody)
	if err != nil {
		return ctx.JSON(500, map[string]interface{}{
			"success": false,
			"message": "server error",
		})
	}
	articleId, _ := strconv.Atoi(jsonBody["articleId"].(string))
	newText := jsonBody["text"].(string)
	article, err := h.usecase.GetArticleById(articleId)
	if err != nil {
		return ctx.JSON(500, map[string]interface{}{
			"success": false,
			"message": "server error",
		})
	}
	if article.NameAuthor != headerName {
		return ctx.JSON(403, map[string]interface{}{
			"success": false,
			"message": "forbidden",
		})
	}

	err = h.usecase.ChangeArticle(articleId, newText)
	if err != nil {
		return ctx.JSON(500, map[string]interface{}{
			"success": false,
			"message": "server error",
		})
	}
	return ctx.JSON(200, map[string]interface{}{
		"success": true,
		"message": "",
	})
}
