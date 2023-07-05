package main

import (
	"github.com/Rosto4eks/eclipse/internal/database"
	"github.com/Rosto4eks/eclipse/internal/handlers"
	"github.com/Rosto4eks/eclipse/internal/server"
	"github.com/Rosto4eks/eclipse/internal/usecase"
	"github.com/sirupsen/logrus"
)

// startup function
func main() {
	database := database.New()
	usecase := usecase.New(database)
	handler := handlers.New(usecase)

	server := server.New(handler)
	if err := server.Start(); err != nil {
		logrus.Fatal(err)
	}
}
