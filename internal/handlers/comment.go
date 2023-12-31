package handlers

import (
	"encoding/json"
	"github.com/Rosto4eks/eclipse/internal/models"
	"github.com/labstack/echo/v4"
	"strconv"
	"time"
)

func (h *handler) GetComments(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		h.logger.Error("handlers", "GetComments", err)
		return ctx.Redirect(302, "/articles")
	}

	comments, err := h.usecase.GetArticleComments(id)
	if err != nil {
		h.logger.Error("handlers", "GetComments", err)
		return ctx.Redirect(302, "/articles")
	}

	article, err := h.usecase.GetArticleById(id)
	if err != nil {
		h.logger.Error("handlers", "GetComments", err)
		return ctx.Redirect(302, "/articles")
	}

	return ctx.Render(200, "article.html", map[string]interface{}{
		"article":  article,
		"comments": comments,
	})
}

func (h *handler) DeleteComment(ctx echo.Context) error {
	userErr := h.auth(ctx, "user")
	authorErr := h.auth(ctx, "author")
	if authorErr != nil {
		if userErr != nil {
			h.logger.Error("handlers", "DeleteComment", authorErr)
			return ctx.JSON(403, map[string]interface{}{
				"success": false,
				"message": "forbidden",
			})
		}
	}

	commentId, err := strconv.Atoi(ctx.Param("comment_id"))
	if err != nil {
		h.logger.Error("handlers", "DeleteComment", err)
		return ctx.JSON(400, map[string]interface{}{
			"success": false,
			"message": "invalid argument",
		})
	}
	comment, err := h.usecase.GetCommentById(commentId)
	if err != nil {
		h.logger.Error("handlers", "DeleteComment", err)
		return ctx.JSON(400, map[string]interface{}{
			"success": false,
			"message": "there are no such comment",
		})
	}
	name := h.authHeader(ctx)
	if name != comment.AuthorName {
		h.logger.Error("handlers", "DeleteComment", err)
		return ctx.JSON(400, map[string]interface{}{
			"success": false,
			"message": "Permission denied",
		})
	}
	err = h.usecase.DeleteComment(commentId)
	if err != nil {
		h.logger.Error("handlers", "DeleteComment", err)
		return ctx.JSON(500, map[string]interface{}{
			"success": false,
			"message": "server error",
		})
	}

	return ctx.JSON(201, map[string]interface{}{
		"success": true,
		"message": "success",
	})
}

func (h *handler) PostNewComment(ctx echo.Context) error {
	if err := h.auth(ctx, "author"); err != nil {
		if err := h.auth(ctx, "user"); err != nil {
			h.logger.Error("handlers", "PostNewComment", err)
			return ctx.JSON(403, map[string]interface{}{
				"success": false,
				"message": "forbidden",
			})
		}
	}
	jsonBody := make(map[string]interface{})
	err := json.NewDecoder(ctx.Request().Body).Decode(&jsonBody)
	if err != nil {
		h.logger.Error("handlers", "PostNewComment", err)
		return ctx.JSON(500, map[string]interface{}{
			"success": false,
			"message": "cant parse JSON",
		})
	}
	user, err := h.usecase.GetUserByName(jsonBody["author"].(string))
	if err != nil {
		h.logger.Error("handlers", "PostNewComment", err)
		return ctx.JSON(400, map[string]interface{}{
			"success": false,
			"message": "cannot find user",
		})
	}
	articleId, _ := strconv.Atoi(ctx.Param("article_id"))
	date := time.Now().Format("2006-01-02 15:04")
	comment := models.Comment{
		UserId:    user.ID,
		ArticleID: articleId,
		Text:      jsonBody["text"].(string),
		Date:      date,
	}
	comment_id, err := h.usecase.AddNewComment(comment)
	if err != nil {
		h.logger.Error("handlers", "PostNewComment", err)
		return ctx.JSON(400, map[string]interface{}{
			"success": false,
			"message": "cannot add new comment",
		})
	}
	return ctx.JSON(201, map[string]interface{}{
		"success":    true,
		"date":       date,
		"comment_id": comment_id,
		"message":    "",
	})
}

func (h *handler) ChangeComment(ctx echo.Context) error {
	if err := h.auth(ctx, "author"); err != nil {
		if err := h.auth(ctx, "user"); err != nil {
			h.logger.Error("handlers", "ChangeComment", err)
			return ctx.JSON(403, map[string]interface{}{
				"success": false,
				"message": "forbidden",
			})
		}
	}
	jsonBody := make(map[string]interface{})
	err := json.NewDecoder(ctx.Request().Body).Decode(&jsonBody)
	if err != nil {
		h.logger.Error("handlers", "ChangeComment", err)
		return ctx.JSON(500, map[string]interface{}{
			"success": false,
			"message": "cant parse json",
		})
	}
	newDate := time.Now().Format("YYYY-MM-DD H24:MI")
	commentId, _ := strconv.Atoi(jsonBody["commentId"].(string))
	text := jsonBody["text"].(string)
	err = h.usecase.ChangeComment(commentId, text)
	if err != nil {
		h.logger.Error("handlers", "ChangeComment", err)
		return ctx.JSON(500, map[string]interface{}{
			"success": false,
			"message": "server error",
		})
	}
	return ctx.JSON(200, map[string]interface{}{
		"success": true,
		"message": "",
		"date":    newDate,
	})
}
