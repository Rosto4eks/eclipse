package main

import (
	"github.com/Rosto4eks/eclipse/internal/handlers"
	"github.com/Rosto4eks/eclipse/internal/server"

	"github.com/sirupsen/logrus"
)

// startup function
func main() {
	handler := handlers.New()
	server := server.New(handler)
	if err := server.Start(); err != nil {
		logrus.Fatal(err)
	}
}
