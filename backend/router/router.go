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

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func Setup(db *gorm.DB, logger *logrus.Logger) {
	r := mux.NewRouter()

	r.Use(loggingMiddleware)
	r.Use(authMiddleware)
	r.Use(rateLimiterMiddleware)

	r.HandleFunc("/api/health-check", healthCheck).Methods("GET")
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/api/health-check", http.StatusFound)
	})

	postService := services.NewPostService(repositories.NewPostRepository(db))
	userService := services.NewUserService(repositories.NewUserRepository(db))

	postHandler := handlers.NewPostHandler(postService, logger)
	userHandler := handlers.NewUserHandler(userService, logger)

	r.HandleFunc("/api/users", userHandler.GetUsers).Methods("GET")
	r.HandleFunc("/api/users/count", userHandler.GetUsersCount).Methods("GET")
	r.HandleFunc("/api/users/{id}", userHandler.GetUserByID).Methods("GET")

	r.HandleFunc("/api/posts", postHandler.GetPostsByUser).Methods("GET")
	r.HandleFunc("/api/posts", postHandler.CreatePost).Methods("POST")
	r.HandleFunc("/api/posts/{id}", postHandler.DeletePost).Methods("DELETE")

	r.HandleFunc("/api/docs", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "docs/index.html")
	})
	r.HandleFunc("/api/docs/raw", getAPIDocs)

	corsHandler := corsMiddleware(r)
	port := config.EnvConfig.PORT
	baseURL := config.EnvConfig.SERVER_BASE_URL
	log.Printf("🚀 Server is up and running on %s:%s/api", baseURL, port)
	log.Printf("📄 API Documentation available at %s:%s/api/docs", baseURL, port)

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), corsHandler)
	if err != nil {
		log.Fatalf("❌ Server failed to start: %v", err)
	}
}
