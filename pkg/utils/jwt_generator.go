package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"github.com/shou-nian/EzCashier/app/models"
)

// GenerateNewJWTAccessToken func for generate a new JWT access (private) token
// with user ID and Role.
func GenerateNewJWTAccessToken(id uint, role models.UserRole) (string, error) {
	// Catch JWT secret key from .env file.
	secret := os.Getenv("JWT_SECRET_KEY")
	expires, _ := strconv.Atoi(os.Getenv("JWT_EXPIRES"))

	// Create a new JWT access token and claims.
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	// Set public claims:
	claims["id"] = id
	claims["expires"] = time.Now().Add(time.Hour * time.Duration(expires)).Unix()

	// Set private token role:
	switch role {
	case models.RoleAdmin:
		claims["admin"] = true
	case models.RoleUser:
		claims["user"] = true
	case models.RoleViewer:
		claims["viewer"] = true
	}

	// Generate token.
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		// Return error, it JWT token generation failed.
		return "", err
	}

	return t, nil
}

// GenerateNewJWTRefreshToken func for generate a new JWT refresh (public) token.
func GenerateNewJWTRefreshToken() (string, error) {
	// Create a new SHA256 hash.
	hash := sha256.New()

	// Create a new now date and time string with salt.
	refresh := os.Getenv("JWT_REFRESH_KEY") + time.Now().String()

	// See: https://pkg.go.dev/io#Writer.Write
	_, err := hash.Write([]byte(refresh))
	if err != nil {
		// Return error, it refresh token generation failed.
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}
