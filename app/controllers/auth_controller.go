package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/shou-nian/EzCashier/app/models"
	"github.com/shou-nian/EzCashier/pkg/utils"
	"github.com/shou-nian/EzCashier/repository/database"
	"net/http"
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

	// Set new JWT token to response header
	jwt, err := utils.GenerateNewJWTAccessToken(user.ID, user.Role)
	if err != nil {
		panic(err)
	}
	c.Header("Authorization", jwt)

	// Return created user
	c.JSON(http.StatusCreated, gin.H{
		"message": "login successful.",
		"user":    user,
	})
}
