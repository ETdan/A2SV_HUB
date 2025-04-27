package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"A2SVHUB/api/routes"
	"A2SVHUB/pkg/config"
	"A2SVHUB/pkg/database"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration

	cfg := config.LoadConfig()

	// Initialize database
	db := database.ConnectDB()

	// Initialize Gin router
	router := gin.Default()
	router.Use(cors.Default())

	// Serve a static JSON file
	router.GET("/", func(c *gin.Context) {
		// Get the current working directory
		baseDir, err := os.Getwd()
		if err != nil {
			log.Fatalf("Failed to get working directory: %v", err)
		}

		// Construct the correct relative path to the Swagger JSON file
		filePath := filepath.Join(baseDir, "..", "docs", "A2SV_HUB_SWAGGER.json")
		fmt.Printf("Serving Swagger JSON file from: %s\n", filePath)

		// Serve the file
		c.File(filePath)
	})

	// Setup routes
	routes.SetupRoutes(router.Group("/api/v0"), &cfg, db)

	// Start the HTTP server
	log.Printf("Server is running on port %s\n", cfg.AppPort)

	// schema update
	// if err := migration.MigrateModels(db); err != nil {
	// 	log.Fatalf("Migration failed: %v", err)
	// }

	if err := router.Run(cfg.AppPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
