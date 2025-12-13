package main

import (
	"log"
	"medical-records-app/internal/auth"
	"medical-records-app/internal/config"
	"medical-records-app/internal/database"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Load configuration
	cfg := config.Load()

	// Initialize database
	db, err := database.Initialize(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Run migrations
	if err := database.RunMigrations(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Get default email to check
	defaultEmail := getEnv("DEFAULT_USER_EMAIL", "admin@medicalrecords.com")
	
	// Check if default user already exists
	var existingUser database.User
	if err := db.Where("email = ?", defaultEmail).First(&existingUser).Error; err == nil {
		log.Println("Default user already exists. Skipping seed.")
		return
	}

	// Create default admin user
	defaultPassword := getEnv("DEFAULT_USER_PASSWORD", "admin123")
	defaultFirstName := getEnv("DEFAULT_USER_FIRSTNAME", "Admin")
	defaultLastName := getEnv("DEFAULT_USER_LASTNAME", "User")

	hashedPassword, err := auth.HashPassword(defaultPassword)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	user := &database.User{
		ID:             uuid.New(),
		Email:          defaultEmail,
		PasswordHash:   hashedPassword,
		FirstName:      defaultFirstName,
		LastName:       defaultLastName,
		Role:           "patient",
		IsEmailVerified: true,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	if err := db.Create(user).Error; err != nil {
		log.Fatalf("Failed to create default user: %v", err)
	}

	log.Println("=")
	log.Println("Default user created successfully!")
	log.Println("=")
	log.Printf("Email: %s", defaultEmail)
	log.Printf("Password: %s", defaultPassword)
	log.Println("=")
	log.Println("⚠️  IMPORTANT: Change this password after first login!")
	log.Println("=")
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

