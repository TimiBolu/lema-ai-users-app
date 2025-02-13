package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/TimiBolu/lema-ai-users-service/config"
	"github.com/TimiBolu/lema-ai-users-service/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect() (*gorm.DB, error) {
	dbName := config.EnvConfig.DB_NAME

	dial := sqlite.Open(dbName)

	// Set GORM logger to log errors and warnings
	logConfig := logger.Config{
		SlowThreshold:             time.Second, // Log queries that take longer than 1s
		LogLevel:                  logger.Warn, // Show only warnings and errors
		IgnoreRecordNotFoundError: true,
		Colorful:                  false,
	}

	// Open connection
	db, err := gorm.Open(dial, &gorm.Config{
		Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logConfig),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Configure connection pool settings
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database instance: %w", err)
	}

	sqlDB.SetMaxIdleConns(10)                 // Max idle connections
	sqlDB.SetMaxOpenConns(100)                // Max open connections
	sqlDB.SetConnMaxLifetime(5 * time.Minute) // Connection max lifetime

	// AutoMigrate
	if err := db.AutoMigrate(&models.User{}, &models.Address{}, &models.Post{}); err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	// Seed the database
	if err := seedDB(db); err != nil {
		return nil, fmt.Errorf("failed to seed database: %w", err)
	}

	return db, nil
}
