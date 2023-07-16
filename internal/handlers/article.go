package handlers

import "github.com/labstack/echo/v4"

func (h *handler) GetArticles(ctx echo.Context) error {
	return ctx.Render(301, "article.html", nil)
}
