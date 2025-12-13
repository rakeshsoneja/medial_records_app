package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Get database configuration
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "medical_user")
	dbPassword := getEnv("DB_PASSWORD", "medical_password")
	dbName := getEnv("DB_NAME", "medical_records")
	sslMode := getEnv("DB_SSLMODE", "disable")

	// Try to get postgres superuser password (optional)
	postgresPassword := getEnv("POSTGRES_PASSWORD", dbPassword)
	
	// Connect to PostgreSQL server (without specifying database)
	// Try with postgres user first, then fall back to specified user
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=%s",
		dbHost, dbPort, "postgres", postgresPassword, sslMode)

	// Try connecting as postgres user first, if that fails, try the specified user
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		// Try with the specified user
		connStr = fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=%s",
			dbHost, dbPort, dbUser, dbPassword, sslMode)
		db, err = sql.Open("postgres", connStr)
		if err != nil {
			log.Fatalf("Failed to connect to PostgreSQL: %v", err)
		}
	}
	defer db.Close()

	// Test connection
	if err = db.Ping(); err != nil {
		log.Fatalf("Failed to ping PostgreSQL: %v", err)
	}

	log.Println("Connected to PostgreSQL server")

	// Check if database exists
	var exists bool
	checkDBQuery := `SELECT EXISTS(SELECT datname FROM pg_catalog.pg_database WHERE datname = $1)`
	err = db.QueryRow(checkDBQuery, dbName).Scan(&exists)
	if err != nil {
		log.Fatalf("Failed to check if database exists: %v", err)
	}

	if !exists {
		// Create database
		// Note: CREATE DATABASE cannot be run in a transaction
		createDBQuery := fmt.Sprintf("CREATE DATABASE %s", dbName)
		_, err = db.Exec(createDBQuery)
		if err != nil {
			log.Fatalf("Failed to create database: %v", err)
		}
		log.Printf("Database '%s' created successfully", dbName)
	} else {
		log.Printf("Database '%s' already exists", dbName)
	}

	// Now connect to the specific database to create user and grant privileges
	connStrWithDB := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, "postgres", dbPassword, dbName, sslMode)
	
	dbWithDB, err := sql.Open("postgres", connStrWithDB)
	if err != nil {
		// If postgres user doesn't work, try without specifying user (use default)
		connStrWithDB = fmt.Sprintf("host=%s port=%s dbname=%s sslmode=%s",
			dbHost, dbPort, dbName, sslMode)
		dbWithDB, err = sql.Open("postgres", connStrWithDB)
		if err != nil {
			log.Printf("Warning: Could not connect to database to set up user. You may need to create user manually.")
			log.Printf("Run these SQL commands:")
			log.Printf("  CREATE USER %s WITH PASSWORD '%s';", dbUser, dbPassword)
			log.Printf("  GRANT ALL PRIVILEGES ON DATABASE %s TO %s;", dbName, dbUser)
			return
		}
	}
	defer dbWithDB.Close()

	// Check if user exists
	var userExists bool
	checkUserQuery := `SELECT EXISTS(SELECT 1 FROM pg_roles WHERE rolname = $1)`
	err = db.QueryRow(checkUserQuery, dbUser).Scan(&userExists)
	if err == nil && !userExists {
		// Try to create user (requires superuser privileges)
		createUserQuery := fmt.Sprintf("CREATE USER %s WITH PASSWORD '%s'", dbUser, dbPassword)
		_, err = db.Exec(createUserQuery)
		if err != nil {
			log.Printf("Warning: Could not create user '%s'. You may need to create it manually.", dbUser)
			log.Printf("Error: %v", err)
			log.Printf("Run this SQL command:")
			log.Printf("  CREATE USER %s WITH PASSWORD '%s';", dbUser, dbPassword)
		} else {
			log.Printf("User '%s' created successfully", dbUser)
		}
	} else if userExists {
		log.Printf("User '%s' already exists", dbUser)
	}

	// Grant privileges
	grantQuery := fmt.Sprintf("GRANT ALL PRIVILEGES ON DATABASE %s TO %s", dbName, dbUser)
	_, err = db.Exec(grantQuery)
	if err != nil {
		log.Printf("Warning: Could not grant privileges. You may need to do it manually.")
		log.Printf("Error: %v", err)
		log.Printf("Run this SQL command:")
		log.Printf("  GRANT ALL PRIVILEGES ON DATABASE %s TO %s;", dbName, dbUser)
	} else {
		log.Printf("Privileges granted to user '%s'", dbUser)
	}

	// Grant schema privileges (for PostgreSQL 15+)
	if dbWithDB.Ping() == nil {
		schemaGrantQuery := "GRANT ALL ON SCHEMA public TO " + dbUser
		_, err = dbWithDB.Exec(schemaGrantQuery)
		if err == nil {
			log.Printf("Schema privileges granted to user '%s'", dbUser)
		}

		alterPrivilegesQuery := "ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL ON TABLES TO " + dbUser
		_, err = dbWithDB.Exec(alterPrivilegesQuery)
		if err == nil {
			log.Printf("Default privileges set for user '%s'", dbUser)
		}
	}

	log.Println("Database setup completed successfully!")
	log.Printf("You can now start the server with: go run cmd/server/main.go")
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

