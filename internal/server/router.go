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
		r.GET("/", h.Home)
		r.GET("/albums", h.Albums)
		r.GET("/albums/:id", h.Album)
		r.GET("/albums/new", h.NewAlbum)
		r.GET("/albums/:id/delete", h.DeleteAlbum)
		r.POST("/albums/new", h.CreateNewAlbum)
	}
	auth := s.router.Group("auth")
	{
		auth.GET("/sign-in", h.SignIn)
		auth.GET("/sign-up", h.SignUp)
		auth.POST("/sign-up", h.NewUser)
		auth.POST("/sign-in", h.PostSignIn)
	}
}
