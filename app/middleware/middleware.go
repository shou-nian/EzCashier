package middleware

import (
	"fmt"
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
// authorization checks for private routes.
//
// This middleware verifies if the user has the necessary credentials to access
// a protected resource. It checks the JWT token (which should be set in the Gin
// context by a previous middleware) for the user's admin status and token expiration.
// If the user is not an admin or the token has expired, the request is aborted
// with a 401 Unauthorized status.
//
// The function doesn't take any parameters directly, but it uses the Gin context
// to access the JWT token and its claims.
//
// Returns:
//
//	gin.HandlerFunc: A middleware function that can be used in a Gin router
//	                 to protect private routes.
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
		expires := claims["expires"].(float64)

		// Set credential `admin` from JWT data of current user.
		credential := claims["admin"].(bool)

		// Only user with `admin` credential can create a new user profile.
		if !credential || now >= int64(expires) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// If the user is authorized, continue to the next middleware or route handler
		c.Next()
	}
}
