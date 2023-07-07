package server

import (
	"fmt"

	"github.com/Rosto4eks/eclipse/internal/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

type Server struct {
	router  *echo.Echo
	logger  *logrus.Logger
	handler handlers.Ihandler
}

func New(handler handlers.Ihandler) *Server {
	return &Server{
		router:  echo.New(),
		logger:  logrus.New(),
		handler: handler,
	}
}

func (s *Server) SetupLogger(cfg *Config) error {
	level, err := logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)
	return nil
}

func (s *Server) init(cfg *Config) error {
	if err := s.SetupLogger(cfg); err != nil {
		return err
	}
	s.initRoutes(s.handler)
	return nil
}

func (s *Server) Start() error {
	s.router.Use(middleware.CORS())
	cfg, err := NewConfig()
	if err != nil {
		return err
	}
	if err = s.init(cfg); err != nil {
		return err
	}
	addr := fmt.Sprintf("%s:%d", cfg.Address, cfg.Port)
	return s.router.Start(addr)
}
