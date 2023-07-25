package main

import (
	"fmt"

	"github.com/Rosto4eks/eclipse/internal/database"
	"github.com/Rosto4eks/eclipse/internal/handlers"
	"github.com/Rosto4eks/eclipse/internal/logger"
	"github.com/Rosto4eks/eclipse/internal/server"
	"github.com/Rosto4eks/eclipse/internal/usecase"
)

// startup function
func main() {

	cfg, err := server.NewConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	logger := logger.New(cfg.LogLevel)
	dbcfg, err := database.NewConfig()
	if err != nil {
		logger.Error("server", "main", err)
		return
	}
	conn, err := database.Connect(dbcfg)
	if err != nil {
		logger.Error("server", "main", err)
		return
	}

	database := database.New(conn, logger)
	usecase := usecase.New(database, logger)
	handler := handlers.New(usecase, logger)

	server := server.New(handler)
	if err := server.Start(cfg); err != nil {
		logger.Error("server", "main", err)
	}
}
