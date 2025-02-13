package router

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/TimiBolu/lema-ai-users-service/config"
	"github.com/TimiBolu/lema-ai-users-service/handlers"
	"github.com/TimiBolu/lema-ai-users-service/repositories"
	"github.com/TimiBolu/lema-ai-users-service/services"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func getAPIDocs(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("docs/api.md")
	if err != nil {
		http.Error(w, "Failed to load documentation", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// Health check handler
func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func Setup(db *gorm.DB, logger *logrus.Logger) {
	// Initialize the router
	r := mux.NewRouter()

	// Add logging middleware
	r.Use(loggingMiddleware)
	r.Use(authMiddleware)
	r.Use(rateLimiterMiddleware)

	// Health check route
	r.HandleFunc("/api/health-check", healthCheck).Methods("GET")

	// Redirect root (/) to /api/health-check
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/api/health-check", http.StatusFound)
	})

	// Initialize services
	postService := services.NewPostService(repositories.NewPostRepository(db))
	userService := services.NewUserService(repositories.NewUserRepository(db))

	// Initialize handlers
	postHandler := handlers.NewPostHandler(postService, logger)
	userHandler := handlers.NewUserHandler(userService, logger)

	// User endpoints
	r.HandleFunc("/api/users", userHandler.GetUsers).Methods("GET")
	r.HandleFunc("/api/users/count", userHandler.GetUsersCount).Methods("GET")
	r.HandleFunc("/api/users/{id}", userHandler.GetUserByID).Methods("GET")

	// Post endpoints
	r.HandleFunc("/api/posts", postHandler.GetPostsByUser).Methods("GET")
	r.HandleFunc("/api/posts", postHandler.CreatePost).Methods("POST")
	r.HandleFunc("/api/posts/{id}", postHandler.DeletePost).Methods("DELETE")

	// API Documentation endpoints
	r.HandleFunc("/api/docs", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "docs/index.html")
	})
	r.HandleFunc("/api/docs/raw", getAPIDocs) // Serve raw Markdown

	corsHandler := corsMiddleware(r)
	// Server configuration
	port := config.EnvConfig.PORT
	baseURL := config.EnvConfig.SERVER_BASE_URL
	log.Printf("🚀 Server is up and running on %s:%s/api", baseURL, port)
	log.Printf("📄 API Documentation available at %s:%s/api/docs", baseURL, port)

	// Start the server
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), corsHandler)
	if err != nil {
		log.Fatalf("❌ Server failed to start: %v", err)
	}
}
