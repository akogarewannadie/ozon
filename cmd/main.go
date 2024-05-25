package main

import (
	"github.com/akogarewannadie/ozon/internal/server"
	"log"
)

func main() {
	if err := server.Run(); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
