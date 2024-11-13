package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/shou-nian/EzCashier/app/models"
	"github.com/shou-nian/EzCashier/pkg/utils"
	"github.com/shou-nian/EzCashier/repository/database"
	"net/http"
)

func CreateUser(c *gin.Context) {
	user := &models.User{}
	request := &models.CreateUserRequest{}

	// Binding request data to struct
	if err := c.ShouldBindJSON(request); err != nil {
		validationErrors, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
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
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	user.Password = hashedPassword
	user.Role = request.Role

	db, err := database.OpenDBConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	created, err := db.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Set JWT token to response header
	jwt, err := utils.GenerateNewJWTAccessToken(created.ID, created.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.Header("jwt", jwt)

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
