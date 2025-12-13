package main

import (
	"log"
	"medical-records-app/internal/config"
	"medical-records-app/internal/database"
	"medical-records-app/internal/router"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
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

	// Log database configuration (without password) for debugging
	log.Printf("Database config - Host: %s, Port: %s, User: %s, DB: %s, SSLMode: %s",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Name, cfg.Database.SSLMode)

	// Initialize database (will create tables automatically via migrations)
	var db *gorm.DB
	var err error
	
	db, err = database.Initialize(cfg)
	if err != nil {
		log.Printf("Database connection failed: %v", err)
		log.Println("")
		log.Println("If this is the first time running, you may need to set up the database.")
		log.Println("Run: go run cmd/setup/main.go")
		log.Println("Or set up PostgreSQL manually with:")
		log.Printf("  CREATE DATABASE %s;", cfg.Database.Name)
		log.Printf("  CREATE USER %s WITH PASSWORD '%s';", cfg.Database.User, cfg.Database.Password)
		log.Printf("  GRANT ALL PRIVILEGES ON DATABASE %s TO %s;", cfg.Database.Name, cfg.Database.User)
		log.Println("")
		log.Println("⚠️  WARNING: Server will start but database-dependent endpoints will fail!")
		log.Println("⚠️  Health endpoint will still work.")
		db = nil // Set to nil so router knows database is not available
	} else {
		// Run migrations only if database is connected
		if err := database.RunMigrations(db); err != nil {
			log.Printf("Failed to run migrations: %v", err)
			log.Println("⚠️  WARNING: Migrations failed, but server will continue")
		}
	}

	// Initialize router (pass nil db if connection failed - health endpoint will still work)
	r := router.Initialize(db, cfg)

	// Start server
	// Render provides PORT environment variable, fallback to SERVER_PORT or 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = os.Getenv("SERVER_PORT")
	}
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

