package config

import (
	"fmt"
	"net/url"
	"os"
	"strings"
)

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
	JWT      JWTConfig
	AWS      AWSConfig
	SMTP     SMTPConfig
	SMS      SMSConfig
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

type ServerConfig struct {
	Port string
	Host string
	Env  string
}

type JWTConfig struct {
	Secret         string
	ExpirationHours int
}

type AWSConfig struct {
	AccessKeyID     string
	SecretAccessKey string
	Region          string
	S3Bucket        string
}

type SMTPConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	From     string
}

type SMSConfig struct {
	Provider        string
	TwilioAccountSID string
	TwilioAuthToken  string
	TwilioPhoneNumber string
}

func Load() *Config {
	// Check if DATABASE_URL is provided (Render sometimes uses this)
	databaseURL := os.Getenv("DATABASE_URL")
	
	var dbConfig DatabaseConfig
	if databaseURL != "" {
		// Parse DATABASE_URL format: postgres://user:password@host:port/dbname?sslmode=require
		dbConfig = parseDatabaseURL(databaseURL)
	} else {
		// Use individual environment variables
		dbConfig = DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "medical_user"),
			Password: getEnv("DB_PASSWORD", "medical_password"),
			Name:     getEnv("DB_NAME", "medical_records"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		}
	}
	
	return &Config{
		Database: dbConfig,
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "8080"),
			Host: getEnv("SERVER_HOST", "localhost"),
			Env:  getEnv("APP_ENV", "development"),
		},
		JWT: JWTConfig{
			Secret:          getEnv("JWT_SECRET", "change-me-in-production"),
			ExpirationHours: getEnvAsInt("JWT_EXPIRATION_HOURS", 24),
		},
		AWS: AWSConfig{
			AccessKeyID:     getEnv("AWS_ACCESS_KEY_ID", ""),
			SecretAccessKey: getEnv("AWS_SECRET_ACCESS_KEY", ""),
			Region:          getEnv("AWS_REGION", "us-east-1"),
			S3Bucket:        getEnv("S3_BUCKET", ""),
		},
		SMTP: SMTPConfig{
			Host:     getEnv("SMTP_HOST", ""),
			Port:     getEnv("SMTP_PORT", "587"),
			User:     getEnv("SMTP_USER", ""),
			Password: getEnv("SMTP_PASSWORD", ""),
			From:     getEnv("SMTP_FROM", "Medical Records App <noreply@medicalrecords.com>"),
		},
		SMS: SMSConfig{
			Provider:        getEnv("SMS_PROVIDER", "twilio"),
			TwilioAccountSID: getEnv("TWILIO_ACCOUNT_SID", ""),
			TwilioAuthToken:  getEnv("TWILIO_AUTH_TOKEN", ""),
			TwilioPhoneNumber: getEnv("TWILIO_PHONE_NUMBER", ""),
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	// Simple conversion - in production, use strconv.Atoi with error handling
	return defaultValue
}

// parseDatabaseURL parses a PostgreSQL connection URL
// Format: postgres://user:password@host:port/dbname?sslmode=require
func parseDatabaseURL(databaseURL string) DatabaseConfig {
	config := DatabaseConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "medical_user",
		Password: "medical_password",
		Name:     "medical_records",
		SSLMode:  "require",
	}

	// Parse the URL
	parsedURL, err := url.Parse(databaseURL)
	if err != nil {
		return config
	}

	// Extract host and port
	config.Host = parsedURL.Hostname()
	if parsedURL.Port() != "" {
		config.Port = parsedURL.Port()
	}

	// Extract user and password
	if parsedURL.User != nil {
		config.User = parsedURL.User.Username()
		if password, ok := parsedURL.User.Password(); ok {
			config.Password = password
		}
	}

	// Extract database name (remove leading slash)
	config.Name = strings.TrimPrefix(parsedURL.Path, "/")

	// Extract SSL mode from query parameters
	if parsedURL.Query().Get("sslmode") != "" {
		config.SSLMode = parsedURL.Query().Get("sslmode")
	}

	return config
}

