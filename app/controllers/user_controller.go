package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/shou-nian/EzCashier/app/models"
	"github.com/shou-nian/EzCashier/pkg/utils"
	"github.com/shou-nian/EzCashier/repository/database"
	"net/http"
)

// CreateUser handles the creation of a new user.
//
// It processes the incoming JSON request to create a new user account,
// hashes the password, stores the user in the database, generates a JWT token,
// and returns the created user information.
//
// Parameters:
//   - c *gin.Context: The Gin context containing the HTTP request and response.
//
// The function doesn't explicitly return a value, but it sends a JSON response
// with the HTTP status code and the created user information. It also sets
// a JWT token in the response header for authentication.
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
	jwt, err := utils.GenerateNewJWTAccessToken(created.ID, created.Role)
	if err != nil {
		panic(err)
	}
	c.Header("Authorization", jwt)

	// Return created user
	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully.",
		"user":    created,
	})
}

func UpdateUser(c *gin.Context) {

}

func UpdateUserInfo(c *gin.Context) {

}

func DeleteUser(c *gin.Context) {

}
