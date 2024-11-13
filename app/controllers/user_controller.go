package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/shou-nian/EzCashier/app/models"
	"github.com/shou-nian/EzCashier/repository/database"
	"net/http"
	"time"
)

func CreateUser(c *gin.Context) {
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
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "You are not authorized to create a new user.",
		})
		return
	}

	user := &models.User{}
	request := &models.UserRequest{}

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
	user.Password = request.Password
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
	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully.",
		"user":    created,
	})
}

func UpdateUser(c *gin.Context) {

}
