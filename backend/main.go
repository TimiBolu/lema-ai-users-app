package main

import (
	"github.com/TimiBolu/lema-ai-users-service/config"
	"github.com/TimiBolu/lema-ai-users-service/database"
	"github.com/TimiBolu/lema-ai-users-service/router"
)

func main() {
	logger := config.InitEnvSchema()

	db, err := database.Connect()
	if err != nil {
		logger.Fatalf("❌ Failed to connect to the database: %v", err)
	}
	logger.Info("✅ Database connection established")

	router.Setup(db, logger)
}
