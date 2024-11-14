package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/shou-nian/EzCashier/app/models"
	"github.com/shou-nian/EzCashier/pkg/redis"
	"github.com/shou-nian/EzCashier/pkg/utils"
	"github.com/shou-nian/EzCashier/repository/database"
	"net/http"
	"os"
	"strconv"
	"time"
)

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
	jwt, err := utils.GenerateNewJWTAccessToken(user.ID, user.Role)
	if err != nil {
		panic(err)
	}

	// Cache new JWT token to Redis
	rds, err := redis.OpenRedisConnection()
	if err != nil {
		panic(err)
	}
	expiration, _ := strconv.Atoi(os.Getenv("JWT_EXPIRES"))
	err = rds.Set(user.PhoneNum, jwt, time.Duration(expiration)*60*60*time.Second)
	if err != nil {
		panic(err)
	}

	// Set new JWT token to response header
	c.Header("Authorization", jwt)

	// Return created user
	c.JSON(http.StatusCreated, gin.H{
		"message": "login successful.",
		"user":    user,
	})
}

func Logout(c *gin.Context) {
	// Clear JWT token from response header
	c.Header("Authorization", "")
	c.JSON(http.StatusOK, gin.H{"message": "logout successful."})
}
