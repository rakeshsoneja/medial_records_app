package main

import (
	"log"
	"medical-records-app/internal/config"
	"medical-records-app/internal/database"
	"medical-records-app/internal/router"
	"os"

	"github.com/joho/godotenv"
)

// @title Medical Records Management API
// @version 1.0
// @description Secure API for managing medical records, prescriptions, appointments, and lab reports
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email support@medicalrecords.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Load configuration
	cfg := config.Load()

	// Initialize database (will create tables automatically via migrations)
	db, err := database.Initialize(cfg)
	if err != nil {
		log.Printf("Database connection failed: %v", err)
		log.Println("")
		log.Println("If this is the first time running, you may need to set up the database.")
		log.Println("Run: go run cmd/setup/main.go")
		log.Println("Or set up PostgreSQL manually with:")
		log.Printf("  CREATE DATABASE %s;", cfg.Database.Name)
		log.Printf("  CREATE USER %s WITH PASSWORD '%s';", cfg.Database.User, cfg.Database.Password)
		log.Printf("  GRANT ALL PRIVILEGES ON DATABASE %s TO %s;", cfg.Database.Name, cfg.Database.User)
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Run migrations
	if err := database.RunMigrations(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Initialize router
	r := router.Initialize(db, cfg)

	// Start server
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

