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

	logConfig := logger.Config{
		SlowThreshold:             time.Second,
		LogLevel:                  logger.Warn,
		IgnoreRecordNotFoundError: true,
		Colorful:                  false,
	}

	db, err := gorm.Open(dial, &gorm.Config{
		Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logConfig),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database instance: %w", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	if err := db.AutoMigrate(&models.User{}, &models.Address{}, &models.Post{}); err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	if err := seedDB(db); err != nil {
		return nil, fmt.Errorf("failed to seed database: %w", err)
	}

	return db, nil
}
