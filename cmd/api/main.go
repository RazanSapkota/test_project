package main

import (
	"log"

	"test_project/internal/routes"
)

func main() {
	router := routes.SetupRouter()

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
