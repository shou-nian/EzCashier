package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/shou-nian/EzCashier/pkg/redis"
	"net/http"
	"os"
	"strconv"
)

// JWTMiddleware returns a Gin middleware function that validates JWT tokens.
//
// This middleware checks for the presence of a JWT token in the "Authorization" header,
// validates it using the JWT_SECRET_KEY environment variable, and sets the parsed token
// in the Gin context if valid. If the token is missing or invalid, it aborts the request
// with a 401 Unauthorized status.
//
// Returns:
//
//	gin.HandlerFunc: A middleware function that can be used in a Gin router.
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("jwt", token)

		c.Next()
	}
}

// PrivateAuthorizationMiddleware returns a Gin middleware function that performs
// additional authorization checks for private routes.
//
// This middleware assumes that a JWT token has already been validated and set in the
// Gin context. It performs the following checks:
//  1. Verifies that the JWT token in the request matches the one cached in Redis.
//  2. For specific routes (e.g., user management), it checks for admin privileges.
//
// If any of these checks fail, the middleware will abort the request with an
// appropriate HTTP status code (401 Unauthorized or 403 Forbidden).
//
// Returns:
//
//	gin.HandlerFunc: A middleware function that can be used in a Gin router to
//	                 enforce private route authorization.
func PrivateAuthorizationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get data from JWT
		token := c.Value("jwt").(*jwt.Token)
		claims := token.Claims.(jwt.MapClaims)

		// Check redis cached jwt equal to request jwt
		rds, err := redis.OpenRedisConnection()
		if err != nil {
			panic(err)
		}
		cachedToken, err := rds.Get(strconv.Itoa(int(claims["id"].(float64))))
		// If not cached token (expiration) or request jwt token not equal to cached token
		// Return Unauthorized code
		if err != nil || cachedToken != token.Raw {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		switch c.Request.URL.Path {
		case "/api/v1/admin":
			// Only with `admin` credential can create & delete user & update user role.
			if !claims["admin"].(bool) {
				c.AbortWithStatus(http.StatusForbidden)
				return
			}
		case "api/v1/user":
			if !claims["user"].(bool) || !claims["admin"].(bool) {
				c.AbortWithStatus(http.StatusForbidden)
				return
			}
		}

		// If the user is authorized, continue to the next middleware or route handler
		c.Next()
	}
}
