package handlers

import (
	"github.com/labstack/echo/v4"
	"strconv"
)

func (h *handler) GetComments(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.Redirect(301, "/articles")
	}

	comments, err := h.usecase.GetArticleComments(id)
	if err != nil {
		return ctx.Redirect(301, "/articles")
	}

	article, err := h.usecase.GetArticleById(id)
	if err != nil {
		return ctx.Redirect(301, "/articles")
	}

	return ctx.Render(500, "article.html", map[string]interface{}{
		"article":  article,
		"comments": comments,
	})
}

func (h *handler) DeleteComment(ctx echo.Context) error {
	userErr := h.auth(ctx, "user")
	authorErr := h.auth(ctx, "author")
	if authorErr != nil {
		if userErr != nil {
			return ctx.JSON(301, map[string]interface{}{
				"success": false,
				"message": "Permission denied",
			})
		}
	}

	commentId, err := strconv.Atoi(ctx.Param("comment_id"))
	if err != nil {
		return ctx.JSON(301, map[string]interface{}{
			"success": false,
			"message": "invalid argument",
		})
	}
	comment, err := h.usecase.GetCommentById(commentId)
	if err != nil {
		return ctx.JSON(301, map[string]interface{}{
			"success": false,
			"message": "there are no such comment",
		})
	}
	name := h.authHeader(ctx)
	if name != comment.AuthorName {
		return ctx.JSON(301, map[string]interface{}{
			"success": false,
			"message": "Permission denied",
		})
	}
	err = h.usecase.DeleteComment(commentId)
	if err != nil {
		return ctx.JSON(301, map[string]interface{}{
			"success": false,
			"message": "server error",
		})
	}

	return ctx.JSON(200, map[string]interface{}{
		"success": true,
		"message": "success",
	})
}
