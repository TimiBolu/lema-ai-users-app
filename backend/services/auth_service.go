package services

import (
	"time"

	"github.com/TimiBolu/lema-ai-users-service/config"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sirupsen/logrus"
)

// merely serves to issue tokens

// Custom claims structure
type CustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// IssueToken generates a new JWT token with the given claims
func IssueToken(username string) (string, error) {
	var jwtSecret = []byte(config.EnvConfig.JWT_SECRET) // Replace with your actual secret key
	// Set custom claims
	claims := CustomClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(), // Token expires in 24 hours
			IssuedAt:  time.Now().Unix(),
			Issuer:    "your-issuer",
		},
	}

	// Create a new token with the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		logrus.WithError(err).Error("Failed to sign token")
		return "", err
	}

	return tokenString, nil
}
