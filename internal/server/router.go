package server

import (
	"html/template"
	"io"

	"github.com/Rosto4eks/eclipse/internal/handlers"
	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func (s *Server) initRoutes(h handlers.Ihandler) {
	// creating renderer for html templates
	s.router.Renderer = &Template{
		templates: template.Must(template.ParseGlob("./public/pages/*.html")),
	}

	// serving static files
	s.router.Static("/public", "./public/")

	r := s.router.Group("")
	{
		r.GET("/", h.GetHome)
		r.GET("/albums", h.GetAlbums)
		r.GET("/albums/:id", h.GetAlbum)
		r.GET("/albums/new", h.GetNewAlbum)
		r.GET("/articles", h.GetArticles)
		r.GET("/articles/new", h.GetNewArticle)
		r.GET("/articles/:id", h.GetArticle)
		r.POST("/albums/new", h.PostNewAlbum)
		r.POST("/articles/new", h.PostNewArticle)
		r.DELETE("/albums/:id/delete", h.DeleteAlbum)
		r.DELETE("/articles/:id/delete-comment/:comment_id", h.DeleteComment)
	}
	auth := s.router.Group("auth")
	{
		auth.GET("/sign-in", h.GetSignIn)
		auth.GET("/sign-up", h.GetSignUp)
		auth.POST("/sign-up", h.PostSignUp)
		auth.POST("/sign-in", h.PostSignIn)
	}
}
