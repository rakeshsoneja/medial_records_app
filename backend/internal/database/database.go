package database

import (
	"fmt"
	"medical-records-app/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Initialize(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.Database.Host,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.Port,
		cfg.Database.SSLMode,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return DB, nil
}

func RunMigrations(db *gorm.DB) error {
	// Auto-migrate all models
	err := db.AutoMigrate(
		&User{},
		&HealthInsurance{},
		&Prescription{},
		&Appointment{},
		&LabReport{},
		&Medication{},
		&Reminder{},
		&SharedRecord{},
		&AuditLog{},
	)

	if err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	return nil
}

