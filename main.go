package main

import (
	"cms/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Setup routes
	routes.SetupRoutes(r)

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
