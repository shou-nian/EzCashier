package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/shou-nian/EzCashier/app/models"
	"github.com/shou-nian/EzCashier/pkg/redis"
	"github.com/shou-nian/EzCashier/pkg/utils"
	"github.com/shou-nian/EzCashier/repository/database"
	"net/http"
	"os"
	"strconv"
	"time"
)

// Login godoc
// @Summary User login
// @Description Authenticate a user and return a JWT token
// @Tags authentication
// @Accept json
// @Produce json
// @Param login body models.LoginRequest true "Login credentials"
// @Success 200 {object} models.User "Successful login"
// @Failure 400 {object} gin.H "Bad request"
// @Failure 401 {object} gin.H "Unauthorized"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /login [post]
func Login(c *gin.Context) {
	request := &models.LoginRequest{}

	// Binding request data to struct
	if err := c.ShouldBindJSON(request); err != nil {
		validationErrors, _ := err.(validator.ValidationErrors)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": validationErrors.Error(),
		})
		return
	}

	db, err := database.OpenDBConnection()
	if err != nil {
		panic(err)
	}

	user, err := db.GetUserByPhoneNum(request.PhoneNum)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Check password is equal
	if !utils.CheckPasswordHash(request.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password."})
		return
	}

	// Generate new JWT token
	token, err := utils.GenerateNewJWTAccessToken(user.ID, user.Role)
	if err != nil {
		panic(err)
	}

	// Cache new JWT token to Redis
	rds, err := redis.OpenRedisConnection()
	if err != nil {
		panic(err)
	}
	expiration, _ := strconv.Atoi(os.Getenv("JWT_EXPIRES"))
	err = rds.Set(strconv.Itoa(int(user.ID)), token, time.Duration(expiration)*60*60*time.Second)
	if err != nil {
		panic(err)
	}

	// Set new JWT token to response header
	c.Header("Authorization", token)

	// Return created user
	c.JSON(http.StatusCreated, gin.H{
		"message": "login successful.",
		"user":    user,
	})
}

// Logout godoc
// @Summary User logout
// @Description Logout a user and invalidate their JWT token
// @Tags authentication
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} gin.H "Successful logout"
// @Failure 401 {object} gin.H "Unauthorized"
// @Router /logout [post]
func Logout(c *gin.Context) {
	token := c.Value("jwt").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	// Remove JWT token from Redis cache
	// Not cover error
	rds, _ := redis.OpenRedisConnection()
	_ = rds.Delete(strconv.Itoa(int(claims["id"].(float64))))

	// Clear JWT token from response header
	c.Header("Authorization", "")
	c.JSON(http.StatusOK, gin.H{"message": "logout successful."})
}
