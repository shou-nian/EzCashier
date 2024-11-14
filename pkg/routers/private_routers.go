package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/shou-nian/EzCashier/app/controllers"
	"github.com/shou-nian/EzCashier/app/middleware"
)

const userRouterPath = "/api/v1/user"

// PrivateRouters func for describe group of private routes.
func PrivateRouters(router *gin.Engine) {
	// Use JWT middleware.
	router.Use(middleware.PrivateAuthorizationMiddleware())
	{
		// user private routers
		router.POST(userRouterPath, controllers.CreateUser)
		router.PUT(userRouterPath, controllers.UpdateUser)
		router.DELETE(userRouterPath, controllers.DeleteUser)

		// auth private routers
		router.POST("/api/v1/logout", controllers.Logout)
	}
}
