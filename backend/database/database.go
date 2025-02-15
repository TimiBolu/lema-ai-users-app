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
		LogLevel:                  logger.Info,
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

	if err := db.Exec("PRAGMA foreign_keys = ON").Error; err != nil {
		return nil, fmt.Errorf("failed to enable foreign keys: %w", err)
	}

	if err := performSafeMigration(db); err != nil {
		return nil, fmt.Errorf("failed to perform migration: %w", err)
	}

	if err := seedDB(db); err != nil {
		return nil, fmt.Errorf("failed to seed database: %w", err)
	}

	return db, nil
}

func performSafeMigration(db *gorm.DB) error {
	// Disable foreign keys during migration
	if err := db.Exec("PRAGMA foreign_keys = OFF").Error; err != nil {
		return err
	}

	tables := []struct {
		name    string
		model   interface{}
		columns []string
	}{
		{
			name:    "users",
			model:   &models.User{},
			columns: []string{"id", "name", "username", "email", "phone"},
		},
		{
			name:    "addresses",
			model:   &models.Address{},
			columns: []string{"id", "user_id", "street", "state", "city", "zipcode"},
		},
		{
			name:    "posts",
			model:   &models.Post{},
			columns: []string{"id", "user_id", "title", "body", "created_at"},
		},
	}

	for _, tbl := range tables {
		if err := migrateTable(db, tbl.name, tbl.model, tbl.columns); err != nil {
			return fmt.Errorf("failed to migrate %s: %w", tbl.name, err)
		}
	}

	// Re-enable foreign keys
	if err := db.Exec("PRAGMA foreign_keys = ON").Error; err != nil {
		return err
	}

	return nil
}

func migrateTable(db *gorm.DB, tableName string, model interface{}, columns []string) error {
	var tableExists bool
	db.Raw(fmt.Sprintf("SELECT COUNT(*) FROM sqlite_master WHERE type='table' AND name='%s'", tableName)).Scan(&tableExists)

	if !tableExists {
		return db.AutoMigrate(model)
	}

	tempTable := tableName + "_old"
	if err := db.Exec(fmt.Sprintf("ALTER TABLE %s RENAME TO %s", tableName, tempTable)).Error; err != nil {
		return err
	}

	if err := db.AutoMigrate(model); err != nil {
		return err
	}

	columnsStr := ""
	for _, col := range columns {
		columnsStr += fmt.Sprintf("COALESCE(%s, '') AS %s,", col, col)
	}
	columnsStr = columnsStr[:len(columnsStr)-1]

	insertSQL := fmt.Sprintf(`
		INSERT INTO %s (%s)
		SELECT %s
		FROM %s
	`, tableName, joinColumns(columns), columnsStr, tempTable)

	if err := db.Exec(insertSQL).Error; err != nil {
		return fmt.Errorf("failed to copy data to %s: %w", tableName, err)
	}

	if err := db.Exec(fmt.Sprintf("DROP TABLE %s", tempTable)).Error; err != nil {
		return err
	}

	return nil
}

func joinColumns(cols []string) string {
	result := ""
	for _, c := range cols {
		result += fmt.Sprintf("`%s`,", c)
	}
	return result[:len(result)-1]
}
