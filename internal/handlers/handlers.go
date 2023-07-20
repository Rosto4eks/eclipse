package handlers

import (
	"github.com/Rosto4eks/eclipse/internal/usecase"
	"github.com/labstack/echo/v4"
)

type Ihandler interface {
	GetHome(echo.Context) error
	GetAlbums(echo.Context) error
	GetAlbum(echo.Context) error
	GetNewAlbum(echo.Context) error
	PostNewAlbum(echo.Context) error
	DeleteAlbum(ctx echo.Context) error
	GetSignIn(echo.Context) error
	GetSignUp(echo.Context) error
	GetLogOut(ctx echo.Context) error
	PostSignUp(echo.Context) error
	PostSignIn(ctx echo.Context) error
	GetArticle(ctx echo.Context) error
	GetArticles(ctx echo.Context) error
	GetNewArticle(ctx echo.Context) error
	PostNewArticle(ctx echo.Context) error
	DeleteArticle(ctx echo.Context) error
	GetComments(ctx echo.Context) error
	DeleteComment(ctx echo.Context) error
	PostNewComment(ctx echo.Context) error
	ChangeArticle(ctx echo.Context) error
	auth(echo.Context, string) error
	writeJWT(token string, ctx echo.Context)
	readJWT(echo.Context) (string, error)
	cleanJWT(ctx echo.Context) error
}

// first layer, handles incoming http requests
type handler struct {
	usecase usecase.Iusecase
}

func New(usecase usecase.Iusecase) *handler {
	return &handler{
		usecase,
	}
}

func (h *handler) GetHome(ctx echo.Context) error {
	headerName := h.authHeader(ctx)
	welcome := ctx.QueryParam("welcome")
	if welcome == "true" {
		welcome = "welcome back, "
	}
	return ctx.Render(200, "home.html", map[string]interface{}{
		"header":  headerName,
		"welcome": welcome,
	})
}
