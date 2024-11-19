package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/shou-nian/EzCashier/app/controllers"
	"github.com/shou-nian/EzCashier/app/middleware"
)

const (
	adminRouterPath = "/api/v1/admin"
	userRouterPath  = "/api/v1/user"
)

// PrivateRouters func for describe group of private routes.
func PrivateRouters(router *gin.Engine) {
	// Use JWT middleware.
	router.Use(middleware.PrivateAuthorizationMiddleware())
	{
		// admin private routers
		router.POST(adminRouterPath, controllers.CreateUser)
		router.PUT(adminRouterPath, controllers.UpdateUserRole)
		router.DELETE(adminRouterPath, controllers.DeleteUser)

		// user private routers
		router.PUT(userRouterPath, controllers.UpdateUserInfo)

		// auth private routers
		router.POST("/api/v1/logout", controllers.Logout)
		router.PUT("api/v1/password", controllers.UpdatePassword)
	}
}
