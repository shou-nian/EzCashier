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

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the provided information
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.CreateUserRequest true "User creation data"
// @Success 201 {object} models.User "User created successfully"
// @Failure 400 {object} gin.H "Bad request"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /api/v1/user [post]
func CreateUser(c *gin.Context) {
	user := &models.User{}
	request := &models.CreateUserRequest{}

	// Binding request data to struct
	if err := c.ShouldBindJSON(request); err != nil {
		validationErrors, _ := err.(validator.ValidationErrors)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": validationErrors.Error(),
		})
		return
	}

	// Create user from request data
	user.Name = request.Name
	user.PhoneNum = request.PhoneNum
	// Hash password
	hashedPassword, err := utils.HashPassword(request.Password)
	if err != nil {
		panic(err)
	}
	user.Password = hashedPassword
	user.Role = request.Role

	db, err := database.OpenDBConnection()
	if err != nil {
		panic(err)
	}
	created, err := db.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	// Set JWT token to response header
	token, err := utils.GenerateNewJWTAccessToken(created.ID, created.Role)
	if err != nil {
		panic(err)
	}
	c.Header("Authorization", token)

	// Return created user
	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully.",
		"user":    created,
	})
}

// UpdateUserRole godoc
// @Summary Update user role
// @Description Update the role of an existing user
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.UpdateUserRoleRequest true "User role update data"
// @Success 200 {object} models.User "User role updated successfully"
// @Failure 400 {object} gin.H "Bad request"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /api/v1/admin [put]
func UpdateUserRole(c *gin.Context) {
	user := &models.User{}
	request := &models.UpdateUserRoleRequest{}

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

	user, err = db.GetUserByPhoneNum(request.PhoneNum)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	// Update user Role
	user.Role = request.Role
	updated, err := db.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	// Generate new JWT token
	token, err := utils.GenerateNewJWTAccessToken(updated.ID, updated.Role)
	if err != nil {
		panic(err)
	}

	// Cache new JWT token to Redis
	rds, err := redis.OpenRedisConnection()
	if err != nil {
		panic(err)
	}
	expiration, _ := strconv.Atoi(os.Getenv("JWT_EXPIRES"))
	err = rds.Set(strconv.Itoa(int(updated.ID)), token, time.Duration(expiration)*60*60*time.Second)
	if err != nil {
		panic(err)
	}

	// Return updated user
	c.JSON(http.StatusOK, gin.H{
		"message": "update successfully.",
		"user":    updated,
	})
}

// UpdateUserInfo godoc
// @Summary Update user information
// @Description Update the information of the authenticated user
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.UpdateUserInfoRequest true "User info update data"
// @Success 200 {object} models.User "User information updated successfully"
// @Failure 400 {object} gin.H "Bad request"
// @Failure 500 {object} gin.H "Internal server error"
// @Security ApiKeyAuth
// @Router /api/v1/user [put]
func UpdateUserInfo(c *gin.Context) {
	user := &models.User{}
	request := &models.UpdateUserInfoRequest{}

	// Binding request data to struct
	if err := c.ShouldBindJSON(request); err != nil {
		validationErrors, _ := err.(validator.ValidationErrors)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": validationErrors.Error(),
		})
		return
	}

	// Get data from JWT
	token := c.Value("jwt").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userId := uint(claims["id"].(float64))

	db, err := database.OpenDBConnection()
	if err != nil {
		panic(err)
	}

	user, err = db.GetUser(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	// Check require updated request data
	if request.Name != "" {
		user.Name = request.Name
	}
	if request.PhoneNum != "" {
		user.PhoneNum = request.PhoneNum
	}

	// Update user info
	updated, err := db.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	// Return updated user
	c.JSON(http.StatusOK, gin.H{
		"message": "update successfully.",
		"user":    updated,
	})
}

// UpdatePassword godoc
// @Summary Update user password
// @Description Update the password of the authenticated user
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.UpdatePassword true "Password update data"
// @Success 200 {object} gin.H "Password updated successfully"
// @Failure 400 {object} gin.H "Bad request"
// @Failure 500 {object} gin.H "Internal server error"
// @Security ApiKeyAuth
// @Router /api/v1/password [put]
func UpdatePassword(c *gin.Context) {
	user := &models.User{}
	request := &models.UpdatePassword{}

	// Binding request data to struct
	if err := c.ShouldBindJSON(request); err != nil {
		validationErrors, _ := err.(validator.ValidationErrors)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": validationErrors.Error(),
		})
		return
	}

	// Get data from JWT
	token := c.Value("jwt").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userId := uint(claims["id"].(float64))

	db, err := database.OpenDBConnection()
	if err != nil {
		panic(err)
	}

	user, err = db.GetUser(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	// Password check
	if !utils.CheckPasswordHash(request.OldPassword, user.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Old password is incorrect.",
		})
		return
	}
	if request.NewPassword != request.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Confirm password does not match.",
		})
		return
	}

	// Hash new password
	hashedPassword, err := utils.HashPassword(request.NewPassword)
	if err != nil {
		panic(err)
	}
	user.Password = hashedPassword

	// Update user password
	updated, err := db.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	// Generate new JWT token
	newToken, err := utils.GenerateNewJWTAccessToken(updated.ID, updated.Role)
	if err != nil {
		panic(err)
	}

	// Cache new JWT token to Redis
	rds, err := redis.OpenRedisConnection()
	if err != nil {
		panic(err)
	}
	expiration, _ := strconv.Atoi(os.Getenv("JWT_EXPIRES"))
	err = rds.Set(strconv.Itoa(int(user.ID)), newToken, time.Duration(expiration)*60*60*time.Second)
	if err != nil {
		panic(err)
	}

	// Set new JWT token to response header
	c.Header("Authorization", newToken)

	c.JSON(http.StatusOK, gin.H{
		"message": "update password successfully.",
	})
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete an existing user
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.DeleteUser true "User deletion data"
// @Success 200 {object} gin.H "User deleted successfully"
// @Failure 400 {object} gin.H "Bad request"
// @Failure 500 {object} gin.H "Internal server error"
// @Security ApiKeyAuth
// @Router /api/v1/admin [delete]
func DeleteUser(c *gin.Context) {
	user := &models.User{}
	request := &models.DeleteUser{}

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

	user, err = db.GetUserByPhoneNum(request.PhoneNum)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	// Set user status to inactive
	user.UserStatus = models.StatusInactive
	_, err = db.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	// Delete user from database
	err = db.DeleteUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	// Delete token from redis cache
	rds, err := redis.OpenRedisConnection()
	if err != nil {
		panic(err)
	}
	err = rds.Delete(strconv.Itoa(int(user.ID)))
	if err != nil {
		panic(err)
	}

	// Clear JWT token from response header
	c.Header("Authorization", "")

	c.JSON(http.StatusOK, gin.H{
		"message": "delete user successfully.",
	})
}
