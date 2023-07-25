package server

import (
	"fmt"

	"github.com/Rosto4eks/eclipse/internal/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	router  *echo.Echo
	handler handlers.Ihandler
}

func New(handler handlers.Ihandler) *Server {
	return &Server{
		router:  echo.New(),
		handler: handler,
	}
}

func (s *Server) Start(cfg *Config) error {
	s.router.Use(middleware.CORS())
	s.initRoutes(s.handler)

	addr := fmt.Sprintf("%s:%d", cfg.Address, cfg.Port)
	return s.router.Start(addr)
}
