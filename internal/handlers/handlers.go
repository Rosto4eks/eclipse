package handlers

import (
	"github.com/Rosto4eks/eclipse/internal/logger"
	"github.com/Rosto4eks/eclipse/internal/usecase"
	"github.com/labstack/echo/v4"
)

type Ihandler interface {
	GetHome(echo.Context) error
	GetAlbums(echo.Context) error
	GetAlbum(echo.Context) error
	GetNewAlbum(echo.Context) error
	PostNewAlbum(echo.Context) error
	LoadAlbums(ctx echo.Context) error
	DeleteAlbum(ctx echo.Context) error
	GetSignIn(echo.Context) error
	GetSignUp(echo.Context) error
	GetLogOut(ctx echo.Context) error
	PostSignUp(echo.Context) error
	PostSignIn(ctx echo.Context) error
	GetArticle(ctx echo.Context) error
	GetArticles(ctx echo.Context) error
	LoadArticles(ctx echo.Context) error
	GetNewArticle(ctx echo.Context) error
	PostNewArticle(ctx echo.Context) error
	SearchArticles(ctx echo.Context) error
	DeleteArticle(ctx echo.Context) error
	GetComments(ctx echo.Context) error
	DeleteComment(ctx echo.Context) error
	PostNewComment(ctx echo.Context) error
	ChangeComment(ctx echo.Context) error
	ChangeArticle(ctx echo.Context) error
	auth(echo.Context, string) error
	writeJWT(token string, ctx echo.Context)
	readJWT(echo.Context) (string, error)
	cleanJWT(ctx echo.Context) error
}

// first layer, handles incoming http requests
type handler struct {
	usecase usecase.Iusecase
	logger  logger.Ilogger
}

func New(usecase usecase.Iusecase, logger logger.Ilogger) *handler {
	return &handler{
		usecase,
		logger,
	}
}

func (h *handler) GetHome(ctx echo.Context) error {
	headerName := h.authHeader(ctx)
	welcome := ctx.QueryParam("welcome")
	if welcome == "true" {
		welcome = "welcome back, "
	}
	articles, err := h.usecase.GetArticles(0, 3)
	if err != nil {
		h.logger.Error("handlers", "GetHome", err)
		return ctx.Render(500, "home.html", map[string]interface{}{
			"header":  headerName,
			"welcome": welcome,
		})
	}
	albums, err := h.usecase.GetAlbums(0, 3)
	if err != nil {
		h.logger.Error("handlers", "GetHome", err)
		return ctx.Render(500, "home.html", map[string]interface{}{
			"header":  headerName,
			"welcome": welcome,
		})
	}
	return ctx.Render(200, "home.html", map[string]interface{}{
		"header":   headerName,
		"welcome":  welcome,
		"articles": articles,
		"albums":   albums,
	})
}
