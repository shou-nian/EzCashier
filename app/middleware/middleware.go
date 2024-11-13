package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"os"
	"time"
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

func PrivateAuthorizationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check if the user is authorized to access the requested resource
		// If not, return a 401 Unauthorized status code
		// Get now time
		now := time.Now().Unix()

		// Get data from JWT
		token := c.Value("jwt").(*jwt.Token)
		claims := token.Claims.(jwt.MapClaims)

		// Set expiration time from JWT data of current user.
		expires := claims["expires"].(int64)

		// Set credential `admin` from JWT data of current user.
		credential := claims["admin"].(bool)

		// Only user with `admin` credential can create a new user profile.
		if !credential || now >= expires {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// If the user is authorized, continue to the next middleware or route handler
		c.Next()
	}
}
