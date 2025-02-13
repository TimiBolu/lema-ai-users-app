package main

import (
	"fmt"

	"github.com/TimiBolu/lema-ai-users-service/config"
	"github.com/TimiBolu/lema-ai-users-service/database"
	"github.com/TimiBolu/lema-ai-users-service/router"
	"github.com/TimiBolu/lema-ai-users-service/services"
)

func main() {
	// Initialize logger
	// Initialize the server environment
	logger := config.InitEnvSchema()

	// Connect to the database
	db, err := database.Connect()
	if err != nil {
		logger.Fatalf("❌ Failed to connect to the database: %v", err)
	}
	logger.Info("✅ Database connection established")

	fmt.Println(services.IssueToken("timi"))
	// Start the router
	router.Setup(db, logger)
}
