package router

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"

	"github.com/TimiBolu/lema-ai-users-service/config"
)

// ANSI color codes for terminal output
const (
	green  = "\033[32m"
	blue   = "\033[34m"
	yellow = "\033[33m"
	red    = "\033[31m"
	reset  = "\033[0m"
)

// ResponseWriter wrapper to capture status code
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// Logging middleware to track API calls and response times with color
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		// Wrap the ResponseWriter to capture the status code
		rw := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		// Process request
		next.ServeHTTP(rw, r)

		// Determine color for method
		var methodColor string
		switch r.Method {
		case "GET":
			methodColor = green
		case "POST":
			methodColor = blue
		case "DELETE":
			methodColor = yellow
		default:
			methodColor = reset
		}

		// Determine color for status code
		var statusColor string
		switch {
		case rw.statusCode >= 200 && rw.statusCode < 300:
			statusColor = green // ✅ Success
		case rw.statusCode >= 400 && rw.statusCode < 500:
			statusColor = yellow // ⚠️ Client Error
		case rw.statusCode >= 500:
			statusColor = red // ❌ Server Error
		default:
			statusColor = reset
		}

		// Log with colors
		duration := time.Since(startTime)
		logrus.Printf("📡 %s%s%s %s | Status: %s%d%s | ⏱️ %v",
			methodColor, r.Method, reset,
			r.URL.Path,
			statusColor, rw.statusCode, reset,
			duration,
		)
	})
}

// Authentication middleware to secure endpoints
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		if !isValidToken(tokenString) {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// Custom claims structure
type CustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func isValidToken(tokenString string) bool {
	jwtSecret := []byte(config.EnvConfig.JWT_SECRET)
	claims := CustomClaims{}

	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		logrus.WithError(err).Error("Failed to parse token")
		return false
	}

	if !token.Valid {
		logrus.Error("Token is not valid")
		return false
	}

	if claims.ExpiresAt < time.Now().Unix() {
		logrus.Error("Token has expired")
		return false
	}

	return true
}

// CORS middleware
func corsMiddleware(r *mux.Router) http.Handler {
	frontendApps := config.EnvConfig.FRONTEND_APPS
	allowedOrigins := strings.Split(frontendApps, ",")

	return cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}).Handler(r)
}

// Helper function to check if a slice contains a string
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// RateLimiter middleware to limit the number of requests per client IP
func rateLimiterMiddleware(next http.Handler) http.Handler {
	// In-memory store to keep track of request counts
	requestCounts := make(map[string]int)
	var mutex sync.Mutex

	rateLimit := config.EnvConfig.XRATE_LIMIT_MAX
	window := time.Minute

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientIP := getClientIP(r)

		mutex.Lock()
		defer mutex.Unlock()

		// Check if the client has exceeded the rate limit
		if requestCounts[clientIP] >= rateLimit {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}

		// Increment the request count for the client
		requestCounts[clientIP]++

		// Start a timer to reset the request count after the window period
		time.AfterFunc(window, func() {
			mutex.Lock()
			defer mutex.Unlock()
			delete(requestCounts, clientIP)
		})

		// Proceed to the next handler
		next.ServeHTTP(w, r)
	})
}

func getClientIP(r *http.Request) string {
	// Extract the client IP from the X-Forwarded-For header if available
	if forwarded := r.Header.Get("X-Forwarded-For"); forwarded != "" {
		return forwarded
	}
	// Fallback to the remote address
	return r.RemoteAddr
}
